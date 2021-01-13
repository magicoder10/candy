package candy

import (
	"candy/game/cell"
	"candy/game/direction"
)

type RangeCutter interface {
	CutRange(start cell.Cell, initialRange int, dir direction.Direction) int
}

var _ RangeCutter = (*noChange)(nil)

type noChange struct {
}

func (n noChange) CutRange(start cell.Cell, initialRange int, dir direction.Direction) int {
	return initialRange
}
