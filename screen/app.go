package screen

import (
	"fmt"
	"time"

	"candy/assets"
	"candy/graphics"
	"candy/input"
	"candy/view"
)

const Width = 1152
const Height = 830

var _ graphics.Sprite = (*App)(nil)

type App struct {
	router *view.Router
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

func NewApp(assets assets.Assets, g graphics.Graphics) (App, error) {
	rt := view.NewRouter()
	routes := []view.Route{
		{Path: "/game", CreateFactory: func(props interface{}) view.View {
			return NewGame(assets, g)
		}},
		{Path: "/", CreateFactory: func(props interface{}) view.View {
			return NewSignIn(assets, g, rt)
		}},
	}
	err := rt.AddRoutes(routes)
	if err != nil {
		return App{}, err
	}
	err = rt.Navigate("/", nil)
	fmt.Println("Please click to get to next screen")
	if err != nil {
		return App{}, err
	}
	return App{
		router: rt,
	}, nil
}
