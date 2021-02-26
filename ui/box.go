package ui

import (
	"sort"

	"candy/ui/ptr"
)

type BoxProps struct {
	OnClick onClickHandler
}

var _ Component = (*Box)(nil)

type Box struct {
	SharedComponent
}

func (b Box) GetName() string {
	return "Box"
}

func (b Box) ComputeLeafSize(_ Constraints) Size {
	padding := b.Style.GetPadding()

	width := 0
	if b.Style.Width != nil {
		width = *b.Style.Width
	}
	width += padding.GetLeft() + padding.GetRight()
	height := 0
	if b.Style.Height != nil {
		height = *b.Style.Height
	}
	height += padding.GetTop() + padding.GetBottom()
	return Size{width: width, height: height}
}

func NewBox(pros *BoxProps, children []Component, style *Style) *Box {
	if pros == nil {
		pros = &BoxProps{}
	}
	if style == nil {
		style = &Style{
			LayoutType: (*LayoutType)(ptr.Int(int(BoxLayoutType))),
		}
	}
	if style.LayoutType == nil {
		style.LayoutType = (*LayoutType)(ptr.Int(int(BoxLayoutType)))
	}
	if children == nil {
		children = make([]Component, 0)
	}
	return &Box{
		SharedComponent: SharedComponent{
			Name:           "Box",
			Layout:         NewLayout(*style.LayoutType),
			Style:          style,
			Children:       children,
			childrenOffset: []Offset{},
			events:         Events{onClick: pros.OnClick},
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
