package main

import (
	"candy/assets"
	"candy/graphics"
	"candy/view"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func main() {
	pixelgl.Run(run)
}

func run() {
	px, err := graphics.NewPixel(pixelgl.WindowConfig{
		Title:       "Candy",
		Icon:        nil,
		Bounds:      pixel.R(0, 0, 1152, 830),
		VSync:       true,
		AlwaysOnTop: true,
	})
	if err != nil {
		panic(err)
	}

	ass, err := assets.LoadAssets("assets")
	if err != nil {
		panic(err)
	}

	gm := view.NewGameScreen(ass, &px)
	gm.Start()

	graphics.StartMainLoop(24, &gm, &px, &px)
}
