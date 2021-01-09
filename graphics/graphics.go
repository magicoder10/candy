package graphics

import (
	"image"
)

type Bound struct {
	X      int
	Y      int
	Width  int
	Height int
}

type Batch interface {
	DrawSprite(x int, y int, z int, imageBound Bound, scale float64)
	RenderBatch()
}

type Graphics interface {
	Clear()
	StartNewBatch(spriteSheet image.Image) Batch
}

type spriteDrawn struct {
	x          int
	y          int
	z          int
	imageBound Bound
	scale      float64
}
