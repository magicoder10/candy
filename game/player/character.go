package player

type character struct {
    walkCycleOffset walkCycleOffset
}

var BlackBoy = character{
    walkCycleOffset: walkCycleOffset{
        x: 0,
        y: spriteColHeight,
    }}
var BlackGirl = character{
    walkCycleOffset: walkCycleOffset{
        x: 0,
        y: 0,
    }}
var BrownBoy = character{
    walkCycleOffset: walkCycleOffset{
        x: spriteRowWidth,
        y: spriteColHeight,
    }}
var BrownGirl = character{
    walkCycleOffset: walkCycleOffset{
        x: spriteRowWidth,
        y: 0,
    }}
var YellowBoy = character{
    walkCycleOffset: walkCycleOffset{
        x: spriteRowWidth * 2,
        y: spriteColHeight,
    }}
var YellowGirl = character{
    walkCycleOffset: walkCycleOffset{
        x: spriteRowWidth * 2,
        y: 0,
    }}
var OrangeBoy = character{
    walkCycleOffset: walkCycleOffset{
        x: spriteRowWidth * 3,
        y: spriteColHeight,
    }}
var OrangeGirl = character{
    walkCycleOffset: walkCycleOffset{
        x: spriteRowWidth * 3,
        y: 0,
    }}
