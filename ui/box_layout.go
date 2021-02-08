package ui

import (
    "math"
)

type BoxLayout struct {
}

func (b BoxLayout) applyConstraintsToChildren(parent Component, parentConstraints Constraints) {
    parentConstraints.maxHeight = math.MaxInt64

    for _, child := range parent.getChildren() {
        applyConstraints(child, parentConstraints)
    }
}

func (b BoxLayout) computeChildrenOffset(parent Component) []offset {
    nextX := 0
    nextY := 0

    offsets := make([]offset, 0)
    for _, child := range parent.getChildren() {
        offsets = append(offsets, offset{
            x: nextX,
            y: nextY,
        })
        nextY += child.getSize().height
    }
    return offsets
}

func (b BoxLayout) computeParentSize(parent Component, parentConstraints Constraints) size {
    height := 0
    children := parent.getChildren()
    length := len(children)

    if length > 0 {
        childrenOffset := parent.getChildrenOffset()
        height = childrenOffset[length-1].y + children[length-1].getSize().height
    }
    return size{width: parentConstraints.maxWidth, height: height}
}
