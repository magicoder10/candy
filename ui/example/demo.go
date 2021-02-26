package main

import (
	"fmt"

	"candy/ui"
	"candy/ui/ptr"
)

var _ ui.Component = (*demo)(nil)

type demo struct {
	*ui.Box
}

func newDemo() *demo {
	return &demo{ui.NewBox(
		&ui.BoxProps{OnClick: func() {
			fmt.Println("Box clicked")
		}},
		[]ui.Component{
			ui.NewButton(
				&ui.ButtonProps{
					Text: ptr.String("Click"),
					OnClick: func() {
						fmt.Println("Button clicked")
					},
				},
				&ui.Style{
					Background: &ui.Background{
						ImagePath: ptr.String("test/image2.jpg"),
					},
				},
			),
			ui.NewText(&ui.TextProps{Text: `
I guess we could discuss the implications of the phrase "meant to be."
That is if we wanted to drown ourselves in a sea of backwardly referential 
semantics and other mumbo-jumbo. Maybe such a discussion would result in the 
determination that "meant to be" is exactly as meaningless a phrase as it 
seems to be, and that none of us is actually meant to be doing anything at all. 
But that's my existential underpants underpinnings showing. It's the way the 
cookie crumbles. And now I want a cookie.
`}, &ui.Style{FontStyle: &ui.FontStyle{
				Family:     ptr.String("Source Code Pro"),
				Weight:     ptr.String("ExtraLight"),
				Italic:     ptr.Bool(false),
				Size:       ptr.Int(20),
				LineHeight: ptr.Int(24),
				Color: &ui.Color{
					Red:   255,
					Green: 255,
					Blue:  255,
					Alpha: 255,
				}}}),
			ui.NewImage(&ui.ImageProps{ImagePath: "test/image3.png"}, nil),
			ui.NewImage(&ui.ImageProps{ImagePath: "test/image1.jpg"}, nil),
		},
		nil)}
}