package audio

import (
    "os"
)

type Audio interface {
    Play()
    Stop()
}

func NewAudio(file *os.File, extension string) (Audio, error) {
    return newBeep(file, extension)
}
