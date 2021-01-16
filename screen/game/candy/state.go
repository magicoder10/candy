package candy

import (
	"time"

	"candy/graphics"
	"candy/screen/game/cell"
)

type state interface {
	update(timeElapsed time.Duration) state
	draw(batch graphics.Batch, x int, y int, z int)
	getCenter() cell.Cell
	setCenter(center cell.Cell)
	cellsHit() []cell.Cell
	exploding() bool
	exploded() bool
	explode()
}

type sharedState struct {
	center        cell.Cell
	powerLevel    int
	remainingTime time.Duration
	lag           int64
	shouldExplode bool
	rangeCutter   RangeCutter
}

func (s *sharedState) setCenter(center cell.Cell) {
	s.center = center
}

func (s sharedState) getCenter() cell.Cell {
	return s.center
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
