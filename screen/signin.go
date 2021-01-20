package screen

import (
	"time"

	"candy/assets"
	"candy/audio"
	"candy/graphics"
	"candy/input"
	"candy/view"
)

var _ view.View = (*SignIn)(nil)

var signInBackgroundBound = graphics.Bound{
	X:      0,
	Y:      0,
	Width:  Width,
	Height: Height,
}

type SignIn struct {
	screen
	backgroundMusic audio.Audio
	batch           graphics.Batch
	router          *view.Router
}

func (s SignIn) Init() {
	s.backgroundMusic.Play()
}

func (s SignIn) Destroy() {
	s.backgroundMusic.Stop()
}

func (s SignIn) Draw() {
	s.batch.DrawSprite(0, 0, 1, signInBackgroundBound, 1)
	s.batch.RenderBatch()
}

func (s SignIn) Update(timeElapsed time.Duration) {
	return
}

func (s SignIn) HandleInput(in input.Input) {
	switch in.Action {
	case input.SinglePress:
		switch in.Device {
		case input.MouseLeftButton:
			s.router.Navigate("/game", nil)
		}
	}
}

func NewSignIn(assets assets.Assets, g graphics.Graphics, router *view.Router) SignIn {
	return SignIn{
		backgroundMusic: assets.GetAudio("screen/signin_bg.mp3"),
		batch:           g.StartNewBatch(assets.GetImage("screen/signin.png")),
		router:          router,
	}
}
