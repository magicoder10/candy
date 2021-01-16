package gamemap

import (
	"candy/screen/game/cell"
)

type eventHandlers struct {
	onCandyExploding func(hitCell cell.Cell)
}
