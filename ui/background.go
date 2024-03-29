package ui

import (
	"image"
	"image/draw"

	"candy/assets"
)

type Background struct {
	Color     *Color
	ImagePath *string

	prevColor     Color
	prevImagePath string

	image      image.Image
	hasChanged bool
}

func (b Background) Paint(painter *Painter, size Size, destLayer draw.Image) {
	if b.image != nil {
		rect := image.Rectangle{
			Max: image.Point{
				X: size.width,
				Y: size.height,
			},
		}
		painter.drawImage(b.image, rect, destLayer, image.Point{})
		return
	}
	if b.Color != nil {
		painter.fillColor(destLayer, *b.Color)
	}
}

func (b *Background) Update(assets *assets.Assets) {
	if b.ImagePath != nil {
		imagePath := *b.ImagePath
		if imagePath != b.prevImagePath {
			b.image = assets.GetImage(imagePath)
			b.hasChanged = true
			b.prevImagePath = *b.ImagePath
		}
	}
	if b.Color != nil {
		if *b.Color == b.prevColor {
			b.hasChanged = true
			b.prevColor = *b.Color
		}
	}
}

func (b *Background) ResetChangeDetection() {
	b.hasChanged = false
}
