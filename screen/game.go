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
	"candy/view"
)

var _ view.View = (*Game)(nil)

const backpackHeight = 94

type Game struct {
	screen
	graphics         graphics.Graphics
	spriteSheetBatch graphics.Batch
	gameMap          *gamemap.Map
	backpack         *game.BackPack
	currPlayerIndex  int
	players          []*player.Player
	rightSideBar     game.RightSideBar
	pubSub           *pubsub.PubSub
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
	g.players[g.currPlayerIndex].HandleInput(in)

	switch in.Action {
	case input.Release:
		switch in.Device {
		case input.RKey:
			g.gameMap.HideItems()
		case input.SpaceKey:
			g.dropCandy()
		}
	case input.Press:
		switch in.Device {
		case input.RKey:
			g.gameMap.RevealItems()
		}
	}
}

func (g Game) dropCandy() {
	currPlayer := g.players[g.currPlayerIndex]
	playerCell := g.getPlayerCell(*currPlayer)
	g.gameMap.AddCandy(playerCell, candy.NewBuilder(currPlayer.GetPowerLevel()))
}

func (g Game) Update(timeElapsed time.Duration) {
	g.gameMap.Update(timeElapsed)

	for _, ply := range g.players {
		ply.Update(timeElapsed)
	}
}

func (g *Game) Init() {
	g.currPlayerIndex = 0
	g.players[g.currPlayerIndex].SetIsTeammate(false)
	g.screen.Init()
}

func (g *Game) onCandyExploding(cell cell.Cell) {
	for _, ply := range g.players {
		playerCell := g.getPlayerCell(*ply)
		if ply.IsNormal() && playerCell == cell {
			ply.Trapped()
		}
	}
}

func (g Game) getPlayerCell(player player.Player) cell.Cell {
	return cell.GetCellLocatedAt(
		player.GetX()-g.gameMap.GetGridXOffset(), player.GetY()-g.gameMap.GetGridYOffset(), player.GetWidth(), player.GetHeight(),
		square.Width, square.Width,
	)
}

func NewGame(
	logger *observability.Logger,
	assets assets.Assets, g graphics.Graphics,
	pubSub *pubsub.PubSub,
) *Game {
	gameMap := gamemap.NewMap(assets, g, pubSub, 0, backpackHeight)
	playerMoveChecker := gameMap.GetPlayerMoveChecker()
	batch := g.StartNewBatch(assets.GetImage("sprite_sheet.png"))
	players := []*player.Player{
		player.NewPlayer(playerMoveChecker, player.BlackBoy, 0, backpackHeight, 1, 2),
		player.NewPlayer(playerMoveChecker, player.BlackGirl, 0, backpackHeight, 1, 3),
		player.NewPlayer(playerMoveChecker, player.BrownBoy, 0, backpackHeight, 1, 4),
		player.NewPlayer(playerMoveChecker, player.BrownGirl, 0, backpackHeight, 1, 5),
		player.NewPlayer(playerMoveChecker, player.YellowBoy, 0, backpackHeight, 1, 6),
		player.NewPlayer(playerMoveChecker, player.YellowGirl, 0, backpackHeight, 1, 7),
		player.NewPlayer(playerMoveChecker, player.OrangeBoy, 0, backpackHeight, 1, 8),
		player.NewPlayer(playerMoveChecker, player.OrangeGirl, 0, backpackHeight, 1, 9),
	}
	backpack := game.NewBackPack(g, 0, 0)
	backpack.AddItem(gameitem.FirstAidKit)
	rightSideBar := game.NewRightSideBar(gamemap.Width, 0, players)
	gm := Game{
		screen: screen{
			name:   "Game",
			logger: logger,
		},
		graphics:         g,
		spriteSheetBatch: batch,
		gameMap:          gameMap,
		players:          players,
		backpack:         &backpack,
		rightSideBar:     rightSideBar,
		pubSub:           pubSub,
	}

	pubSub.Subscribe(pubsub.OnCandyExploding, func(payload interface{}) {
		c := payload.(cell.Cell)
		gm.onCandyExploding(c)
	})
	return &gm
}
