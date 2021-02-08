package gamestate

import (
	"bytes"
	"encoding/gob"
)

type Cell struct {
	Row int
	Col int
}

func GetCell(payloadBytes []byte) (Cell, error) {
	var payload Cell

	dec := gob.NewDecoder(bytes.NewReader(payloadBytes))
	err := dec.Decode(&payload)
	if err != nil {
		return Cell{}, err
	}
	return payload, nil
}
