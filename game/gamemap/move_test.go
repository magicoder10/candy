package gamemap

import (
	"candy/game/direction"
	"candy/game/square"
	"candy/graphics"
	"github.com/stretchr/testify/assert"
	"testing"
)

var _ square.Square = (*obstacle)(nil)

type obstacle struct {
	canEnter bool
}

func (o obstacle) Draw(batch graphics.Batch, x int, y int) {
	return
}

func (o obstacle) IsBreakable() bool {
	return false
}

func (o obstacle) CanEnter() bool {
	return o.canEnter
}

func Test_moveCheckerCanMove(t *testing.T) {
	checker := moveChecker{
		gridXOffset: 0,
		gridYOffset: 0,
		maxRow:      2,
		maxCol:      3,
		grid: &[][]square.Square{

			{
				&obstacle{canEnter: false},
				&obstacle{canEnter: false},
				&obstacle{canEnter: false},
				&obstacle{canEnter: false},
			},
			{
				&obstacle{canEnter: false},
				&obstacle{canEnter: false},
				&obstacle{canEnter: false},
				&obstacle{canEnter: false},
			},
			{
				&obstacle{canEnter: false},
				nil,
				&obstacle{canEnter: false},
				&obstacle{canEnter: false},
			},
		},
		tiles: &[]*square.Tile{},
	}

	type testCase struct {
		name            string
		currX           int
		currY           int
		objectWidth     int
		objectHeight    int
		dir             direction.Direction
		stepSize        int
		moveChecker     moveChecker
		expectedCanMove bool
	}

	testSuites := []struct {
		name      string
		testCases []testCase
	}{
		{
			name: "Border",
			testCases: []testCase{
				{
					name:            "move down at bottom border",
					currX:           1,
					currY:           1,
					objectWidth:     10,
					objectHeight:    10,
					dir:             direction.Down,
					stepSize:        2,
					moveChecker:     checker,
					expectedCanMove: false,
				},
				{
					name:            "move down at bottom border",
					currX:           1,
					currY:           1,
					objectWidth:     10,
					objectHeight:    10,
					dir:             direction.Up,
					stepSize:        2,
					moveChecker:     checker,
					expectedCanMove: true,
				},
				{
					name:            "move left at left border",
					currX:           1,
					currY:           1,
					objectWidth:     10,
					objectHeight:    10,
					dir:             direction.Left,
					stepSize:        2,
					moveChecker:     checker,
					expectedCanMove: false,
				},
				{
					name:            "move right at left border",
					currX:           1,
					currY:           1,
					objectWidth:     10,
					objectHeight:    10,
					dir:             direction.Right,
					stepSize:        2,
					moveChecker:     checker,
					expectedCanMove: true,
				},
			},
		},
	}

	for _, testSuite := range testSuites {
		t.Run(testSuite.name, func(t *testing.T) {
			for _, tc := range testSuite.testCases {
				t.Run(tc.name, func(t *testing.T) {
					canMove := tc.moveChecker.CanMove(
						tc.currX, tc.currY, tc.objectWidth, tc.objectHeight,
						tc.dir, tc.stepSize,
					)
					assert.Equal(t, tc.expectedCanMove, canMove)
				})
			}
		})
	}
}
