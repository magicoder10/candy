package screen

import (
	"fmt"
	"time"

	"candy/assets"
	"candy/graphics"
	"candy/input"
	"candy/pubsub"
	"candy/view"
)

const Width = 1152
const Height = 830

var _ graphics.Sprite = (*App)(nil)

type App struct {
	assets assets.Assets
	router *view.Router
	pubSub *pubsub.PubSub
}

func (a App) Draw() {
	currView := a.router.CurrentView()
	if currView == nil {
		return
	}
	currView.Draw()
}

func (a App) Update(timeElapsed time.Duration) {
	currView := a.router.CurrentView()
	if currView == nil {
		return
	}
	currView.Update(timeElapsed)
}

func (a App) HandleInput(in input.Input) {
	currView := a.router.CurrentView()
	if currView == nil {
		return
	}
	currView.HandleInput(in)
}

func (a *App) Launch() error {
	a.pubSub.Start()
	err := a.router.Navigate("/", nil)
	if err != nil {
		return err
	}
	fmt.Println("Please click to get to next screen")
	return nil
}

func NewApp(assets assets.Assets, g graphics.Graphics) (App, error) {
	rt := view.NewRouter()
	pubSub := pubsub.NewPubSub()

	routes := []view.Route{
		{Path: "/game", CreateFactory: func(props interface{}) view.View {
			gm := NewGame(assets, g, pubSub)
			return gm
		}},
		{Path: "/", CreateFactory: func(props interface{}) view.View {
			return NewSignIn(assets, g, rt)
		}},
	}
	err := rt.AddRoutes(routes)
	return App{
		assets: assets,
		pubSub: pubSub,
		router: rt,
	}, err
}
