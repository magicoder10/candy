package ui

import (
	"image"
	"image/draw"
	"sort"
)

type BoxProps struct {
}

var _ Component = (*Box)(nil)

type Box struct {
	SharedComponent
	props BoxProps
}

func (b Box) Paint(painter *Painter, destLayer draw.Image, offset Offset) {
	contentLayer := image.NewRGBA(image.Rectangle{
		Max: image.Point{
			X: b.size.width,
			Y: b.size.height,
		},
	})

	sortedChildren := Children{
		children:       b.children,
		childrenOffset: b.childrenOffset,
	}
	sort.Sort(&sortedChildren)

	for index, child := range sortedChildren.children {
		childOffset := sortedChildren.childrenOffset[index]
		child.Paint(painter, contentLayer, childOffset)
	}

	painter.drawImage(contentLayer, image.Rectangle{
		Min: image.Point{},
		Max: contentLayer.Bounds().Max,
	}, destLayer, image.Point{
		X: offset.x,
		Y: offset.y,
	})
}

func (b Box) ComputeLeafSize(_ Constraints) Size {
	width := 0
	if b.style.Width != nil {
		width = *b.style.Width
	}
	height := 0
	if b.style.Height != nil {
		height = *b.style.Height
	}
	return Size{width: width, height: height}
}

func NewBox(props *BoxProps, children []Component, style *Style) *Box {
	if props == nil {
		props = &BoxProps{}
	}
	if style == nil {
		style = &Style{
			LayoutType: BoxLayoutType,
		}
	}
	if children == nil {
		children = make([]Component, 0)
	}
	return &Box{
		props: *props,
		SharedComponent: SharedComponent{
			layout:         newLayout(style.LayoutType),
			style:          *style,
			children:       children,
			childrenOffset: []Offset{},
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
