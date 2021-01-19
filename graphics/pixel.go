package graphics

import (
	"bufio"
	"bytes"
	"candy/input"
	"image"
	"io/ioutil"
	"sort"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font"
)

var _ Graphics = (*Pixel)(nil)
var _ Window = (*Pixel)(nil)

type Pixel struct {
	window *pixelgl.Window
	texts  *[]*pixelText
}

func (p *Pixel) SetCursorVisible(isVisible bool) {
	p.window.SetCursorVisible(isVisible)
}

// Graphics
func (p *Pixel) StartNewBatch(spriteSheet image.Image) Batch {
	pixelImg := pixel.PictureDataFromImage(spriteSheet)
	return newPixelBatch(p.window, pixelImg)
}

func (p *Pixel) NewText(face font.Face, x int, y int, width int, height int, scale float64, alignment alignment) Text {
	atlas := text.NewAtlas(face, text.ASCII)
	buf := bytes.Buffer{}
	return &pixelText{
		buf:       bufio.NewReadWriter(bufio.NewReader(&buf), bufio.NewWriter(&buf)),
		graphics:  p,
		text:      text.New(pixel.V(float64(x), float64(y)), atlas),
		width:     width,
		height:    height,
		scale:     scale,
		alignment: alignment,
	}
}

func (p *Pixel) RenderTexts() {
	for _, t := range *p.texts {
		t.buf.Flush()
		buf, _ := ioutil.ReadAll(t.buf)
		bound := t.text.BoundsOf(string(buf))
		halfWidth := bound.W() * t.scale / 2.0
		halfHeight := bound.H() * t.scale / 2.0

		switch t.alignment {
		case AlignCenter:
			t.text.Dot.X += float64(t.width)/2 - halfWidth
			t.text.Dot.Y += float64(t.height)/2 - halfHeight
		}
		t.text.Write(buf)
		t.text.Draw(p.window, pixel.IM.Scaled(t.text.Orig, t.scale))
		t.text.Clear()
	}
	texts := make([]*pixelText, 0)
	p.texts = &texts
}

// Window
func (p Pixel) Clear() {
	p.window.Clear(colornames.Black)
}

func (p Pixel) IsClosed() bool {
	return p.window.Closed()
}

func (p Pixel) Redraw() {
	p.window.Update()
}

func (p Pixel) PollEvents() []input.Input {
	inputs := make([]input.Input, 0)
	if p.window.Pressed(pixelgl.KeyLeft) {
		inputs = append(inputs, input.Input{
			Action: input.Press,
			Device: input.LeftArrowKey,
		})
	}
	if p.window.Pressed(pixelgl.KeyRight) {
		inputs = append(inputs, input.Input{
			Action: input.Press,
			Device: input.RightArrowKey,
		})
	}
	if p.window.Pressed(pixelgl.KeyUp) {
		inputs = append(inputs, input.Input{
			Action: input.Press,
			Device: input.UpArrowKey,
		})
	}
	if p.window.Pressed(pixelgl.KeyDown) {
		inputs = append(inputs, input.Input{
			Action: input.Press,
			Device: input.DownArrowKey,
		})
	}
	if p.window.Pressed(pixelgl.KeyR) {
		inputs = append(inputs, input.Input{
			Action: input.Press,
			Device: input.RKey,
		})
	}
	if p.window.JustReleased(pixelgl.KeyLeft) {
		inputs = append(inputs, input.Input{
			Action: input.Release,
			Device: input.LeftArrowKey,
		})
	}
	if p.window.JustReleased(pixelgl.KeyRight) {
		inputs = append(inputs, input.Input{
			Action: input.Release,
			Device: input.RightArrowKey,
		})
	}
	if p.window.JustReleased(pixelgl.KeyUp) {
		inputs = append(inputs, input.Input{
			Action: input.Release,
			Device: input.UpArrowKey,
		})
	}
	if p.window.JustReleased(pixelgl.KeyDown) {
		inputs = append(inputs, input.Input{
			Action: input.Release,
			Device: input.DownArrowKey,
		})
	}
	if p.window.JustReleased(pixelgl.KeyR) {
		inputs = append(inputs, input.Input{
			Action: input.Release,
			Device: input.RKey,
		})
	}
	if p.window.JustReleased(pixelgl.KeySpace) {
		inputs = append(inputs, input.Input{
			Action: input.Release,
			Device: input.SpaceKey,
		})
	}
	if p.window.JustPressed(pixelgl.MouseButtonLeft) {
		inputs = append(inputs, input.Input{
			Action: input.SinglePress,
			Device: input.MouseLeftButton,
		})
	}
	return inputs
}

func NewPixel(config pixelgl.WindowConfig) (Pixel, error) {
	win, err := pixelgl.NewWindow(config)
	if err != nil {
		return Pixel{}, err
	}
	texts := make([]*pixelText, 0)
	return Pixel{window: win, texts: &texts}, nil
}

var _ Batch = (*pixelBatch)(nil)

type pixelBatch struct {
	atlas        text.Atlas
	window       *pixelgl.Window
	spriteSheet  *pixel.PictureData
	batch        *pixel.Batch
	spritesDrawn []spriteDrawn
}

func (p *pixelBatch) DrawSprite(x int, y int, z int, imageBound Bound, scale float64) {
	p.spritesDrawn = append(p.spritesDrawn, spriteDrawn{
		x:          x,
		y:          y,
		z:          z,
		imageBound: imageBound,
		scale:      scale,
	})
}

func (p *pixelBatch) RenderBatch() {
	p.batch.Clear()

	sort.SliceStable(p.spritesDrawn, func(i, j int) bool {
		return p.spritesDrawn[i].z > p.spritesDrawn[j].z
	})

	for _, spriteDrawn := range p.spritesDrawn {
		sprite, matrix := prepareDrawing(spriteDrawn.x, spriteDrawn.y, p.spriteSheet, spriteDrawn.imageBound, spriteDrawn.scale)
		sprite.Draw(p.batch, matrix)
	}

	p.batch.Draw(p.window)
	p.spritesDrawn = make([]spriteDrawn, 0)
}

func newPixelBatch(windows *pixelgl.Window, spriteSheet *pixel.PictureData) *pixelBatch {
	return &pixelBatch{
		window:       windows,
		spriteSheet:  spriteSheet,
		batch:        pixel.NewBatch(&pixel.TrianglesData{}, spriteSheet),
		spritesDrawn: make([]spriteDrawn, 0),
	}
}

func prepareDrawing(x int, y int, image *pixel.PictureData, imageBound Bound, scale float64) (*pixel.Sprite, pixel.Matrix) {
	maxX := imageBound.X + imageBound.Width
	maxY := imageBound.Y + imageBound.Height
	bound := pixel.R(
		float64(imageBound.X),
		float64(imageBound.Y),
		float64(maxX),
		float64(maxY),
	)
	sprite := pixel.NewSprite(image, bound)

	scaledImgWidth := float64(imageBound.Width) * scale
	scaledImgHeight := float64(imageBound.Height) * scale
	pos := pixel.V(float64(x)+scaledImgWidth/2, float64(y)+scaledImgHeight/2)
	return sprite, pixel.IM.Moved(pos).Scaled(pos, scale)
}

var _ Text = (*pixelText)(nil)

type pixelText struct {
	buf       *bufio.ReadWriter
	graphics  *Pixel
	text      *text.Text
	width     int
	height    int
	scale     float64
	alignment alignment
}

func (t pixelText) Write(p []byte) (int, error) {
	return t.buf.Write(p)
}

func (t *pixelText) Draw() {
	*t.graphics.texts = append(*t.graphics.texts, t)
}
