package gamemap

import (
    "candy/game/candy"
    "candy/game/cell"
    "candy/game/direction"
    "candy/game/square"
)

var _ candy.RangeCutter = (*candyRangeCutter)(nil)

type candyRangeCutter struct {
    maxRow int
    maxCol int
    grid   *[][]square.Square
}

func (c candyRangeCutter) CutRange(start cell.Cell, initialRange int, dir direction.Direction) int {
    for currRange := 1; currRange <= initialRange; currRange++ {
        nc := nextCell(start, currRange, dir)
        if !inGrid(nc, c.maxRow, c.maxCol) {
            return currRange - 1
        }
        sq := (*c.grid)[nc.Row][nc.Col]
        if sq == nil || sq.CanEnter() {
            continue
        }
        if sq.IsBroken() || sq.IsBreakable() {
            return currRange
        } else {
            return currRange - 1
        }
    }
    return initialRange
}

func nextCell(start cell.Cell, offset int, dir direction.Direction) cell.Cell {
    switch dir {
    case direction.Up:
        return cell.Cell{
            Row: start.Row + offset,
            Col: start.Col,
        }
    case direction.Down:
        return cell.Cell{
            Row: start.Row - offset,
            Col: start.Col,
        }
    case direction.Left:
        return cell.Cell{
            Row: start.Row,
            Col: start.Col - offset,
        }
    case direction.Right:
        return cell.Cell{
            Row: start.Row,
            Col: start.Col + offset,
        }
    }
    return start
}
