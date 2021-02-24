package candy

import (
	"candy/pubsub"
	"time"

	"candy/game/cell"

	"github.com/teamyapp/ui/graphics"
)

type state interface {
	update(timeElapsed time.Duration) state
	draw(batch graphics.Batch, x int, y int, z int)
	cellsHit() []cell.Cell
	exploding() bool
	exploded() bool
	explode()
}

type sharedState struct {
	droppedBy     int
	center        cell.Cell
	powerLevel    int
	remainingTime time.Duration
	lag           int64
	shouldExplode bool
	rangeCutter   RangeCutter
	pubSub        *pubsub.PubSub
}

func (s *sharedState) explode() {
	s.shouldExplode = true
}

func (s sharedState) cellsHit() []cell.Cell {
	return []cell.Cell{}
}

func (s sharedState) exploding() bool {
	return false
}

func (s sharedState) update(timeElapsed time.Duration) state {
	return &s
}

func (s sharedState) draw(batch graphics.Batch, x int, y int, z int) {
	return
}

func (s sharedState) exploded() bool {
	return false
}
