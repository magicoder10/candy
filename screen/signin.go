package screen

import (
	"os"
	"time"

	"candy/assets"
	"candy/audio"
	"candy/client"
	"candy/graphics"
	"candy/input"
	"candy/observability"
	"candy/pubsub"
	"candy/server/gamestate"
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
	remotePubSub    *pubsub.Remote
	playerID        string
	isCreator       bool
	gameID          string
	client          *client.Client
}

func (s *SignIn) Init() {
	s.backgroundMusic.Play()
	s.screen.Init()

	args := os.Args[1:]

	var err error

	if len(args) < 1 {
		s.gameID, err = s.client.CreateGame(8)
		if err != nil {
			s.logger.Errorf("%w\n", err)
			return
		}
		s.isCreator = true
		s.logger.Infof("Created game:%s\n", s.gameID)
		s.logger.Infof("Invite others to your game: \n go run main.go %s\n", s.gameID)
		s.logger.Infoln("Click on sign in screen to start the game")
	} else {
		s.gameID = args[0]
	}

	s.remotePubSub.Subscribe(pubsub.NewStartGame(s.gameID), func(payload []byte) {
		gameSetup, err := gamestate.GetSetup(payload)
		if err != nil {
			s.logger.Errorln(err)
			return
		}
		s.router.Navigate("/game", gameRouteProps{
			gameSetup: gameSetup,
			gameID:    s.gameID,
			playerID:  s.playerID,
		})
	})

	s.playerID, err = s.client.JoinGame(s.gameID)
	if err != nil {
		s.logger.Errorln(err)
		return
	}
	s.logger.Infof("Joined game:%s\n", s.gameID)
	s.logger.Infof("Player id:%s\n", s.playerID)
}

func (s SignIn) Destroy() {
	s.backgroundMusic.Stop()
	s.screen.Destroy()
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
			if s.isCreator {
				s.client.StartGame(s.gameID)
			}
		}
	}
}

func NewSignIn(
	logger *observability.Logger,
	assets assets.Assets, g graphics.Graphics,
	router *view.Router,
	remotePubSub *pubsub.Remote,
	client *client.Client,
) *SignIn {
	return &SignIn{
		screen: screen{
			name:   "Sign In",
			logger: logger,
		},
		backgroundMusic: assets.GetAudio("screen/signin_bg.mp3"),
		batch:           g.StartNewBatch(assets.GetImage("screen/signin.png")),
		router:          router,
		remotePubSub:    remotePubSub,
		client:          client,
	}
}
