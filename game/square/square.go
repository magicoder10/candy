package square

import (
	"candy/game/gameitem"

	"github.com/teamyapp/ui/graphics"
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
}
