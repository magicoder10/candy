package ui

import (
	"image/draw"
	"time"

	"candy/input"
)

type Component interface {
	HandleInput(in input.Input)
	Update(timeElapsed time.Duration)
	getLayout() layout
	getChildren() []Component
	ComputeLeafSize() Size
	getSize() Size
	setSize(size Size)
	getChildrenOffset() []Offset
	setChildrenOffset(childrenOffsets []Offset)
	Paint(painter *Painter, destLayer draw.Image, offset Offset)
}

type Style struct {
	Width      *int
	Height     *int
	LayoutType LayoutType
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
	layout         layout
	style          Style
	size           Size
	childrenOffset []Offset
	children       []Component
}

func (s SharedComponent) HandleInput(in input.Input) {
	return
}

func (s SharedComponent) Update(timeElapsed time.Duration) {
	return
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

func (s *SharedComponent) setChildrenOffset(childrenOffsets []Offset) {
	s.childrenOffset = childrenOffsets
}

func (s *SharedComponent) ComputeLeafSize() Size {
	return Size{}
}
