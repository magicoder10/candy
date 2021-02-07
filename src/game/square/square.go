package square

import (
	"candy/game/gameitem"
	"candy/graphics"
	"candy/server/gamestate"
)

const Width = 60

type Square interface {
	Draw(batch graphics.Batch, x int, y int)
	IsBreakable() bool
	Break()
	UnblockFire()
	IsBroken() bool
	ShouldRemove() bool
	CanEnter() bool
	HasRevealedItem() bool
	RetrieveGameItem() gameitem.Type
	RevealItem()
	HideItem()
}

func New(square gamestate.Square) Square {
	switch square.SquareType {
	case gamestate.YellowTile:
		return newYellowTile(gameitem.NewType(square.GameItemType))
	case gamestate.GreenTile:
		return newGreenTile(gameitem.NewType(square.GameItemType))
	}
	return nil
}
