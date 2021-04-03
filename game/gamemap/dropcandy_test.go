package gamemap

import (
	"testing"

	"candy/game/square"

	"github.com/stretchr/testify/assert"
)

func TestDropCandyChecker_CanDropCandy(t *testing.T) {
	testCases := []struct {
		name                 string
		gameMap              *Map
		playerX              int
		playerY              int
		expectedCanDropCandy bool
	}{
		{
			name: "map empty",
			gameMap: &Map{
				batch:  nil,
				maxRow: 4,
				maxCol: 4,
				grid: &[][]square.Square{
					{nil, nil, nil, nil},
					{nil, nil, nil, nil},
					{nil, nil, nil, nil},
					{nil, nil, nil, nil},
				},
			},
			playerX:              1,
			playerY:              2,
			expectedCanDropCandy: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			dropCandyChecker := NewDropCandyChecker(testCase.gameMap)
			gotCanDropCandyCheck := dropCandyChecker.CanDropCandy(
				testCase.playerX,
				testCase.playerY,
				1,
				1)
			assert.Equal(t, testCase.expectedCanDropCandy, gotCanDropCandyCheck)
		})
	}

}
