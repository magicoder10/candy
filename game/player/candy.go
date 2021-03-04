package player

type DropCandyChecker interface {
	CanDropCandy(playerX int, playerY int, playerWidth int, playerHeight int) bool
}
