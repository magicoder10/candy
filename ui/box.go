package ui

import (
	"sort"
)

type BoxProps struct {
	OnClick      onClickHandler
	OnMouseEnter onMouseEnterHandler
	OnMouseLeave onMouseLeaveHandler
}

var _ Component = (*Box)(nil)

type Box struct {
	SharedComponent
}

func (b Box) GetName() string {
	return "Box"
}

func (b Box) ComputeLeafSize(_ Constraints) Size {
	style := b.getStyle()
	padding := style.GetPadding()

	width := 0
	if style.Width != nil {
		width = *style.Width
	}
	width += padding.GetLeft() + padding.GetRight()
	height := 0
	if style.Height != nil {
		height = *style.Height
	}
	height += padding.GetTop() + padding.GetBottom()
	return Size{width: width, height: height}
}

func NewBox(props *BoxProps, children []Component, statefulStyle *StatefulStyle) *Box {
	if props == nil {
		props = &BoxProps{}
	}
	if statefulStyle == nil {
		statefulStyle = NewStatefulStyleWithLayout(BoxLayoutType)
	}
	normalStyle := statefulStyle.Styles[NormalState]
	if normalStyle.LayoutType == nil {
		normalStyle.LayoutType = LayoutTypePtr(BoxLayoutType)
	}
	if children == nil {
		children = make([]Component, 0)
	}
	return &Box{
		SharedComponent: SharedComponent{
			Name:           "Box",
			States:        map[State]struct{}{},
			StatefulStyle:  statefulStyle,
			Children:       children,
			childrenOffset: []Offset{},
			events: Events{
				onClick:      props.OnClick,
				onMouseEnter: props.OnMouseEnter,
				onMouseLeave: props.OnMouseLeave,
			},
		}}
}

var _ sort.Interface = (*Children)(nil)

type Children struct {
	children       []Component
	childrenOffset []Offset
}

func (c Children) Len() int {
	return len(c.children)
}

func (c Children) Less(i, j int) bool {
	return c.childrenOffset[i].z < c.childrenOffset[j].z
}

func (c *Children) Swap(i, j int) {
	c.children[i], c.children[j] = c.children[j], c.children[i]
	c.childrenOffset[i], c.childrenOffset[j] = c.childrenOffset[j], c.childrenOffset[i]
}
