package state

import (
	"time"

	"candy/game/cell"
	"candy/game/cutter"
	"candy/graphics"
)

type State interface {
	Update(timeElapsed time.Duration) State
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
