package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"

	"candy/screen"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/teamyapp/ui/assets"
	"github.com/teamyapp/ui/graphics"
	"github.com/teamyapp/ui/observability"
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

	eb := graphics.NewEbiten(true)

	ass, err := assets.LoadAssets("public")
	if err != nil {
		panic(err)
	}

	logger := observability.NewLogger(observability.Info)

	app, err := screen.NewApp(&logger, ass, &eb)
	if err != nil {
		panic(err)
	}
	err = app.Launch()
	if err != nil {
		panic(err)
	}

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
