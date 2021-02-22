package player

import (
	"candy/game/direction"

	"github.com/teamyapp/ui/graphics"
)

func draw(
	batch graphics.Batch,
	regionOffset regionOffset,
	character character,
	x int, y int, z int, scale float64,
	direction direction.Direction, step int,
) {
	bound := graphics.Bound{
		X:      regionOffset.x + character.walkCycleOffset.x + step*spriteWidth,
		Y:      regionOffset.y + character.walkCycleOffset.y + int(direction)*spriteHeight,
		Width:  spriteWidth,
		Height: spriteHeight,
	}
	batch.DrawSprite(x, y, z, bound, scale)
}
