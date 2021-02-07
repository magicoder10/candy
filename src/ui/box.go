package ui

import (
	"image"
	"image/draw"
	"sort"
)

var _ sort.Interface = (*Children)(nil)

type Children struct {
	children       []Component
	childrenOffset []offset
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

var _ Component = (*Box)(nil)

type Box struct {
	sharedComponent
}

func (b Box) paint(painter *painter, destLayer draw.Image, offset offset) {
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
		child.paint(painter, contentLayer, childOffset)
	}

	painter.drawImage(contentLayer, image.Rectangle{
		Min: image.Point{},
		Max: contentLayer.Bounds().Max,
	}, destLayer, image.Point{
		X: offset.x,
		Y: offset.y,
	})
}

func (b Box) computeLeafSize() size {
	width := 0
	if b.style.Width != nil {
		width = *b.style.Width
	}
	height := 0
	if b.style.Height != nil {
		height = *b.style.Height
	}
	return size{width: width, height: height}
}

type BoxBuilder struct {
	children []Component
	componentBuilder
}

func (b *BoxBuilder) Children(children []Component) *BoxBuilder {
	b.children = children
	return b
}

func (b *BoxBuilder) Build() *Box {
	if b.style == nil {
		b.style = &Style{}
	}
	if b.layout == nil {
		b.layout = BoxLayout{}
	}
	return &Box{sharedComponent{
		layout:   b.layout,
		style:    *b.style,
		children: b.children,
	}}
}

func NewBoxBuilder() *BoxBuilder {
	return &BoxBuilder{}
}
