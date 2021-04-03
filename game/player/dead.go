package player

import (
	"candy/game/direction"
	"candy/graphics"
	"candy/input"
	"time"
)

var _ state = (*deadState)(nil)
var TombstoneImageDuration = (3 * time.Second).Nanoseconds()

type deadState struct {
	*sharedState
	tombstone *Tombstone
	lag       int64
}

func (d *deadState) handleInput(in input.Input) state {
	return d
}

func (d deadState) draw(batch graphics.Batch) {
	if d.lag <= TombstoneImageDuration {
		d.tombstone.draw(batch, d.x+d.playerWidth/2-d.tombstone.width/2, d.y, d.y)
	}
}

func (d *deadState) update(timeElapsed time.Duration) state {
	d.lag += timeElapsed.Nanoseconds()
	return d
}

func newDeadState(state *sharedState) *deadState {
	ts := newTombstone()
	state.direction = direction.Down
	return &deadState{
		sharedState: state,
		tombstone:   &ts,
		lag:         0,
	}
}
