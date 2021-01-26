package player

import (
	"candy/server/gamestate"
)

type character struct {
	walkCycleOffset walkCycleOffset
}

var blackBoy = character{
	walkCycleOffset: walkCycleOffset{
		x: 0,
		y: spriteColHeight,
	}}
var blackGirl = character{
	walkCycleOffset: walkCycleOffset{
		x: 0,
		y: 0,
	}}
var brownBoy = character{
	walkCycleOffset: walkCycleOffset{
		x: spriteRowWidth,
		y: spriteColHeight,
	}}
var brownGirl = character{
	walkCycleOffset: walkCycleOffset{
		x: spriteRowWidth,
		y: 0,
	}}
var yellowBoy = character{
	walkCycleOffset: walkCycleOffset{
		x: spriteRowWidth * 2,
		y: spriteColHeight,
	}}
var yellowGirl = character{
	walkCycleOffset: walkCycleOffset{
		x: spriteRowWidth * 2,
		y: 0,
	}}
var orangeBoy = character{
	walkCycleOffset: walkCycleOffset{
		x: spriteRowWidth * 3,
		y: spriteColHeight,
	}}
var orangeGirl = character{
	walkCycleOffset: walkCycleOffset{
		x: spriteRowWidth * 3,
		y: 0,
	}}

var characterMap = map[gamestate.Character]character{
	gamestate.BlackBoy:   blackBoy,
	gamestate.BlackGirl:  blackGirl,
	gamestate.BrownBoy:   brownBoy,
	gamestate.BrownGirl:  brownGirl,
	gamestate.YellowBoy:  yellowBoy,
	gamestate.YellowGirl: yellowGirl,
	gamestate.OrangeBoy:  orangeBoy,
	gamestate.OrangeGirl: orangeGirl,
}

func NewCharacter(character gamestate.Character) character {
	val, ok := characterMap[character]
	if !ok {
		return blackBoy
	}
	return val
}
