package ui

import (
	"image"
)

type onClickHandler func(target Component)

type onMouseMoveHandler func(target Component, cursorPosition image.Point)

type onMouseEnterHandler func(target Component)

type onMouseLeaveHandler func(target Component)

type Events struct {
	onClick      onClickHandler
	onMouseMove  onMouseMoveHandler
	onMouseEnter onMouseEnterHandler
	onMouseLeave onMouseLeaveHandler
}

func (e *Events) tryOnClick(target Component) {
	if e.onClick == nil {
		return
	}
	e.onClick(target)
}

func (e *Events) tryOnMouseMove(target Component, cursorPosition image.Point) {
	if e.onMouseMove == nil {
		return
	}
	e.onMouseMove(target, cursorPosition)
}

func (e *Events) tryOnMouseLeave(target Component) {
	if e.onMouseLeave == nil {
		return
	}
	e.onMouseLeave(target)
}

func (e *Events) tryOnMouseEnter(target Component) {
	if e.onMouseEnter == nil {
		return
	}
	e.onMouseEnter(target)
}
