package cell

func GetLeftNeighborCells(cornerCells CornerCells, minCol int) []Cell {
	neighborCells := make([]Cell, 0)
	if cornerCells.TopLeft.Col-1 >= minCol {
		// Add left col
		for row := cornerCells.BottomLeft.Row; row <= cornerCells.TopLeft.Row; row++ {
			cell := Cell{
				Row: row,
				Col: cornerCells.TopLeft.Col - 1,
			}
			neighborCells = append(neighborCells, cell)
		}
	}
	return neighborCells
}

func GetRightNeighborCells(cornerCells CornerCells, maxCol int) []Cell {
	neighborCells := make([]Cell, 0)
	if cornerCells.BottomRight.Col+1 <= maxCol {
		// Add right col
		for row := cornerCells.BottomLeft.Row; row <= cornerCells.TopLeft.Row; row++ {
			cell := Cell{
				Row: row,
				Col: cornerCells.BottomRight.Col + 1,
			}
			neighborCells = append(neighborCells, cell)
		}
	}
	return neighborCells
}

func GetTopNeighborCells(cornerCells CornerCells, maxRow int) []Cell {
	neighborCells := make([]Cell, 0)
	if cornerCells.TopLeft.Row+1 <= maxRow {
		// Add top row
		for col := cornerCells.TopLeft.Col; col <= cornerCells.TopRight.Col; col++ {
			cell := Cell{
				Row: cornerCells.TopLeft.Row + 1,
				Col: col,
			}
			neighborCells = append(neighborCells, cell)
		}
	}
	return neighborCells
}

func GetBottomNeighborCells(cornerCells CornerCells, minRow int) []Cell {
	neighborCells := make([]Cell, 0)
	if cornerCells.BottomRight.Row-1 >= minRow {
		// Add bottom row
		for col := cornerCells.BottomLeft.Col; col <= cornerCells.BottomRight.Col; col++ {
			cell := Cell{
				Row: cornerCells.BottomRight.Row - 1,
				Col: col,
			}
			neighborCells = append(neighborCells, cell)
		}
	}
	return neighborCells
}
