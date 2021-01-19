package view

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
)

var _ view = (*GameScreen)(nil)

const backpackHeight = 94

type GameScreen struct {
	graphics         graphics.Graphics
	spriteSheetBatch graphics.Batch
	gameMap          *gamemap.Map
	backpack         *game.BackPack
	currPlayer       int
	players          []*player.Player
	rightSideBar     game.RightSideBar
}

func (g GameScreen) Draw() {
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

func (g GameScreen) HandleInput(in input.Input) {
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

func (g GameScreen) dropCandy() {
	currPlayer := g.players[g.currPlayer]
	playerCell := g.getPlayerCell(*currPlayer)
	g.gameMap.AddCandy(playerCell, candy.NewBuilder(currPlayer.GetPowerLevel()))
}

func (g GameScreen) Update(timeElapsed time.Duration) {
	g.gameMap.Update(timeElapsed)

	for _, ply := range g.players {
		ply.Update(timeElapsed)
	}
}

func (g *GameScreen) Start() {
	g.currPlayer = 0
}

func (g *GameScreen) onCandyExploding(cell cell.Cell) {
	for _, ply := range g.players {
		playerCell := g.getPlayerCell(*ply)
		if ply.IsNormal() && playerCell == cell {
			ply.Trapped()
		}
	}
}

func (g GameScreen) getPlayerCell(player player.Player) cell.Cell {
	return cell.GetCellLocatedAt(
		player.GetX()-g.gameMap.GetGridXOffset(), player.GetY()-g.gameMap.GetGridYOffset(), player.GetWidth(), player.GetHeight(),
		square.Width, square.Width,
	)
}

func NewGameScreen(assets assets.Assets, g graphics.Graphics) GameScreen {
	gameMap := gamemap.NewMap(assets, g, 0, backpackHeight)
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
	gm := GameScreen{
		graphics:         g,
		spriteSheetBatch: batch,
		gameMap:          gameMap,
		players:          players,
		backpack:         &backpack,
		rightSideBar:     rightSideBar,
	}
	gameMap.OnCandyExploding(gm.onCandyExploding)
	return gm
}
