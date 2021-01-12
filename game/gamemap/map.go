package gamemap

import (
	"math"
	"math/rand"
	"time"

	"candy/assets"
	"candy/game/candy"
	"candy/game/cell"
	"candy/game/direction"
	"candy/game/gameitem"
	"candy/game/square"
	"candy/game/tile"
	"candy/graphics"
)

const defaultRows = 12
const defaultCols = 15

type Map struct {
	batch       graphics.Batch
	maxRow      int
	maxCol      int
	gridXOffset int
	gridYOffset int
	grid        *[defaultRows][defaultCols]square.Square
	tiles       *[]*tile.Tile
	candies     *[]*candy.Candy
}

func (m Map) DrawMap() {
	bound := graphics.Bound{
		X:      0,
		Y:      0,
		Width:  900,
		Height: 736,
	}
	m.batch.DrawSprite(0, 0, math.MaxInt32, bound, 1)
	m.batch.RenderBatch()
}

func (m Map) DrawGrid(batch graphics.Batch) {
	for row, gridRow := range m.grid {
		for col := len(gridRow) - 1; col >= 0; col-- {
			if gridRow[col] == nil {
				continue
			}
			gridRow[col].Draw(batch, m.gridXOffset+col*square.Width, m.gridYOffset+row*square.Width)
		}
	}
}

func (m Map) RevealItems() {
	for _, t := range *m.tiles {
		t.RevealItem()
	}
}

func (m Map) HideItems() {
	for _, t := range *m.tiles {
		t.HideItem()
	}
}

func (m *Map) Update(timeElapsed time.Duration) {
	for _, c := range *m.candies {
		c.Update(timeElapsed)
	}
	m.removeExplodedCandies()
}

func (m *Map) AddCandy(cell cell.Cell, candy candy.Candy) bool {
	if cell.Row < 0 || cell.Row > m.maxRow || cell.Col < 0 || cell.Col > m.maxCol {
		return false
	}
	if m.grid[cell.Row][cell.Col] != nil {
		return false
	}
	*m.candies = append(*m.candies, &candy)
	m.grid[cell.Row][cell.Col] = &candy

	candy.MoveTo(cell)
	return true
}

func (m *Map) removeExplodedCandies() {
	newCandies := make([]*candy.Candy, 0)

	for _, c := range *m.candies {
		if c.Exploded() {
			cellOn := c.GetCellOn()
			m.grid[cellOn.Row][cellOn.Col] = nil
		} else {
			newCandies = append(newCandies, c)
		}
	}
	m.candies = &newCandies
}

func (m Map) CanMove(currX int, currY int, objectWidth int, objectHeight int, dir direction.Direction, stepSize int) bool {
	if !m.inBound(currX, currY, objectWidth, objectHeight, dir, stepSize) {
		return false
	}
	cornerCells := cell.GetCornerCells(currX, currY, objectWidth, objectHeight, square.Width, square.Width)
	neighborCells := m.getNeighborCells(cornerCells, dir)
	blockingCells := m.getBlockingCells(neighborCells)

	for _, blockingCell := range blockingCells {
		if isBlocked(blockingCell, currX, currY, objectWidth, objectHeight, dir, stepSize) {
			return false
		}
	}
	return true
}

func (m Map) inBound(currX int, currY int, objectWidth, objectHeight int, dir direction.Direction, stepSize int) bool {
	switch dir {
	case direction.Up:
		nextY := currY + objectHeight + stepSize
		return nextY <= m.gridYOffset+(m.maxRow+1)*square.Width
	case direction.Down:
		nextY := currY - stepSize
		return nextY >= m.gridYOffset
	case direction.Left:
		nextX := currX - stepSize
		return nextX >= m.gridXOffset
	case direction.Right:
		nextX := currX + objectWidth + stepSize
		return nextX <= m.gridXOffset+(m.maxCol+1)*square.Width
	}
	return true
}

func isBlocked(blockingCell cell.Cell, currX int, currY int, objectWidth int, objectHeight int, dir direction.Direction, stepSize int) bool {
	switch dir {
	case direction.Up:
		nextY := currY + objectHeight + stepSize
		cellBottom := blockingCell.Row * square.Width
		return nextY > cellBottom
	case direction.Down:
		nextY := currY - stepSize
		cellTop := blockingCell.Row*square.Width + square.Width
		return nextY < cellTop
	case direction.Left:
		nextX := currX - stepSize
		cellRight := blockingCell.Col*square.Width + square.Width
		return nextX < cellRight
	case direction.Right:
		nextX := currX + objectWidth + stepSize
		cellLeft := blockingCell.Col * square.Width
		return nextX > cellLeft
	}
	return false
}

func (m Map) getBlockingCells(cells []cell.Cell) []cell.Cell {
	newCells := make([]cell.Cell, 0)
	for _, c := range cells {
		if len(*m.tiles) <= c.Row || len(m.grid[c.Row]) <= c.Col {
			continue
		}
		if m.grid[c.Row][c.Col] == nil {
			continue
		}
		if m.grid[c.Row][c.Col].CanEnter() {
			continue
		}
		newCells = append(newCells, c)
	}
	return newCells
}

func (m Map) getNeighborCells(cornerCells cell.CornerCells, dir direction.Direction) []cell.Cell {
	switch dir {
	case direction.Up:
		return cell.GetTopNeighborCells(cornerCells, m.maxRow)
	case direction.Down:
		return cell.GetBottomNeighborCells(cornerCells, 0)
	case direction.Left:
		return cell.GetLeftNeighborCells(cornerCells, 0)
	case direction.Right:
		return cell.GetRightNeighborCells(cornerCells, m.maxCol)
	}
	return []cell.Cell{}
}

func randomGameItem() gameitem.GameItem {
	index := rand.Intn(len(gameitem.GameItems))
	return gameitem.GameItems[index]
}

func NewMap(assets assets.Assets, g graphics.Graphics) *Map {
	rand.Seed(time.Now().UnixNano())
	var grid [defaultRows][defaultCols]square.Square

	mapConfig := [][]rune{
		{},
		{},
		{'G', 'Y'},
		{},
		{'G', 'Y', ' ', 'Y', ' ', ' ', ' ', 'G'},
		{' ', ' ', ' ', 'Y', ' ', ' ', ' ', 'G'},
		{'G', 'Y', ' ', 'Y', ' ', ' ', ' ', 'G'},
		{},
		{'G', 'Y'},
		{},
		{'G', 'Y'},
		{},
	}

	tiles := make([]*tile.Tile, 0)

	for row, rowConfig := range mapConfig {
		for col, colConfig := range rowConfig {
			if colConfig == ' ' {
				continue
			}
			t := tile.NewTile(colConfig, randomGameItem())
			tiles = append(tiles, &t)
			grid[row][col] = &t
		}
	}

	candies := make([]*candy.Candy, 0)
	return &Map{
		batch:       g.StartNewBatch(assets.GetImage("map/default.png")),
		maxRow:      defaultRows - 1,
		maxCol:      defaultCols - 1,
		gridXOffset: 0,
		gridYOffset: 0,
		grid:        &grid,
		tiles:       &tiles,
		candies:     &candies,
	}
}
