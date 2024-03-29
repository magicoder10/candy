package cell

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const maxRow = 11
const maxCol = 14

func TestGetLeftNeighborCells(t *testing.T) {
	type testCase struct {
		name          string
		cornerCells   CornerCells
		minCol        int
		expectedCells []Cell
	}
	testSuites := []struct {
		name      string
		testCases []testCase
	}{
		{
			name: "Within Border",
			testCases: []testCase{
				{
					name: "Object corners overlap with grid corners",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 1, Col: 1},
						TopRight:    Cell{Row: 1, Col: 1},
						BottomLeft:  Cell{Row: 1, Col: 1},
						BottomRight: Cell{Row: 1, Col: 1},
					},
					minCol: 0,
					expectedCells: []Cell{
						{
							Row: 1,
							Col: 0,
						},
					},
				},
				{
					name: "Object occupies two horizontal cells",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 1, Col: 1},
						TopRight:    Cell{Row: 1, Col: 2},
						BottomLeft:  Cell{Row: 1, Col: 1},
						BottomRight: Cell{Row: 1, Col: 2},
					},
					minCol: 0,
					expectedCells: []Cell{
						{
							Row: 1,
							Col: 0,
						},
					},
				},
				{
					name: "Object occupies two vertical cells",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 2, Col: 1},
						TopRight:    Cell{Row: 2, Col: 1},
						BottomLeft:  Cell{Row: 1, Col: 1},
						BottomRight: Cell{Row: 1, Col: 1},
					},
					minCol: 0,
					expectedCells: []Cell{
						{
							Row: 1,
							Col: 0,
						},
						{
							Row: 2,
							Col: 0,
						},
					},
				},
				{
					name: "Object occupies four cells",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 2, Col: 1},
						TopRight:    Cell{Row: 2, Col: 2},
						BottomLeft:  Cell{Row: 1, Col: 1},
						BottomRight: Cell{Row: 1, Col: 2},
					},
					minCol: 0,
					expectedCells: []Cell{
						{
							Row: 1,
							Col: 0,
						},
						{
							Row: 2,
							Col: 0,
						},
					},
				},
				{
					name: "Object is big and occupies 6 cells",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 2, Col: 1},
						TopRight:    Cell{Row: 2, Col: 3},
						BottomLeft:  Cell{Row: 1, Col: 1},
						BottomRight: Cell{Row: 1, Col: 3},
					},
					minCol: 0,
					expectedCells: []Cell{
						{
							Row: 1,
							Col: 0,
						},
						{
							Row: 2,
							Col: 0,
						},
					},
				},
			},
		},
		{
			name: "At Border",
			testCases: []testCase{
				{
					name: "left bottom",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 0, Col: 0},
						TopRight:    Cell{Row: 0, Col: 0},
						BottomLeft:  Cell{Row: 0, Col: 0},
						BottomRight: Cell{Row: 0, Col: 0},
					},
					minCol:        0,
					expectedCells: []Cell{},
				},
				{
					name: "bottom occupies one cell",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 0, Col: 1},
						TopRight:    Cell{Row: 0, Col: 1},
						BottomLeft:  Cell{Row: 0, Col: 1},
						BottomRight: Cell{Row: 0, Col: 1},
					},
					minCol: 0,
					expectedCells: []Cell{
						{
							Row: 0,
							Col: 0,
						},
					},
				},
				{
					name: "bottom occupies two cells",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 0, Col: 1},
						TopRight:    Cell{Row: 0, Col: 2},
						BottomLeft:  Cell{Row: 0, Col: 1},
						BottomRight: Cell{Row: 0, Col: 2},
					},
					minCol: 0,
					expectedCells: []Cell{
						{
							Row: 0,
							Col: 0,
						},
					},
				},
				{
					name: "right bottom",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 0, Col: maxCol},
						TopRight:    Cell{Row: 0, Col: maxCol},
						BottomLeft:  Cell{Row: 0, Col: maxCol},
						BottomRight: Cell{Row: 0, Col: maxCol},
					},
					minCol: 0,
					expectedCells: []Cell{
						{
							Row: 0,
							Col: maxCol - 1,
						},
					},
				},
				{
					name: "right occupies one cell",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 1, Col: maxCol},
						TopRight:    Cell{Row: 1, Col: maxCol},
						BottomLeft:  Cell{Row: 1, Col: maxCol},
						BottomRight: Cell{Row: 1, Col: maxCol},
					},
					minCol: 0,
					expectedCells: []Cell{
						{
							Row: 1,
							Col: maxCol - 1,
						},
					},
				},
				{
					name: "right occupies two cells",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 3, Col: maxCol},
						TopRight:    Cell{Row: 3, Col: maxCol},
						BottomLeft:  Cell{Row: 2, Col: maxCol},
						BottomRight: Cell{Row: 2, Col: maxCol},
					},
					minCol: 0,
					expectedCells: []Cell{
						{
							Row: 2,
							Col: maxCol - 1,
						},
						{
							Row: 3,
							Col: maxCol - 1,
						},
					},
				},
				{
					name: "right top",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: maxRow, Col: maxCol},
						TopRight:    Cell{Row: maxRow, Col: maxCol},
						BottomLeft:  Cell{Row: maxRow, Col: maxCol},
						BottomRight: Cell{Row: maxRow, Col: maxCol},
					},
					minCol: 0,
					expectedCells: []Cell{
						{
							Row: maxRow,
							Col: maxCol - 1,
						},
					},
				},
				{
					name: "top occupies one cell",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: maxRow, Col: 3},
						TopRight:    Cell{Row: maxRow, Col: 3},
						BottomLeft:  Cell{Row: maxRow, Col: 3},
						BottomRight: Cell{Row: maxRow, Col: 3},
					},
					minCol: 0,
					expectedCells: []Cell{
						{
							Row: maxRow,
							Col: 2,
						},
					},
				},
				{
					name: "top occupies two cells",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: maxRow, Col: 1},
						TopRight:    Cell{Row: maxRow, Col: 2},
						BottomLeft:  Cell{Row: maxRow, Col: 1},
						BottomRight: Cell{Row: maxRow, Col: 2},
					},
					minCol: 0,
					expectedCells: []Cell{
						{
							Row: maxRow,
							Col: 0,
						},
					},
				},
				{
					name: "left top",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: maxRow, Col: 0},
						TopRight:    Cell{Row: maxRow, Col: 0},
						BottomLeft:  Cell{Row: maxRow, Col: 0},
						BottomRight: Cell{Row: maxRow, Col: 0},
					},
					minCol:        0,
					expectedCells: []Cell{},
				},
				{
					name: "left occupies one cell",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 3, Col: 0},
						TopRight:    Cell{Row: 3, Col: 0},
						BottomLeft:  Cell{Row: 3, Col: 0},
						BottomRight: Cell{Row: 3, Col: 0},
					},
					minCol:        0,
					expectedCells: []Cell{},
				},
				{
					name: "left occupies two cells",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 3, Col: 0},
						TopRight:    Cell{Row: 3, Col: 0},
						BottomLeft:  Cell{Row: 2, Col: 0},
						BottomRight: Cell{Row: 2, Col: 0},
					},
					minCol:        0,
					expectedCells: []Cell{},
				},
			},
		},
	}
	for _, testSuite := range testSuites {
		t.Run(testSuite.name, func(t *testing.T) {
			for _, testCase := range testSuite.testCases {
				t.Run(testCase.name, func(t *testing.T) {
					cells := GetLeftNeighborCells(
						testCase.cornerCells, testCase.minCol,
					)
					assert.Equal(t, testCase.expectedCells, cells)
				})
			}
		})
	}
}

func TestGetRightNeighborCells(t *testing.T) {
	type testCase struct {
		name          string
		cornerCells   CornerCells
		maxCol        int
		expectedCells []Cell
	}
	testSuites := []struct {
		name      string
		testCases []testCase
	}{
		{
			name: "Within Border",
			testCases: []testCase{
				{
					name: "Object corners overlap with grid corners",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 1, Col: 1},
						TopRight:    Cell{Row: 1, Col: 1},
						BottomLeft:  Cell{Row: 1, Col: 1},
						BottomRight: Cell{Row: 1, Col: 1},
					},
					maxCol: maxCol,
					expectedCells: []Cell{
						{
							Row: 1,
							Col: 2,
						},
					},
				},
				{
					name: "Object occupies two horizontal cells",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 1, Col: 1},
						TopRight:    Cell{Row: 1, Col: 2},
						BottomLeft:  Cell{Row: 1, Col: 1},
						BottomRight: Cell{Row: 1, Col: 2},
					},
					maxCol: maxCol,
					expectedCells: []Cell{
						{
							Row: 1,
							Col: 3,
						},
					},
				},
				{
					name: "Object occupies two vertical cells",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 2, Col: 1},
						TopRight:    Cell{Row: 2, Col: 1},
						BottomLeft:  Cell{Row: 1, Col: 1},
						BottomRight: Cell{Row: 1, Col: 1},
					},
					maxCol: maxCol,
					expectedCells: []Cell{
						{
							Row: 1,
							Col: 2,
						},
						{
							Row: 2,
							Col: 2,
						},
					},
				},
				{
					name: "Object occupies four cells",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 2, Col: 1},
						TopRight:    Cell{Row: 2, Col: 2},
						BottomLeft:  Cell{Row: 1, Col: 1},
						BottomRight: Cell{Row: 1, Col: 2},
					},
					maxCol: maxCol,
					expectedCells: []Cell{
						{
							Row: 1,
							Col: 3,
						},
						{
							Row: 2,
							Col: 3,
						},
					},
				},
				{
					name: "Object is big and occupies 6 cells",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 2, Col: 1},
						TopRight:    Cell{Row: 2, Col: 3},
						BottomLeft:  Cell{Row: 1, Col: 1},
						BottomRight: Cell{Row: 1, Col: 3},
					},
					maxCol: maxCol,
					expectedCells: []Cell{
						{
							Row: 1,
							Col: 4,
						},
						{
							Row: 2,
							Col: 4,
						},
					},
				},
			},
		},
		{
			name: "At Border",
			testCases: []testCase{
				{
					name: "left bottom",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 0, Col: 0},
						TopRight:    Cell{Row: 0, Col: 0},
						BottomLeft:  Cell{Row: 0, Col: 0},
						BottomRight: Cell{Row: 0, Col: 0},
					},
					maxCol: maxCol,
					expectedCells: []Cell{
						{
							Row: 0,
							Col: 1,
						},
					},
				},
				{
					name: "bottom occupies one cell",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 0, Col: 1},
						TopRight:    Cell{Row: 0, Col: 1},
						BottomLeft:  Cell{Row: 0, Col: 1},
						BottomRight: Cell{Row: 0, Col: 1},
					},
					maxCol: maxCol,
					expectedCells: []Cell{
						{
							Row: 0,
							Col: 2,
						},
					},
				},
				{
					name: "bottom occupies two cells",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 0, Col: 1},
						TopRight:    Cell{Row: 0, Col: 2},
						BottomLeft:  Cell{Row: 0, Col: 1},
						BottomRight: Cell{Row: 0, Col: 2},
					},
					maxCol: maxCol,
					expectedCells: []Cell{
						{
							Row: 0,
							Col: 3,
						},
					},
				},
				{
					name: "right bottom",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 0, Col: maxCol},
						TopRight:    Cell{Row: 0, Col: maxCol},
						BottomLeft:  Cell{Row: 0, Col: maxCol},
						BottomRight: Cell{Row: 0, Col: maxCol},
					},
					maxCol:        maxCol,
					expectedCells: []Cell{},
				},
				{
					name: "right occupies one cell",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 1, Col: maxCol},
						TopRight:    Cell{Row: 1, Col: maxCol},
						BottomLeft:  Cell{Row: 1, Col: maxCol},
						BottomRight: Cell{Row: 1, Col: maxCol},
					},
					maxCol:        maxCol,
					expectedCells: []Cell{},
				},
				{
					name: "right occupies two cells",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 3, Col: maxCol},
						TopRight:    Cell{Row: 3, Col: maxCol},
						BottomLeft:  Cell{Row: 2, Col: maxCol},
						BottomRight: Cell{Row: 2, Col: maxCol},
					},
					maxCol:        maxCol,
					expectedCells: []Cell{},
				},
				{
					name: "right top",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: maxRow, Col: maxCol},
						TopRight:    Cell{Row: maxRow, Col: maxCol},
						BottomLeft:  Cell{Row: maxRow, Col: maxCol},
						BottomRight: Cell{Row: maxRow, Col: maxCol},
					},
					maxCol:        maxCol,
					expectedCells: []Cell{},
				},
				{
					name: "top occupies one cell",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: maxRow, Col: 3},
						TopRight:    Cell{Row: maxRow, Col: 3},
						BottomLeft:  Cell{Row: maxRow, Col: 3},
						BottomRight: Cell{Row: maxRow, Col: 3},
					},
					maxCol: maxCol,
					expectedCells: []Cell{
						{
							Row: maxRow,
							Col: 4,
						},
					},
				},
				{
					name: "top occupies two cells",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: maxRow, Col: 1},
						TopRight:    Cell{Row: maxRow, Col: 2},
						BottomLeft:  Cell{Row: maxRow, Col: 1},
						BottomRight: Cell{Row: maxRow, Col: 2},
					},
					maxCol: maxCol,
					expectedCells: []Cell{
						{
							Row: maxRow,
							Col: 3,
						},
					},
				},
				{
					name: "left top",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: maxRow, Col: 0},
						TopRight:    Cell{Row: maxRow, Col: 0},
						BottomLeft:  Cell{Row: maxRow, Col: 0},
						BottomRight: Cell{Row: maxRow, Col: 0},
					},
					maxCol: maxCol,
					expectedCells: []Cell{
						{
							Row: maxRow,
							Col: 1,
						},
					},
				},
				{
					name: "left occupies one cell",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 3, Col: 0},
						TopRight:    Cell{Row: 3, Col: 0},
						BottomLeft:  Cell{Row: 3, Col: 0},
						BottomRight: Cell{Row: 3, Col: 0},
					},
					maxCol: maxCol,
					expectedCells: []Cell{
						{
							Row: 3,
							Col: 1,
						},
					},
				},
				{
					name: "left occupies two cells",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 3, Col: 0},
						TopRight:    Cell{Row: 3, Col: 0},
						BottomLeft:  Cell{Row: 2, Col: 0},
						BottomRight: Cell{Row: 2, Col: 0},
					},
					maxCol: maxCol,
					expectedCells: []Cell{
						{
							Row: 2,
							Col: 1,
						},
						{
							Row: 3,
							Col: 1,
						},
					},
				},
			},
		},
	}

	for _, testSuite := range testSuites {
		t.Run(testSuite.name, func(t *testing.T) {
			for _, testCase := range testSuite.testCases {
				t.Run(testCase.name, func(t *testing.T) {
					cells := GetRightNeighborCells(
						testCase.cornerCells, testCase.maxCol,
					)
					assert.Equal(t, testCase.expectedCells, cells)
				})
			}
		})
	}
}

func TestGetTopNeighborCellsCells(t *testing.T) {
	type testCase struct {
		name          string
		cornerCells   CornerCells
		maxRow        int
		expectedCells []Cell
	}
	testSuites := []struct {
		name      string
		testCases []testCase
	}{
		{
			name: "Within Border",
			testCases: []testCase{
				{
					name: "Object corners overlap with grid corners",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 1, Col: 1},
						TopRight:    Cell{Row: 1, Col: 1},
						BottomLeft:  Cell{Row: 1, Col: 1},
						BottomRight: Cell{Row: 1, Col: 1},
					},
					maxRow: maxRow,
					expectedCells: []Cell{
						{
							Row: 2,
							Col: 1,
						},
					},
				},
				{
					name: "Object occupies two horizontal cells",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 1, Col: 1},
						TopRight:    Cell{Row: 1, Col: 2},
						BottomLeft:  Cell{Row: 1, Col: 1},
						BottomRight: Cell{Row: 1, Col: 2},
					},
					maxRow: maxRow,
					expectedCells: []Cell{
						{
							Row: 2,
							Col: 1,
						},
						{
							Row: 2,
							Col: 2,
						},
					},
				},
				{
					name: "Object occupies two vertical cells",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 2, Col: 1},
						TopRight:    Cell{Row: 2, Col: 1},
						BottomLeft:  Cell{Row: 1, Col: 1},
						BottomRight: Cell{Row: 1, Col: 1},
					},
					maxRow: maxRow,
					expectedCells: []Cell{
						{
							Row: 3,
							Col: 1,
						},
					},
				},
				{
					name: "Object occupies four cells",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 2, Col: 1},
						TopRight:    Cell{Row: 2, Col: 2},
						BottomLeft:  Cell{Row: 1, Col: 1},
						BottomRight: Cell{Row: 1, Col: 2},
					},
					maxRow: maxRow,
					expectedCells: []Cell{
						{
							Row: 3,
							Col: 1,
						},
						{
							Row: 3,
							Col: 2,
						},
					},
				},
				{
					name: "Object is big and occupies 6 cells",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 2, Col: 1},
						TopRight:    Cell{Row: 2, Col: 3},
						BottomLeft:  Cell{Row: 1, Col: 1},
						BottomRight: Cell{Row: 1, Col: 3},
					},
					maxRow: maxRow,
					expectedCells: []Cell{
						{
							Row: 3,
							Col: 1,
						},
						{
							Row: 3,
							Col: 2,
						},
						{
							Row: 3,
							Col: 3,
						},
					},
				},
			},
		},
		{
			name: "At Border",
			testCases: []testCase{
				{
					name: "left bottom",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 0, Col: 0},
						TopRight:    Cell{Row: 0, Col: 0},
						BottomLeft:  Cell{Row: 0, Col: 0},
						BottomRight: Cell{Row: 0, Col: 0},
					},
					maxRow: maxRow,
					expectedCells: []Cell{
						{
							Row: 1,
							Col: 0,
						},
					},
				},
				{
					name: "bottom occupies one cell",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 0, Col: 1},
						TopRight:    Cell{Row: 0, Col: 1},
						BottomLeft:  Cell{Row: 0, Col: 1},
						BottomRight: Cell{Row: 0, Col: 1},
					},
					maxRow: maxRow,
					expectedCells: []Cell{
						{
							Row: 1,
							Col: 1,
						},
					},
				},
				{
					name: "bottom occupies two cells",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 0, Col: 1},
						TopRight:    Cell{Row: 0, Col: 2},
						BottomLeft:  Cell{Row: 0, Col: 1},
						BottomRight: Cell{Row: 0, Col: 2},
					},
					maxRow: maxRow,
					expectedCells: []Cell{
						{
							Row: 1,
							Col: 1,
						},
						{
							Row: 1,
							Col: 2,
						},
					},
				},
				{
					name: "right bottom",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 0, Col: maxCol},
						TopRight:    Cell{Row: 0, Col: maxCol},
						BottomLeft:  Cell{Row: 0, Col: maxCol},
						BottomRight: Cell{Row: 0, Col: maxCol},
					},
					maxRow: maxRow,
					expectedCells: []Cell{
						{
							Row: 1,
							Col: maxCol,
						},
					},
				},
				{
					name: "right occupies one cell",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 1, Col: maxCol},
						TopRight:    Cell{Row: 1, Col: maxCol},
						BottomLeft:  Cell{Row: 1, Col: maxCol},
						BottomRight: Cell{Row: 1, Col: maxCol},
					},
					maxRow: maxRow,
					expectedCells: []Cell{
						{
							Row: 2,
							Col: maxCol,
						},
					},
				},
				{
					name: "right occupies two cells",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 3, Col: maxCol},
						TopRight:    Cell{Row: 3, Col: maxCol},
						BottomLeft:  Cell{Row: 2, Col: maxCol},
						BottomRight: Cell{Row: 2, Col: maxCol},
					},
					maxRow: maxRow,
					expectedCells: []Cell{
						{
							Row: 4,
							Col: maxCol,
						},
					},
				},
				{
					name: "right top",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: maxRow, Col: maxCol},
						TopRight:    Cell{Row: maxRow, Col: maxCol},
						BottomLeft:  Cell{Row: maxRow, Col: maxCol},
						BottomRight: Cell{Row: maxRow, Col: maxCol},
					},
					maxRow:        maxRow,
					expectedCells: []Cell{},
				},
				{
					name: "top occupies one cell",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: maxRow, Col: 3},
						TopRight:    Cell{Row: maxRow, Col: 3},
						BottomLeft:  Cell{Row: maxRow, Col: 3},
						BottomRight: Cell{Row: maxRow, Col: 3},
					},
					maxRow:        maxRow,
					expectedCells: []Cell{},
				},
				{
					name: "top occupies two cells",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: maxRow, Col: 1},
						TopRight:    Cell{Row: maxRow, Col: 2},
						BottomLeft:  Cell{Row: maxRow, Col: 1},
						BottomRight: Cell{Row: maxRow, Col: 2},
					},
					maxRow:        maxRow,
					expectedCells: []Cell{},
				},
				{
					name: "left top",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: maxRow, Col: 0},
						TopRight:    Cell{Row: maxRow, Col: 0},
						BottomLeft:  Cell{Row: maxRow, Col: 0},
						BottomRight: Cell{Row: maxRow, Col: 0},
					},
					maxRow:        maxRow,
					expectedCells: []Cell{},
				},
				{
					name: "left occupies one cell",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 3, Col: 0},
						TopRight:    Cell{Row: 3, Col: 0},
						BottomLeft:  Cell{Row: 3, Col: 0},
						BottomRight: Cell{Row: 3, Col: 0},
					},
					maxRow: maxRow,
					expectedCells: []Cell{
						{
							Row: 4,
							Col: 0,
						},
					},
				},
				{
					name: "left occupies two cells",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 3, Col: 0},
						TopRight:    Cell{Row: 3, Col: 0},
						BottomLeft:  Cell{Row: 2, Col: 0},
						BottomRight: Cell{Row: 2, Col: 0},
					},
					maxRow: maxRow,
					expectedCells: []Cell{
						{
							Row: 4,
							Col: 0,
						},
					},
				},
			},
		},
	}

	for _, testSuite := range testSuites {
		t.Run(testSuite.name, func(t *testing.T) {
			for _, testCase := range testSuite.testCases {
				t.Run(testCase.name, func(t *testing.T) {
					cells := GetTopNeighborCells(
						testCase.cornerCells, testCase.maxRow,
					)
					assert.Equal(t, testCase.expectedCells, cells)
				})
			}
		})
	}
}

func TestGetBottomNeighborCellsCells(t *testing.T) {
	type testCase struct {
		name          string
		cornerCells   CornerCells
		minRow        int
		expectedCells []Cell
	}
	testSuites := []struct {
		name      string
		testCases []testCase
	}{
		{
			name: "Within Border",
			testCases: []testCase{
				{
					name: "Object corners overlap with grid corners",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 1, Col: 1},
						TopRight:    Cell{Row: 1, Col: 1},
						BottomLeft:  Cell{Row: 1, Col: 1},
						BottomRight: Cell{Row: 1, Col: 1},
					},
					minRow: 0,
					expectedCells: []Cell{
						{
							Row: 0,
							Col: 1,
						},
					},
				},
				{
					name: "Object occupies two horizontal cells",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 1, Col: 1},
						TopRight:    Cell{Row: 1, Col: 2},
						BottomLeft:  Cell{Row: 1, Col: 1},
						BottomRight: Cell{Row: 1, Col: 2},
					},
					minRow: 0,
					expectedCells: []Cell{
						{
							Row: 0,
							Col: 1,
						},
						{
							Row: 0,
							Col: 2,
						},
					},
				},
				{
					name: "Object occupies two vertical cells",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 2, Col: 1},
						TopRight:    Cell{Row: 2, Col: 1},
						BottomLeft:  Cell{Row: 1, Col: 1},
						BottomRight: Cell{Row: 1, Col: 1},
					},
					minRow: 0,
					expectedCells: []Cell{
						{
							Row: 0,
							Col: 1,
						},
					},
				},
				{
					name: "Object occupies four cells",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 2, Col: 1},
						TopRight:    Cell{Row: 2, Col: 2},
						BottomLeft:  Cell{Row: 1, Col: 1},
						BottomRight: Cell{Row: 1, Col: 2},
					},
					minRow: 0,
					expectedCells: []Cell{
						{
							Row: 0,
							Col: 1,
						},
						{
							Row: 0,
							Col: 2,
						},
					},
				},
				{
					name: "Object is big and occupies 6 cells",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 2, Col: 1},
						TopRight:    Cell{Row: 2, Col: 3},
						BottomLeft:  Cell{Row: 1, Col: 1},
						BottomRight: Cell{Row: 1, Col: 3},
					},
					minRow: 0,
					expectedCells: []Cell{
						{
							Row: 0,
							Col: 1,
						},
						{
							Row: 0,
							Col: 2,
						},
						{
							Row: 0,
							Col: 3,
						},
					},
				},
			},
		},
		{
			name: "At Border",
			testCases: []testCase{
				{
					name: "left bottom",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 0, Col: 0},
						TopRight:    Cell{Row: 0, Col: 0},
						BottomLeft:  Cell{Row: 0, Col: 0},
						BottomRight: Cell{Row: 0, Col: 0},
					},
					minRow:        0,
					expectedCells: []Cell{},
				},
				{
					name: "bottom occupies one cell",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 0, Col: 1},
						TopRight:    Cell{Row: 0, Col: 1},
						BottomLeft:  Cell{Row: 0, Col: 1},
						BottomRight: Cell{Row: 0, Col: 1},
					},
					minRow:        0,
					expectedCells: []Cell{},
				},
				{
					name: "bottom occupies two cells",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 0, Col: 1},
						TopRight:    Cell{Row: 0, Col: 2},
						BottomLeft:  Cell{Row: 0, Col: 1},
						BottomRight: Cell{Row: 0, Col: 2},
					},
					minRow:        0,
					expectedCells: []Cell{},
				},
				{
					name: "right bottom",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 0, Col: maxCol},
						TopRight:    Cell{Row: 0, Col: maxCol},
						BottomLeft:  Cell{Row: 0, Col: maxCol},
						BottomRight: Cell{Row: 0, Col: maxCol},
					},
					minRow:        0,
					expectedCells: []Cell{},
				},
				{
					name: "right occupies one cell",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 1, Col: maxCol},
						TopRight:    Cell{Row: 1, Col: maxCol},
						BottomLeft:  Cell{Row: 1, Col: maxCol},
						BottomRight: Cell{Row: 1, Col: maxCol},
					},
					minRow: 0,
					expectedCells: []Cell{
						{
							Row: 0,
							Col: maxCol,
						},
					},
				},
				{
					name: "right occupies two cells",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 3, Col: maxCol},
						TopRight:    Cell{Row: 3, Col: maxCol},
						BottomLeft:  Cell{Row: 2, Col: maxCol},
						BottomRight: Cell{Row: 2, Col: maxCol},
					},
					minRow: 0,
					expectedCells: []Cell{
						{
							Row: 1,
							Col: maxCol,
						},
					},
				},
				{
					name: "right top",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: maxRow, Col: maxCol},
						TopRight:    Cell{Row: maxRow, Col: maxCol},
						BottomLeft:  Cell{Row: maxRow, Col: maxCol},
						BottomRight: Cell{Row: maxRow, Col: maxCol},
					},
					minRow: 0,
					expectedCells: []Cell{
						{
							Row: maxRow - 1,
							Col: maxCol,
						},
					},
				},
				{
					name: "top occupies one cell",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: maxRow, Col: 3},
						TopRight:    Cell{Row: maxRow, Col: 3},
						BottomLeft:  Cell{Row: maxRow, Col: 3},
						BottomRight: Cell{Row: maxRow, Col: 3},
					},
					minRow: 0,
					expectedCells: []Cell{
						{
							Row: maxRow - 1,
							Col: 3,
						},
					},
				},
				{
					name: "top occupies two cells",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: maxRow, Col: 1},
						TopRight:    Cell{Row: maxRow, Col: 2},
						BottomLeft:  Cell{Row: maxRow, Col: 1},
						BottomRight: Cell{Row: maxRow, Col: 2},
					},
					minRow: 0,
					expectedCells: []Cell{
						{
							Row: maxRow - 1,
							Col: 1,
						},
						{
							Row: maxRow - 1,
							Col: 2,
						},
					},
				},
				{
					name: "left top",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: maxRow, Col: 0},
						TopRight:    Cell{Row: maxRow, Col: 0},
						BottomLeft:  Cell{Row: maxRow, Col: 0},
						BottomRight: Cell{Row: maxRow, Col: 0},
					},
					minRow: 0,
					expectedCells: []Cell{
						{
							Row: maxRow - 1,
							Col: 0,
						},
					},
				},
				{
					name: "left occupies one cell",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 3, Col: 0},
						TopRight:    Cell{Row: 3, Col: 0},
						BottomLeft:  Cell{Row: 3, Col: 0},
						BottomRight: Cell{Row: 3, Col: 0},
					},
					minRow: 0,
					expectedCells: []Cell{
						{
							Row: 2,
							Col: 0,
						},
					},
				},
				{
					name: "left occupies two cells",
					cornerCells: CornerCells{
						TopLeft:     Cell{Row: 3, Col: 0},
						TopRight:    Cell{Row: 3, Col: 0},
						BottomLeft:  Cell{Row: 2, Col: 0},
						BottomRight: Cell{Row: 2, Col: 0},
					},
					minRow: 0,
					expectedCells: []Cell{
						{
							Row: 1,
							Col: 0,
						},
					},
				},
			},
		},
	}

	for _, testSuite := range testSuites {
		t.Run(testSuite.name, func(t *testing.T) {
			for _, testCase := range testSuite.testCases {
				t.Run(testCase.name, func(t *testing.T) {
					cells := GetBottomNeighborCells(
						testCase.cornerCells, testCase.minRow,
					)
					assert.Equal(t, testCase.expectedCells, cells)
				})
			}
		})
	}
}
