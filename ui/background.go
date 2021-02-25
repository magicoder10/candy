package ui

import (
	"image/draw"
)

type Background struct {
	Color *Color
}

func (b Background) Paint(painter *Painter, destLayer draw.Image) {
	if b.Color != nil {
		painter.fillColor(destLayer, *b.Color)
	}
}
