package gamemap

import (
	"testing"

	"candy/game/direction"
	"candy/game/square"

	"github.com/stretchr/testify/assert"
)

var _ square.Square = (*obstacle)(nil)

type obstacle struct {
	square.Tile
	canEnter bool
}

func (o obstacle) CanEnter() bool {
	return o.canEnter
}

func Test_moveCheckerCanMove(t *testing.T) {
	defaultMoveChecker := moveChecker{
		gridXOffset: 0,
		gridYOffset: 0,
		maxRow:      3,
		maxCol:      3,
		grid: &[][]square.Square{
			{&obstacle{canEnter: false}, nil, nil, nil},
			{nil, nil, &obstacle{canEnter: false}, nil},
			{&obstacle{canEnter: false}, nil, &obstacle{canEnter: false}, nil},
			{nil, nil, nil, &obstacle{canEnter: false}},
		},
	}

	type testCase struct {
		name            string
		currX           int
		currY           int
		objectWidth     int
		objectHeight    int
		dir             direction.Direction
		stepSize        int
		expectedCanMove bool
	}

	testSuites := []struct {
		name        string
		moveChecker moveChecker
		testCases   []testCase
	}{
		{
			name:        "Border",
			moveChecker: defaultMoveChecker,
			testCases: []testCase{
				{
					name:            "move down at bottom border",
					currX:           1,
					currY:           1,
					objectWidth:     10,
					objectHeight:    10,
					dir:             direction.Down,
					stepSize:        2,
					expectedCanMove: false,
				},
				{
					name:            "move up at bottom border",
					currX:           1,
					currY:           1,
					objectWidth:     10,
					objectHeight:    10,
					dir:             direction.Up,
					stepSize:        2,
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
					expectedCanMove: true,
				},
			},
		},
		{
			name:        "Has obstacle in the direction of movement",
			moveChecker: defaultMoveChecker,
			testCases: []testCase{
				{
					name:            "move left",
					currX:           1 * square.Width,
					currY:           0,
					objectWidth:     square.Width,
					objectHeight:    square.Width,
					dir:             direction.Left,
					stepSize:        square.Width,
					expectedCanMove: false,
				},

				{
					name:            "move down",
					currX:           0,
					currY:           3 * square.Width,
					objectWidth:     square.Width,
					objectHeight:    square.Width,
					dir:             direction.Down,
					stepSize:        square.Width,
					expectedCanMove: false,
				},
				{
					name:            "move right",
					currX:           1 * square.Width,
					currY:           2 * square.Width,
					objectWidth:     square.Width,
					objectHeight:    square.Width,
					dir:             direction.Right,
					stepSize:        square.Width,
					expectedCanMove: false,
				},
				{
					name:            "move up",
					currX:           2 * square.Width,
					currY:           0,
					objectWidth:     10,
					objectHeight:    10,
					dir:             direction.Up,
					stepSize:        square.Width,
					expectedCanMove: false,
				},
			},
		},
		{
			name: "Has nothing in the direction of movement",
			moveChecker: moveChecker{
				gridXOffset: 0,
				gridYOffset: 0,
				maxRow:      3,
				maxCol:      3,
				grid: &[][]square.Square{
					{nil, nil, nil, nil},
					{nil, nil, nil, nil},
					{nil, nil, nil, nil},
					{nil, nil, nil, nil},
				},
			},
			testCases: []testCase{
				{
					name:            "move left",
					currX:           1 * square.Width,
					currY:           0,
					objectWidth:     square.Width,
					objectHeight:    square.Width,
					dir:             direction.Left,
					stepSize:        square.Width,
					expectedCanMove: true,
				},

				{
					name:            "move down",
					currX:           0,
					currY:           3 * square.Width,
					objectWidth:     square.Width,
					objectHeight:    square.Width,
					dir:             direction.Down,
					stepSize:        square.Width,
					expectedCanMove: true,
				},
				{
					name:            "move right",
					currX:           1 * square.Width,
					currY:           2 * square.Width,
					objectWidth:     square.Width,
					objectHeight:    square.Width,
					dir:             direction.Right,
					stepSize:        square.Width,
					expectedCanMove: true,
				},
				{
					name:            "move up",
					currX:           2 * square.Width,
					currY:           0,
					objectWidth:     10,
					objectHeight:    10,
					dir:             direction.Up,
					stepSize:        square.Width,
					expectedCanMove: true,
				},
			},
		},
		{
			name:        "Blockers around the player but not on moving direction",
			moveChecker: defaultMoveChecker,
			testCases: []testCase{
				{
					name:            "move up blocker in left side",
					currX:           1 * square.Width,
					currY:           0,
					objectWidth:     square.Width,
					objectHeight:    square.Width,
					dir:             direction.Up,
					stepSize:        square.Width,
					expectedCanMove: true,
				},

				{
					name:            "move right blocker in down side",
					currX:           0,
					currY:           3 * square.Width,
					objectWidth:     square.Width,
					objectHeight:    square.Width,
					dir:             direction.Right,
					stepSize:        square.Width,
					expectedCanMove: true,
				},
				{
					name:            "move down blocker in left and right side",
					currX:           1 * square.Width,
					currY:           2 * square.Width,
					objectWidth:     square.Width,
					objectHeight:    square.Width,
					dir:             direction.Down,
					stepSize:        square.Width,
					expectedCanMove: true,
				},
				{
					name:            "move left facing blocker",
					currX:           2 * square.Width,
					currY:           0,
					objectWidth:     10,
					objectHeight:    10,
					dir:             direction.Left,
					stepSize:        square.Width,
					expectedCanMove: true,
				},
			},
		},
	}

	for _, testSuite := range testSuites {
		t.Run(testSuite.name, func(t *testing.T) {
			for _, tc := range testSuite.testCases {
				t.Run(tc.name, func(t *testing.T) {
					canMove := testSuite.moveChecker.CanMove(
						tc.currX, tc.currY, tc.objectWidth, tc.objectHeight,
						tc.dir, tc.stepSize,
					)
					assert.Equal(t, tc.expectedCanMove, canMove)
				})
			}
		})
	}
}
