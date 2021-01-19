package view

import (
	"time"

	"candy/input"
)

type view interface {
	Draw()
	Update(timeElapsed time.Duration)
	HandleInput(in input.Input)
}
