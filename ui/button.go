package ui

import (
	"fmt"

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

func NewButton(props *ButtonProps, style *Style) *Button {
	if props == nil {
		props = &ButtonProps{}
	}
	if style == nil {
		style = &Style{
			LayoutType: (*LayoutType)(ptr.Int(int(InlineLayoutType))),
		}
	}
	if style.LayoutType == nil {
		style.LayoutType = (*LayoutType)(ptr.Int(int(InlineLayoutType)))
	}
	if style.Background == nil {
		style.Background = &Background{Color: &Color{
			Red:   87,
			Green: 37,
			Blue:  229,
			Alpha: 255,
		}}
	}
	if style.Padding == nil {
		style.Padding = &EdgeSpacing{
			Top:    ptr.Int(6),
			Bottom: ptr.Int(6),
			Left:   ptr.Int(20),
			Right:  ptr.Int(20),
		}
	}

	textStyle := copyStyle(style)
	textStyle.Padding = nil
	textStyle.Background = nil

	boxStyle := copyStyle(style)

	style.Background = nil
	style.Padding = nil

	return &Button{
		SharedComponent: SharedComponent{
			Name:   "Button",
			Layout: NewLayout(InlineLayoutType),
			Style:  style,
			Children: []Component{
				NewBox(
					&BoxProps{
						OnClick: props.OnClick,
						OnMouseEnter: func() {
							fmt.Println("Mouse enter")
						},
						OnMouseLeave: func() {
							fmt.Println("Mouse leave")
						},
					}, []Component{
						NewText(&TextProps{Text: props.getText()}, textStyle),
					},
					boxStyle,
				),
			},
			childrenOffset: make([]Offset, 0),
		},
	}
}

func copyStyle(src *Style) *Style {
	target := Style{}
	target.FontStyle = src.FontStyle
	target.LayoutType = src.LayoutType
	target.Padding = src.Padding
	target.Alignment = src.Alignment
	target.Background = src.Background
	target.Width = src.Width
	target.Height = src.Height
	return &target
}
