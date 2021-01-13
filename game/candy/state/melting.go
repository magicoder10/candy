package state

import (
	"time"

	"candy/game/cell"
	"candy/game/cutter"
	"candy/graphics"
)

const meltingTime = 7 * time.Second

var meltingImageDuration = (180 * time.Millisecond).Nanoseconds()

const maxMeltingImages = 4

const width = 60
const height = 60

var _ State = (*Melting)(nil)

type Melting struct {
	shared
	meltingImageIndex int
}

func (m *Melting) CellsHit() []cell.Cell {
	return []cell.Cell{}
}

func (m *Melting) Exploding() bool {
	return false
}

func (m *Melting) Update(timeElapsed time.Duration) State {
	m.remainingTime -= timeElapsed
	if m.remainingTime <= 0 || m.shouldExplode {
		return newExploding(m.shared)
	}
	m.lag += timeElapsed.Nanoseconds()

	imageJumps := int(m.lag / meltingImageDuration)
	m.meltingImageIndex = (m.meltingImageIndex + imageJumps) % maxMeltingImages
	m.shared.lag %= meltingImageDuration
	return m
}

func (m Melting) Draw(batch graphics.Batch, x int, y int, z int) {
	bound := graphics.Bound{
		X:      640,
		Y:      323 - m.meltingImageIndex*height,
		Width:  width,
		Height: height,
	}
	batch.DrawSprite(x, y, z, bound, 1)
}

func (m Melting) Exploded() bool {
	return false
}

func NewMelting(powerLevel int, center cell.Cell, rangeCutter cutter.Range) *Melting {
	return &Melting{
		shared: shared{
			center:        center,
			powerLevel:    powerLevel,
			remainingTime: meltingTime,
			rangeCutter:   rangeCutter,
		},
	}
}
