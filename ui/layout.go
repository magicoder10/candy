package ui

import (
	"candy/ui/ptr"
)

type Layout interface {
	applyConstraintsToChildren(parent Component, parentConstraints Constraints)
	computeParentSize(parent Component, parentConstraints Constraints) Size
	computeChildrenOffset(parent Component) []Offset
}

type LayoutType int

const (
	BoxLayoutType LayoutType = iota
	InlineLayoutType
)

func NewLayout(layoutType LayoutType) Layout {
	switch layoutType {
	case BoxLayoutType:
		return BoxLayout{}
	case InlineLayoutType:
		return InlineLayout{}
	}
	return BoxLayout{}
}

func LayoutTypePtr(layoutType LayoutType) *LayoutType {
	return (*LayoutType)(ptr.Int(int(layoutType)))
}

func getFullHeight(parent Component) int {
	children := parent.getChildren()
	style := parent.getStyle()
	padding := style.GetPadding()

	length := len(children)

	height := 0

	if style.Height != nil {
		height = *style.Height
	} else if length > 0 {
		childrenOffset := parent.getChildrenOffset()
		height = childrenOffset[length-1].y + children[length-1].getSize().height
	} else {
		height = padding.GetTop()
	}
	return height + padding.GetBottom()
}
