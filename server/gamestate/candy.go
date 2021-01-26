package gamestate

import (
	"bytes"
	"encoding/gob"
)

type Candy struct {
	Cell
	PowerLevel int
}

func GetCandy(payloadBytes []byte) (Candy, error) {
	var payload Candy

	dec := gob.NewDecoder(bytes.NewReader(payloadBytes))
	err := dec.Decode(&payload)
	if err != nil {
		return Candy{}, err
	}
	return payload, nil
}
