package graphics

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const rgbaPixelByteCount = 4

var _ Graphics = (*Ebiten)(nil)

type Ebiten struct {
}

func (e *Ebiten) NewCanvas(width int, height int) *EbitenCanvas {
	return &EbitenCanvas {
		buf: make([]byte, rgbaPixelByteCount*width*height),
	}
}

func (e Ebiten) SetCursorVisible(isVisible bool) {
	if isVisible {
		ebiten.SetCursorMode(ebiten.CursorModeVisible)
	} else {
		ebiten.SetCursorMode(ebiten.CursorModeHidden)
	}
}

func (e *Ebiten) GetCursorPosition() image.Point {
	x, y := ebiten.CursorPosition()
	return image.Point{X: x, Y: y}
}

func NewEbiten(autoClearScreen bool) Ebiten {
	ebiten.SetScreenClearedEveryFrame(autoClearScreen)
	return Ebiten{}
}

var _ Canvas = (*EbitenCanvas)(nil)

type EbitenCanvas struct {
	buf []byte
}

func (e *EbitenCanvas) OverrideContent(img image.Image) {
	bound := img.Bounds()
	width := bound.Max.X - bound.Min.X
	height := bound.Max.Y - bound.Min.Y

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			index := (y*width + x) * rgbaPixelByteCount
			setPixel(e.buf, index, img.At(x, y))
		}
	}
}

func (e *EbitenCanvas) DrawOn(img *ebiten.Image) {
	if e.buf == nil {
		return
	}
	img.ReplacePixels(e.buf)
}

func setPixel(buf []byte, index int, c color.Color) {
	r, g, b, a := c.RGBA()
	buf[index], buf[index+1], buf[index+2], buf[index+3] = byte(r), byte(g), byte(b), byte(a)
}
