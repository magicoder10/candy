package ui

import (
	"image"
	"image/draw"
	"time"

	"candy/input"
)

type Component interface {
	HandleInput(in input.Input)
	Update(timeElapsed time.Duration)
	getLayout() layout
	getChildren() []Component
	computeLeafSize() size
	getSize() size
	setSize(size size)
	getChildrenOffset() []offset
	setChildrenOffset(childrenOffsets []offset)
	paint(painter *painter, destLayer draw.Image, offset offset)
	rasterize(painter *painter, renderLayer draw.Image, position image.Point)
}

type Style struct {
	Width  *int
	Height *int
}

type size struct {
	width  int
	height int
}

type offset struct {
	x int
	y int
	z int
}

type sharedComponent struct {
	layout         layout
	style          Style
	size           size
	childrenOffset []offset
	children       []Component
}

func (s sharedComponent) HandleInput(in input.Input) {
	return
}

func (s sharedComponent) Update(timeElapsed time.Duration) {
	return
}

func (s sharedComponent) getChildren() []Component {
	return s.children
}

func (s sharedComponent) getSize() size {
	return s.size
}

func (s sharedComponent) getChildrenOffset() []offset {
	return s.childrenOffset
}

func (s sharedComponent) getLayout() layout {
	return s.layout
}

func (s *sharedComponent) setSize(size size) {
	s.size = size
}

func (s *sharedComponent) setChildrenOffset(childrenOffsets []offset) {
	s.childrenOffset = childrenOffsets
}

func (s *sharedComponent) rasterize(painter *painter, renderLayer draw.Image, position image.Point) {

}
