package player

import (
	"candy/pubsub"
	"time"

	"candy/game/direction"
	"candy/game/square"
	"candy/input"
)

var nanoPerStep = (100 * time.Millisecond).Nanoseconds()

const totalSteps = 3

var _ state = (*walkingState)(nil)

type walkingState struct {
	sharedState
	stepSize int
	lag      int64
}

func (w *walkingState) update(timeElapsed time.Duration) {
	w.lag += timeElapsed.Nanoseconds()
	steps := int(w.lag / nanoPerStep)
	w.sharedState.currStep = nextStep(w.sharedState.currStep, steps)

	w.lag %= nanoPerStep
}

func (w walkingState) handleInput(in input.Input) state {
	if in.Action == input.Release {
		switch in.Device {
		case input.UpArrowKey, input.DownArrowKey, input.LeftArrowKey, input.RightArrowKey:
			return newStandingState(w.sharedState)
		}
	} else if in.Action == input.Press {
		switch in.Device {
		case input.UpArrowKey:
			return w.nextWalking(direction.Up)
		case input.DownArrowKey:
			return w.nextWalking(direction.Down)
		case input.LeftArrowKey:
			return w.nextWalking(direction.Left)
		case input.RightArrowKey:
			return w.nextWalking(direction.Right)
		}
	}
	return &w
}

func (w walkingState) nextWalking(direction direction.Direction) *walkingState {
	w.currStep = nextStep(w.sharedState.currStep, 1)
	w.sharedState = resetStepIfChangeDirection(w.sharedState, direction)

	if w.moveChecker.CanMove(w.sharedState.x, w.sharedState.y, w.playerWidth, w.playerHeight, direction, w.stepSize) {
		w.sharedState.x, w.sharedState.y = w.nextPosition(w.sharedState.x, w.sharedState.y, direction)
	}
	return newWalkingState(w.sharedState, w.lag, direction)
}

func (w walkingState) nextPosition(currX int, currY int, dir direction.Direction) (int, int) {
	switch dir {
	case direction.Up:
		return currX, currY + w.stepSize
	case direction.Down:
		return currX, currY - w.stepSize
	case direction.Left:
		return currX - w.stepSize, currY
	case direction.Right:
		return currX + w.stepSize, currY
	}
	return currX, currY
}

func resetStepIfChangeDirection(shared sharedState, direction direction.Direction) sharedState {
	if shared.direction != direction {
		shared.currStep = 1
	}
	return shared
}

func nextStep(currStep int, steps int) int {
	return (currStep + steps) % totalSteps
}

func newWalkingState(shared sharedState, lag int64, direction direction.Direction) *walkingState {
	// Check change of direction
	shared.direction = direction
	shared.pubSub.Publish(pubsub.OnPlayerWalking, pubsub.OnPlayerWalkingPayload{
		X:      shared.x,
		Y:      shared.y,
		Width:  shared.playerWidth,
		Height: shared.playerHeight,
	})
	return &walkingState{
		sharedState: shared,
		lag:         lag,
		stepSize:    square.Width / 10,
	}
}

func newWalkingStateFromStanding(shared sharedState, lag int64, direction direction.Direction) *walkingState {
	// Check change of direction
	shared.currStep = nextStep(shared.currStep, 1)
	return newWalkingState(shared, lag, direction)
}
