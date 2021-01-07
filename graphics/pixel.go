package graphics

import (
	"image"
	"sort"

	"candy/input"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var _ Graphics = (*Pixel)(nil)
var _ Window = (*Pixel)(nil)

type Pixel struct {
	window *pixelgl.Window
}

// Graphics
func (p *Pixel) StartNewBatch(spriteSheet image.Image) Batch {
	pixelImg := pixel.PictureDataFromImage(spriteSheet)
	return newPixelBatch(p.window, pixelImg)
}

func (p *Pixel) DrawImage(x int, y int, image image.Image, imageBound Bound, scale float64) {
	pixelImg := pixel.PictureDataFromImage(image)
	sprite, matrix := prepareDrawing(x, y, pixelImg, imageBound, scale)
	sprite.Draw(p.window, matrix)
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

func (p Pixel) PollEvents() *input.Input {
	if p.window.Pressed(pixelgl.KeyLeft) {
		return &input.Input{
			Action: input.Press,
			Device: input.LeftArrowKey,
		}
	} else if p.window.Pressed(pixelgl.KeyRight) {
		return &input.Input{
			Action: input.Press,
			Device: input.RightArrowKey,
		}
	} else if p.window.Pressed(pixelgl.KeyUp) {
		return &input.Input{
			Action: input.Press,
			Device: input.UpArrowKey,
		}
	} else if p.window.Pressed(pixelgl.KeyDown) {
		return &input.Input{
			Action: input.Press,
			Device: input.DownArrowKey,
		}
	} else if p.window.Pressed(pixelgl.KeyR) {
		return &input.Input{
			Action: input.Press,
			Device: input.RKey,
		}
	} else if p.window.JustReleased(pixelgl.KeyLeft) {
		return &input.Input{
			Action: input.Release,
			Device: input.LeftArrowKey,
		}
	} else if p.window.JustReleased(pixelgl.KeyRight) {
		return &input.Input{
			Action: input.Release,
			Device: input.RightArrowKey,
		}
	} else if p.window.JustReleased(pixelgl.KeyUp) {
		return &input.Input{
			Action: input.Release,
			Device: input.UpArrowKey,
		}
	} else if p.window.JustReleased(pixelgl.KeyDown) {
		return &input.Input{
			Action: input.Release,
			Device: input.DownArrowKey,
		}
	} else if p.window.JustReleased(pixelgl.KeyR) {
		return &input.Input{
			Action: input.Release,
			Device: input.RKey,
		}
	}
	return nil
}

func NewPixel(config pixelgl.WindowConfig) (Pixel, error) {
	win, err := pixelgl.NewWindow(config)
	if err != nil {
		return Pixel{}, err
	}
	return Pixel{window: win}, nil
}

var _ Batch = (*PixelBatch)(nil)

type PixelBatch struct {
	window       *pixelgl.Window
	spriteSheet  *pixel.PictureData
	batch        *pixel.Batch
	spritesDrawn []spriteDrawn
}

func (p *PixelBatch) DrawSprite(x int, y int, z int, imageBound Bound, scale float64) {
	p.spritesDrawn = append(p.spritesDrawn, spriteDrawn{
		x:          x,
		y:          y,
		z:          z,
		imageBound: imageBound,
		scale:      scale,
	})
}

func (p *PixelBatch) RenderBatch() {
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

func newPixelBatch(windows *pixelgl.Window, spriteSheet *pixel.PictureData) *PixelBatch {
	return &PixelBatch{
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
