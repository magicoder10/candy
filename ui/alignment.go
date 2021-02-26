package ui

type AlignHorizontal int

func (a AlignHorizontal) Ptr() *AlignHorizontal {
	return &a
}

const (
	AlignLeft AlignHorizontal = iota
	AlignHorizontalCenter
	AlignRight
)

type AlignVertical int

const (
	AlignTop AlignVertical = iota
	AlignVerticalCenter
	AlignBottom
)

type Alignment struct {
	Horizontal *AlignHorizontal
	Vertical   *AlignVertical
}

func (a Alignment) getHorizontal() AlignHorizontal {
	if a.Horizontal == nil {
		return AlignLeft
	}
	return *a.Horizontal
}

func (a Alignment) getVertical() AlignVertical {
	if a.Vertical == nil {
		return AlignTop
	}
	return *a.Vertical
}

func (a Alignment) AlignHorizontal(parent Component, child Component) int {
	padding := parent.getStyle().GetPadding()
	margin := child.getStyle().GetMargin()
	switch a.getHorizontal() {
	case AlignLeft:
		return margin.GetLeft() + padding.GetLeft()
	case AlignRight:
		return parent.getSize().width - padding.GetRight() - margin.GetRight() - child.getSize().width
	case AlignHorizontalCenter:
		return (parent.getSize().width - child.getSize().width) / 2
	}
	return margin.GetLeft() + padding.GetLeft()
}
