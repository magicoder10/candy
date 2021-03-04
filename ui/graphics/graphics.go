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

type Canvas interface {
	OverrideContent(img image.Image)
}

type Graphics interface {
	SetCursorVisible(isVisible bool)
	GetCursorPosition() image.Point
}
