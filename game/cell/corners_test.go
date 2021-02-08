package cell

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestGetCornerCells(t *testing.T) {
    testCases := []struct {
        name                string
        bottomLeftX         int
        bottomLeftY         int
        objectWidth         int
        objectHeight        int
        cellWidth           int
        cellHeight          int
        expectedCornerCells CornerCells
    }{
        {
            name:        "Object corners overlap with grid corners",
            bottomLeftX: 4, bottomLeftY: 2, objectWidth: 2, objectHeight: 2,
            cellWidth: 2, cellHeight: 2,
            expectedCornerCells: CornerCells{
                TopLeft:     Cell{Row: 1, Col: 2},
                TopRight:    Cell{Row: 1, Col: 2},
                BottomLeft:  Cell{Row: 1, Col: 2},
                BottomRight: Cell{Row: 1, Col: 2},
            },
        },
        {
            name:        "Object corners on vertical grid lines",
            bottomLeftX: 8, bottomLeftY: 3, objectWidth: 2, objectHeight: 2,
            cellWidth: 2, cellHeight: 2,
            expectedCornerCells: CornerCells{
                TopLeft:     Cell{Row: 2, Col: 4},
                TopRight:    Cell{Row: 2, Col: 4},
                BottomLeft:  Cell{Row: 1, Col: 4},
                BottomRight: Cell{Row: 1, Col: 4},
            },
        },
        {
            name:        "Object corners on horizontal grid lines",
            bottomLeftX: 5, bottomLeftY: 2, objectWidth: 2, objectHeight: 2,
            cellWidth: 2, cellHeight: 2,
            expectedCornerCells: CornerCells{
                TopLeft:     Cell{Row: 1, Col: 2},
                TopRight:    Cell{Row: 1, Col: 3},
                BottomLeft:  Cell{Row: 1, Col: 2},
                BottomRight: Cell{Row: 1, Col: 3},
            },
        },
        {
            name:        "Object corners not on grid lines",
            bottomLeftX: 9, bottomLeftY: 3, objectWidth: 2, objectHeight: 2,
            cellWidth: 2, cellHeight: 2,
            expectedCornerCells: CornerCells{
                TopLeft:     Cell{Row: 2, Col: 4},
                TopRight:    Cell{Row: 2, Col: 5},
                BottomLeft:  Cell{Row: 1, Col: 4},
                BottomRight: Cell{Row: 1, Col: 5},
            },
        },
        {
            name:        "Object width occupy multiple columns and height occupy multiple columns",
            bottomLeftX: 4, bottomLeftY: 2, objectWidth: 4, objectHeight: 6,
            cellWidth: 2, cellHeight: 2,
            expectedCornerCells: CornerCells{
                TopLeft:     Cell{Row: 3, Col: 2},
                TopRight:    Cell{Row: 3, Col: 3},
                BottomLeft:  Cell{Row: 1, Col: 2},
                BottomRight: Cell{Row: 1, Col: 3},
            },
        },
    }

    for _, testCase := range testCases {
        t.Run(testCase.name, func(t *testing.T) {
            cornerCells := GetCornerCells(
                testCase.bottomLeftX, testCase.bottomLeftY, testCase.objectWidth, testCase.objectHeight,
                testCase.cellWidth, testCase.cellHeight,
            )
            assert.Equal(t, testCase.expectedCornerCells, cornerCells)
        })
    }
}
