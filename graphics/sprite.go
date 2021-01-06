package graphics

import (
	"time"

	"candy/input"
)

type Sprite interface {
	Draw(graphics Graphics)
	Update(timeElapsed time.Duration)
	HandleInput(in input.Input)
}
