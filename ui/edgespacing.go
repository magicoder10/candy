package ui

type EdgeSpacing struct {
	All    *int
	Top    *int
	Bottom *int
	Left   *int
	Right  *int
}

func (e EdgeSpacing) GetTop() int {
	if e.Top != nil {
		return *e.Top
	} else if e.All != nil {
		return *e.All
	} else {
		return 0
	}
}

func (e EdgeSpacing) GetBottom() int {
	if e.Bottom != nil {
		return *e.Bottom
	} else if e.All != nil {
		return *e.All
	} else {
		return 0
	}
}

func (e EdgeSpacing) GetLeft() int {
	if e.Left != nil {
		return *e.Left
	} else if e.All != nil {
		return *e.All
	} else {
		return 0
	}
}

func (e EdgeSpacing) GetRight() int {
	if e.Right != nil {
		return *e.Right
	} else if e.All != nil {
		return *e.All
	} else {
		return 0
	}
}
