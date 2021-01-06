package tile

import (
	"candy/graphics"
)

const Width = 60
const Height = 60

type Tile struct {
	imageXOffset int
	imageYOffset int
	xOffset      int
	yOffset      int
	canEnter     bool
}

func (t Tile) Draw(batch graphics.Batch, x int, y int) {
	bound := graphics.Bound{
		X:      t.imageXOffset,
		Y:      t.imageYOffset,
		Width:  64,
		Height: 80,
	}
	batch.DrawSprite(x+t.xOffset, y+t.yOffset, y, bound, 1)
}

func (t Tile) CanEnter() bool {
	return t.canEnter
}

func newYellow() Tile {
	return Tile{
		imageXOffset: 576,
		imageYOffset: 304,
		xOffset:      -4,
		yOffset:      -2,
		canEnter:     false,
	}
}

func newGreen() Tile {
	return Tile{
		imageXOffset: 576,
		imageYOffset: 224,
		xOffset:      -4,
		yOffset:      -2,
		canEnter:     false,
	}
}

func NewTile(tileType rune) *Tile {
	switch tileType {
	case 'Y':
		tile := newYellow()
		return &tile
	case 'G':
		tile := newGreen()
		return &tile
	case ' ':
		return nil
	default:
		tile := newGreen()
		return &tile
	}
}
