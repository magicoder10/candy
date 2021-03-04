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
		&ui.BoxProps{OnClick: func(target ui.Component) {
			fmt.Println("Box clicked")
		}},
		[]ui.Component{
			ui.NewButton(
				&ui.ButtonProps{
					Text: ptr.String("Click"),
					OnClick: func(target ui.Component) {
						fmt.Println("Button clicked")
					},
				},
				&demoButtonStyles,
			),
			ui.NewText(&ui.TextProps{Text: `
I guess we could discuss the implications of the phrase "meant to be."
That is if we wanted to drown ourselves in a sea of backwardly referential 
semantics and other mumbo-jumbo. Maybe such a discussion would result in the 
determination that "meant to be" is exactly as meaningless a phrase as it 
seems to be, and that none of us is actually meant to be doing anything at all. 
But that's my existential underpants underpinnings showing. It's the way the 
cookie crumbles. And now I want a cookie.
`}, &textStyles),
			ui.NewImage(&ui.ImageProps{ImagePath: "test/image3.png"}, nil),
			ui.NewImage(&ui.ImageProps{ImagePath: "test/image1.jpg"}, nil),
		},
		nil)}
}
