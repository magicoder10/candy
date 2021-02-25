package ui

type layout interface {
	applyConstraintsToChildren(parent Component, parentConstraints Constraints)
	computeParentSize(parent Component, parentConstraints Constraints) Size
	computeChildrenOffset(parent Component) []Offset
}

type LayoutType int

const (
	BoxLayoutType LayoutType = iota
	InlineLayoutType
)

func newLayout(layoutType LayoutType) layout {
	switch layoutType {
	case BoxLayoutType:
		return BoxLayout{}
	case InlineLayoutType:
		return InlineLayout{}
	}
	return BoxLayout{}
}
