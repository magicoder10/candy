package state

import (
	"time"

	"candy/game/cell"
	"candy/graphics"
)

var _ State = (*Exploded)(nil)

type Exploded struct {
	shared
}

func (e Exploded) CellsHit() []cell.Cell {
	return []cell.Cell{}
}

func (e Exploded) Exploding() bool {
	return false
}

func (e Exploded) Update(timeElapsed time.Duration) State {
	return &e
}

func (e Exploded) Draw(batch graphics.Batch, x int, y int, z int) {
	return
}

func (e Exploded) Exploded() bool {
	return true
}
