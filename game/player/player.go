package player

import (
	"time"

	"candy/game/gamemap"
	"candy/game/player/state"
	"candy/game/tile"
	"candy/graphics"
	"candy/input"
)

const spriteWidth = 48
const spriteHeight = 48
const spriteRowWidth = 3 * spriteWidth
const spriteColHeight = 4 * spriteHeight

type Player struct {
	regionXOffset    int
	regionYOffset    int
	walkCycleXOffset int
	walkCycleYOffset int
	state            state.State
}

func (p Player) Draw(batch graphics.Batch) {
	bound := graphics.Bound{
		X:      p.regionXOffset + p.walkCycleXOffset + p.state.GetCurrentStep()*spriteWidth,
		Y:      p.regionYOffset + p.walkCycleYOffset + int(p.state.GetDirection())*spriteHeight,
		Width:  spriteWidth,
		Height: spriteHeight,
	}
	y := p.state.GetY()
	batch.DrawSprite(p.state.GetX()-tile.Width/6, y, y, bound, float64(tile.Width)/spriteWidth)
}

func (p *Player) HandleInput(in input.Input) {
	p.state = p.state.HandleInput(in)
}

func (p Player) Update(timeElapsed time.Duration) {
	p.state.Update(timeElapsed)
}

func newPlayer(
	gameMap gamemap.Map,
	walkCycleXOffset int,
	walkCycleYOffset int,
	row int,
	col int,
) Player {
	return Player{
		regionXOffset:    0,
		regionYOffset:    0,
		walkCycleXOffset: walkCycleXOffset,
		walkCycleYOffset: walkCycleYOffset,
		state:            state.NewStandingOnTile(gameMap, tile.Width-2*tile.Width/6, tile.Height/4, row, col),
	}
}

func NewBlackBoy(gameMap gamemap.Map, row int, col int) Player {
	return newPlayer(gameMap, 0, spriteColHeight, row, col)
}

func NewBlackGirl(gameMap gamemap.Map, row int, col int) Player {
	return newPlayer(gameMap, 0, 0, row, col)
}

func NewBrownBoy(gameMap gamemap.Map, row int, col int) Player {
	return newPlayer(gameMap, spriteRowWidth, spriteColHeight, row, col)
}

func NewBrownGirl(gameMap gamemap.Map, row int, col int) Player {
	return newPlayer(gameMap, spriteRowWidth, 0, row, col)
}
func NewYellowBoy(gameMap gamemap.Map, row int, col int) Player {
	return newPlayer(gameMap, spriteRowWidth*2, spriteColHeight, row, col)
}

func NewYellowGirl(gameMap gamemap.Map, row int, col int) Player {
	return newPlayer(gameMap, spriteRowWidth*2, 0, row, col)
}

func NewOrangeBoy(gameMap gamemap.Map, row int, col int) Player {
	return newPlayer(gameMap, spriteRowWidth*3, spriteColHeight, row, col)
}

func NewOrangeGirl(gameMap gamemap.Map, row int, col int) Player {
	return newPlayer(gameMap, spriteRowWidth*3, 0, row, col)
}
