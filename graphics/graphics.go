package graphics

import (
	"image"
	"io"

	"golang.org/x/image/font"
)

type alignment int

const (
	AlignCenter alignment = iota
)

type Bound struct {
	X      int
	Y      int
	Width  int
	Height int
}

type Text interface {
	io.Writer
	Draw()
}

type Batch interface {
	DrawSprite(x int, y int, z int, imageBound Bound, scale float64)
	RenderBatch()
}

type Graphics interface {
	Clear()
	NewText(face font.Face, x int, y int, width int, height int, scale float64, alignment alignment) Text
	StartNewBatch(spriteSheet image.Image) Batch
	SetCursorVisible(isVisible bool)
}

type spriteDrawn struct {
	x          int
	y          int
	z          int
	imageBound Bound
	scale      float64
}
