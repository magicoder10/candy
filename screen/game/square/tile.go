package square

import (
	"candy/graphics"
	"candy/screen/game/gameitem"
)

const revealItemXOffset = 16
const revealItemYOffset = 36
const revealItemZOffset = -1
const brokenTileXOffset = 4
const brokenTileYOffset = 2

var _ Square = (*Tile)(nil)

type Tile struct {
	state tileState
}

func (t Tile) ShouldRemove() bool {
	return t.state.shouldRemove()
}

func (t *Tile) UnblockFire() {
	t.state = t.state.unblockFire()
}

func (t Tile) IsBlocking() bool {
	return true
}

func (t *Tile) Break() {
	t.state = t.state.breakTile()
}

func (t Tile) IsBreakable() bool {
	return true
}

func (t Tile) Draw(batch graphics.Batch, x int, y int) {
	t.state.draw(batch, x, y)
}

func (t Tile) CanEnter() bool {
	return t.state.canEnter()
}

func (t *Tile) RevealItem() {
	t.state.revealItem()
}

func (t *Tile) HideItem() {
	t.state.hideItem()
}

func (t *Tile) RemoveItem() gameitem.GameItem {
	return t.state.removeItem()
}

func newYellow(gameItem gameitem.GameItem) Tile {
	return Tile{
		state: &tileSolidState{
			tileSharedState{
				imageXOffset: 576,
				imageYOffset: 304,
				xOffset:      -4,
				yOffset:      -2,
				gameItem:     gameItem,
			},
		},
	}
}

func newGreen(gameItem gameitem.GameItem) Tile {
	return Tile{
		state: &tileSolidState{
			tileSharedState{
				imageXOffset: 576,
				imageYOffset: 224,
				xOffset:      -4,
				yOffset:      -2,
				gameItem:     gameItem,
			},
		},
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
