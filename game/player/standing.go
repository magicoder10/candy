package player

import (
	"candy/game/direction"
	"candy/game/square"
	"candy/input"
	"candy/pubsub"
)

var _ state = (*standingState)(nil)

type standingState struct {
	*sharedState
}

func (s standingState) handleInput(in input.Input) state {
	switch in.Action {
	case input.Press:
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
	case input.SinglePress:
		switch in.Device {
		case input.SpaceKey:
			s.dropCandy()
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
	pubSub *pubsub.PubSub,
) standingState {
	return standingState{
		&sharedState{
			moveChecker:    moveChecker,
			currStep:       1,
			direction:      direction.Down,
			playerWidth:    playerWidth,
			playerHeight:   playerHeight,
			x:              gridX + col*square.Width,
			y:              gridY + row*square.Width,
			regionOffset:   regionOffset,
			powerLevel:     character.initialPower,
			stepSize:       character.initialStepSize,
			candyLimit:     character.initialCandyLimit,
			candyAvailable: character.initialCandyLimit,
			character:      character,
			pubSub:         pubSub,
		},
	}
}

func newStandingState(shared *sharedState) *standingState {
	shared.currStep = 1
	return &standingState{sharedState: shared}
}
