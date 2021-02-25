package ui

type Constraints struct {
	minWidth  int
	maxWidth  int
	minHeight int
	maxHeight int
}

func NewScreenConstraint(screenWidth int, screenHeight int) Constraints {
	return Constraints{
		minWidth:  screenWidth,
		maxWidth:  screenWidth,
		minHeight: screenHeight,
		maxHeight: screenHeight,
	}
}

func applyConstraints(component Component, constraints Constraints) {
	if len(component.getChildren()) == 0 {
		component.setSize(component.ComputeLeafSize(constraints))
		return
	}
	layout := component.getLayout()
	layout.applyConstraintsToChildren(component, constraints)
	component.setChildrenOffset(layout.computeChildrenOffset(component))
	component.setSize(layout.computeParentSize(component, constraints))
}
