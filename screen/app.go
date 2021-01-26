package screen

import (
	"fmt"
	"time"

	"candy/assets"
	"candy/client"
	"candy/graphics"
	"candy/input"
	"candy/observability"
	"candy/pubsub"
	"candy/server/gamestate"
	"candy/view"
)

const Width = 1152
const Height = 830

var _ graphics.Sprite = (*App)(nil)

type App struct {
	logger       *observability.Logger
	assets       assets.Assets
	router       *view.Router
	pubSub       *pubsub.PubSub
	remotePubSub *pubsub.Remote
	gameClient   *client.Client
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
	a.logger.Debugf("User action: %s\n", in)
	currView := a.router.CurrentView()
	if currView == nil {
		return
	}
	currView.HandleInput(in)
}

func (a *App) Launch() error {
	a.pubSub.Start()
	err := a.remotePubSub.Start("localhost", 8081)
	if err != nil {
		return err
	}

	err = a.gameClient.Start("localhost", 8082)
	if err != nil {
		return err
	}

	err = a.router.Navigate("/", nil)
	if err != nil {
		return err
	}
	fmt.Println("Please click to get to next screen")
	return nil
}

func NewApp(logger *observability.Logger, assets assets.Assets, g graphics.Graphics) (App, error) {
	rt := view.NewRouter(logger)
	pubSub := pubsub.NewPubSub(logger)
	remotePubSub := pubsub.NewRemote(logger)
	gameClient := client.New(logger)

	routes := []view.Route{
		{Path: "/game", CreateFactory: func(props interface{}) view.View {
			parsedProps := props.(gameRouteProps)
			gm := NewGame(
				logger, assets, g, pubSub, remotePubSub,
				parsedProps.gameID, parsedProps.gameSetup, parsedProps.playerID,
			)
			return gm
		}},
		{Path: "/", CreateFactory: func(props interface{}) view.View {
			return NewSignIn(logger, assets, g, rt, remotePubSub, gameClient)
		}},
	}
	err := rt.AddRoutes(routes)
	return App{
		logger:       logger,
		assets:       assets,
		pubSub:       pubSub,
		remotePubSub: remotePubSub,
		gameClient:   gameClient,
		router:       rt,
	}, err
}

type gameRouteProps struct {
	gameSetup gamestate.Setup
	gameID    string
	playerID  string
}
