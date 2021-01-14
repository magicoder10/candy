package player

import (
	"time"

	"candy/game/direction"
	"candy/game/square"
	"candy/graphics"
	"candy/input"
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
}

type sharedState struct {
	currStep         int
	direction        direction.Direction
	playerWidth      int
	playerHeight     int
	x                int
	y                int
	moveChecker      MoveChecker
	regionXOffset    int
	regionYOffset    int
	walkCycleXOffset int
	walkCycleYOffset int

}

func (s sharedState) update(timeElapsed time.Duration) {
	return
}

func (s sharedState) isNormal() bool {
	return true
}

func (s sharedState) trapped() state {
	return trappedState{sharedState: s}
}

func (s sharedState) draw(batch graphics.Batch) {
	bound := graphics.Bound{
		X:      s.regionXOffset + s.walkCycleXOffset + s.currStep*spriteWidth,
		Y:      s.regionYOffset + s.walkCycleYOffset + int(s.direction)*spriteHeight,
		Width:  spriteWidth,
		Height: spriteHeight,
	}
	batch.DrawSprite(s.x-square.Width/6, s.y, s.y, bound, float64(square.Width)/spriteWidth)
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
