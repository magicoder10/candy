package ui

import (
	"math"
)

var _ Layout = (*InlineLayout)(nil)

type InlineLayout struct {
	BoxLayout
}

func (i InlineLayout) applyConstraintsToChildren(parent Component, parentConstraints Constraints) {
	parentConstraints.maxHeight = math.MaxInt64

	style := parent.getStyle()
	if style.Width != nil {
		parentConstraints.maxWidth = *style.Width
	}
	parent.setSize(Size{
		width:  parentConstraints.maxWidth,
		height: parentConstraints.maxHeight,
	})

	for _, child := range parent.getChildren() {
		applyConstraints(child, parentConstraints)
	}
}

func (i InlineLayout) computeParentSize(parent Component, _ Constraints) Size {
	children := parent.getChildren()
	style := parent.getStyle()

	width := 0.0
	if style.Width != nil {
		width = float64(*style.Width)
	} else {
		for _, child := range children {
			width = math.Max(width, float64(child.getSize().width))
		}
	}
	padding := parent.getStyle().GetPadding()
	fullWidth := int(width) + padding.GetLeft() + padding.GetRight()
	return Size{width: fullWidth, height: getFullHeight(parent)}
}
