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

var _ graphics.Sprite = (*Game)(nil)

type Game struct {
	spriteSheet image.Image
	gameMap     gamemap.Map
	currPlayer  int
	players     []player.Player
}

func (g Game) Draw(graphics graphics.Graphics) {
	g.gameMap.DrawMap(graphics)

	batch := graphics.StartNewBatch(g.spriteSheet)
	g.gameMap.DrawTiles(batch)
	for _, ply := range g.players {
		ply.Draw(batch)
	}
	batch.RenderBatch()
}

func (g Game) HandleInput(in input.Input) {
	g.players[g.currPlayer].HandleInput(in)
}

func (g Game) Update(timeElapsed time.Duration) {
	for _, ply := range g.players {
		ply.Update(timeElapsed)
	}
}

func (g *Game) Start() {
	g.currPlayer = 0
}

func NewGame(assets assets.Assets) Game {
	gameMap := gamemap.NewMap(assets)
	return Game{
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
