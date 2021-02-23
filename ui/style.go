package ui

import (
	"image"
	"image/color"
)

type Style struct {
	Width      *int
	Height     *int
	LayoutType LayoutType
	FontStyle  FontStyle
}

type FontStyle struct {
	Family     string
	Weight     string
	Italic     bool
	LineHeight int
	Color      Color
	Size       int
}

var _ color.Color = (*Color)(nil)

type Color struct {
	Red   uint8
	Green uint8
	Blue  uint8
	Alpha uint8
}

func (c Color) RGBA() (r, g, b, a uint32) {
	return highBits(c.Red), highBits(c.Green), highBits(c.Blue), highBits(c.Alpha)
}

func highBits(num uint8) uint32 {
	red := uint32(num)
	red |= red << 8
	return red
}

func (c Color) toUniform() *image.Uniform {
	return image.NewUniform(c)
}
