package player

import (
	"candy/graphics"
)

type Tombstone struct {
	lag        int64
	width      int
	imageBound graphics.Bound
}

func (ts Tombstone) getWidth() int {
	return ts.width
}
func (ts Tombstone) draw(batch graphics.Batch, x int, y int, z int) {
	bound := ts.imageBound
	batch.DrawSprite(x, y, z, bound, 1)
}

func newTombstone() Tombstone {
	return Tombstone{
		lag:   0,
		width: 60,
		imageBound: graphics.Bound{
			X:      960,
			Y:      933,
			Width:  54,
			Height: 77,
		},
	}
}
