package client

import (
	"context"
	"fmt"

	"candy/observability"
	"candy/server/proto"

	"google.golang.org/grpc"
)

type Client struct {
	logger  *observability.Logger
	conn    *grpc.ClientConn
	client  proto.GameClient
	context context.Context
}

func (c Client) CreateGame(maxPlayers int) (string, error) {
	res, err := c.client.CreateGame(c.context, &proto.CreateGameRequest{
		MaxPlayers: int32(maxPlayers),
	})
	if err != nil {
		return "", err
	}
	return res.GameId, nil
}

func (c Client) JoinGame(gameID string) (string, error) {
	res, err := c.client.JoinGame(c.context, &proto.JoinGameRequest{
		GameId: gameID,
	})
	if err != nil {
		return "", err
	}
	return res.AssignedPlayerId, nil
}

func (c Client) StartGame(gameID string) error {
	_, err := c.client.StartGame(c.context, &proto.StartGameRequest{
		GameId: gameID,
	})
	return err
}

func (c *Client) Start(ipAddress string, port int) error {
	var err error

	c.conn, err = grpc.Dial(fmt.Sprintf("%s:%d", ipAddress, port), grpc.WithInsecure())
	if err != nil {
		return err
	}
	c.client = proto.NewGameClient(c.conn)
	return nil
}

func (c Client) Stop() {
	c.conn.Close()
}

func New(logger *observability.Logger) *Client {
	return &Client{
		logger:  logger,
		context: context.Background(),
	}
}
