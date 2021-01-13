package candy

import (
	"time"

	"candy/game/cell"
	"candy/game/cutter"
	"candy/graphics"
)

type state interface {
	Update(timeElapsed time.Duration) state
	Draw(batch graphics.Batch, x int, y int, z int)
	GetCenter() cell.Cell
	SetCenter(center cell.Cell)
	CellsHit() []cell.Cell
	Exploding() bool
	Exploded() bool
	Explode()
}

type sharedState struct {
	center        cell.Cell
	powerLevel    int
	remainingTime time.Duration
	lag           int64
	shouldExplode bool
	rangeCutter   cutter.Range
}

func (s *sharedState) SetCenter(center cell.Cell) {
	s.center = center
}

func (s sharedState) GetCenter() cell.Cell {
	return s.center
}

func (s *sharedState) Explode() {
	s.shouldExplode = true
}

func (s sharedState) CellsHit() []cell.Cell {
	return []cell.Cell{}
}

func (s sharedState) Exploding() bool {
	return false
}

func (s sharedState) Update(timeElapsed time.Duration) state {
	return &s
}

func (s sharedState) Draw(batch graphics.Batch, x int, y int, z int) {
	return
}

func (s sharedState) Exploded() bool {
	return false
}
