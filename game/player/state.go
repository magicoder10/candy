package player

import (
	"time"

	"candy/game/direction"
	"candy/game/gamemap"
	"candy/game/square"
	"candy/graphics"
	"candy/input"
)

type state interface {
	handleInput(in input.Input) state
	update(timeElapsed time.Duration)
	draw(batch graphics.Batch,
		regionXOffset int,
		regionYOffset int,
		walkCycleXOffset int,
		walkCycleYOffset int,
	)
	getX() int
	getY() int
	getWidth() int
	getHeight() int
}

type sharedState struct {
	gameMap   *gamemap.Map
	currStep  int
	direction direction.Direction
	width     int
	height    int
	x         int
	y         int
}

func (s sharedState) update(timeElapsed time.Duration) {
	return
}

func (s sharedState) draw(
	batch graphics.Batch,
	regionXOffset int, regionYOffset int,
	walkCycleXOffset int,
	walkCycleYOffset int,
) {
	bound := graphics.Bound{
		X:      regionXOffset + walkCycleXOffset + s.currStep*spriteWidth,
		Y:      regionYOffset + walkCycleYOffset + int(s.direction)*spriteHeight,
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
	return s.width
}

func (s sharedState) getHeight() int {
	return s.height
}
