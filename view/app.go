package view

import (
	"fmt"
	"time"

	"candy/assets"
	"candy/graphics"
	"candy/input"
	"candy/router"
)

const ScreenWidth = 1152
const ScreenHeight = 830

var _ graphics.Sprite = (*App)(nil)

type App struct {
	router *router.Router
}

func (a App) Draw() {
	currRoute := a.router.CurrentRoute()
	if currRoute == nil {
		return
	}
	view := currRoute.Object.(view)
	view.Draw()
}

func (a App) Update(timeElapsed time.Duration) {
	currRoute := a.router.CurrentRoute()
	if currRoute == nil {
		return
	}
	view := currRoute.Object.(view)
	view.Update(timeElapsed)
}

func (a App) HandleInput(in input.Input) {
	currRoute := a.router.CurrentRoute()
	if currRoute == nil {
		return
	}
	view := currRoute.Object.(view)
	view.HandleInput(in)
}

func NewApp(assets assets.Assets, g graphics.Graphics) (App, error) {
	rt := router.NewRouter()
	routes := []router.Route{
		{Path: "/game", Object: NewGameScreen(assets, g)},
		{Path: "/", Object: NewSignInScreen(assets, g, rt)},
	}
	err := rt.AddRoutes(routes)
	if err != nil {
		return App{}, err
	}
	err = rt.Navigate("/")
	fmt.Println("Please click to get to next screen")
	if err != nil {
		return App{}, err
	}
	return App{
		router: rt,
	}, nil
}
