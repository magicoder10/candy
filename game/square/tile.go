package square

import (
	"candy/game/gameitem"
	"candy/graphics"
)

const revealItemXOffset = 16
const revealItemYOffset = 36
const revealItemZOffset = -1
const brokenTileXOffset = 4
const brokenTileYOffset = 2

var _ Square = (*Tile)(nil)

type Tile struct {
	imageXOffset int
	imageYOffset int
	xOffset      int
	yOffset      int
	showItem     bool
	isBroken     bool
	canEnter     bool
	gameItem     gameitem.GameItem
}

func (t *Tile) ShouldRemove() bool {
	return t.canEnter && t.gameItem == gameitem.None
}

func (t *Tile) UnblockFire() {
	t.canEnter = true
}

func (t *Tile) IsBlocking() bool {
	return true
}

func (t *Tile) Break() {
	t.isBroken = true
}

func (t Tile) IsBreakable() bool {
	return true
}

func (t Tile) Draw(batch graphics.Batch, x int, y int) {
	bound := graphics.Bound{
		X:      t.imageXOffset,
		Y:      t.imageYOffset,
		Width:  64,
		Height: 80,
	}

	newX := x + t.xOffset
	newY := y + t.yOffset
	if !t.isBroken {
		batch.DrawSprite(newX, newY, y, bound, 1)
	}

	if t.gameItem != gameitem.None {
		if t.isBroken {
			batch.DrawSprite(
				newX+brokenTileXOffset, newY+brokenTileYOffset, y,
				t.gameItem.GetBound(), 1)
		} else if t.showItem {
			batch.DrawSprite(
				newX+revealItemXOffset, newY+revealItemYOffset, y+revealItemZOffset,
				t.gameItem.GetBound(), 0.6)
		}
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
		gameItem:     gameItem,
	}
}

func newGreen(gameItem gameitem.GameItem) Tile {
	return Tile{
		imageXOffset: 576,
		imageYOffset: 224,
		xOffset:      -4,
		yOffset:      -2,
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
