package player

import (
	"candy/input"
	"candy/screen/game/direction"
	"candy/screen/game/square"
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
	gridX int, gridY int,
	row int, col int,
	regionOffset regionOffset,
	character character,
) standingState {
	return standingState{
		sharedState{
			moveChecker:  moveChecker,
			currStep:     1,
			direction:    direction.Down,
			playerWidth:  playerWidth,
			playerHeight: playerHeight,
			x:            gridX + col*square.Width,
			y:            gridY + row*square.Width,
			regionOffset: regionOffset,
			character:    character,
		},
	}
}

func newStandingState(shared sharedState) standingState {
	shared.currStep = 1
	return standingState{shared}
}
