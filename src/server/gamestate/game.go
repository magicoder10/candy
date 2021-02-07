package gamestate

import (
	"bytes"
	"encoding/gob"
	"errors"
	"math/rand"
	"time"

	"candy/pubsub"

	uuid "github.com/satori/go.uuid"
)

type Game struct {
	pubSubRemote *pubsub.Remote
	ID           string
	MaxPlayers   int
	players      map[string]Player
	gameMap      Map
}

func (g *Game) JoinGame() (Player, error) {
	if len(g.players) >= g.MaxPlayers {
		return Player{}, errors.New("game is full")
	}
	playerID := randomID()
	randomCharacter := characters[rand.Intn(len(characters))]

	ply := Player{
		ID:        playerID,
		Character: randomCharacter,
	}

	g.players[playerID] = ply
	g.pubSubRemote.Publish(pubsub.NewGameSetup(g.ID), Setup{Players: g.players})
	return ply, nil
}

func (g *Game) Start() {
	rand.Seed(time.Now().UnixNano())

	gMap := newMap()
	emptyPositions := make([]Position, 0)

	for row, squareRow := range gMap.Squares {
		for col, sq := range squareRow {
			if sq.SquareType == EmptySquare {
				emptyPositions = append(emptyPositions, Position{
					X: col * squareWidth,
					Y: row * squareWidth,
				})
			}
		}
	}
	shuffle(emptyPositions)

	index := 0

	for id, ply := range g.players {
		if index >= len(emptyPositions) {
			continue
		}
		ply.Position = emptyPositions[index]
		g.players[id] = ply
		index++
	}

	g.gameMap = gMap

	g.pubSubRemote.Publish(pubsub.NewStartGame(g.ID), Setup{
		Players: g.players,
		Map:     g.gameMap,
	})
}

func shuffle(positions []Position) {
	for i := range positions {
		choice := i + int(float32(len(positions)-i)*rand.Float32())
		swap(positions, i, choice)
	}
}

func swap(positions []Position, a int, b int) {
	positions[a], positions[b] = positions[b], positions[a]
}

func randomID() string {
	return uuid.Must(uuid.NewV4(), nil).String()
}

func NewGame(pubSubRemote *pubsub.Remote, maxPlayers int) Game {
	return Game{
		pubSubRemote: pubSubRemote,
		ID:           randomID(),
		MaxPlayers:   maxPlayers,
		players:      make(map[string]Player),
	}
}

type Setup struct {
	Players map[string]Player
	Map     Map
}

func GetSetup(payloadBytes []byte) (Setup, error) {
	var payload Setup
	dec := gob.NewDecoder(bytes.NewReader(payloadBytes))
	err := dec.Decode(&payload)
	if err != nil {
		return Setup{}, err
	}
	return payload, nil
}
