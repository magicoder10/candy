package player

import (
	"time"

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
	p.state.draw(batch, p.regionXOffset, p.regionYOffset, p.walkCycleXOffset, p.walkCycleYOffset)
}

func (p *Player) HandleInput(in input.Input) {
	p.state = p.state.handleInput(in)
}

func (p Player) Update(timeElapsed time.Duration) {
	p.state.update(timeElapsed)
}

func (p Player) GetX() int {
	return p.state.getX()
}

func (p Player) GetY() int {
	return p.state.getY()
}

func (p Player) GetWidth() int {
	return p.state.getWidth()
}

func (p Player) GetHeight() int {
	return p.state.getHeight()
}

func (p Player) GetPowerLevel() int {
	return 3
}

func newPlayer(
	moveChecker MoveChecker,
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
		state:            newStandingStateOnSquare(moveChecker, square.Width-2*square.Width/6, square.Width/4, row, col),
	}
}

func NewBlackBoy(moveChecker MoveChecker, row int, col int) Player {
	return newPlayer(moveChecker, 0, spriteColHeight, row, col)
}

func NewBlackGirl(moveChecker MoveChecker, row int, col int) Player {
	return newPlayer(moveChecker, 0, 0, row, col)
}

func NewBrownBoy(moveChecker MoveChecker, row int, col int) Player {
	return newPlayer(moveChecker, spriteRowWidth, spriteColHeight, row, col)
}

func NewBrownGirl(moveChecker MoveChecker, row int, col int) Player {
	return newPlayer(moveChecker, spriteRowWidth, 0, row, col)
}
func NewYellowBoy(moveChecker MoveChecker, row int, col int) Player {
	return newPlayer(moveChecker, spriteRowWidth*2, spriteColHeight, row, col)
}

func NewYellowGirl(moveChecker MoveChecker, row int, col int) Player {
	return newPlayer(moveChecker, spriteRowWidth*2, 0, row, col)
}

func NewOrangeBoy(moveChecker MoveChecker, row int, col int) Player {
	return newPlayer(moveChecker, spriteRowWidth*3, spriteColHeight, row, col)
}

func NewOrangeGirl(moveChecker MoveChecker, row int, col int) Player {
	return newPlayer(moveChecker, spriteRowWidth*3, 0, row, col)
}
