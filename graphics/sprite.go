package graphics

import (
    "time"

    "candy/input"
)

type Sprite interface {
    Draw()
    Update(timeElapsed time.Duration)
    HandleInput(in input.Input)
}
