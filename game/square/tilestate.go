package square

import (
	"candy/game/gameitem"
	"candy/graphics"
)

type tileState interface {
	draw(batch graphics.Batch, x int, y int)
	unblockFire() tileState
	shouldRemove() bool
	canEnter() bool
	breakTile() tileState
	revealItem()
	hideItem()
	isBroken() bool
	breakable() bool
	removeItem() gameitem.GameItem
}

type tileSharedState struct {
	xOffset      int
	yOffset      int
	imageXOffset int
	imageYOffset int
	showItem     bool
	gameItem     gameitem.GameItem
}

func (t *tileSharedState) revealItem() {
	t.showItem = true
}

func (t *tileSharedState) hideItem() {
	t.showItem = false
}

func (t *tileSharedState) removeItem() gameitem.GameItem {
	return gameitem.None
}

func (t tileSharedState) shouldRemove() bool {
	return false
}

func (t tileSharedState) canEnter() bool {
	return false
}

func (t tileSharedState) breakable() bool {
	return false
}

func (t tileSharedState) isBroken() bool {
	return false
}

// tileSolidState

var _ tileState = (*tileSolidState)(nil)

type tileSolidState struct {
	tileSharedState
}

func (t tileSolidState) breakTile() tileState {
	return &tileBrokenState{t.tileSharedState}
}

func (t *tileSolidState) unblockFire() tileState {
	return t
}

func (t tileSolidState) breakable() bool {
	return true
}

func (t tileSolidState) draw(batch graphics.Batch, x int, y int) {
	bound := graphics.Bound{
		X:      t.imageXOffset,
		Y:      t.imageYOffset,
		Width:  64,
		Height: 80,
	}
	newX := x + t.xOffset
	newY := y + t.yOffset
	batch.DrawSprite(newX, newY, y, bound, 1)

	if t.gameItem != gameitem.None && t.showItem {
		batch.DrawSprite(
			newX+revealItemXOffset, newY+revealItemYOffset, y+revealItemZOffset,
			t.gameItem.GetBound(), 0.6)
	}
}

// tileBrokenState

var _ tileState = (*tileBrokenState)(nil)

type tileBrokenState struct {
	tileSharedState
}

func (t tileBrokenState) breakTile() tileState {
	return &t
}

func (t tileBrokenState) draw(batch graphics.Batch, x int, y int) {
	if t.gameItem == gameitem.None {
		return
	}
	batch.DrawSprite(
		x+t.xOffset+brokenTileXOffset, y+t.yOffset+brokenTileYOffset, y,
		t.gameItem.GetBound(), 1)
}

func (t tileBrokenState) unblockFire() tileState {
	return &tileCollectItemState{t}
}

func (t tileBrokenState) isBroken() bool {
	return true
}

// tileCollectItemState

var _ tileState = (*tileCollectItemState)(nil)

type tileCollectItemState struct {
	tileBrokenState
}

func (t tileCollectItemState) shouldRemove() bool {
	return t.gameItem == gameitem.None
}

func (t tileCollectItemState) removeItem() gameitem.GameItem {
	gameItem := t.gameItem
	t.gameItem = gameitem.None
	return gameItem
}

func (t tileCollectItemState) canEnter() bool {
	return true
}
