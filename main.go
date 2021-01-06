package main

import (
	"fmt"
	"time"

	"candy/assets"
	"candy/game"
	"candy/graphics"
	"candy/input"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

var milliPerUpdate = (time.Second / 60).Milliseconds()

func main() {
	pixelgl.Run(run)
}

func run() {
	ass, err := assets.LoadAssets("assets")
	if err != nil {
		panic(err)
	}

	cfg := pixelgl.WindowConfig{
		Title:       "Candy",
		Icon:        nil,
		Bounds:      pixel.R(0, 0, 900, 736),
		VSync:       true,
		AlwaysOnTop: true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	gh := graphics.NewPixel(win)
	inputs := make(chan input.Input)

	gm := game.NewGame(ass, &gh, inputs)
	gm.Start(func() {
		fmt.Println("Game Over!")
	})

	prevTime := time.Now()
	var lag int64

	for !win.Closed() {
		now := time.Now()
		elapsed := now.Sub(prevTime)
		lag += elapsed.Milliseconds()
		prevTime = now

		pollEvents(win, inputs)

		for lag >= milliPerUpdate {
			gm.Update(elapsed)
			lag -= milliPerUpdate
		}
		gh.Clear()
		gm.Draw()

		win.Update()
	}
}

func pollEvents(win *pixelgl.Window, inputs chan<- input.Input) {
	if win.Pressed(pixelgl.KeyLeft) {
		inputs <- input.Input{
			Action: input.Press,
			Device: input.LeftArrowKey,
		}
	} else if win.Pressed(pixelgl.KeyRight) {
		inputs <- input.Input{
			Action: input.Press,
			Device: input.RightArrowKey,
		}
	} else if win.Pressed(pixelgl.KeyUp) {
		inputs <- input.Input{
			Action: input.Press,
			Device: input.UpArrowKey,
		}
	} else if win.Pressed(pixelgl.KeyDown) {
		inputs <- input.Input{
			Action: input.Press,
			Device: input.DownArrayKey,
		}
	} else if win.JustReleased(pixelgl.KeyLeft) {
		inputs <- input.Input{
			Action: input.Release,
			Device: input.LeftArrowKey,
		}
	} else if win.JustReleased(pixelgl.KeyRight) {
		inputs <- input.Input{
			Action: input.Release,
			Device: input.RightArrowKey,
		}
	} else if win.JustReleased(pixelgl.KeyUp) {
		inputs <- input.Input{
			Action: input.Release,
			Device: input.UpArrowKey,
		}
	} else if win.JustReleased(pixelgl.KeyDown) {
		inputs <- input.Input{
			Action: input.Release,
			Device: input.DownArrayKey,
		}
	}
}
