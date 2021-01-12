package state

import (
	"time"

	"candy/game/direction"
	"candy/game/gamemap"
	"candy/game/square"
	"candy/input"
)

var _ State = (*Standing)(nil)

type Standing struct {
	sharedState
}

func (s Standing) Update(timeElapsed time.Duration) {
	return
}

func (s Standing) HandleInput(in input.Input) State {
	if in.Action == input.Press {
		switch in.Device {
		case input.UpArrowKey:
			return newWalkingFromStanding(s.sharedState, 0, direction.Up)
		case input.DownArrowKey:
			return newWalkingFromStanding(s.sharedState, 0, direction.Down)
		case input.LeftArrowKey:
			return newWalkingFromStanding(s.sharedState, 0, direction.Left)
		case input.RightArrowKey:
			return newWalkingFromStanding(s.sharedState, 0, direction.Right)
		}
	}
	return s
}

func NewStandingOnSquare(gameMap *gamemap.Map, width int, height int, row int, col int) Standing {
	return Standing{
		sharedState{
			gameMap:   gameMap,
			currStep:  1,
			direction: direction.Down,
			width:     width,
			height:    height,
			x:         col*square.Width + square.Width/2,
			y:         row * square.Width,
		},
	}
}

func NewStanding(shared sharedState) Standing {
	shared.currStep = 1
	return Standing{shared}
}
