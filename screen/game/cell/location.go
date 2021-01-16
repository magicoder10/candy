package cell

func GetCellLocatedAt(bottomLeftX int, bottomLeftY int, objectWidth int, objectHeight int, cellWidth int, cellHeight int) Cell {
	centerX := bottomLeftX + objectWidth/2
	centerY := bottomLeftY + objectHeight/2
	// cell width: 3
	//
	// 0 1 2 3 4 5 6 7 8 9 10
	// ------
	//       ------
	//             -------
	row := centerY / cellHeight
	if centerY%cellHeight == 0 {
		row++
	}

	col := centerX / cellWidth
	if centerX%cellWidth == 0 {
		col++
	}
	return Cell{
		Row: row,
		Col: col,
	}
}
