package ui

import (
	"image"
	"image/draw"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

type Painter struct {
}

func (Painter) drawImage(src image.Image, srcRect image.Rectangle, dest draw.Image, destPoint image.Point) {
	width := srcRect.Max.X - srcRect.Min.X
	height := srcRect.Max.Y - srcRect.Min.Y

	destRect := image.Rectangle{
		Min: destPoint,
		Max: image.Point{
			X: destPoint.X + width,
			Y: destPoint.Y + height,
		},
	}
	draw.Draw(dest, destRect, src, srcRect.Min, draw.Over)
}

func (Painter) fillColor(destImg draw.Image, color Color) {
	draw.Draw(destImg, destImg.Bounds(), color.toUniform(), image.Point{}, draw.Src)
}

func (Painter) drawString(
	dest draw.Image, destPoint image.Point,
	ft *truetype.Font, text string, fontSize int, color Color,
) {
	// TODO: support anti-aliasing
	options := truetype.Options{Size: float64(fontSize), DPI: float64(72)}
	face := truetype.NewFace(ft, &options)

	drawer := font.Drawer{
		Dst:  dest,
		Src:  color.toUniform(),
		Face: face,
		Dot:  freetype.Pt(destPoint.X, destPoint.Y+fontSize),
	}
	drawer.DrawString(text)
}
