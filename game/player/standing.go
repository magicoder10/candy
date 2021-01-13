package player

import (
	"time"

	"candy/game/direction"
	"candy/game/gamemap"
	"candy/game/square"
	"candy/input"
)

var _ state = (*standingState)(nil)

type standingState struct {
	sharedState
}

func (s standingState) Update(timeElapsed time.Duration) {
	return
}

func (s standingState) HandleInput(in input.Input) state {
	if in.Action == input.Press {
		switch in.Device {
		case input.UpArrowKey:
			return newWalkingStateFromStanding(s.sharedState, 0, direction.Up)
		case input.DownArrowKey:
			return newWalkingStateFromStanding(s.sharedState, 0, direction.Down)
		case input.LeftArrowKey:
			return newWalkingStateFromStanding(s.sharedState, 0, direction.Left)
		case input.RightArrowKey:
			return newWalkingStateFromStanding(s.sharedState, 0, direction.Right)
		}
	}
	return s
}

func newStandingOnSquare(gameMap *gamemap.Map, width int, height int, row int, col int) standingState {
	return standingState{
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

func newStandingState(shared sharedState) standingState {
	shared.currStep = 1
	return standingState{shared}
}
