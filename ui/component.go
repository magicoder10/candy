package ui

import (
	"image"
	"image/draw"
	"sort"
	"time"

	"candy/assets"
	"candy/graphics"
	"candy/input"
)

type UpdateDeps struct {
	assets   *assets.Assets
	fonts    *Fonts
	graphics graphics.Graphics
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
	SetParent(parent Component)
	GetParent() Component
	SetState(state State)
	ResetState(state State)
	Update(timeElapsed time.Duration, screenOffset Offset, deps *UpdateDeps)
	Paint(painter *Painter, destLayer draw.Image, offset Offset)
	ComputeLeafSize(constraints Constraints) Size
	changeDetector
	lifeCycle
	HandleInput(in input.Input, screenOffset Offset)
	getLayout() Layout
	getChildren() []Component
	getSize() Size
	setSize(size Size)
	getStyle() *Style
	getChildrenOffset() []Offset
	setChildrenOffset(childrenOffsets []Offset)
}

type State int

const (
	NormalState State = iota
	HoverState
	FocusState
)

var statePriority = []State{
	NormalState,
	HoverState,
	FocusState,
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
	Name               string
	States             map[State]struct{}
	StatefulStyle      *StatefulStyle
	style              *Style
	size               Size
	childrenOffset     []Offset
	contentLayer       draw.Image
	parent             Component
	Children           []Component
	events             Events
	hasChanged         bool
	prevCursorPosition *image.Point
}

func (s *SharedComponent) Init() {
	s.MarkChanged()
	s.StatefulStyle.Init()
	for _, child := range s.Children {
		child.Init()
	}
}
func (s *SharedComponent) Destroy() {}

func (s *SharedComponent) Paint(painter *Painter, destLayer draw.Image, offset Offset) {
	if s.hasChanged {
		s.initContentLayer()

		style := s.getStyle()
		if style.Background != nil {
			style.Background.Paint(painter, s.size, s.contentLayer)
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

	if s.contentLayer == nil {
		return
	}

	painter.drawImage(s.contentLayer, image.Rectangle{
		Min: image.Point{},
		Max: s.contentLayer.Bounds().Max,
	}, destLayer, image.Point{
		X: offset.x,
		Y: offset.y,
	})
}

func (s *SharedComponent) initContentLayer() {
	s.contentLayer = image.NewRGBA(image.Rectangle{
		Max: image.Point{
			X: s.size.width,
			Y: s.size.height,
		},
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
				s.events.tryOnClick(s)
			}
		}
	}
}

func (s *SharedComponent) BoundingBox(offset Offset) image.Rectangle {
	return image.Rect(offset.x, offset.y, offset.x+s.size.width, offset.y+s.size.height)
}

func (s *SharedComponent) Update(timeElapsed time.Duration, screenOffset Offset, deps *UpdateDeps) {
	for index, child := range s.Children {
		offset := screenOffset
		if index < len(s.childrenOffset) {
			relativeOffset := s.childrenOffset[index]
			offset = screenOffset.Add(relativeOffset)
		}
		child.Update(timeElapsed, offset, deps)
		if child.HasChanged() {
			s.hasChanged = true
		}
	}
	style := s.getStyle()
	style.Update(deps)
	if s.StatefulStyle.HasChanged() {
		s.hasChanged = true
	}

	cursorPos := deps.graphics.GetCursorPosition()
	if s.prevCursorPosition != nil && !cursorPos.Eq(*s.prevCursorPosition) {
		s.events.tryOnMouseMove(s, cursorPos)

		boundingBox := s.BoundingBox(screenOffset)
		if s.prevCursorPosition.In(boundingBox) && !cursorPos.In(boundingBox) {
			s.events.tryOnMouseLeave(s)
		}
		if !s.prevCursorPosition.In(boundingBox) && cursorPos.In(boundingBox) {
			s.events.tryOnMouseEnter(s)
		}
	}
	s.prevCursorPosition = &cursorPos
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

	s.StatefulStyle.ResetChangeDetection()
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
	style := s.getStyle()
	if style.LayoutType == nil {
		return NewLayout(BoxLayoutType)
	}
	return NewLayout(*style.LayoutType)
}

func (s *SharedComponent) setSize(size Size) {
	s.size = size
}

func (s *SharedComponent) getStyle() *Style {
	if s.HasChanged() {
		s.style = s.StatefulStyle.ComputeStyle(s.States)
	}
	if s.style == nil {
		return &Style{}
	}
	return s.style
}

func (s SharedComponent) GetParent() Component {
	return s.parent
}

func (s *SharedComponent) SetParent(parent Component) {
	s.parent = parent
}

func (s *SharedComponent) SetState(state State) {
	s.States[state] = struct{}{}
	s.MarkChanged()
}
func (s *SharedComponent) ResetState(state State) {
	delete(s.States, state)
	s.MarkChanged()
}

func (s *SharedComponent) setChildrenOffset(childrenOffsets []Offset) {
	s.childrenOffset = childrenOffsets
}

func (s *SharedComponent) ComputeLeafSize(_ Constraints) Size {
	return Size{}
}
