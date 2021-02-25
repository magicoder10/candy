package ui

type Changeable interface {
	hasChanged() bool
}
