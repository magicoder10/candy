package player

import (
	"time"

	"candy/game/direction"
	"candy/graphics"
	"candy/input"
)

var trapTimeOut = (5 * time.Second).Nanoseconds()

var _ state = (*trappedState)(nil)

type trappedState struct {
	*sharedState
	jelly         *Jelly
	prevDirection direction.Direction
	trappedLag    int64
}

func (t trappedState) isNormal() bool {
	return false
}

func (t *trappedState) handleInput(_ input.Input) state {
	return t
}

func (t trappedState) draw(batch graphics.Batch) {
	t.sharedState.draw(batch)
	t.jelly.draw(batch, t.x+t.playerWidth/2-t.jelly.width/2, t.y, t.y+jellyZOffset)
}

func (t *trappedState) update(timeElapsed time.Duration) state {
	t.trappedLag += timeElapsed.Nanoseconds()
	if t.trappedLag >= trapTimeOut {
		t.showMarker = false
		return newDeadState(t.sharedState)
	}

	t.jelly.update(timeElapsed)
	return t
}

func newTrapState(state *sharedState) *trappedState {
	jl := newJelly()
	prevDirection := state.direction
	state.direction = direction.Down
	return &trappedState{
		sharedState:   state,
		jelly:         &jl,
		prevDirection: prevDirection,
		trappedLag:    0,
	}
}
