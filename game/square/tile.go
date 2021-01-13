package square

import (
	"candy/game/gameitem"
	"candy/graphics"
)

const revealItemXOffset = 16
const revealItemYOffset = 36
const revealItemZOffset = -1

var _ Square = (*Tile)(nil)

type Tile struct {
	imageXOffset int
	imageYOffset int
	xOffset      int
	yOffset      int
	canEnter     bool
	showItem     bool
	gameItem     gameitem.GameItem
}

func (t Tile) IsBreakable() bool {
	return false
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
		batch.DrawSprite(x+t.xOffset+revealItemXOffset, y+t.yOffset+revealItemYOffset, y+revealItemZOffset, t.gameItem.GetBound(), 0.6)
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
	gameItem := t.gameItem
	t.gameItem = gameitem.None
	return gameItem
}

func newYellow(gameItem gameitem.GameItem) Tile {
	return Tile{
		imageXOffset: 576,
		imageYOffset: 304,
		xOffset:      -4,
		yOffset:      -2,
		canEnter:     false,
		gameItem:     gameItem,
	}
}

func newGreen(gameItem gameitem.GameItem) Tile {
	return Tile{
		imageXOffset: 576,
		imageYOffset: 224,
		xOffset:      -4,
		yOffset:      -2,
		canEnter:     false,
		gameItem:     gameItem,
	}
}

func NewTile(tileType rune, gameItem gameitem.GameItem) Tile {
	switch tileType {
	case 'Y':
		tile := newYellow(gameItem)
		return tile
	case 'G':
		tile := newGreen(gameItem)
		return tile
	default:
		tile := newGreen(gameItem)
		return tile
	}
}
