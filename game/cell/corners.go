package cell

import (
    "fmt"
)

type CornerCells struct {
    TopLeft     Cell
    TopRight    Cell
    BottomLeft  Cell
    BottomRight Cell
}

func (c CornerCells) String() string {
    return fmt.Sprintf("[TopLeft:%v,TopRight:%v,BottomLeft:%v,BottomRight:%v]", c.TopLeft, c.TopRight, c.BottomLeft, c.BottomRight)
}

// Detailed analysis: https://docs.google.com/document/d/1FF49lQrCCNLEuqInmXRPnbl-jvDLLtqor6Q4OgGUIy4/edit
func GetCornerCells(
    bottomLeftX int,
    bottomLeftY int,
    objectWidth int,
    objectHeight int,
    cellWidth int,
    cellHeight int,
) CornerCells {
    topLeft := getTopLeftCell(bottomLeftX, bottomLeftY, objectHeight, cellWidth, cellHeight)
    topRight := getTopRightCell(bottomLeftX, bottomLeftY, objectWidth, objectHeight, cellWidth, cellHeight)
    bottomLeft := Cell{
        Row: bottomLeftY / cellHeight,
        Col: bottomLeftX / cellWidth,
    }
    bottomRight := getBottomRightCell(bottomLeftX, bottomLeftY, objectWidth, cellWidth, cellHeight)
    return CornerCells{
        TopLeft:     topLeft,
        TopRight:    topRight,
        BottomLeft:  bottomLeft,
        BottomRight: bottomRight,
    }
}

func getTopLeftCell(bottomLeftX int, bottomLeftY int, objectHeight int, cellWidth int, cellHeight int) Cell {
    topY := bottomLeftY + objectHeight
    row := topY / cellHeight
    if topY%cellHeight == 0 {
        row--
    }
    column := bottomLeftX / cellWidth
    return Cell{
        Row: row,
        Col: column,
    }
}

func getTopRightCell(bottomLeftX int, bottomLeftY int, objectWidth int, objectHeight int, cellWidth int, cellHeight int) Cell {
    topY := bottomLeftY + objectHeight
    row := topY / cellHeight
    if topY%cellHeight == 0 {
        row--
    }
    rightX := bottomLeftX + objectWidth
    column := rightX / cellWidth
    if rightX%cellWidth == 0 {
        column--
    }
    return Cell{
        Row: row,
        Col: column,
    }
}

func getBottomRightCell(bottomLeftX int, bottomLeftY int, objectWidth int, cellWidth int, cellHeight int) Cell {
    row := bottomLeftY / cellHeight
    rightX := bottomLeftX + objectWidth
    column := rightX / cellWidth
    if rightX%cellWidth == 0 {
        column--
    }
    return Cell{
        Row: row,
        Col: column,
    }
}
