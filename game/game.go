package game

import (
	"image"
	"time"

	"candy/assets"
	"candy/game/gamemap"
	"candy/game/player"
	"candy/graphics"
	"candy/input"
)

type Game struct {
	spriteSheet image.Image
	graphics    graphics.Graphics
	inputs      <-chan input.Input
	gameMap     gamemap.Map
	currPlayer  int
	players     []player.Player
}

func (g *Game) Start(onGameEnd func()) {
	g.currPlayer = 0
	go func() {
		for in := range g.inputs {
			g.handleInput(in)
		}
	}()
}

func (g Game) handleInput(in input.Input) {
	g.players[g.currPlayer].HandleInput(in)
}

func (g Game) Update(timeElapsed time.Duration) {
	for _, ply := range g.players {
		ply.Update(timeElapsed)
	}
}

func (g Game) Draw() {
	g.gameMap.DrawMap()

	batch := g.graphics.StartNewBatch(g.spriteSheet)
	g.gameMap.DrawTiles(batch)
	for _, ply := range g.players {
		ply.Draw(batch)
	}
	batch.RenderBatch()
}

func NewGame(assets assets.Assets, graphics graphics.Graphics, inputs <-chan input.Input) Game {
	gameMap := gamemap.NewMap(assets, graphics)
	return Game{
		graphics:    graphics,
		inputs:      inputs,
		spriteSheet: assets.GetImage("sprite_sheet.png"),
		gameMap:     gameMap,
		players: []player.Player{
			player.NewBlackBoy(gameMap, 2, 1),
			player.NewBlackGirl(gameMap, 3, 1),
			player.NewBrownBoy(gameMap, 4, 1),
			player.NewBrownGirl(gameMap, 5, 1),
			player.NewYellowBoy(gameMap, 7, 1),
			player.NewOrangeBoy(gameMap, 8, 1),
			player.NewOrangeGirl(gameMap, 9, 1),
		},
	}
}
