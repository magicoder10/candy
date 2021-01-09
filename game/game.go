package game

import (
	"time"

	"candy/assets"
	"candy/game/gamemap"
	"candy/game/player"
	"candy/graphics"
	"candy/input"
)

var _ graphics.Sprite = (*Game)(nil)

type Game struct {
	spriteSheetBatch graphics.Batch
	gameMap          gamemap.Map
	currPlayer       int
	players          []player.Player
}

func (g Game) Draw() {
	g.gameMap.DrawMap()

	g.gameMap.DrawTiles(g.spriteSheetBatch)
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
		}
	case input.Press:
		switch in.Device {
		case input.RKey:
			g.gameMap.RevealItems()
		}
	}
}

func (g Game) Update(timeElapsed time.Duration) {
	for _, ply := range g.players {
		ply.Update(timeElapsed)
	}
}

func (g *Game) Start() {
	g.currPlayer = 0
}

func NewGame(assets assets.Assets, g graphics.Graphics) Game {
	gameMap := gamemap.NewMap(assets, g)
	batch := g.StartNewBatch(assets.GetImage("sprite_sheet.png"))
	return Game{
		spriteSheetBatch: batch,
		gameMap:          gameMap,
		players: []player.Player{
			player.NewBlackBoy(gameMap, 1, 2),
			player.NewBlackGirl(gameMap, 1, 3),
			player.NewBrownBoy(gameMap, 1, 4),
			player.NewBrownGirl(gameMap, 1, 5),
			player.NewYellowBoy(gameMap, 1, 6),
			player.NewYellowGirl(gameMap, 1, 7),
			player.NewOrangeBoy(gameMap, 1, 8),
			player.NewOrangeGirl(gameMap, 1, 9),
		},
	}
}
