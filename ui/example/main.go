package main

import (
	"log"

	"candy/assets"
	"candy/graphics"
	"candy/observability"
	"candy/screen"
	"candy/ui"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	screenWidth := 1000
	screenHeight := 1000

	ass, err := assets.LoadAssets("public")
	if err != nil {
		panic(err)
	}

	eb := graphics.NewEbiten(false)

	logger := observability.NewLogger(observability.Debug)
	rootConstraint := ui.NewScreenConstraint(screenWidth, screenHeight)
	renderEngine := ui.NewRenderEngine(&logger, &eb, rootConstraint)

	app := newApp(&ass, &renderEngine)

	g := graphics.NewEbitenWindow(graphics.WindowConfig{
		Width:  screen.Width,
		Height: screen.Height,
		Title:  "Candy",
	}, app, 24, &eb)
	g.Init()

	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
