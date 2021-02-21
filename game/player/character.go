package player

type character struct {
	walkCycleOffset walkCycleOffset
	initialPower    int
}

var BlackBoy = character{
	walkCycleOffset: walkCycleOffset{
		x: 0,
		y: spriteColHeight,
	},
	initialPower: 1,
}
var BlackGirl = character{
	walkCycleOffset: walkCycleOffset{
		x: 0,
		y: 0,
	},
	initialPower: 1,
}
var BrownBoy = character{
	walkCycleOffset: walkCycleOffset{
		x: spriteRowWidth,
		y: spriteColHeight,
	}}
var BrownGirl = character{
	walkCycleOffset: walkCycleOffset{
		x: spriteRowWidth,
		y: 0,
	},
	initialPower: 1,
}
var YellowBoy = character{
	walkCycleOffset: walkCycleOffset{
		x: spriteRowWidth * 2,
		y: spriteColHeight,
	},
	initialPower: 1,
}
var YellowGirl = character{
	walkCycleOffset: walkCycleOffset{
		x: spriteRowWidth * 2,
		y: 0,
	},
	initialPower: 1,
}
var OrangeBoy = character{
	walkCycleOffset: walkCycleOffset{
		x: spriteRowWidth * 3,
		y: spriteColHeight,
	},
	initialPower: 1,
}
var OrangeGirl = character{
	walkCycleOffset: walkCycleOffset{
		x: spriteRowWidth * 3,
		y: 0,
	},
	initialPower: 1,
}
