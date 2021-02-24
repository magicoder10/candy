package player

import (
	"time"

	"candy/game/direction"
	"candy/game/square"
	"candy/pubsub"

	"github.com/teamyapp/ui/graphics"
	"github.com/teamyapp/ui/input"
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
	getCurCandyAmount() int
	getMaxCandyAmount() int
	increasePowerLevel(amountIncrease int)
	increaseStepSize(amountIncrease int)
	incrementCandyAvailable()
	isNormal() bool
}

type sharedState struct {
	currStep       int
	direction      direction.Direction
	playerWidth    int
	playerHeight   int
	x              int
	y              int
	moveChecker    MoveChecker
	regionOffset   regionOffset
	powerLevel     int
	stepSize       int
	candyAvailable int
	candyLimit     int
	character      character
	pubSub         *pubsub.PubSub
}

func (s sharedState) update(timeElapsed time.Duration) {
	return
}

func (s sharedState) isNormal() bool {
	return true
}

func (s *sharedState) trapped() state {
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

func (s sharedState) getCurCandyAmount() int {
	return s.candyAvailable
}

func (s sharedState) getMaxCandyAmount() int {
	return s.candyLimit
}

func (s *sharedState) increasePowerLevel(amountIncrease int) {
	s.powerLevel += amountIncrease
}

func (s *sharedState) increaseStepSize(amountIncrease int) {
	s.stepSize += amountIncrease
}

func (s *sharedState) incrementCandyAvailable() {
	if s.candyAvailable < s.candyLimit {
		s.candyAvailable++
	}
}

func (s *sharedState) dropCandy() {
	if s.candyAvailable == 0 {
		return
	}
	s.candyAvailable--
	s.pubSub.Publish(pubsub.OnDropCandy, pubsub.OnDropCandyPayload{
		X:          s.x,
		Y:          s.y,
		Width:      s.playerWidth,
		Height:     s.playerHeight,
		PowerLevel: s.powerLevel,
	})
}
