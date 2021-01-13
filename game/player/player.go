package player

import (
	"time"

	"candy/game/gamemap"
	"candy/game/square"
	"candy/graphics"
	"candy/input"
)

const spriteWidth = 48
const spriteHeight = 48
const spriteRowWidth = 3 * spriteWidth
const spriteColHeight = 4 * spriteHeight

type Player struct {
	// bottom left X of the whole section of walk cycles for all 8 players
	regionXOffset int
	// bottom right Y of the whole section of walk cycles for all 8 players
	regionYOffset int
	// bottom left X of the walk cycle for the current player
	walkCycleXOffset int
	// bottom left Y of the walk cycle for the current player
	walkCycleYOffset int
	state            state
}

func (p Player) Draw(batch graphics.Batch) {
	bound := graphics.Bound{
		X:      p.regionXOffset + p.walkCycleXOffset + p.state.GetCurrentStep()*spriteWidth,
		Y:      p.regionYOffset + p.walkCycleYOffset + int(p.state.GetDirection())*spriteHeight,
		Width:  spriteWidth,
		Height: spriteHeight,
	}
	y := p.state.GetY()
	batch.DrawSprite(p.state.GetX()-square.Width/6, y, y, bound, float64(square.Width)/spriteWidth)
}

func (p *Player) HandleInput(in input.Input) {
	p.state = p.state.HandleInput(in)
}

func (p Player) Update(timeElapsed time.Duration) {
	p.state.Update(timeElapsed)
}

func (p Player) GetX() int {
	return p.state.GetX()
}

func (p Player) GetY() int {
	return p.state.GetY()
}

func (p Player) GetWidth() int {
	return p.state.GetWidth()
}

func (p Player) GetHeight() int {
	return p.state.GetHeight()
}

func (p Player) GetPowerLevel() int {
	return 3
}

func newPlayer(
	gameMap *gamemap.Map,
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
		state:            newStandingOnSquare(gameMap, square.Width-2*square.Width/6, square.Width/4, row, col),
	}
}

func NewBlackBoy(gameMap *gamemap.Map, row int, col int) Player {
	return newPlayer(gameMap, 0, spriteColHeight, row, col)
}

func NewBlackGirl(gameMap *gamemap.Map, row int, col int) Player {
	return newPlayer(gameMap, 0, 0, row, col)
}

func NewBrownBoy(gameMap *gamemap.Map, row int, col int) Player {
	return newPlayer(gameMap, spriteRowWidth, spriteColHeight, row, col)
}

func NewBrownGirl(gameMap *gamemap.Map, row int, col int) Player {
	return newPlayer(gameMap, spriteRowWidth, 0, row, col)
}
func NewYellowBoy(gameMap *gamemap.Map, row int, col int) Player {
	return newPlayer(gameMap, spriteRowWidth*2, spriteColHeight, row, col)
}

func NewYellowGirl(gameMap *gamemap.Map, row int, col int) Player {
	return newPlayer(gameMap, spriteRowWidth*2, 0, row, col)
}

func NewOrangeBoy(gameMap *gamemap.Map, row int, col int) Player {
	return newPlayer(gameMap, spriteRowWidth*3, spriteColHeight, row, col)
}

func NewOrangeGirl(gameMap *gamemap.Map, row int, col int) Player {
	return newPlayer(gameMap, spriteRowWidth*3, 0, row, col)
}
