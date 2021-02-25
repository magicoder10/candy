package ui

type onClickHandler func()

type Events struct {
	onClick onClickHandler
}

func (e *Events) tryOnClick() {
	if e.onClick == nil {
		return
	}
	e.onClick()
}
