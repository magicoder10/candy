package player

import (
	"candy/game/direction"
	"candy/game/square"
	"candy/input"
)

var _ state = (*standingState)(nil)

type standingState struct {
	sharedState
}

func (s standingState) handleInput(in input.Input) state {
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

func newStandingStateOnSquare(
	moveChecker MoveChecker,
	playerWidth int, playerHeight int,
	row int, col int,
	regionXOffset int,
	regionYOffset int,
	walkCycleXOffset int,
	walkCycleYOffset int,
) standingState {
	return standingState{
		sharedState{
			moveChecker:      moveChecker,
			currStep:         1,
			direction:        direction.Down,
			playerWidth:      playerWidth,
			playerHeight:     playerHeight,
			x:                col*square.Width,
			y:                row * square.Width,
			regionXOffset:    regionXOffset,
			regionYOffset:    regionYOffset,
			walkCycleXOffset: walkCycleXOffset,
			walkCycleYOffset: walkCycleYOffset,
		},
	}
}

func newStandingState(shared sharedState) standingState {
	shared.currStep = 1
	return standingState{shared}
}
