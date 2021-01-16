package gamemap

import (
	"testing"

	"candy/game/cell"
	"candy/game/direction"
	"candy/game/gameitem"
	"candy/game/square"

	"github.com/stretchr/testify/assert"
)

func TestCandyRangeCutter_CutRangeBySolidTile(t *testing.T) {
	rangeCutter := candyRangeCutter{
		maxRow: 4,
		maxCol: 4,
		grid: &[][]square.Square{
			{nil, nil, nil, nil},
			{square.NewTile('Y', gameitem.Power), nil, nil, nil},
			{nil, nil, nil, nil},
			{nil, nil, nil, nil},
		},
	}

	testCases := []struct {
		name          string
		start         cell.Cell
		initialRange  int
		dir           direction.Direction
		expectedRange int
	}{
		{
			start:         cell.Cell{Row: 0, Col: 0},
			initialRange:  3,
			dir:           direction.Up,
			expectedRange: 1,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			newRange := rangeCutter.CutRange(testCase.start, testCase.initialRange, testCase.dir)
			assert.Equal(t, testCase.expectedRange, newRange)
		})
	}
}

func TestCandyRangeCutter_CutRangeByBrokenTile(t *testing.T) {
	tile := square.NewTile('Y', gameitem.Power)
	rangeCutter := candyRangeCutter{
		maxRow: 4,
		maxCol: 4,
		grid: &[][]square.Square{
			{nil, nil, nil, nil},
			{tile, nil, nil, nil},
			{nil, nil, nil, nil},
			{nil, nil, nil, nil},
		},
	}

	tile.Break()
	newRange := rangeCutter.CutRange(cell.Cell{Row: 0, Col: 0}, 3, direction.Up)
	assert.Equal(t, 1, newRange)
}

func TestCandyRangeCutter_CutRangeByGameItem(t *testing.T) {
	tile := square.NewTile('Y', gameitem.Power)
	rangeCutter := candyRangeCutter{
		maxRow: 4,
		maxCol: 4,
		grid: &[][]square.Square{
			{nil, nil, nil, nil},
			{tile, nil, nil, nil},
			{nil, nil, nil, nil},
			{nil, nil, nil, nil},
		},
	}

	tile.Break()
	tile.UnblockFire()
	newRange := rangeCutter.CutRange(cell.Cell{Row: 0, Col: 0}, 3, direction.Up)
	assert.Equal(t, 3, newRange)
}
