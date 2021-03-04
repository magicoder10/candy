package graphics

import (
	"bufio"
	"bytes"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

const rgbaPixelByteCount = 4

var _ Graphics = (*Ebiten)(nil)

type Ebiten struct {
	reverseY bool
	texts    []*ebitenText
	batches  []*ebitenBatch
	buffer   *ebiten.Image
	imageBytes []byte
}

func (e *Ebiten) NewCanvas() *EbitenCanvas {
	return &EbitenCanvas {}
}

func (e *Ebiten) initBuffer(width int, height int) {
	e.buffer = ebiten.NewImage(width, height)
}

func (e *Ebiten) NewText(face font.Face, x int, y int, width int, height int, _ float64, alignment alignment) Text {
	buf := bytes.Buffer{}
	txt := &ebitenText{
		x:         x,
		y:         y,
		width:     width,
		height:    height,
		fontFace:  face,
		buf:       bufio.NewReadWriter(bufio.NewReader(&buf), bufio.NewWriter(&buf)),
		graphics:  e,
		alignment: alignment,
	}
	e.texts = append(e.texts, txt)
	return txt
}

func (e *Ebiten) RenderTexts(target *ebiten.Image) {
	for _, t := range e.texts {
		bound := text.BoundString(t.fontFace, t.textContent)
		width := float64(bound.Max.X - bound.Min.X)
		height := float64(bound.Max.Y - bound.Min.Y)

		x := float64(t.x) + float64(t.width)/2 - width/2
		y := float64(t.y) + float64(t.height)/2 - height/2

		adjustedY := adjustY(e.reverseY, target, int(y), 0)

		switch t.alignment {
		case AlignCenter:
			text.Draw(target, t.textContent, t.fontFace, int(x), adjustedY, color.White)
		}
	}
}

func (e *Ebiten) StartNewBatch(spriteSheet image.Image) Batch {
	batch := newEbitenBatch(ebiten.NewImageFromImage(spriteSheet), e)
	e.batches = append(e.batches, batch)
	return batch
}

func (e Ebiten) SetCursorVisible(isVisible bool) {
	if isVisible {
		ebiten.SetCursorMode(ebiten.CursorModeVisible)
	} else {
		ebiten.SetCursorMode(ebiten.CursorModeHidden)
	}
}

func (e *Ebiten) Render(screen *ebiten.Image) {
	for _, batch := range e.batches {
		batch.renderBatch(e.buffer)
	}
	e.RenderTexts(e.buffer)
	screen.DrawImage(e.buffer, nil)
}

func (e *Ebiten) GetCursorPosition() image.Point {
	x, y := ebiten.CursorPosition()
	return image.Point{X: x, Y: y}
}

func NewEbiten(autoClearScreen bool, reverseY bool) Ebiten {
	if autoClearScreen {
		ebiten.SetScreenClearedEveryFrame(autoClearScreen)
	}
	return Ebiten{
		texts:    make([]*ebitenText, 0),
		batches:  make([]*ebitenBatch, 0),
		reverseY: reverseY,
	}
}

func adjustY(shouldAdjust bool, screen *ebiten.Image, originalY int, scaledHeight float64) int {
	if !shouldAdjust {
		return originalY
	} else {
		screenHeight := screen.Bounds().Max.Y - screen.Bounds().Min.Y
		return screenHeight - originalY - int(scaledHeight)
	}
}

var _ Canvas = (*EbitenCanvas)(nil)

type EbitenCanvas struct {
	buf []byte
}

func (e *EbitenCanvas) OverrideContent(img image.Image) {
	e.buf = toBytes(img)
}

func (e *EbitenCanvas) Render(img *ebiten.Image) {
	if len(e.buf) == 0 {
		return
	}
	img.ReplacePixels(e.buf)
}

func toBytes(img image.Image) []byte {
	bound := img.Bounds()
	width := bound.Max.X - bound.Min.X
	height := bound.Max.Y - bound.Min.Y

	buf := make([]byte, rgbaPixelByteCount*width*height)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			index := (x*width + y) * rgbaPixelByteCount
			setPixel(buf, index, img.At(x, y))
		}
	}
	return buf
}

func setPixel(buf []byte, index int, c color.Color) {
	r, g, b, a := c.RGBA()
	buf[index], buf[index+1], buf[index+2], buf[index+2] = byte(r), byte(g), byte(b), byte(a)
}
