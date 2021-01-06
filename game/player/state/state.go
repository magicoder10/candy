package state

import (
	"time"

	"candy/game/direction"
	"candy/game/gamemap"
	"candy/input"
)

type State interface {
	HandleInput(in input.Input) State
	Update(timeElapsed time.Duration)
	GetCurrentStep() int
	GetDirection() direction.Direction
	GetX() int
	GetY() int
}

type sharedState struct {
	gameMap   gamemap.Map
	currStep  int
	direction direction.Direction
	width     int
	height    int
	x         int
	y         int
}

func (s sharedState) GetCurrentStep() int {
	return s.currStep
}

func (s sharedState) GetDirection() direction.Direction {
	return s.direction
}

func (s sharedState) GetX() int {
	return s.x
}

func (s sharedState) GetY() int {
	return s.y
}
