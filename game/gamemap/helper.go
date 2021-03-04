package gamemap

import (
	"candy/game/cell"
	"candy/game/square"
)

func (m Map) GetObjectCell(objectX int, objectY int, objectWidth int, objectHeight int) cell.Cell {
	return cell.GetCellLocatedAt(
		objectX-m.gridXOffset, objectY-m.gridYOffset, objectWidth, objectHeight,
		square.Width, square.Width,
	)
}
