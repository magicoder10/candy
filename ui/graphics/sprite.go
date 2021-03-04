package graphics

import (
	"time"

	"candy/input"
)

type Sprite interface {
	Update(timeElapsed time.Duration)
	HandleInput(in input.Input)
}
