package main

import (
	"image/draw"

	"candy/assets"
	"candy/observability"
	"candy/pubsub"
	"candy/ui"
)

const Height = 830

var _ ui.Component = (*App)(nil)

type App struct {
	ui.SharedComponent
	logger *observability.Logger
	assets assets.Assets
	router *ui.Router
	pubSub *pubsub.PubSub
}

func (a *App) Paint(painter *ui.Painter, destLayer draw.Image, offset ui.Offset) {
	if len(a.Children) > 0 {
		a.Children[0].Paint(painter, destLayer, offset)
	}
}

func (a *App) Init() {
	a.pubSub.Start()
	err := a.router.Navigate("/", nil)
	if err != nil {
		return
	}
}

func NewApp(logger *observability.Logger, assets assets.Assets) (*App, error) {
	rt := ui.NewRouter(logger)
	pubSub := pubsub.NewPubSub(logger)

	routes := []ui.Route{
		{Path: "/demo", CreateFactory: func(props interface{}) ui.Component {
			return newDemo()
		}},
		{Path: "/", CreateFactory: func(props interface{}) ui.Component {
			return NewSignIn(rt, assets)
		}},
	}
	err := rt.AddRoutes(routes)
	if err != nil {
		return nil, err
	}
	statefulStyle := ui.NewStatefulStyle()
	statefulStyle.Styles[ui.NormalState].LayoutType = ui.LayoutTypePtr(ui.BoxLayoutType)

	app := &App{
		logger: logger,
		assets: assets,
		pubSub: pubSub,
		router: rt,
		SharedComponent: ui.SharedComponent{
			Name:          "App",
			StatefulStyle: statefulStyle,
			Children:      []ui.Component{},
		},
	}
	rt.OnCurrentChange(func(curr ui.Component) {
		app.Children = []ui.Component{curr}
		app.MarkChanged()
	})
	return app, err
}
