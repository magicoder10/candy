package player

import (
	"time"

	"candy/game/direction"
	"candy/graphics"
	"candy/input"
)

var _ state = (*trappedState)(nil)

type trappedState struct {
	*sharedState
	jelly         *Jelly
	prevDirection direction.Direction
}

func (t trappedState) isNormal() bool {
	return false
}

func (t *trappedState) handleInput(in input.Input) state {
	return t
}

func (t trappedState) draw(batch graphics.Batch) {
	t.sharedState.draw(batch)
	t.jelly.draw(batch, t.x+t.playerWidth/2-t.jelly.width/2, t.y, t.y+jellyZOffset)
}

func (t trappedState) update(timeElapsed time.Duration) {
	t.sharedState.update(timeElapsed)
	t.jelly.update(timeElapsed)
}

func newTrapState(state *sharedState) *trappedState {
	jl := newJelly()
	prevDirection := state.direction
	state.direction = direction.Down
	return &trappedState{
		sharedState:   state,
		jelly:         &jl,
		prevDirection: prevDirection,
	}
}
