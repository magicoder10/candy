package main

import (
	"log"

	"candy/assets"
	"candy/graphics"
	"candy/observability"
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

	renderEngine := ui.NewRenderEngine(rootConstraint, &logger, &eb, &ass)
	renderEngine.Render(newApp())

	g := graphics.NewEbitenWindow(graphics.WindowConfig{
		Width:  screenWidth,
		Height: screenHeight,
		Title:  "Example",
	}, renderEngine, 24, &eb)
	g.Init()

	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
