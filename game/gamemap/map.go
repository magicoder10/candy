package gamemap

import (
	"candy/game/gameitem"
	"image"
	"math/rand"
	"time"

	"candy/assets"
	"candy/game/cell"
	"candy/game/direction"
	"candy/game/tile"
	"candy/graphics"
)

type Map struct {
	backgroundImage image.Image
	maxRow          int
	maxCol          int
	tileXOffset     int
	tileYOffset     int
	tiles           [][]*tile.Tile
}

func (m Map) DrawMap(g graphics.Graphics) {
	bound := graphics.Bound{
		X:      0,
		Y:      0,
		Width:  900,
		Height: 736,
	}
	g.DrawImage(0, 0, m.backgroundImage, bound, 1)
}

func (m Map) DrawTiles(batch graphics.Batch) {
	for row, rowTiles := range m.tiles {
		for col := len(rowTiles) - 1; col >= 0; col-- {
			if rowTiles[col] == nil {
				continue
			}
			rowTiles[col].Draw(batch, m.tileXOffset+col*tile.Width, m.tileYOffset+row*tile.Height)
		}
	}
}

func (m Map) RevealItems() {
	for _, rowTiles := range m.tiles {
		for _, t := range rowTiles {
			if t == nil {
				continue
			}
			t.RevealItem()
		}
	}
}

func (m Map) HideItems(){
	for _, rowTiles := range m.tiles {
		for _, t := range rowTiles {
			if t == nil {
				continue
			}
			t.HideItem()
		}
	}
}

func (m Map) CanMove(currX int, currY int, objectWidth int, objectHeight int, dir direction.Direction, stepSize int) bool {
	if !m.inBound(currX, currY, objectWidth, objectHeight, dir, stepSize) {
		return false
	}
	cornerCells := cell.GetCornerCells(currX, currY, objectWidth, objectHeight, tile.Width, tile.Height)
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
		return nextY <= m.tileYOffset+(m.maxRow+1)*tile.Height
	case direction.Down:
		nextY := currY - stepSize
		return nextY >= m.tileYOffset
	case direction.Left:
		nextX := currX - stepSize
		return nextX >= m.tileXOffset
	case direction.Right:
		nextX := currX + objectWidth + stepSize
		return nextX <= m.tileXOffset+(m.maxCol+1)*tile.Width
	}
	return true
}

func isBlocked(blockingCell cell.Cell, currX int, currY int, objectWidth int, objectHeight int, dir direction.Direction, stepSize int) bool {
	switch dir {
	case direction.Up:
		nextY := currY + objectHeight + stepSize
		cellBottom := blockingCell.Row * tile.Height
		return nextY > cellBottom
	case direction.Down:
		nextY := currY - stepSize
		cellTop := blockingCell.Row*tile.Height + tile.Height
		return nextY < cellTop
	case direction.Left:
		nextX := currX - stepSize
		cellRight := blockingCell.Col*tile.Width + tile.Width
		return nextX < cellRight
	case direction.Right:
		nextX := currX + objectWidth + stepSize
		cellLeft := blockingCell.Col * tile.Width
		return nextX > cellLeft
	}
	return false
}

func (m Map) getBlockingCells(cells []cell.Cell) []cell.Cell {
	newCells := make([]cell.Cell, 0)
	for _, c := range cells {
		if len(m.tiles) <= c.Row || len(m.tiles[c.Row]) <= c.Col {
			continue
		}
		if m.tiles[c.Row][c.Col] == nil {
			continue
		}
		if m.tiles[c.Row][c.Col].CanEnter() {
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

func NewMap(assets assets.Assets) Map {
	rand.Seed(time.Now().UnixNano())
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

	tiles := make([][]*tile.Tile, 0)
	for _, rowConfig := range mapConfig {
		tileRow := make([]*tile.Tile, 0)

		for _, colConfig := range rowConfig {
			tileRow = append(tileRow, tile.NewTile(colConfig, randomGameItem()))
		}
		tiles = append(tiles, tileRow)
	}
	return Map{
		backgroundImage: assets.GetImage("map/default.png"),
		maxRow:          11,
		maxCol:          14,
		tileXOffset:     0,
		tileYOffset:     0,
		tiles:           tiles,
	}
}
