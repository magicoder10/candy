package graphics

import (
	"image"
	"io"

	"golang.org/x/image/font"
)

type alignment int

const (
	AlignLeft alignment = iota
	AlignRight
	AlignTop
	AlignDown
	AlignCenter
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
	RenderTexts()
	StartNewBatch(spriteSheet image.Image) Batch
}

type spriteDrawn struct {
	x          int
	y          int
	z          int
	imageBound Bound
	scale      float64
}
