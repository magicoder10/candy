package audio

import (
    "os"
    "time"

    "github.com/faiface/beep"
    "github.com/faiface/beep/mp3"
    "github.com/faiface/beep/speaker"
    "github.com/faiface/beep/wav"
)

const mp3Extension = "mp3"
const wavExtension = "wav"

type audioDecoder = func(f *os.File) (s beep.StreamSeekCloser, format beep.Format, err error)

var beepDecoders = map[string]audioDecoder{
    mp3Extension: func(f *os.File) (s beep.StreamSeekCloser, format beep.Format, err error) {
        return mp3.Decode(f)
    },
    wavExtension: func(f *os.File) (s beep.StreamSeekCloser, format beep.Format, err error) {
        return wav.Decode(f)
    }}

var _ Audio = (*Beep)(nil)

type Beep struct {
    streamer beep.StreamSeekCloser
    format   beep.Format
}

func (b Beep) Play() {
    err := speaker.Init(b.format.SampleRate, b.format.SampleRate.N(time.Second/10))
    if err != nil {
        return
    }
    speaker.Play(b.streamer)
}

func (b Beep) Stop() {
    b.streamer.Close()
}

func newBeep(file *os.File, extension string) (Beep, error) {
    streamer, format, err := beepDecoders[extension](file)
    if err != nil {
        return Beep{}, err
    }
    return Beep{
        streamer: streamer,
        format:   format,
    }, nil
}
