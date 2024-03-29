package gamemap

import (
	"math"
	"math/rand"
	"sync"
	"time"

	"candy/assets"
	"candy/game/candy"
	"candy/game/cell"
	"candy/game/gameitem"
	"candy/game/player"
	"candy/game/square"
	"candy/graphics"
	"candy/pubsub"
)

const defaultRows = 12
const defaultCols = 15

const Width = 900

var mapBound = graphics.Bound{
	X:      0,
	Y:      0,
	Width:  900,
	Height: 736,
}

type brokenSquare struct {
	cell   cell.Cell
	square square.Square
}

type Map struct {
	screenX     int
	screenY     int
	batch       graphics.Batch
	maxRow      int
	maxCol      int
	gridXOffset int
	gridYOffset int
	grid        *[][]square.Square
	// this is only used for debugging
	tiles             *[]*square.Tile
	candiesMutex      *sync.Mutex
	candies           *map[cell.Cell]*candy.Candy
	brokenSquares     map[*candy.Candy][]brokenSquare
	candyRangeCutter  candy.RangeCutter
	playerMoveChecker player.MoveChecker
	pubSub            *pubsub.PubSub
}

func (m Map) DrawMap() {
	m.batch.DrawSprite(m.screenX, m.screenY, math.MaxInt32, mapBound, 1)
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
	// this is only used for debugging
	for _, t := range *m.tiles {
		t.RevealItem()
	}
}

func (m Map) HideItems() {
	// this is only used for debugging
	for _, t := range *m.tiles {
		t.HideItem()
	}
}

func (m *Map) Update(timeElapsed time.Duration) {
	m.candiesMutex.Lock()
	defer m.candiesMutex.Unlock()
	for _, c := range *m.candies {
		c.Update(timeElapsed)
	}

	onHitNextCell := func(currCandy *candy.Candy, nextHitCell cell.Cell) {
		m.collectBrokenSquares(nextHitCell, currCandy)
		m.pubSub.Publish(pubsub.OnCandyExploding, nextHitCell)
	}
	m.propagateExplosion(onHitNextCell)
	m.removeExplodedCandies()
}

func (m Map) propagateExplosion(onHitNextCell func(currCandy *candy.Candy, nextHitCell cell.Cell)) {
	queue := make([]cell.Cell, 0)
	visited := make(map[cell.Cell]struct{})

	for candyCell, c := range *m.candies {
		if c.Exploding() {
			visited[candyCell] = struct{}{}
			queue = append(queue, candyCell)
			onHitNextCell(c, candyCell)
		}
	}

	for len(queue) > 0 {
		currCell := queue[0]
		queue = queue[1:]

		if currCandy, ok := (*m.candies)[currCell]; ok {
			for _, nextCell := range currCandy.CellsHit() {
				if !inGrid(nextCell, m.maxRow, m.maxCol) {
					continue
				}
				if _, ok := visited[nextCell]; ok {
					continue
				}
				visited[nextCell] = struct{}{}
				queue = append(queue, nextCell)
				onHitNextCell(currCandy, nextCell)
			}
		}
	}
}

func (m Map) collectBrokenSquares(cell cell.Cell, candy *candy.Candy) {
	sq := (*m.grid)[cell.Row][cell.Col]
	if sq != nil && sq.IsBreakable() && !sq.IsBroken() {
		sq.Break()
		if _, ok := m.brokenSquares[candy]; !ok {
			m.brokenSquares[candy] = make([]brokenSquare, 0)
		}

		m.brokenSquares[candy] = append(m.brokenSquares[candy],
			brokenSquare{
				cell:   cell,
				square: sq,
			})
	}
}

func (m *Map) removeExplodedCandies() {
	newCandies := make(map[cell.Cell]*candy.Candy)

	for candyCell, c := range *m.candies {
		if c.Exploded() {
			(*m.grid)[candyCell.Row][candyCell.Col] = nil
			m.removeBrokenSquares(c)
		} else {
			newCandies[candyCell] = c
		}
	}
	m.candies = &newCandies
}

func (m *Map) removeBrokenSquares(cd *candy.Candy) {
	if brokenSquares, ok := m.brokenSquares[cd]; ok {
		for _, bs := range brokenSquares {
			bs.square.UnblockFire()
			if bs.square.ShouldRemove() {
				(*m.grid)[bs.cell.Row][bs.cell.Col] = nil
			}
		}
		delete(m.brokenSquares, cd)
	}
}

func (m Map) GetPlayerMoveChecker() player.MoveChecker {
	return m.playerMoveChecker
}

func (m *Map) AddCandy(cell cell.Cell, candyBuilder candy.Builder) bool {
	m.candiesMutex.Lock()
	defer m.candiesMutex.Unlock()

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

func randomGameItem() gameitem.Type {
	index := rand.Intn(len(gameitem.Types))
	return gameitem.Types[index]
}

func (m Map) HasRevealedItem(c cell.Cell) bool {
	if c.Row >= len(*m.grid) || (len(*m.grid) > 0 && c.Col >= len((*m.grid)[0])) {
		return false
	}
	sq := (*m.grid)[c.Row][c.Col]
	if sq == nil {
		return false
	}
	return sq.HasRevealedItem()
}

func (m Map) RetrieveGameItem(c cell.Cell) gameitem.Type {
	gameItemType := (*m.grid)[c.Row][c.Col].RetrieveGameItem()
	(*m.grid)[c.Row][c.Col] = nil
	return gameItemType
}

func NewMap(assets assets.Assets, g graphics.Graphics, pubSub *pubsub.PubSub, screenX int, screenY int) *Map {
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
			tiles = append(tiles, t)
			grid[row][col] = t
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
		candiesMutex:     &sync.Mutex{},
		candies:          &candies,
		candyRangeCutter: &cdRangeCutter,
		playerMoveChecker: &moveChecker{
			gridXOffset: gridXOffset,
			gridYOffset: gridYOffset,
			maxRow:      maxRow,
			maxCol:      maxCol,
			grid:        &grid,
		},
		brokenSquares: make(map[*candy.Candy][]brokenSquare),
		pubSub:        pubSub,
	}
}
