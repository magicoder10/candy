package ui

import (
	"image"
)

type onClickHandler func()

type onMouseMoveHandler func(cursorPosition image.Point)

type onMouseEnterHandler func()

type onMouseLeaveHandler func()

type Events struct {
	onClick      onClickHandler
	onMouseMove  onMouseMoveHandler
	onMouseEnter onMouseEnterHandler
	onMouseLeave onMouseLeaveHandler
}

func (e *Events) tryOnClick() {
	if e.onClick == nil {
		return
	}
	e.onClick()
}

func (e *Events) tryOnMouseMove(cursorPosition image.Point) {
	if e.onMouseMove == nil {
		return
	}
	e.onMouseMove(cursorPosition)
}

func (e *Events) tryOnMouseLeave() {
	if e.onMouseLeave == nil {
		return
	}
	e.onMouseLeave()
}

func (e *Events) tryOnMouseEnter() {
	if e.onMouseEnter == nil {
		return
	}
	e.onMouseEnter()
}
