package input

type Input struct {
	Action action
	Device device
}

type action int

const (
	Press action = iota
	Release
)

type device int

const (
	LeftArrowKey device = iota
	RightArrowKey
	UpArrowKey
	DownArrayKey
)
