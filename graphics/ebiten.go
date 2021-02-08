package graphics

import (
	"bufio"
	"bytes"
	"image"
	"image/color"
	"io/ioutil"
	"sort"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

var _ Graphics = (*Ebiten)(nil)

type Ebiten struct {
	reverseY bool
	texts    []*ebitenText
	batches  []*ebitenBatch
}

func (e Ebiten) Clear() {
	e.texts = make([]*ebitenText, 0)
	e.batches = make([]*ebitenBatch, 0)
}

func (e *Ebiten) NewText(face font.Face, x int, y int, width int, height int, scale float64, alignment alignment) Text {
	buf := bytes.Buffer{}
	return &ebitenText{
		x:         x,
		y:         y,
		width:     width,
		height:    height,
		fontFace:  face,
		buf:       bufio.NewReadWriter(bufio.NewReader(&buf), bufio.NewWriter(&buf)),
		graphics:  e,
		alignment: alignment,
	}
}

func (e *Ebiten) RenderTexts(screen *ebiten.Image) {
	for _, t := range e.texts {
		bound := text.BoundString(t.fontFace, t.textContent)
		width := float64(bound.Max.X - bound.Min.X)
		height := float64(bound.Max.Y - bound.Min.Y)

		x := float64(t.x) + float64(t.width)/2 - width/2
		y := float64(t.y) + float64(t.height)/2 - height/2

		adjustedY := adjustY(e.reverseY, screen, int(y), 0)

		switch t.alignment {
		case AlignCenter:
			text.Draw(screen, t.textContent, t.fontFace, int(x), adjustedY, color.White)
		}
	}
}

func (e *Ebiten) StartNewBatch(spriteSheet image.Image) Batch {
	return newEbitenBatch(ebiten.NewImageFromImage(spriteSheet), e)
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
		batch.renderBatch(screen)
	}
	e.RenderTexts(screen)
}

func NewEbiten(reverseY bool) Ebiten {
	return Ebiten{
		texts:    make([]*ebitenText, 0),
		batches:  make([]*ebitenBatch, 0),
		reverseY: reverseY,
	}
}

type ebitenText struct {
	buf         *bufio.ReadWriter
	textContent string
	fontFace    font.Face
	graphics    *Ebiten
	x           int
	y           int
	width       int
	height      int
	alignment   alignment
}

func (t *ebitenText) Write(p []byte) (int, error) {
	return t.buf.Write(p)
}

func (t *ebitenText) Draw() {
	t.buf.Flush()
	buf, _ := ioutil.ReadAll(t.buf)
	t.textContent = string(buf)

	t.graphics.texts = append(t.graphics.texts, t)
}

var _ Batch = (*ebitenBatch)(nil)

type ebitenBatch struct {
	ebiten       *Ebiten
	spriteSheet  *ebiten.Image
	spritesDrawn []*spriteDrawn
}

func (e *ebitenBatch) RenderBatch() {
	e.ebiten.batches = append(e.ebiten.batches, e)
}

func (e *ebitenBatch) renderBatch(screen *ebiten.Image) {
	sort.SliceStable(e.spritesDrawn, func(i, j int) bool {
		if e.ebiten.reverseY {
			return e.spritesDrawn[i].z > e.spritesDrawn[j].z
		} else {
			return e.spritesDrawn[i].z < e.spritesDrawn[j].z
		}
	})

	spSheetHeight := e.spriteSheet.Bounds().Max.Y - e.spriteSheet.Bounds().Min.Y

	for _, spriteDrawn := range e.spritesDrawn {

		maxX := spriteDrawn.imageBound.X + spriteDrawn.imageBound.Width
		maxY := spriteDrawn.imageBound.Y + spriteDrawn.imageBound.Height
		bound := image.Rectangle{
			Min: image.Point{
				X: spriteDrawn.imageBound.X,
				Y: spSheetHeight - maxY,
			},
			Max: image.Point{
				X: maxX,
				Y: spSheetHeight - spriteDrawn.imageBound.Y,
			},
		}

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(spriteDrawn.scale, spriteDrawn.scale)

		scaledHeight := float64(spriteDrawn.imageBound.Height) * spriteDrawn.scale
		adjustedY := adjustY(e.ebiten.reverseY, screen, spriteDrawn.y, scaledHeight)
		op.GeoM.Translate(
			float64(spriteDrawn.x),
			float64(adjustedY),
		)

		screen.DrawImage(e.spriteSheet.SubImage(bound).(*ebiten.Image), op)
	}
	e.spritesDrawn = make([]*spriteDrawn, 0)
}

func (e *ebitenBatch) DrawSprite(x int, y int, z int, imageBound Bound, scale float64) {
	e.spritesDrawn = append(e.spritesDrawn, &spriteDrawn{
		x:          x,
		y:          y,
		z:          z,
		imageBound: imageBound,
		scale:      scale,
	})
}

func newEbitenBatch(spriteSheet *ebiten.Image, ebiten *Ebiten) *ebitenBatch {
	return &ebitenBatch{
		spriteSheet:  spriteSheet,
		spritesDrawn: make([]*spriteDrawn, 0),
		ebiten:       ebiten,
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
