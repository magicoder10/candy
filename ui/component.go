package ui

import (
	"image"
	"image/draw"
	"time"

	"candy/assets"
	"candy/input"
)

type UpdateDeps struct {
	assets *assets.Assets
	fonts  *Fonts
}

type Component interface {
	GetName() string
	HandleInput(in input.Input)
	handleInput(in input.Input, offset Offset)
	Update(timeElapsed time.Duration, deps *UpdateDeps)
	getLayout() layout
	getChildren() []Component
	ComputeLeafSize(constraints Constraints) Size
	getSize() Size
	setSize(size Size)
	getStyle() Style
	getChildrenOffset() []Offset
	setChildrenOffset(childrenOffsets []Offset)
	HasChanged() bool
	ResetChangeDetection()
	Paint(painter *Painter, destLayer draw.Image, offset Offset)
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

type SharedComponent struct {
	name           string
	layout         layout
	style          *Style
	size           Size
	childrenOffset []Offset
	children       []Component
	events         Events
	hasChanged     bool
}

func (s *SharedComponent) HandleInput(in input.Input) {
	s.handleInput(in, Offset{
		x: 0,
		y: 0,
		z: 0,
	})
}

func (s *SharedComponent) handleInput(in input.Input, offset Offset) {
	for index, child := range s.children {
		child.handleInput(in, s.childrenOffset[index])
	}

	switch in.Action {
	case input.SinglePress:
		switch in.Device {
		case input.MouseLeftButton:
			rect := s.BoundingBox(offset)
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
	for _, child := range s.children {
		child.Update(timeElapsed, deps)
		if child.HasChanged() {
			s.hasChanged = true
		}
	}
	if s.style != nil {
		s.style.Update(deps)
		if s.style.hasChanged {
			s.hasChanged = true
		}
	}
}

func (s SharedComponent) HasChanged() bool {
	return s.hasChanged
}

func (s *SharedComponent) ResetChangeDetection() {
	for _, child := range s.children {
		child.ResetChangeDetection()
	}

	if s.style != nil {
		s.style.ResetChangeDetection()
	}
	s.hasChanged = false
}

func (s SharedComponent) GetName() string {
	return s.name
}

func (s SharedComponent) getChildren() []Component {
	return s.children
}

func (s SharedComponent) getSize() Size {
	return s.size
}

func (s SharedComponent) getChildrenOffset() []Offset {
	return s.childrenOffset
}

func (s SharedComponent) getLayout() layout {
	return s.layout
}

func (s *SharedComponent) setSize(size Size) {
	s.size = size
}

func (s SharedComponent) getStyle() Style {
	if s.style == nil {
		return Style{}
	}
	return *s.style
}

func (s *SharedComponent) setChildrenOffset(childrenOffsets []Offset) {
	s.childrenOffset = childrenOffsets
}

func (s *SharedComponent) ComputeLeafSize(_ Constraints) Size {
	return Size{}
}
