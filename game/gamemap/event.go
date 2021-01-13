package gamemap

import (
	"candy/game/cell"
)

type eventHandlers struct {
	onCandyExploding func(hitCell cell.Cell)
}
