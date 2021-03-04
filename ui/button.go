package ui

import (
	"candy/ui/ptr"
)

type ButtonProps struct {
	Text    *string
	OnClick onClickHandler
}

func (b ButtonProps) getText() string {
	if b.Text == nil {
		return ""
	} else {
		return *b.Text
	}
}

var _ Component = (*Button)(nil)

type Button struct {
	SharedComponent
}

func NewButton(props *ButtonProps, statefulStyle *StatefulStyle) *Button {
	if props == nil {
		props = &ButtonProps{}
	}
	if statefulStyle == nil {
		statefulStyle = NewStatefulStyleWithLayout(InlineLayoutType)
	}
	normalStyle := statefulStyle.Styles[NormalState]
	if normalStyle.LayoutType == nil {
		normalStyle.LayoutType = LayoutTypePtr(InlineLayoutType)
	}
	if normalStyle.Background == nil {
		normalStyle.Background = &Background{Color: &Color{
			Red:   87,
			Green: 37,
			Blue:  229,
			Alpha: 255,
		}}
	}
	if normalStyle.Padding == nil {
		normalStyle.Padding = &EdgeSpacing{
			Top:    ptr.Int(6),
			Bottom: ptr.Int(6),
			Left:   ptr.Int(20),
			Right:  ptr.Int(20),
		}
	}

	textStatefulStyle := copyStatefulStyle(statefulStyle, false)
	textStatefulStyle.Styles[NormalState].Padding = nil
	textStatefulStyle.Styles[NormalState].Background = nil

	boxStatefulStyle := copyStatefulStyle(statefulStyle, false)

	normalStyle.Background = nil
	normalStyle.Padding = nil

	return &Button{
		SharedComponent: SharedComponent{
			Name:          "Button",
			StatefulStyle: statefulStyle,
			States:        map[State]struct{}{},
			Children: []Component{
				NewBox(
					&BoxProps{
						OnClick: props.OnClick,
						OnMouseEnter: func(target Component) {
							target.SetState(HoverState)
						},
						OnMouseLeave: func(target Component) {
							target.ResetState(HoverState)
						},
					}, []Component{
						//NewText(
						//	&TextProps{Text: props.getText()},
						//	textStatefulStyle),
					},
					boxStatefulStyle,
				),
			},
			childrenOffset: make([]Offset, 0),
		},
	}
}
