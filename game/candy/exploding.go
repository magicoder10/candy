package candy

import (
	"candy/pubsub"
	"time"

	"candy/game/cell"
	"candy/game/direction"
	"candy/game/square"
	"candy/graphics"
)

const explodingTimeShort = 250 * time.Millisecond
const explodingTimeMedium = 500 * time.Millisecond
const explodingTimeLong = 750 * time.Millisecond

const animationDelay = 300 * time.Millisecond

type explodeDirection struct {
	cell      cell.Cell
	edge      graphics.Bound
	end       graphics.Bound
	direction direction.Direction
}

var directions = []explodeDirection{
	{
		cell:      cell.Cell{Row: 1, Col: 0},
		edge:      explosionVerticalEdge,
		end:       explosionTopEnd,
		direction: direction.Up,
	},
	{
		cell:      cell.Cell{Row: -1, Col: 0},
		edge:      explosionVerticalEdge,
		end:       explosionBottomEnd,
		direction: direction.Down,
	},
	{
		cell:      cell.Cell{Row: 0, Col: 1},
		edge:      explosionHorizontalEdge,
		end:       explosionRightEnd,
		direction: direction.Right,
	},
	{
		cell:      cell.Cell{Row: 0, Col: -1},
		edge:      explosionHorizontalEdge,
		end:       explosionLeftEnd,
		direction: direction.Left,
	},
}
var _ state = (*explodingState)(nil)

type explodingState struct {
	sharedState
	explodingTimeNano int64
	hitRange          int
}

func (e explodingState) exploding() bool {
	return true
}

func (e explodingState) cellsHit() []cell.Cell {
	cells := make([]cell.Cell, 0)
	e.hitRanges(func(dir explodeDirection, currRange int) {
		nextRow := e.center.Row + currRange*dir.cell.Row
		nextCol := e.center.Col + currRange*dir.cell.Col
		cells = append(cells, cell.Cell{Row: nextRow, Col: nextCol})
	})
	return cells
}

func (e *explodingState) update(timeElapsed time.Duration) state {
	e.remainingTime -= timeElapsed
	if e.remainingTime <= 0 {
		return &explodedState{}
	}
	e.lag += timeElapsed.Nanoseconds()

	rangeIncreaseDuration := e.explodingTimeNano / int64(e.powerLevel)
	for e.hitRange < e.powerLevel && e.lag > rangeIncreaseDuration {
		e.hitRange += 1
		e.lag -= rangeIncreaseDuration
	}
	return e
}

func (e explodingState) draw(batch graphics.Batch, x int, y int, z int) {
	e.drawEnds(batch, x, y, z)
	e.drawEdges(batch, x, y, z)
	batch.DrawSprite(x, y, z, explosionCenter, 1)
}

func (e explodingState) drawEdges(batch graphics.Batch, x int, y int, z int) {
	e.hitRanges(func(dir explodeDirection, currRange int) {
		shift := currRange * square.Width
		nextX := x + dir.cell.Col*shift
		nextY := y + dir.cell.Row*shift
		batch.DrawSprite(nextX, nextY, z+nextY, dir.edge, 1)
	})
}

func (e explodingState) drawEnds(batch graphics.Batch, x int, y int, z int) {
	if e.hitRange < 1 {
		return
	}

	for _, dir := range directions {
		hitRange := e.rangeCutter.CutRange(e.center, e.hitRange, dir.direction)
		if hitRange < 1 {
			continue
		}
		shift := hitRange * square.Width
		nextX := x + dir.cell.Col*shift
		nextY := y + dir.cell.Row*shift
		batch.DrawSprite(nextX, nextY, nextY, dir.end, 1)
	}
}

func (e explodingState) hitRanges(processRange func(dir explodeDirection, currRange int)) {
	for _, dir := range directions {
		hitRange := e.rangeCutter.CutRange(e.center, e.hitRange, dir.direction)
		for currRange := 1; currRange <= hitRange; currRange++ {
			processRange(dir, currRange)
		}
	}
}

func newExplodingState(sharedState sharedState) *explodingState {
	defer func() {
		sharedState.pubSub.Publish(pubsub.OnCandyStartExploding, sharedState.droppedBy)
	}()
	explodingTime := getExplodingTime(sharedState.powerLevel)
	sharedState.remainingTime = explodingTime + animationDelay
	return &explodingState{
		sharedState:       sharedState,
		hitRange:          0,
		explodingTimeNano: explodingTime.Nanoseconds(),
	}
}

func getExplodingTime(powerLevel int) time.Duration {
	if powerLevel <= 3 {
		return explodingTimeShort
	} else if powerLevel <= 6 {
		return explodingTimeMedium
	} else {
		return explodingTimeLong
	}
}
