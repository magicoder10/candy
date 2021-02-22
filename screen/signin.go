package screen

import (
	"time"

	"candy/view"

	"github.com/teamyapp/ui/assets"
	"github.com/teamyapp/ui/audio"
	"github.com/teamyapp/ui/graphics"
	"github.com/teamyapp/ui/input"
	"github.com/teamyapp/ui/observability"
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
	s.screen.Init()
}

func (s SignIn) Destroy() {
	s.backgroundMusic.Stop()
	s.screen.Destroy()
}

func (s SignIn) Draw() {
	s.batch.DrawSprite(0, 0, 1, signInBackgroundBound, 1)
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

func NewSignIn(
	logger *observability.Logger,
	assets assets.Assets, g graphics.Graphics,
	router *view.Router,
) SignIn {
	return SignIn{
		screen: screen{
			name:   "Sign In",
			logger: logger,
		},
		backgroundMusic: assets.GetAudio("screen/signin_bg.mp3"),
		batch:           g.StartNewBatch(assets.GetImage("screen/signin.png")),
		router:          router,
	}
}
