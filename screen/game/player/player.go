package player

import (
	"time"

	"candy/graphics"
	"candy/input"
	"candy/screen/game/direction"
	"candy/screen/game/square"
)

const spriteWidth = 48
const spriteHeight = 48
const spriteRowWidth = 3 * spriteWidth
const spriteColHeight = 4 * spriteHeight
const bodyWidth = 2 * square.Width / 3
const feetLength = square.Width / 4

type regionOffset struct {
	x int
	y int
}

type walkCycleOffset struct {
	// bottom left X of the walk cycle for the current player
	x int
	// bottom left Y of the walk cycle for the current player
	y int
}

type Player struct {
	state        state
	regionOffset regionOffset
	character    character
}

func (p Player) Draw(batch graphics.Batch) {
	p.state.draw(batch)
}

func (p Player) DrawStand(batch graphics.Batch, x int, y int, z int, scale float64) {
	draw(batch, p.regionOffset, p.character, x, y, z, scale, direction.Down, 1)
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

func (p Player) IsNormal() bool {
	return p.state.isNormal()
}

func (p *Player) Trapped() {
	p.state = p.state.trapped()
}

func NewPlayer(
	moveChecker MoveChecker,
	character character,
	gridX int,
	gridY int,
	row int,
	col int,
) *Player {
	return &Player{
		regionOffset: regionOffset{
			x: 0,
			y: 0,
		},
		character: character,
		state: newStandingStateOnSquare(
			moveChecker, bodyWidth, feetLength,
			gridX, gridY,
			row, col,
			regionOffset{
				x: 0,
				y: 0,
			},
			character,
		),
	}
}
