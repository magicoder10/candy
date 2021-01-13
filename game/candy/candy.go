package candy

import (
	"errors"
	"time"

	"candy/game/cell"
	"candy/game/cutter"
	"candy/game/square"
	"candy/graphics"
)

var _ square.Square = (*Candy)(nil)

type Candy struct {
	state State
}

func (c Candy) IsBreakable() bool {
	return true
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
	c.state.Explode()
}

func (c Candy) Exploded() bool {
	return c.state.Exploded()
}

func (c Candy) Exploding() bool {
	return c.state.Exploding()
}

func (c Candy) CellsHit() []cell.Cell {
	return c.state.CellsHit()
}

func (c *Candy) MoveTo(cell cell.Cell) {
	c.state.SetCenter(cell)
}

func (c Candy) GetCellOn() cell.Cell {
	return c.state.GetCenter()
}

func newCandy(powerLevel int, center cell.Cell, rangeCutter cutter.Range) Candy {
	return Candy{state: newMeltingState(powerLevel, center, rangeCutter)}
}

type Builder struct {
	powerLevel  int
	center      *cell.Cell
	rangeCutter cutter.Range
}

func (b *Builder) Center(center cell.Cell) *Builder {
	b.center = &center
	return b
}

func (b *Builder) RangeCutter(rangeCutter cutter.Range) *Builder {
	b.rangeCutter = rangeCutter
	return b
}

func (b *Builder) Build() (Candy, error) {
	if b.center == nil {
		return Candy{}, errors.New("center cannot be empty")
	}
	return newCandy(b.powerLevel, *b.center, b.rangeCutter), nil
}

func NewBuilder(powerLevel int) Builder {
	return Builder{powerLevel: powerLevel, rangeCutter: cutter.NoChange{}}
}
