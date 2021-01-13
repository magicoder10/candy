package state

import (
	"time"

	"candy/game/cell"
	"candy/game/direction"
	"candy/game/square"
	"candy/graphics"
)

const explodingTime = 400 * time.Millisecond

var explodingTimeNano = explodingTime.Nanoseconds()

type explodeDirection struct {
	cell      cell.Cell
	edge      graphics.Bound
	end       graphics.Bound
	direction direction.Direction
}

var directions = []explodeDirection{
	{
		cell:      cell.Cell{Row: 1, Col: 0},
		edge:      explosionVerticalEdge,
		end:       explosionTopEnd,
		direction: direction.Up,
	},
	{
		cell:      cell.Cell{Row: -1, Col: 0},
		edge:      explosionVerticalEdge,
		end:       explosionBottomEnd,
		direction: direction.Down,
	},
	{
		cell:      cell.Cell{Row: 0, Col: 1},
		edge:      explosionHorizontalEdge,
		end:       explosionRightEnd,
		direction: direction.Right,
	},
	{
		cell:      cell.Cell{Row: 0, Col: -1},
		edge:      explosionHorizontalEdge,
		end:       explosionLeftEnd,
		direction: direction.Left,
	},
}
var _ State = (*Exploding)(nil)

type Exploding struct {
	shared
	hitRange int
}

func (e Exploding) Exploding() bool {
	return true
}

func (e Exploding) CellsHit() []cell.Cell {
	cells := make([]cell.Cell, 0)
	for currRange := 1; currRange <= e.hitRange; currRange++ {
		for _, dir := range directions {
			nextRow := e.center.Row + currRange*dir.cell.Row
			nextCol := e.center.Col + currRange*dir.cell.Col
			cells = append(cells, cell.Cell{Row: nextRow, Col: nextCol})
		}
	}
	return cells
}

func (e *Exploding) Update(timeElapsed time.Duration) State {
	e.remainingTime -= timeElapsed
	if e.remainingTime <= 0 {
		return &Exploded{}
	}
	e.lag += timeElapsed.Nanoseconds()

	rangeIncreaseDuration := explodingTimeNano / int64(e.powerLevel)
	for e.hitRange < e.powerLevel && e.lag > rangeIncreaseDuration {
		e.hitRange += 1
		e.lag -= rangeIncreaseDuration
	}
	return e
}

func (e Exploding) Draw(batch graphics.Batch, x int, y int, z int) {
	e.drawEnds(batch, x, y, z)
	e.drawEdges(batch, x, y, z)
	batch.DrawSprite(x, y, z, explosionCenter, 1)
}

func (e Exploding) drawEdges(batch graphics.Batch, x int, y int, z int) {
	for _, dir := range directions {
		hitRange := e.rangeCutter.CutRange(e.center, e.hitRange, dir.direction)
		for i := 1; i < hitRange; i++ {
			shift := i * square.Width
			nextX := x + dir.cell.Col*shift
			nextY := y + dir.cell.Row*shift
			batch.DrawSprite(nextX, nextY, nextY, dir.edge, 1)
		}
	}
}

func (e Exploding) drawEnds(batch graphics.Batch, x int, y int, z int) {
	if e.hitRange < 1 {
		return
	}

	for _, dir := range directions {
		hitRange := e.rangeCutter.CutRange(e.center, e.hitRange, dir.direction)
		if hitRange < 1 {
			continue
		}
		shift := hitRange * square.Width
		nextX := x + dir.cell.Col*shift
		nextY := y + dir.cell.Row*shift
		batch.DrawSprite(nextX, nextY, nextY, dir.end, 1)
	}
}

func (e Exploding) Exploded() bool {
	return false
}

func newExploding(shared shared) *Exploding {
	shared.remainingTime = explodingTime + 500*time.Millisecond
	return &Exploding{shared: shared, hitRange: 0}
}
