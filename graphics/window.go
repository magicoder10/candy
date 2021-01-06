package graphics

import (
	"candy/input"
)

type Window interface {
	IsClosed() bool
	PollEvents() *input.Input
	Redraw()
}
