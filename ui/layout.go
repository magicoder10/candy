package ui

type layout interface {
    applyConstraintsToChildren(parent Component, parentConstraints Constraints)
    computeParentSize(parent Component, parentConstraints Constraints) size
    computeChildrenOffset(parent Component) []offset
}

var _ layout = (*BoxLayout)(nil)
