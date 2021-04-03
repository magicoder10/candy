package gamemap

import (
	"candy/game/square"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestDropCandyChecker_CanDropCandy(t *testing.T) {

	gamemap := &Map {
		screenX:3,
		screenY:3,
		batch: nil,
		maxRow: 4,
		maxCol: 4,
		gridXOffset: 3,
		gridYOffset: 3,
		grid: &[][]square.Square{
			{nil, nil, nil, nil},
			{nil, nil, nil, nil},
			{nil, nil, nil, nil},
			{nil, nil, nil, nil},
		},

	}
	dropCandyChecker := NewDropCandyChecker(gamemap)

	testCases := []struct{
		name string
		playerX int
		playerY int
		playerWidth int
		playerHeight int
		threshold	bool
	}{
		{   name: "testcase1",
			playerX: 1,
			playerY: 2,
			playerWidth: 1,
			playerHeight: 1,
			threshold: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T){
			canDropCandyChecker :=
				dropCandyChecker.CanDropCandy(
					testCase.playerX,
					testCase.playerY,
					testCase.playerWidth,
					testCase.playerHeight)
			assert.Equal(t, testCase.threshold, canDropCandyChecker)
		})
	}

}














