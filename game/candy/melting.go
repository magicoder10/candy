package candy

import (
	"time"

	"candy/game/cell"
	"candy/graphics"
)

const meltingTime = 7 * time.Second

var meltingImageDuration = (180 * time.Millisecond).Nanoseconds()

const maxMeltingImages = 4

const width = 60
const height = 60

var _ state = (*meltingState)(nil)

type meltingState struct {
	sharedState
	meltingImageIndex int
}

func (m *meltingState) update(timeElapsed time.Duration) state {
	m.remainingTime -= timeElapsed
	if m.remainingTime <= 0 || m.shouldExplode {
		return newExplodingState(m.sharedState)
	}
	m.lag += timeElapsed.Nanoseconds()

	imageJumps := int(m.lag / meltingImageDuration)
	m.meltingImageIndex = (m.meltingImageIndex + imageJumps) % maxMeltingImages
	m.sharedState.lag %= meltingImageDuration
	return m
}

func (m meltingState) draw(batch graphics.Batch, x int, y int, z int) {
	bound := graphics.Bound{
		X:      640,
		Y:      323 - m.meltingImageIndex*height,
		Width:  width,
		Height: height,
	}
	batch.DrawSprite(x, y, z, bound, 1)
}

func newMeltingState(powerLevel int, center cell.Cell, rangeCutter RangeCutter) *meltingState {
	return &meltingState{
		sharedState: sharedState{
			center:        center,
			powerLevel:    powerLevel,
			remainingTime: meltingTime,
			rangeCutter:   rangeCutter,
		},
	}
}
