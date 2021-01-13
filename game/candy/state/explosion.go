package state

import (
	"candy/game/square"
	"candy/graphics"
)

var explosionCenter = graphics.Bound{
	X:      761,
	Y:      324,
	Width:  square.Width,
	Height: square.Width,
}
var explosionVerticalEdge = graphics.Bound{
	X:      701,
	Y:      24,
	Width:  square.Width,
	Height: square.Width,
}
var explosionHorizontalEdge = graphics.Bound{
	X:      701,
	Y:      84,
	Width:  square.Width,
	Height: square.Width,
}
var explosionTopEnd = graphics.Bound{
	X:      701,
	Y:      324,
	Width:  square.Width,
	Height: square.Width,
}
var explosionBottomEnd = graphics.Bound{
	X:      701,
	Y:      144,
	Width:  square.Width,
	Height: square.Width,
}
var explosionLeftEnd = graphics.Bound{
	X:      701,
	Y:      264,
	Width:  square.Width,
	Height: square.Width,
}
var explosionRightEnd = graphics.Bound{
	X:      701,
	Y:      204,
	Width:  square.Width,
	Height: square.Width,
}
