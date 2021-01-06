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
	DrawImage(x int, y int, image image.Image, imageBound Bound, scale float64)
	StartNewBatch(spriteSheet image.Image) Batch
}

type imageDrawn struct {
	x          int
	y          int
	z          int
	image      image.Image
	imageBound Bound
	scale      float64
}

type spriteDrawn struct {
	x          int
	y          int
	z          int
	imageBound Bound
	scale      float64
}
