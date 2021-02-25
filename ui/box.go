package ui

import (
	"image"
	"image/draw"
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

func (b Box) Paint(painter *Painter, destLayer draw.Image, offset Offset) {
	if !b.hasChanged {
		return
	}

	contentLayer := image.NewRGBA(image.Rectangle{
		Max: image.Point{
			X: b.size.width,
			Y: b.size.height,
		},
	})
	if b.style.Background != nil {
		b.style.Background.Paint(painter, contentLayer)
	}

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
	padding := b.style.GetPadding()

	width := 0
	if b.style.Width != nil {
		width = *b.style.Width
	}
	width += padding.GetLeft() + padding.GetRight()
	height := 0
	if b.style.Height != nil {
		height = *b.style.Height
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
			name:           "Box",
			layout:         newLayout(*style.LayoutType),
			style:          style,
			children:       children,
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
