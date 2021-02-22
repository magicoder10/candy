package square

import (
	"candy/game/gameitem"

	"github.com/teamyapp/ui/graphics"
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

func (t Tile) HasRevealedItem() bool {
	return t.state.hasItem()
}

func (t *Tile) RetrieveGameItem() gameitem.Type {
	return t.state.removeItem()
}

func (t Tile) IsBroken() bool {
	return t.state.isBroken()
}

func (t Tile) ShouldRemove() bool {
	return t.state.shouldRemove()
}

func (t *Tile) UnblockFire() {
	t.state = t.state.unblockFire()
}

func (t *Tile) Break() {
	t.state = t.state.breakTile()
}

func (t Tile) IsBreakable() bool {
	return t.state.breakable()
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

func (t *Tile) RemoveItem() gameitem.Type {
	return t.state.removeItem()
}

func newYellow(gameItemType gameitem.Type) Tile {
	return Tile{
		state: &tileSolidState{
			tileSharedState{
				imageXOffset: 576,
				imageYOffset: 304,
				xOffset:      -4,
				yOffset:      -2,
				gameItemType: gameItemType,
			},
		},
	}
}

func newGreen(gameItemType gameitem.Type) Tile {
	return Tile{
		state: &tileSolidState{
			tileSharedState{
				imageXOffset: 576,
				imageYOffset: 224,
				xOffset:      -4,
				yOffset:      -2,
				gameItemType: gameItemType,
			},
		},
	}
}

func NewTile(tileType rune, gameItemType gameitem.Type) *Tile {
	switch tileType {
	case 'Y':
		tile := newYellow(gameItemType)
		return &tile
	case 'G':
		tile := newGreen(gameItemType)
		return &tile
	default:
		tile := newGreen(gameItemType)
		return &tile
	}
}
