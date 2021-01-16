package candy

import (
	"errors"
	"time"

	"candy/game/cell"
	"candy/game/square"
	"candy/graphics"
)

var _ square.Square = (*Candy)(nil)

type Candy struct {
	state state
}

func (c Candy) ShouldRemove() bool {
	return false
}

func (c Candy) UnblockFire() {
	return
}

func (c Candy) Break() {
	c.state.explode()
}

func (c Candy) IsBreakable() bool {
	return true
}

func (c Candy) Draw(batch graphics.Batch, x int, y int) {
	c.state.draw(batch, x, y, y+square.Width-1)
}

func (c Candy) CanEnter() bool {
	return false
}

func (c *Candy) Update(timeElapsed time.Duration) {
	c.state = c.state.update(timeElapsed)
}

func (c Candy) Exploded() bool {
	return c.state.exploded()
}

func (c Candy) Exploding() bool {
	return c.state.exploding()
}

func (c Candy) CellsHit() []cell.Cell {
	return c.state.cellsHit()
}

func (c *Candy) MoveTo(cell cell.Cell) {
	c.state.setCenter(cell)
}

func (c Candy) GetCellOn() cell.Cell {
	return c.state.getCenter()
}

func newCandy(powerLevel int, center cell.Cell, rangeCutter RangeCutter) Candy {
	return Candy{state: newMeltingState(powerLevel, center, rangeCutter)}
}

type Builder struct {
	powerLevel  int
	center      *cell.Cell
	rangeCutter RangeCutter
}

func (b *Builder) Center(center cell.Cell) *Builder {
	b.center = &center
	return b
}

func (b *Builder) RangeCutter(rangeCutter RangeCutter) *Builder {
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
	return Builder{powerLevel: powerLevel, rangeCutter: noChange{}}
}
