package player

import (
	"time"

	"candy/game/direction"
	"candy/input"
	"candy/pubsub"
)

var nanoPerStep = (100 * time.Millisecond).Nanoseconds()

const totalSteps = 3

var _ state = (*walkingState)(nil)

type walkingState struct {
	sharedState
	lag int64
}

func (w *walkingState) update(timeElapsed time.Duration) {
	w.lag += timeElapsed.Nanoseconds()
	steps := int(w.lag / nanoPerStep)
	w.sharedState.currStep = nextStep(w.sharedState.currStep, steps)

	for step := 0; step < steps; step++ {
		w.walk()
	}

	w.lag %= nanoPerStep
}

func (w *walkingState) handleInput(in input.Input) state {
	switch in.Action {
	case input.Release:
		switch in.Device {
		case input.UpArrowKey, input.DownArrowKey, input.LeftArrowKey, input.RightArrowKey:
			return newStandingState(w.sharedState)
		}
	case input.SinglePress:
		switch in.Device {
		case input.UpArrowKey:
			w.direction = direction.Up
		case input.DownArrowKey:
			w.direction = direction.Down
		case input.LeftArrowKey:
			w.direction = direction.Left
		case input.RightArrowKey:
			w.direction = direction.Right
		case input.SpaceKey:
			w.dropCandy()
		}
	}
	return w
}

func (w *walkingState) walk() {
	w.currStep = nextStep(w.sharedState.currStep, 1)
	w.sharedState = resetStepIfChangeDirection(w.sharedState, w.direction)

	if w.moveChecker.CanMove(w.sharedState.x, w.sharedState.y, w.playerWidth, w.playerHeight, w.direction, w.stepSize) {
		w.sharedState.x, w.sharedState.y = w.nextPosition(w.sharedState.x, w.sharedState.y, w.direction)
	}
}

func (w *walkingState) nextPosition(currX int, currY int, dir direction.Direction) (int, int) {
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
	w := &walkingState{
		sharedState: shared,
		lag:         lag,
	}
	w.walk()
	return w
}

func newWalkingStateFromStanding(shared sharedState, lag int64, direction direction.Direction) *walkingState {
	// Check change of direction
	shared.currStep = nextStep(shared.currStep, 1)
	return newWalkingState(shared, lag, direction)
}
