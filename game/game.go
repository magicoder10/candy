package game

import (
	"time"

	"candy/assets"
	"candy/game/candy"
	"candy/game/cell"
	"candy/game/gamemap"
	"candy/game/player"
	"candy/game/square"
	"candy/graphics"
	"candy/input"
)

var _ graphics.Sprite = (*Game)(nil)

type Game struct {
	spriteSheetBatch graphics.Batch
	gameMap          *gamemap.Map
	currPlayer       int
	players          []*player.Player
}

func (g Game) Draw() {
	g.gameMap.DrawMap()

	g.gameMap.DrawGrid(g.spriteSheetBatch)
	for _, ply := range g.players {
		ply.Draw(g.spriteSheetBatch)
	}
	g.spriteSheetBatch.RenderBatch()
}

func (g Game) HandleInput(in input.Input) {
	g.players[g.currPlayer].HandleInput(in)

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
	currPlayer := g.players[g.currPlayer]
	playerCell := getPlayerCell(*currPlayer)
	g.gameMap.AddCandy(playerCell, candy.NewBuilder(currPlayer.GetPowerLevel()))
}

func (g Game) Update(timeElapsed time.Duration) {
	g.gameMap.Update(timeElapsed)

	for _, ply := range g.players {
		ply.Update(timeElapsed)
	}
}

func (g *Game) Start() {
	g.currPlayer = 0
}

func (g *Game) onCandyExploding(cell cell.Cell) {
	for _, ply := range g.players {
		playerCell := getPlayerCell(*ply)
		if ply.IsNormal() && playerCell == cell {
			ply.Trapped()
		}
	}
}

func getPlayerCell(player player.Player) cell.Cell {
	return cell.GetCellLocatedAt(
		player.GetX(), player.GetY(), player.GetWidth(), player.GetHeight(),
		square.Width, square.Width,
	)
}

func NewGame(assets assets.Assets, g graphics.Graphics) Game {
	gameMap := gamemap.NewMap(assets, g)
	playerMoveChecker := gameMap.GetPlayerMoveChecker()
	batch := g.StartNewBatch(assets.GetImage("sprite_sheet.png"))
	players := []*player.Player{
		player.NewBlackBoy(playerMoveChecker, 1, 2),
		player.NewBlackGirl(playerMoveChecker, 1, 3),
		player.NewBrownBoy(playerMoveChecker, 1, 4),
		player.NewBrownGirl(playerMoveChecker, 1, 5),
		player.NewYellowBoy(playerMoveChecker, 1, 6),
		player.NewYellowGirl(playerMoveChecker, 1, 7),
		player.NewOrangeBoy(playerMoveChecker, 1, 8),
		player.NewOrangeGirl(playerMoveChecker, 1, 9),
	}
	gm := Game{
		spriteSheetBatch: batch,
		gameMap:          gameMap,
		players:          players,
	}
	gameMap.OnCandyExploding(gm.onCandyExploding)
	return gm
}
