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

func (Painter) drawString(
	dest draw.Image, destPoint image.Point,
	ft *truetype.Font, text string, fontSize int, color Color,
) {
	// TODO: support anti-aliasing
	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(ft)
	c.SetFontSize(float64(fontSize))
	c.SetClip(dest.Bounds())
	c.SetDst(dest)
	c.SetSrc(color.toUniform())
	c.SetHinting(font.HintingFull)
	pt := freetype.Pt(destPoint.X, destPoint.Y+c.PointToFixed(float64(fontSize)).Round())

	c.DrawString(text, pt)
}
