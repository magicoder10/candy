package main

import (
	"candy/assets"
	"candy/graphics"
	"candy/screen"

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
		Bounds:      pixel.R(0, 0, screen.Width, screen.Height),
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

	app, err := screen.NewApp(ass, &px)
	if err != nil {
		panic(err)
	}
	err = app.Launch()
	if err != nil {
		panic(err)
	}
	graphics.StartMainLoop(24, &app, &px, &px)
}
