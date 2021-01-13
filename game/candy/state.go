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

type shared struct {
	center        cell.Cell
	powerLevel    int
	remainingTime time.Duration
	lag           int64
	shouldExplode bool
	rangeCutter   cutter.Range
}

func (s *shared) SetCenter(center cell.Cell) {
	s.center = center
}

func (s shared) GetCenter() cell.Cell {
	return s.center
}

func (s *shared) Explode() {
	s.shouldExplode = true
}

func (s shared) CellsHit() []cell.Cell {
	return []cell.Cell{}
}

func (s shared) Exploding() bool {
	return false
}

func (s shared) Update(timeElapsed time.Duration) state {
	return &s
}

func (s shared) Draw(batch graphics.Batch, x int, y int, z int) {
	return
}

func (s shared) Exploded() bool {
	return false
}
