package state

import (
	"time"

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

func (m *Melting) Update(timeElapsed time.Duration) State {
	m.remainingTime -= timeElapsed
	if m.remainingTime <= 0 {
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

func NewMelting(powerLevel int) *Melting {
	return &Melting{
		shared: shared{
			powerLevel:    powerLevel,
			remainingTime: meltingTime,
		},
	}
}
