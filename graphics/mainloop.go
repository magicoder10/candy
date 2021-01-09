package graphics

import (
	"time"
)

func StartMainLoop(framesPerSeconds int64, sp Sprite, window Window, g Graphics) {
	milliPerUpdate := (time.Second / time.Duration(framesPerSeconds)).Milliseconds()

	prevTime := time.Now()
	var lag int64

	for !window.IsClosed() {
		now := time.Now()
		elapsed := now.Sub(prevTime)
		lag += elapsed.Milliseconds()
		prevTime = now

		inputs := window.PollEvents()
		for _, in := range inputs {
			sp.HandleInput(in)
		}

		for lag >= milliPerUpdate {
			sp.Update(elapsed)
			lag -= milliPerUpdate
		}
		g.Clear()
		sp.Draw()

		window.Redraw()
	}
}
