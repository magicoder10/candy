package marker

import (
	"candy/graphics"
	"time"
)

const YOffset = 5
var markerImageDuration = (180 * time.Millisecond).Nanoseconds()

type Marker struct {
	lag int64
	width int
	imageIndex int
	imageSet []graphics.Bound
}

func (m Marker) GetWidth()  int {
	return m.width
}
func (m Marker) Draw(batch graphics.Batch, x int, y int, z int) {
	bound := m.imageSet[m.imageIndex]
	batch.DrawSprite(x, y, z, bound, 1)
}

func (m *Marker) Update(timeElapsed time.Duration) {
	m.lag += timeElapsed.Nanoseconds()
	imageJumps := int(m.lag / markerImageDuration)
	m.imageIndex = (m.imageIndex + imageJumps) % len(m.imageSet)
	m.lag -= int64(imageJumps) * markerImageDuration
}

func newCurrentPlayer() Marker {
	return Marker{
		lag:        0,
		width:      14,
		imageIndex: 0,
		imageSet:   []graphics.Bound{
			{
				X:      960,
				Y:      1032,
				Width:  14,
				Height: 22,
			},
			{
				X:      974,
				Y:      1032,
				Width:  14,
				Height: 20,
			},
			{
				X:      988,
				Y:      1032,
				Width:  14,
				Height: 18,
			},
			{
				X:      1002,
				Y:      1032,
				Width:  14,
				Height: 20,
			},
		},
	}
}


// TODO: isTeammate feature needs to be added later
func NewMarker(isTeammate bool) Marker {
	return newCurrentPlayer()
}




