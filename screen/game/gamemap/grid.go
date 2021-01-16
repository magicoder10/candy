package gamemap

import (
	"candy/screen/game/cell"
)

func inGrid(cell cell.Cell, maxRow int, maxCol int) bool {
	return cell.Row >= 0 && cell.Row <= maxRow &&
		cell.Col >= 0 && cell.Col <= maxCol

}
