package graphics

import (
	"time"
)

func StartMainLoop(framesPerSeconds int64, sp Sprite, window Window, g Graphics) {
	nanoPerUpdate := time.Second.Nanoseconds() / framesPerSeconds
	updateTime := time.Duration(nanoPerUpdate)

	prevTime := time.Now()
	var lag int64

	for !window.IsClosed() {
		now := time.Now()
		elapsed := now.Sub(prevTime)
		lag += elapsed.Nanoseconds()
		prevTime = now

		inputs := window.PollEvents()
		for _, in := range inputs {
			sp.HandleInput(in)
		}

		for lag >= nanoPerUpdate {
			sp.Update(updateTime)
			lag -= nanoPerUpdate
		}
		g.Clear()
		sp.Draw()

		window.Redraw()
	}
}
