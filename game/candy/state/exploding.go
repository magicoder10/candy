package state

import (
	"time"

	"candy/graphics"
)

const explodingTime = 2 * time.Second

var _ State = (*Exploding)(nil)

type Exploding struct {
	shared
	hitRange int
}

func (e *Exploding) Update(timeElapsed time.Duration) State {
	e.remainingTime -= timeElapsed
	if e.remainingTime <= 0 {
		return Exploded{}
	}
	//if e.hitRange {
	//
	//}
	return e
}

func (e Exploding) Draw(batch graphics.Batch, x int, y int, z int) {

}

func (e Exploding) Exploded() bool {
	return false
}

func newExploding(shared shared) *Exploding {
	shared.remainingTime = explodingTime
	return &Exploding{shared: shared}
}
