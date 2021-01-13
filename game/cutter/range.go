package cutter

import (
	"candy/game/cell"
	"candy/game/direction"
)

type Range interface {
	CutRange(start cell.Cell, initialRange int, dir direction.Direction) int
}

var _ Range = (*NoChange)(nil)

type NoChange struct {
}

func (n NoChange) CutRange(start cell.Cell, initialRange int, dir direction.Direction) int {
	return initialRange
}
