package player

import (
	"candy/screen/game/direction"
)

type MoveChecker interface {
	CanMove(currX int, currY int, objectWidth int, objectHeight int, dir direction.Direction, stepSize int) bool
}
