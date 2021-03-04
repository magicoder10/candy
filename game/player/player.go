package player

import (
	"time"

	"candy/game/direction"
	"candy/game/marker"
	"candy/game/square"
	"candy/graphics"
	"candy/input"
	"candy/pubsub"
)

const spriteWidth = 48
const spriteHeight = 48
const spriteRowWidth = 3 * spriteWidth
const spriteColHeight = 4 * spriteHeight
const bodyWidth = (2*square.Width - 10) / 3
const feetLength = square.Width / 4

const jellyZOffset = -1

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
	marker       *marker.Marker
}

func (p Player) Draw(batch graphics.Batch) {
	p.state.draw(batch)
	if p.marker != nil {
		markerX := p.state.getX() + p.state.getWidth()/2 - p.marker.GetWidth()/2
		// TODO: player.state.getHeight() should rename to getDepth()
		// TODO: square.Width need to be replaced with new getHeight()
		markerY := p.state.getY() + square.Width + marker.YOffset
		p.marker.Draw(batch, markerX, markerY, p.state.getY()+jellyZOffset-1)
	}
}

func (p Player) DrawStand(batch graphics.Batch, x int, y int, z int, scale float64) {
	draw(batch, p.regionOffset, p.character, x, y, z, scale, direction.Down, 1)
}

func (p *Player) HandleInput(in input.Input) {
	p.state = p.state.handleInput(in)
}

func (p Player) Update(timeElapsed time.Duration) {
	p.state.update(timeElapsed)
	if p.marker != nil {
		p.marker.Update(timeElapsed)
	}
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

func (p Player) IsNormal() bool {
	return p.state.isNormal()
}

func (p *Player) Trapped() {
	p.state = p.state.trapped()
}

func (p *Player) ShowMarker(isTeammate bool) {
	mk := marker.NewMarker(isTeammate)
	p.marker = &mk
}

func (p *Player) IncreasePowerLevel(amountIncrease int) {
	p.state.increasePowerLevel(amountIncrease)
}

func (p *Player) IncreaseStepSize(amountIncrease int) {
	p.state.increaseStepSize(amountIncrease)
}

func (p *Player) IncreaseCandyLimit(amountIncrease int) {
	p.state.increaseCandyLimit(amountIncrease)
}

func (p *Player) IncrementAvailableCandy() {
	p.state.incrementAvailableCandy()
}

func NewPlayer(
	dropCandyChecker DropCandyChecker,
	moveChecker MoveChecker,
	character character,
	gridX int,
	gridY int,
	row int,
	col int,
	pubSub *pubsub.PubSub,
) *Player {
	return &Player{
		regionOffset: regionOffset{
			x: 0,
			y: 0,
		},
		character: character,
		state: newStandingStateOnSquare(
			dropCandyChecker, moveChecker,
			bodyWidth, feetLength,
			gridX, gridY,
			row, col,
			regionOffset{
				x: 0,
				y: 0,
			},
			character,
			pubSub,
		),
	}
}
