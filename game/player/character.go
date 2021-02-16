package player

type character struct {
	walkCycleOffset walkCycleOffset
	initialStepSize int
}

var BlackBoy = character{
	walkCycleOffset: walkCycleOffset{
		x: 0,
		y: spriteColHeight,
	},
	initialStepSize: 10,
}

var BlackGirl = character{
	walkCycleOffset: walkCycleOffset{
		x: 0,
		y: 0,
	},
	initialStepSize: 4,
}

var BrownBoy = character{
	walkCycleOffset: walkCycleOffset{
		x: spriteRowWidth,
		y: spriteColHeight,
	},
	initialStepSize: 6,
}

var BrownGirl = character{
	walkCycleOffset: walkCycleOffset{
		x: spriteRowWidth,
		y: 0,
	},
	initialStepSize: 6,
}

var YellowBoy = character{
	walkCycleOffset: walkCycleOffset{
		x: spriteRowWidth * 2,
		y: spriteColHeight,
	},
	initialStepSize: 6,
}

var YellowGirl = character{
	walkCycleOffset: walkCycleOffset{
		x: spriteRowWidth * 2,
		y: 0,
	},
	initialStepSize: 6,
}

var OrangeBoy = character{
	walkCycleOffset: walkCycleOffset{
		x: spriteRowWidth * 3,
		y: spriteColHeight,
	},
	initialStepSize: 6,
}

var OrangeGirl = character{
	walkCycleOffset: walkCycleOffset{
		x: spriteRowWidth * 3,
		y: 0,
	},
	initialStepSize: 6,
}
