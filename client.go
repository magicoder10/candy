package main

import (
	"candy/assets"
	"candy/env"
	"candy/graphics"
	"candy/observability"
	"candy/screen"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func main() {
	pixelgl.Run(run)
}

func run() {
	env.AutoLoad()

	px, err := graphics.NewPixel(pixelgl.WindowConfig{
		Title:       "Candy",
		Icon:        nil,
		Bounds:      pixel.R(0, 0, screen.Width, screen.Height),
		VSync:       true,
		AlwaysOnTop: false,
	})
	if err != nil {
		panic(err)
	}

	ass, err := assets.LoadAssets("assets")
	if err != nil {
		panic(err)
	}

	logger := observability.NewLogger(observability.Info)

	app, err := screen.NewApp(&logger, ass, &px)
	if err != nil {
		panic(err)
	}
	err = app.Launch()
	if err != nil {
		panic(err)
	}
	graphics.StartMainLoop(24, &app, &px, &px)
}