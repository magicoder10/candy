package graphics

import (
	"image"
	"sort"

	"github.com/hajimehoshi/ebiten/v2"
)

var _ Batch = (*ebitenBatch)(nil)

type ebitenBatch struct {
	ebiten       *Ebiten
	spriteSheet  *ebiten.Image
	spritesDrawn []*spriteDrawn
}

func (e *ebitenBatch) renderBatch(target *ebiten.Image) {
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
		adjustedY := adjustY(e.ebiten.reverseY, target, spriteDrawn.y, scaledHeight)
		op.GeoM.Translate(
			float64(spriteDrawn.x),
			float64(adjustedY),
		)

		target.DrawImage(e.spriteSheet.SubImage(bound).(*ebiten.Image), op)
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
