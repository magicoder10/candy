package ui

import (
	"math"
)

var _ Layout = (*BoxLayout)(nil)

type BoxLayout struct {
}

func (b BoxLayout) applyConstraintsToChildren(parent Component, parentConstraints Constraints) {
	parentConstraints.maxHeight = math.MaxInt64
	parentConstraints.minWidth = parentConstraints.maxWidth

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

func (b BoxLayout) computeChildrenOffset(parent Component) []Offset {
	style := parent.getStyle()
	padding := style.GetPadding()

	aligner := style.GetAlignment()
	nextY := padding.GetTop()

	offsets := make([]Offset, 0)
	for _, child := range parent.getChildren() {
		nextY += child.getStyle().GetMargin().GetTop()

		offset := Offset{
			x: aligner.AlignHorizontal(parent, child),
			y: nextY,
		}
		offsets = append(offsets, offset)
		nextY += child.getSize().height
	}
	return offsets
}

func (b BoxLayout) computeParentSize(parent Component, parentConstraints Constraints) Size {
	return Size{width: parentConstraints.maxWidth, height: getFullHeight(parent)}
}
