package ui

import (
	"image"
	"image/draw"
	"sort"
	"time"

	"candy/assets"
	"candy/input"
)

type UpdateDeps struct {
	assets *assets.Assets
	fonts  *Fonts
}

type changeDetector interface {
	MarkChanged()
	HasChanged() bool
	ResetChangeDetection()
}

type lifeCycle interface {
	Init()
	Destroy()
}

type Component interface {
	GetName() string
	Update(timeElapsed time.Duration, deps *UpdateDeps)
	Paint(painter *Painter, destLayer draw.Image, offset Offset)
	ComputeLeafSize(constraints Constraints) Size
	changeDetector
	lifeCycle

	HandleInput(in input.Input, screenOffset Offset)
	getLayout() Layout
	getChildren() []Component
	getSize() Size
	setSize(size Size)
	getStyle() Style
	getChildrenOffset() []Offset
	setChildrenOffset(childrenOffsets []Offset)
}

type Size struct {
	width  int
	height int
}

type Offset struct {
	x int
	y int
	z int
}

func (o Offset) Add(relativeOffset Offset) Offset {
	o.x += relativeOffset.x
	o.y += relativeOffset.y
	o.z += relativeOffset.z
	return o
}

type SharedComponent struct {
	Name           string
	Layout         Layout
	Style          *Style
	size           Size
	childrenOffset []Offset
	contentLayer   draw.Image
	Children       []Component
	events         Events
	hasChanged     bool
}

func (s *SharedComponent) Init() {
	s.MarkChanged()
	return
}
func (s *SharedComponent) Destroy() {}

func (s *SharedComponent) Paint(painter *Painter, destLayer draw.Image, offset Offset) {
	if s.hasChanged || s.contentLayer == nil {
		s.contentLayer = image.NewRGBA(image.Rectangle{
			Max: image.Point{
				X: s.size.width,
				Y: s.size.height,
			},
		})
		if s.Style != nil && s.Style.Background != nil {
			s.Style.Background.Paint(painter, s.size, s.contentLayer)
		}

		sortedChildren := Children{
			children:       s.Children,
			childrenOffset: s.childrenOffset,
		}
		sort.Sort(&sortedChildren)

		for index, child := range sortedChildren.children {
			childOffset := sortedChildren.childrenOffset[index]
			child.Paint(painter, s.contentLayer, childOffset)
		}
	}

	painter.drawImage(s.contentLayer, image.Rectangle{
		Min: image.Point{},
		Max: s.contentLayer.Bounds().Max,
	}, destLayer, image.Point{
		X: offset.x,
		Y: offset.y,
	})
}

func (s *SharedComponent) HandleInput(in input.Input, screenOffset Offset) {
	for index, child := range s.Children {
		relativeOffset := s.childrenOffset[index]
		child.HandleInput(in, screenOffset.Add(relativeOffset))
	}

	switch in.Action {
	case input.SinglePress:
		switch in.Device {
		case input.MouseLeftButton:
			rect := s.BoundingBox(screenOffset)
			if in.CursorPosition.In(rect) {
				s.events.tryOnClick()
			}
		}
	}
}

func (s *SharedComponent) BoundingBox(offset Offset) image.Rectangle {
	return image.Rect(offset.x, offset.y, offset.x+s.size.width, offset.y+s.size.height)
}

func (s *SharedComponent) Update(timeElapsed time.Duration, deps *UpdateDeps) {
	for _, child := range s.Children {
		child.Update(timeElapsed, deps)
		if child.HasChanged() {
			s.hasChanged = true
		}
	}
	if s.Style != nil {
		s.Style.Update(deps)
		if s.Style.hasChanged {
			s.hasChanged = true
		}
	}
}

func (s *SharedComponent) MarkChanged() {
	s.hasChanged = true
}

func (s SharedComponent) HasChanged() bool {
	return s.hasChanged
}

func (s *SharedComponent) ResetChangeDetection() {
	for _, child := range s.Children {
		child.ResetChangeDetection()
	}

	if s.Style != nil {
		s.Style.ResetChangeDetection()
	}
	s.hasChanged = false
}

func (s SharedComponent) GetName() string {
	return s.Name
}

func (s SharedComponent) getChildren() []Component {
	return s.Children
}

func (s SharedComponent) getSize() Size {
	return s.size
}

func (s SharedComponent) getChildrenOffset() []Offset {
	return s.childrenOffset
}

func (s SharedComponent) getLayout() Layout {
	return s.Layout
}

func (s *SharedComponent) setSize(size Size) {
	s.size = size
}

func (s SharedComponent) getStyle() Style {
	if s.Style == nil {
		return Style{}
	}
	return *s.Style
}

func (s *SharedComponent) setChildrenOffset(childrenOffsets []Offset) {
	s.childrenOffset = childrenOffsets
}

func (s *SharedComponent) ComputeLeafSize(_ Constraints) Size {
	return Size{}
}
