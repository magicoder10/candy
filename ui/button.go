package ui

import (
	"image/draw"

	"candy/ui/ptr"
)

type ButtonProps struct {
	Text *string
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

func (b Button) GetName() string {
	return "Button"
}

func (b Button) Paint(painter *Painter, destLayer draw.Image, offset Offset) {
	b.children[0].Paint(painter, destLayer, offset)
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
	return &Button{
		SharedComponent: SharedComponent{
			layout: newLayout(InlineLayoutType),
			children: []Component{
				NewBox([]Component{
					NewText(&TextProps{Text: props.getText()}, style),
				}, style),
			},
			childrenOffset: make([]Offset, 0),
		},
	}
}
