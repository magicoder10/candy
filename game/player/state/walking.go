package state

import (
	"time"

	"candy/game/direction"
	"candy/input"
)

const milliPerStep = 300
const totalSteps = 3

var _ State = (*Walking)(nil)

type Walking struct {
	sharedState
	stepSize int
	lag      int
}

func (w *Walking) Update(timeElapsed time.Duration) {
	w.lag += int(timeElapsed.Milliseconds())
	steps := w.lag / milliPerStep
	w.sharedState.currStep = nextStep(w.sharedState.currStep, steps)

	w.lag = w.lag % milliPerStep
}

func (w Walking) HandleInput(in input.Input) State {
	if in.Action == input.Release {
		switch in.Device {
		case input.UpArrowKey, input.DownArrowKey, input.LeftArrowKey, input.RightArrowKey:
			return NewStanding(w.sharedState)
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

func (w Walking) nextWalking(direction direction.Direction) *Walking {
	w.currStep = nextStep(w.sharedState.currStep, 1)
	w.sharedState = resetStepIfChangeDirection(w.sharedState, direction)

	if w.gameMap.CanMove(w.sharedState.x, w.sharedState.y, w.width, w.height, direction, w.stepSize) {
		w.sharedState.x, w.sharedState.y = w.nextPosition(w.sharedState.x, w.sharedState.y, direction)
	}
	return newWalking(w.sharedState, w.lag, direction)
}

func (w Walking) nextPosition(currX int, currY int, dir direction.Direction) (int, int) {
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

func newWalking(shared sharedState, lag int, direction direction.Direction) *Walking {
	// Check change of direction
	shared.direction = direction
	return &Walking{
		sharedState: shared,
		lag:         lag,
		stepSize:    10,
	}
}
