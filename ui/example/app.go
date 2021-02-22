package main

import (
	"candy/assets"
	"candy/ui"
)

var _ ui.Component = (*app)(nil)

type app struct {
	*ui.Box
}

func newApp(assets *assets.Assets) *app {
	return &app{ui.NewBox(nil, []ui.Component{
		ui.NewImage(assets, &ui.ImageProps{ImagePath: "test/image3.png"}, nil),
		ui.NewImage(assets, &ui.ImageProps{ImagePath: "test/image1.jpg"}, nil),
		ui.NewImage(assets, &ui.ImageProps{ImagePath: "test/image2.jpg"}, nil),
	},
		nil)}
}
