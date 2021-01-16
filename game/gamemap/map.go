package gamemap

import (
	"math"
	"math/rand"
	"time"

	"candy/assets"
	"candy/game/candy"
	"candy/game/cell"
	"candy/game/gameitem"
	"candy/game/player"
	"candy/game/square"
	"candy/graphics"
)

const defaultRows = 12
const defaultCols = 15

type Map struct {
	screenX           int
	screenY           int
	batch             graphics.Batch
	maxRow            int
	maxCol            int
	gridXOffset       int
	gridYOffset       int
	grid              *[][]square.Square
	tiles             *[]*square.Tile
	candies           *map[cell.Cell]*candy.Candy
	candyRangeCutter  candy.RangeCutter
	playerMoveChecker player.MoveChecker
	eventHandlers     eventHandlers
}

func (m Map) DrawMap() {
	bound := graphics.Bound{
		X:      0,
		Y:      0,
		Width:  900,
		Height: 736,
	}
	m.batch.DrawSprite(m.screenX, m.screenY, math.MaxInt32, bound, 1)
	m.batch.RenderBatch()
}

func (m Map) DrawGrid(batch graphics.Batch) {
	for row, gridRow := range *m.grid {
		for col := len(gridRow) - 1; col >= 0; col-- {
			if gridRow[col] == nil {
				continue
			}
			x := m.gridXOffset + col*square.Width
			y := m.gridYOffset + row*square.Width
			gridRow[col].Draw(batch, x, y)
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
	queue := make([]cell.Cell, 0)
	visited := make(map[cell.Cell]struct{})

	for candyCell, c := range *m.candies {
		if c.Exploding() {
			visited[candyCell] = struct{}{}
			queue = append(queue, candyCell)
		}
	}

	for len(queue) > 0 {
		currCell := queue[0]
		queue = queue[1:]

		if c, ok := (*m.candies)[currCell]; ok {
			c.Explode()

			for _, nextCell := range c.CellsHit() {
				if !inGrid(nextCell, m.maxRow, m.maxCol) {
					continue
				}
				if _, ok := visited[nextCell]; ok {
					continue
				}
				visited[nextCell] = struct{}{}
				queue = append(queue, nextCell)
				m.eventHandlers.onCandyExploding(nextCell)
			}
		}
	}
	m.removeExplodedCandies()
}

func (m *Map) removeExplodedCandies() {
	newCandies := make(map[cell.Cell]*candy.Candy)

	for candyCell, c := range *m.candies {
		if c.Exploded() {
			(*m.grid)[candyCell.Row][candyCell.Col] = nil
		} else {
			newCandies[candyCell] = c
		}
	}
	m.candies = &newCandies
}

func (m Map) GetPlayerMoveChecker() player.MoveChecker {
	return m.playerMoveChecker
}

func (m *Map) AddCandy(cell cell.Cell, candyBuilder candy.Builder) bool {
	if cell.Row < 0 || cell.Row > m.maxRow || cell.Col < 0 || cell.Col > m.maxCol {
		return false
	}
	if (*m.grid)[cell.Row][cell.Col] != nil {
		return false
	}
	c, err := candyBuilder.
		Center(cell).
		RangeCutter(m.candyRangeCutter).
		Build()
	if err != nil {
		return false
	}
	(*m.candies)[cell] = &c
	(*m.grid)[cell.Row][cell.Col] = &c
	return true
}

func randomGameItem() gameitem.GameItem {
	index := rand.Intn(len(gameitem.GameItems))
	return gameitem.GameItems[index]
}

func (m *Map) OnCandyExploding(callback func(hitCell cell.Cell)) {
	m.eventHandlers.onCandyExploding = callback
}

func (m Map) GetGridXOffset() int {
	return m.gridXOffset
}

func (m Map) GetGridYOffset() int {
	return m.gridYOffset
}

func NewMap(assets assets.Assets, g graphics.Graphics, screenX int, screenY int) *Map {
	rand.Seed(time.Now().UnixNano())
	grid := make([][]square.Square, 0)

	for row := 0; row < defaultRows; row++ {
		gridRow := make([]square.Square, defaultCols)
		grid = append(grid, gridRow)
	}

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

	tiles := make([]*square.Tile, 0)

	for row, rowConfig := range mapConfig {
		for col, colConfig := range rowConfig {
			if colConfig == ' ' {
				continue
			}
			t := square.NewTile(colConfig, randomGameItem())
			tiles = append(tiles, &t)
			grid[row][col] = &t
		}
	}

	candies := make(map[cell.Cell]*candy.Candy, 0)
	maxRow := defaultRows - 1
	maxCol := defaultCols - 1
	cdRangeCutter := candyRangeCutter{
		maxRow: maxRow,
		maxCol: maxCol,
		grid:   &grid,
	}
	gridXOffset := screenX
	gridYOffset := screenY

	return &Map{
		screenX:          screenX,
		screenY:          screenY,
		batch:            g.StartNewBatch(assets.GetImage("map/default.png")),
		maxRow:           maxRow,
		maxCol:           maxCol,
		gridXOffset:      gridXOffset,
		gridYOffset:      gridYOffset,
		grid:             &grid,
		tiles:            &tiles,
		candies:          &candies,
		candyRangeCutter: &cdRangeCutter,
		playerMoveChecker: &moveChecker{
			gridXOffset: gridXOffset,
			gridYOffset: gridYOffset,
			maxRow:      maxRow,
			maxCol:      maxCol,
			grid:        &grid,
		},
	}
}
