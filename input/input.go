package input

type Input struct {
	Action action
	Device device
}

type action int

const (
	Press action = iota
	SinglePress
	Release
)

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
