package pubsub

import (
	"fmt"

	"candy/game/direction"
)

type Topic string

const (
	OnCandyExploding     Topic = "onCandyExploding"
	OnPlayerWalking      Topic = "onPlayerWalking"
	IncrementPlayerPower Topic = "IncrementPlayerPower"
	OnDropCandy          Topic = "onDropCandy"
)

func NewGameSetup(gameID string) Topic {
	return Topic(fmt.Sprintf("game/%s/setup", gameID))
}

func NewStartGame(gameID string) Topic {
	return Topic(fmt.Sprintf("game/%s/start", gameID))
}

func NewSyncPlayerMove(gameID string, playerID string) Topic {
	return Topic(fmt.Sprintf("game/%s/move/%s", gameID, playerID))
}

func NewSyncDropCandy(gameID string) Topic {
	return Topic(fmt.Sprintf("game/%s/drop-candy", gameID))
}

func NewSyncRetrieveGameItem(gameID string) Topic {
	return Topic(fmt.Sprintf("game/%s/gameitem/retrieve", gameID))
}

type OnDropCandyPayload struct {
	X      int
	Y      int
	Width  int
	Height int
}

type OnPlayerWalkingPayload struct {
	PlayerID  string
	X         int
	Y         int
	Width     int
	Height    int
	Direction direction.Direction
	CurrStep  int
}
