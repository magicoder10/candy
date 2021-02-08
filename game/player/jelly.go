package player

import (
    "time"

    "candy/graphics"
)

var JellyImageDuration = (180 * time.Millisecond).Nanoseconds()

type Jelly struct {
    lag        int64
    width      int
    imageIndex int
    imageSet   []graphics.Bound
}

func (j Jelly) getWidth() int {
    return j.width
}
func (j Jelly) draw(batch graphics.Batch, x int, y int, z int) {
    bound := j.imageSet[j.imageIndex]
    batch.DrawSprite(x, y, z, bound, 1)
}

func (j *Jelly) update(timeElapsed time.Duration) {
    j.lag += timeElapsed.Nanoseconds()
    imageJumps := int(j.lag / JellyImageDuration)
    j.imageIndex = (j.imageIndex + imageJumps) % len(j.imageSet)
    j.lag -= int64(imageJumps) * JellyImageDuration
}

func newJelly() Jelly {
    return Jelly{
        lag:        0,
        width:      60,
        imageIndex: 0,
        imageSet: []graphics.Bound{
            {
                X:      900,
                Y:      1134,
                Width:  60,
                Height: 72,
            },
            {
                X:      900,
                Y:      1054,
                Width:  60,
                Height: 65,
            },
            {
                X:      900,
                Y:      974,
                Width:  60,
                Height: 58,
            },
            {
                X:      900,
                Y:      894,
                Width:  60,
                Height: 65,
            },
            {
                X:      960,
                Y:      1134,
                Width:  60,
                Height: 72,
            },
            {
                X:      960,
                Y:      1054,
                Width:  60,
                Height: 78,
            },
        },
    }
}
