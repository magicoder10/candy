package ui

import (
	"math"
)

var _ layout = (*BoxLayout)(nil)

type BoxLayout struct {
}

func (b BoxLayout) applyConstraintsToChildren(parent Component, parentConstraints Constraints) {
	parentConstraints.maxHeight = math.MaxInt64
	parentConstraints.minWidth = parentConstraints.maxWidth

	for _, child := range parent.getChildren() {
		applyConstraints(child, parentConstraints)
	}
}

func (b BoxLayout) computeChildrenOffset(parent Component) []Offset {
	padding := parent.getStyle().GetPadding()

	nextX := padding.GetLeft()
	nextY := padding.GetTop()

	offsets := make([]Offset, 0)
	for _, child := range parent.getChildren() {
		offsets = append(offsets, Offset{
			x: nextX,
			y: nextY,
		})
		nextY += child.getSize().height
	}
	return offsets
}

func (b BoxLayout) computeParentSize(parent Component, parentConstraints Constraints) Size {
	height := 0
	children := parent.getChildren()
	length := len(children)

	if length > 0 {
		childrenOffset := parent.getChildrenOffset()
		height = childrenOffset[length-1].y + children[length-1].getSize().height
	}
	return Size{width: parentConstraints.maxWidth, height: height}
}
