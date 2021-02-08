package gamestate

import (
	"bytes"
	"encoding/gob"

	"candy/game/direction"
)

type Character int

const (
	BlackBoy Character = iota
	BlackGirl
	BrownBoy
	BrownGirl
	YellowBoy
	YellowGirl
	OrangeBoy
	OrangeGirl
)

var characters = []Character{
	BlackBoy,
	BlackGirl,
	BrownBoy,
	BrownGirl,
	YellowBoy,
	YellowGirl,
	OrangeBoy,
	OrangeGirl,
}

type Position struct {
	X int
	Y int
}

type Player struct {
	ID        string
	Character Character
	Position  Position
}

type PlayerState struct {
	X           int
	Y           int
	Direction   direction.Direction
	CurrentStep int
}

func GetSyncPlayerMovePayload(payloadBytes []byte) (PlayerState, error) {
	var payload PlayerState

	dec := gob.NewDecoder(bytes.NewReader(payloadBytes))
	err := dec.Decode(&payload)
	if err != nil {
		return PlayerState{}, err
	}
	return payload, nil
}
