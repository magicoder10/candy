package square

import (
	"candy/graphics"
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
}
