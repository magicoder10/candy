package gamemap

import (
	"candy/game/cell"
	"candy/game/player"
	"candy/game/square"
)

var _ player.DropCandyChecker = (*DropCandyChecker)(nil)

type DropCandyChecker struct {
	gameMap *Map
}

func (d DropCandyChecker) CanDropCandy(playerX int, playerY int, playerWidth int, playerHeight int) bool {
	playerCell := d.getPlayerCell(playerX, playerY, playerWidth, playerHeight)
	return (*d.gameMap.grid)[playerCell.Row][playerCell.Col] == nil
}

func (d DropCandyChecker) getPlayerCell(playerX int, playerY int, playerWidth int, playerHeight int) cell.Cell {
	return d.getObjectCell(playerX, playerY, playerWidth, playerHeight)
}

func (d DropCandyChecker) getObjectCell(objectX int, objectY int, objectWidth int, objectHeight int) cell.Cell {
	return cell.GetCellLocatedAt(
		objectX-d.gameMap.GetGridXOffset(), objectY-d.gameMap.GetGridYOffset(), objectWidth, objectHeight,
		square.Width, square.Width,
	)
}

func NewDropCandyChecker(gameMap *Map) DropCandyChecker {
	return DropCandyChecker{
		gameMap: gameMap,
	}
}
