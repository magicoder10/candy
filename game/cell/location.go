package cell

func GetCellLocatedAt(bottomLeftX int, bottomLeftY int, objectWidth int, objectHeight int, cellWidth int, cellHeight int) Cell {
    centerX := bottomLeftX + objectWidth/2
    centerY := bottomLeftY + objectHeight/2
    row := centerY / cellHeight
    col := centerX / cellWidth
    return Cell{
        Row: row,
        Col: col,
    }
}
