package player

import (
	"candy/input"
)

var _ state = (*trappedState)(nil)

type trappedState struct {
	sharedState
}

func (t trappedState) isNormal() bool {
	return false
}

func (t trappedState) handleInput(in input.Input) state {
	return t
}
