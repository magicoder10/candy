package main

import (
	"candy/assets"
	"candy/game"
	"candy/graphics"
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
		Bounds:      pixel.R(0, 0, 900, 736),
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

	gm := game.NewGame(ass, &px)
	gm.Start()

	graphics.StartMainLoop(24, &gm, &px, &px)
}
