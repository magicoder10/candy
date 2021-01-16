package game

import (
	"candy/game/player"
	"candy/graphics"
)

const playerStatusBarLeft = 56
const playerStatusBarBottom = 130
const playerStatusBarHeight = 54
const playerStatusBarGap = 4

var rightBarBackpackImageBound = graphics.Bound{
	X:      900,
	Y:      0,
	Width:  252,
	Height: 830,
}

type rightSideBar struct {
	screenX          int
	screenY          int
	playerStatusList *playerStatusList
}

func (r rightSideBar) draw(batch graphics.Batch) {
	batch.DrawSprite(r.screenX, r.screenY, r.screenY, rightBarBackpackImageBound, 1)
	r.playerStatusList.draw(batch)
}

func newRightSideBar(screenX int, screenY int, players []*player.Player) rightSideBar {
	return rightSideBar{
		screenX: screenX,
		screenY: screenY,
		playerStatusList: &playerStatusList{
			screenX: screenX + playerStatusBarLeft,
			screenY: screenY + playerStatusBarBottom,
			z:       screenY,
			players: players,
		},
	}
}

type playerStatusList struct {
	screenX int
	screenY int
	z       int
	players []*player.Player
}

func (p playerStatusList) draw(batch graphics.Batch) {
	for index, ply := range p.players {
		y := p.screenY + index*(playerStatusBarHeight+playerStatusBarGap) + 5
		ply.DrawStand(batch, p.screenX+5, y, p.z, 1)
	}
}
