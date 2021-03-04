package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"

	"candy/assets"
	"candy/ui/graphics"
	"candy/observability"
	"candy/ui"

	"github.com/hajimehoshi/ebiten/v2"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var memprofile = flag.String("memprofile", "", "write memory profile to this file")

func main() {
	flag.Parse()
	// Enable CPU profiling
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	// Enable memory profiling
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		pprof.WriteHeapProfile(f)
		return
	}

	screenWidth := 1152
	screenHeight := 830

	ass, err := assets.LoadAssets("public")
	if err != nil {
		panic(err)
	}

	eb := graphics.NewEbiten(false)
	ebitenCanvas := eb.NewCanvas(screenWidth, screenHeight)

	logger := observability.NewLogger(observability.Debug)
	rootConstraint := ui.NewScreenConstraint(screenWidth, screenHeight)

	renderEngine := ui.NewRenderEngine(rootConstraint, &logger, &ass, &eb, ebitenCanvas)

	app, err := NewApp(&logger, ass)
	if err != nil {
		panic(err)
	}
	renderEngine.Render(app)

	g := graphics.NewEbitenWindow(&graphics.WindowConfig{
		Width:  screenWidth,
		Height: screenHeight,
		Title:  "Example",
	}, renderEngine, 24, ebitenCanvas)
	g.Init()

	if err = ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
