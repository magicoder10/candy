package gamemap

import (
	"candy/screen/game/cell"
	"candy/screen/game/direction"
	"candy/screen/game/player"
	"candy/screen/game/square"
)

var _ player.MoveChecker = (*moveChecker)(nil)

type moveChecker struct {
	gridXOffset int
	gridYOffset int
	maxRow      int
	maxCol      int
	grid        *[][]square.Square
}

func (m moveChecker) CanMove(currX int, currY int, objectWidth int, objectHeight int, dir direction.Direction, stepSize int) bool {
	if !m.inBound(currX, currY, objectWidth, objectHeight, dir, stepSize) {
		return false
	}

	currX -= m.gridXOffset
	currY -= m.gridYOffset

	cornerCells := cell.GetCornerCells(currX, currY, objectWidth, objectHeight, square.Width, square.Width)
	neighborCells := m.getNeighborCells(cornerCells, dir)
	blockingCells := m.getBlockingCells(neighborCells)

	for _, blockingCell := range blockingCells {
		if isBlocked(blockingCell, currX, currY, objectWidth, objectHeight, dir, stepSize) {
			return false
		}
	}
	return true
}

func (m moveChecker) inBound(currX int, currY int, objectWidth, objectHeight int, dir direction.Direction, stepSize int) bool {
	switch dir {
	case direction.Up:
		nextY := currY + objectHeight + stepSize
		return nextY <= m.gridYOffset+(m.maxRow+1)*square.Width
	case direction.Down:
		nextY := currY - stepSize
		return nextY >= m.gridYOffset
	case direction.Left:
		nextX := currX - stepSize
		return nextX >= m.gridXOffset
	case direction.Right:
		nextX := currX + objectWidth + stepSize
		return nextX <= m.gridXOffset+(m.maxCol+1)*square.Width
	}
	return true
}

func (m moveChecker) getBlockingCells(cells []cell.Cell) []cell.Cell {
	newCells := make([]cell.Cell, 0)
	for _, c := range cells {
		if m.maxRow < c.Row || m.maxCol < c.Col {
			continue
		}
		if (*m.grid)[c.Row][c.Col] == nil {
			continue
		}
		if (*m.grid)[c.Row][c.Col].CanEnter() {
			continue
		}
		newCells = append(newCells, c)
	}
	return newCells
}

func (m moveChecker) getNeighborCells(cornerCells cell.CornerCells, dir direction.Direction) []cell.Cell {
	switch dir {
	case direction.Up:
		return cell.GetTopNeighborCells(cornerCells, m.maxRow)
	case direction.Down:
		return cell.GetBottomNeighborCells(cornerCells, 0)
	case direction.Left:
		return cell.GetLeftNeighborCells(cornerCells, 0)
	case direction.Right:
		return cell.GetRightNeighborCells(cornerCells, m.maxCol)
	}
	return []cell.Cell{}
}

func isBlocked(blockingCell cell.Cell, currX int, currY int, objectWidth int, objectHeight int, dir direction.Direction, stepSize int) bool {
	switch dir {
	case direction.Up:
		nextY := currY + objectHeight + stepSize
		cellBottom := blockingCell.Row * square.Width
		return nextY > cellBottom
	case direction.Down:
		nextY := currY - stepSize
		cellTop := blockingCell.Row*square.Width + square.Width
		return nextY < cellTop
	case direction.Left:
		nextX := currX - stepSize
		cellRight := blockingCell.Col*square.Width + square.Width
		return nextX < cellRight
	case direction.Right:
		nextX := currX + objectWidth + stepSize
		cellLeft := blockingCell.Col * square.Width
		return nextX > cellLeft
	}
	return false
}
