package candy

import (
	"time"

	"candy/game/candy/state"
	"candy/game/cell"
	"candy/game/square"
	"candy/graphics"
)

var _ square.Square = (*Candy)(nil)

type Candy struct {
	cell  cell.Cell
	state state.State
}

func (c Candy) Draw(batch graphics.Batch, x int, y int) {
	c.state.Draw(batch, x, y, y+square.Width-1)
}

func (c Candy) CanEnter() bool {
	return false
}

func (c *Candy) Update(timeElapsed time.Duration) {
	c.state = c.state.Update(timeElapsed)
}

func (c Candy) Explode() {
}

func (c Candy) Exploded() bool {
	return c.state.Exploded()
}

func (c *Candy) MoveTo(cell cell.Cell) {
	c.cell = cell
}

func (c Candy) GetCellOn() cell.Cell {
	return c.cell
}

func NewCandy(powerLevel int) Candy {
	return Candy{state: state.NewMelting(powerLevel)}
}
