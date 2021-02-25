package input

import (
	"fmt"
	"image"
)

type Input struct {
	Action         action
	Device         device
	CursorPosition image.Point
}

func (in Input) String() string {
	return fmt.Sprintf("[Input(Action=%s Device=%s)]", actionNames[in.Action], deviceNames[in.Device])
}

var actionNames = map[action]string{
	Press:       "Press",
	SinglePress: "SinglePress",
	Release:     "Release",
}

type action int

const (
	Press action = iota
	SinglePress
	Release
)

var deviceNames = map[device]string{
	LeftArrowKey:    "LeftArrowKey",
	RightArrowKey:   "RightArrowKey",
	UpArrowKey:      "UpArrowKey",
	DownArrowKey:    "DownArrowKey",
	RKey:            "RKey",
	SpaceKey:        "SpaceKey",
	MouseLeftButton: "MouseLeftButton",
}

type device int

const (
	LeftArrowKey device = iota
	RightArrowKey
	UpArrowKey
	DownArrowKey
	RKey
	SpaceKey
	MouseLeftButton
)
