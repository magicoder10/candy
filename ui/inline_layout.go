package ui

import (
	"math"
)

var _ layout = (*InlineLayout)(nil)

type InlineLayout struct {
	BoxLayout
}

func (b InlineLayout) applyConstraintsToChildren(parent Component, parentConstraints Constraints) {
	parentConstraints.maxHeight = math.MaxInt64

	for _, child := range parent.getChildren() {
		applyConstraints(child, parentConstraints)
	}
}

func (b InlineLayout) computeParentSize(parent Component, parentConstraints Constraints) Size {
	height := 0
	children := parent.getChildren()
	length := len(children)

	if length > 0 {
		childrenOffset := parent.getChildrenOffset()
		height = childrenOffset[length-1].y + children[length-1].getSize().height
	}
	width := 0.0
	for _, child := range children {
		width = math.Max(width, float64(child.getSize().width))
	}
	padding := parent.getStyle().GetPadding()
	fullWidth := int(width) + padding.GetLeft() + padding.GetRight()
	fullHeight := height + padding.GetBottom()
	return Size{width: fullWidth, height: fullHeight}
}
