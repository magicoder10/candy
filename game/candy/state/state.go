package state

import (
	"time"

	"candy/graphics"
)

type State interface {
	Update(timeElapsed time.Duration) State
	Draw(batch graphics.Batch, x int, y int, z int)
	Exploded() bool
}

type shared struct {
	powerLevel    int
	remainingTime time.Duration
	lag           int64
}
