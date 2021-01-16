package gameitem

import (
	"candy/graphics"
)

type GameItem int

const (
	None GameItem = iota
	Speed
	Power
	Candy
	FirstAidKit
)

var GameItems = []GameItem{
	None,
	Speed,
	Power,
	Candy,
	FirstAidKit,
}

func (g GameItem) GetBound() graphics.Bound {
	switch g {
	case Speed:
		return graphics.Bound{
			X:      761,
			Y:      204,
			Width:  60,
			Height: 60,
		}
	case Power:
		return graphics.Bound{
			X:      761,
			Y:      264,
			Width:  60,
			Height: 60,
		}
	case Candy:
		return graphics.Bound{
			X:      761,
			Y:      144,
			Width:  60,
			Height: 60,
		}
	case FirstAidKit:
		return graphics.Bound{
			X:      761,
			Y:      84,
			Width:  60,
			Height: 60,
		}
	default:
		return graphics.Bound{}
	}
}
