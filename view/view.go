package view

import (
	"time"

	"github.com/teamyapp/ui/input"
)

type View interface {
	Draw()
	Update(timeElapsed time.Duration)
	HandleInput(in input.Input)
	Init()
	Destroy()
}

type CreateFactory func(props interface{}) View
