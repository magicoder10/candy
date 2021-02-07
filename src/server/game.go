package server

import (
	"context"
	"errors"
	"fmt"

	"candy/observability"
	"candy/pubsub"
	"candy/server/gamestate"
	"candy/server/proto"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

var _ proto.GameServer = (*Game)(nil)

type Game struct {
	GRPCServer
	pubSubRemote *pubsub.Remote
	games        map[string]gamestate.Game
}

func (g *Game) CreateGame(ctx context.Context, request *proto.CreateGameRequest) (*proto.CreateGameResponse, error) {
	gm := gamestate.NewGame(g.pubSubRemote, int(request.MaxPlayers))
	g.games[gm.ID] = gm
	return &proto.CreateGameResponse{
		GameId: gm.ID,
	}, nil
}

func (g *Game) JoinGame(ctx context.Context, request *proto.JoinGameRequest) (*proto.JoinGameResponse, error) {
	gm, ok := g.games[request.GameId]
	if !ok {
		return nil, errors.New(fmt.Sprintf("game not found:%s", request.GameId))
	}
	ply, err := gm.JoinGame()
	if err != nil {
		return nil, err
	}
	return &proto.JoinGameResponse{
		AssignedPlayerId: ply.ID,
	}, nil
}

func (g *Game) StartGame(ctx context.Context, request *proto.StartGameRequest) (*empty.Empty, error) {
	gm, ok := g.games[request.GameId]
	if !ok {
		return nil, errors.New(fmt.Sprintf("game not found:%s", request.GameId))
	}
	gm.Start()
	return &empty.Empty{}, nil
}

func (g *Game) Start(port int) error {
	g.grpcServer = grpc.NewServer()
	proto.RegisterGameServer(g.grpcServer, g)

	g.GRPCServer.Start(port)
	return nil
}

func NewGame(logger *observability.Logger, pubSubRemote *pubsub.Remote) Game {
	return Game{
		GRPCServer: GRPCServer{
			name:   "Game",
			logger: logger,
		},
		pubSubRemote: pubSubRemote,
		games:        make(map[string]gamestate.Game),
	}
}
