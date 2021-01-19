package view

import (
	"time"

	"candy/assets"
	"candy/graphics"
	"candy/input"
	"candy/router"
)

var _ view = (*SignInScreen)(nil)

var signInBackgroundBound = graphics.Bound{
	X:      0,
	Y:      0,
	Width:  ScreenWidth,
	Height: ScreenHeight,
}

type SignInScreen struct {
	batch  graphics.Batch
	router *router.Router
}

func (s SignInScreen) Draw() {

	s.batch.DrawSprite(0, 0, 1, signInBackgroundBound, 1)
	s.batch.RenderBatch()
}

func (s SignInScreen) Update(timeElapsed time.Duration) {
	return
}

func (s SignInScreen) HandleInput(in input.Input) {
	switch in.Action {
	case input.SinglePress:
		switch in.Device {
		case input.MouseLeftButton:
			s.router.Navigate("/game")
		}
	}
}

func NewSignInScreen(assets assets.Assets, g graphics.Graphics, router *router.Router) SignInScreen {
	return SignInScreen{
		batch:  g.StartNewBatch(assets.GetImage("screen/signin.png")),
		router: router,
	}
}
