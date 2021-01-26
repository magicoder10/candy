package screen

import (
	"time"

	"candy/assets"
	"candy/game"
	"candy/game/candy"
	"candy/game/cell"
	"candy/game/gameitem"
	"candy/game/gamemap"
	"candy/game/player"
	"candy/game/square"
	"candy/graphics"
	"candy/input"
	"candy/observability"
	"candy/pubsub"
	"candy/server/gamestate"
	"candy/view"
)

var _ view.View = (*Game)(nil)

const backpackHeight = 94

type Game struct {
	gameID string
	screen
	graphics         graphics.Graphics
	spriteSheetBatch graphics.Batch
	gameMap          *gamemap.Map
	backpack         *game.BackPack
	currPlayerID     string
	players          map[string]*player.Player
	rightSideBar     game.RightSideBar
	pubSub           *pubsub.PubSub
	remotePubSub     *pubsub.Remote
}

func (g Game) Draw() {
	g.gameMap.DrawMap()

	g.gameMap.DrawGrid(g.spriteSheetBatch)
	for _, ply := range g.players {
		ply.Draw(g.spriteSheetBatch)
	}

	g.backpack.Draw(g.spriteSheetBatch)
	g.rightSideBar.Draw(g.spriteSheetBatch)

	g.spriteSheetBatch.RenderBatch()
	g.graphics.RenderTexts()
}

func (g Game) HandleInput(in input.Input) {
	if currPlayer, ok := g.players[g.currPlayerID]; ok {
		currPlayer.HandleInput(in)
	}

	switch in.Action {
	case input.Release:
		switch in.Device {
		case input.RKey:
			g.gameMap.HideItems()
		}
	case input.Press:
		switch in.Device {
		case input.RKey:
			g.gameMap.RevealItems()
		}
	}
}

func (g Game) dropCandy(payload pubsub.OnDropCandyPayload) {
	playerCell := g.getObjectCell(payload.X, payload.Y, payload.Width, payload.Height)
	currPlayer := g.players[g.currPlayerID]

	succeed := g.gameMap.AddCandy(playerCell, candy.NewBuilder(currPlayer.GetPowerLevel()))
	if succeed {
		g.remotePubSub.Publish(pubsub.NewSyncDropCandy(g.gameID), gamestate.Candy{
			Cell: gamestate.Cell{
				Row: playerCell.Row,
				Col: playerCell.Col,
			},
			PowerLevel: currPlayer.GetPowerLevel(),
		})
	}
}

func (g Game) Update(timeElapsed time.Duration) {
	g.gameMap.Update(timeElapsed)

	for _, ply := range g.players {
		ply.Update(timeElapsed)
	}
}

func (g *Game) Init() {
	g.players[g.currPlayerID].ShowMarker(false)
	g.screen.Init()

	g.pubSub.Subscribe(pubsub.OnCandyExploding, func(payload interface{}) {
		c := payload.(cell.Cell)
		g.onCandyExploding(c)
	})
	g.pubSub.Subscribe(pubsub.OnDropCandy, func(payload interface{}) {
		pl := payload.(pubsub.OnDropCandyPayload)
		g.dropCandy(pl)
	})
	g.pubSub.Subscribe(pubsub.OnPlayerWalking, func(payload interface{}) {
		p := payload.(pubsub.OnPlayerWalkingPayload)
		g.onPlayerWalking(p)
	})
	g.pubSub.Subscribe(pubsub.IncrementPlayerPower, func(_ interface{}) {
		g.incrementPlayerPower()
	})
	g.remotePubSub.Subscribe(pubsub.NewSyncDropCandy(g.gameID), func(payload []byte) {
		cdState, err := gamestate.GetCandy(payload)
		if err != nil {
			return
		}
		g.gameMap.AddCandy(cell.Cell{Row: cdState.Row, Col: cdState.Col}, candy.NewBuilder(cdState.PowerLevel))
	})
	g.remotePubSub.Subscribe(pubsub.NewSyncRetrieveGameItem(g.gameID), func(payload []byte) {
		c, err := gamestate.GetCell(payload)
		if err != nil {
			return
		}
		g.gameMap.RetrieveGameItem(cell.Cell{Row: c.Row, Col: c.Col})
	})
	g.syncPlayerMoves()
}

func (g *Game) onCandyExploding(cell cell.Cell) {
	for _, ply := range g.players {
		playerCell := g.getPlayerCell(ply)
		if ply.IsNormal() && playerCell == cell {
			ply.Trapped()
		}
	}
}

func (g Game) getPlayerCell(player *player.Player) cell.Cell {
	return g.getObjectCell(player.GetX(), player.GetY(), player.GetWidth(), player.GetHeight())
}

func (g Game) getObjectCell(objectX int, objectY int, objectWidth int, objectHeight int) cell.Cell {
	return cell.GetCellLocatedAt(
		objectX-g.gameMap.GetGridXOffset(), objectY-g.gameMap.GetGridYOffset(), objectWidth, objectHeight,
		square.Width, square.Width,
	)
}

func (g Game) onPlayerWalking(payload pubsub.OnPlayerWalkingPayload) {
	c := g.getObjectCell(payload.X, payload.Y, payload.Width, payload.Height)

	if g.gameMap.HasRevealedItem(c) {
		gameItemType := g.gameMap.RetrieveGameItem(c)
		item := gameitem.WithPubSub(gameItemType, g.pubSub)

		g.backpack.AddItem(item)
		g.remotePubSub.Publish(pubsub.NewSyncRetrieveGameItem(g.gameID), gamestate.Cell{
			Row: c.Row,
			Col: c.Col,
		})
	}
	g.remotePubSub.Publish(pubsub.NewSyncPlayerMove(g.gameID, payload.PlayerID), gamestate.PlayerState{
		X:           payload.X,
		Y:           payload.Y,
		CurrentStep: payload.CurrStep,
		Direction:   payload.Direction,
	})
}

func (g Game) incrementPlayerPower() {
	g.players[g.currPlayerID].IncrementPower()
}

func (g Game) syncPlayerMoves() {
	for id, ply := range g.players {
		if id == g.currPlayerID {
			continue
		}
		ply.SyncMove(g.gameID)
	}
}

func NewGame(
	logger *observability.Logger,
	assets assets.Assets, g graphics.Graphics,
	pubSub *pubsub.PubSub,
	remotePubSub *pubsub.Remote,
	gameID string,
	gameSetup gamestate.Setup,
	currPlayerID string,
) *Game {
	gameMap := gamemap.NewMap(assets, g, pubSub, gameSetup, 0, backpackHeight)
	playerMoveChecker := gameMap.GetPlayerMoveChecker()
	batch := g.StartNewBatch(assets.GetImage("sprite_sheet.png"))

	players := make(map[string]*player.Player)

	for playerID, plyState := range gameSetup.Players {
		ply := player.NewPlayer(
			playerID,
			playerMoveChecker,
			player.NewCharacter(plyState.Character),
			0,
			backpackHeight,
			pubSub, remotePubSub,
			plyState.Position.X,
			plyState.Position.Y,
		)
		players[playerID] = ply
	}
	backpack := game.NewBackPack(g, 0, 0)
	rightSideBar := game.NewRightSideBar(gamemap.Width, 0, players)
	gm := Game{
		gameID: gameID,
		screen: screen{
			name:   "Game",
			logger: logger,
		},
		graphics:         g,
		spriteSheetBatch: batch,
		gameMap:          gameMap,
		currPlayerID:     currPlayerID,
		players:          players,
		backpack:         &backpack,
		rightSideBar:     rightSideBar,
		pubSub:           pubSub,
		remotePubSub:     remotePubSub,
	}
	return &gm
}
