package state

import (
	"time"

	"candy/graphics"
)

var _ State = (*Exploded)(nil)

type Exploded struct {
}

func (e Exploded) Update(timeElapsed time.Duration) State {
	return e
}

func (e Exploded) Draw(batch graphics.Batch, x int, y int, z int) {
	return
}

func (e Exploded) Exploded() bool {
	return true
}
