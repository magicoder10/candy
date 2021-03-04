package gamemap

import (
	"candy/game/player"
)

var _ player.DropCandyChecker = (*DropCandyChecker)(nil)

type DropCandyChecker struct {
	gameMap *Map
}

func (d DropCandyChecker) CanDropCandy(playerX int, playerY int, playerWidth int, playerHeight int) bool {
	playerCell := d.gameMap.GetObjectCell(playerX, playerY, playerWidth, playerHeight)
	return (*d.gameMap.grid)[playerCell.Row][playerCell.Col] == nil
}

func NewDropCandyChecker(gameMap *Map) DropCandyChecker {
	return DropCandyChecker{
		gameMap: gameMap,
	}
}
