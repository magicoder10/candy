package cell

import (
	"fmt"
)

type Cell struct {
	Row int
	Col int
}

func (c Cell) String() string {
	return fmt.Sprintf("[Row:%d,Col:%d]", c.Row, c.Col)
}
