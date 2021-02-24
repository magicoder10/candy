package player

import (
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
	getCurCandyAmount() int
	getMaxCandyAmount() int
	increasePowerLevel(amountIncrease int)
	increaseStepSize(amountIncrease int)
	incrementCurCandyAmount() // after a candy explodes, curCandyAmount should increase by 1, means player can drop one more candy
	decrementCurCandyAmount() // after a candy is exploded, curCandyAmount should decrease by 1
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
	curCandyAmount int
	maxCandyAmount int
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
	return s.curCandyAmount
}

func (s sharedState) getMaxCandyAmount() int {
	return s.maxCandyAmount
}

func (s *sharedState) increasePowerLevel(amountIncrease int) {
	s.powerLevel += amountIncrease
}

func (s *sharedState) increaseStepSize(amountIncrease int) {
	s.stepSize += amountIncrease
}

func (s *sharedState) incrementCurCandyAmount() {
	s.curCandyAmount++
}

func (s *sharedState) decrementCurCandyAmount() {
	s.curCandyAmount--
}

func (s sharedState) dropCandy() {
	s.pubSub.Publish(pubsub.OnDropCandy, pubsub.OnDropCandyPayload{
		X:              s.x,
		Y:              s.y,
		Width:          s.playerWidth,
		Height:         s.playerHeight,
		PowerLevel:     s.powerLevel,
		MaxCandyAmount: s.maxCandyAmount,
		//		CandyCurAmount: s.candyCurAmount,
	})
}
