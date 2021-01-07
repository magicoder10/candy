package tile

import (
	"candy/game/gameitem"
	"candy/graphics"
)

const Width = 60
const Height = 60
const GameItemXOffset = 16
const GameItemYOffset = 36
const GameItemHeightOffset = -1

type Tile struct {
	imageXOffset int
	imageYOffset int
	xOffset      int
	yOffset      int
	canEnter     bool
	showItem     bool
	gameItem     gameitem.GameItem
}

func (t Tile) Draw(batch graphics.Batch, x int, y int) {
	bound := graphics.Bound{
		X:      t.imageXOffset,
		Y:      t.imageYOffset,
		Width:  64,
		Height: 80,
	}
	batch.DrawSprite(x+t.xOffset, y+t.yOffset, y, bound, 1)

	if t.gameItem != gameitem.None && t.showItem {
		batch.DrawSprite(x+t.xOffset + GameItemXOffset, y+t.yOffset + GameItemYOffset, y + GameItemHeightOffset, t.gameItem.GetBound(), 0.6)
	}
}

func (t Tile) CanEnter() bool {
	return t.canEnter
}

func (t *Tile) RevealItem() gameitem.GameItem {
	t.showItem = true
	return t.gameItem
}

func (t *Tile) HideItem() gameitem.GameItem {
	t.showItem = false
	return t.gameItem
}

func (t *Tile) RemoveItem() gameitem.GameItem {
	item := t
	t.gameItem = gameitem.None
	return item.gameItem
}

func newYellow(gameItem gameitem.GameItem) Tile {
	return Tile{
		imageXOffset: 576,
		imageYOffset: 304,
		xOffset:      -4,
		yOffset:      -2,
		canEnter:     false,
		gameItem: gameItem,
	}
}

func newGreen(gameItem gameitem.GameItem) Tile {
	return Tile{
		imageXOffset: 576,
		imageYOffset: 224,
		xOffset:      -4,
		yOffset:      -2,
		canEnter:     false,
		gameItem: gameItem,
	}
}

func NewTile(tileType rune, gameItem gameitem.GameItem) *Tile {
	switch tileType {
	case 'Y':
		tile := newYellow(gameItem)
		return &tile
	case 'G':
		tile := newGreen(gameItem)
		return &tile
	case ' ':
		return nil
	default:
		tile := newGreen(gameItem)
		return &tile
	}
}
