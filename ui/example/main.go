package main

import (
	"candy/assets"
	"candy/graphics"
	"candy/observability"
	"candy/ui"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func main() {
	pixelgl.Run(run)
}

func run() {
	screenWidth := 1000
	screenHeight := 1000

	px, err := graphics.NewPixel(pixelgl.WindowConfig{
		Title:       "Candy",
		Icon:        nil,
		Bounds:      pixel.R(0, 0, float64(screenWidth), float64(screenHeight)),
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

	logger := observability.NewLogger(observability.Debug)
	rootConstraint := ui.NewScreenConstraint(screenWidth, screenHeight)
	renderEngine := ui.NewRenderEngine(&logger, &px, rootConstraint)

	app := newApp(&ass, &renderEngine)
	graphics.StartMainLoop(24, app, &px, &px)
}
