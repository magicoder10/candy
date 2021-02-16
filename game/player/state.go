package player

import (
	"candy/observability"
	"time"

	"candy/game/direction"
	"candy/game/square"
	"candy/graphics"
	"candy/input"
	"candy/pubsub"
)

type state interface {
	handleInput(in input.Input) state
	update(timeElapsed time.Duration)
	draw(batch graphics.Batch)
	trapped() state
	getX() int
	getY() int
	getWidth() int
	getHeight() int
	isNormal() bool
	increaseSpeed(stepSizeDelta int)
}

type sharedState struct {
	logger       *observability.Logger
	currStep     int
	direction    direction.Direction
	playerWidth  int
	playerHeight int
	x            int
	y            int
	moveChecker  MoveChecker
	regionOffset regionOffset
	character    character
	pubSub       *pubsub.PubSub
	stepSize     int
}

func (s sharedState) update(timeElapsed time.Duration) {
	return
}

func (s sharedState) isNormal() bool {
	return true
}

func (s sharedState) trapped() state {
	return newTrapState(s)
}

func (s sharedState) draw(batch graphics.Batch) {
	draw(batch,
		s.regionOffset,
		s.character,
		s.x-square.Width/6, s.y, s.y, float64(square.Width)/spriteWidth,
		s.direction, s.currStep)
}

func (s sharedState) getX() int {
	return s.x
}

func (s sharedState) getY() int {
	return s.y
}

func (s sharedState) getWidth() int {
	return s.playerWidth
}

func (s sharedState) getHeight() int {
	return s.playerHeight
}

func (s sharedState) dropCandy() {
	s.pubSub.Publish(pubsub.OnDropCandy, pubsub.OnDropCandyPayload{
		X:      s.x,
		Y:      s.y,
		Width:  s.playerWidth,
		Height: s.playerHeight,
	})
}

func (s *sharedState) increaseSpeed(stepSizeDelta int) {
	s.logger.Infof("Previous player step size: %d\n", s.stepSize)
	s.stepSize += stepSizeDelta
	s.logger.Infof("After increase player step size: %d\n", s.stepSize)
}
