package graphics

import (
	"time"
)

func StartMainLoop(framesPerSeconds int64, sp Sprite, window Window, g Graphics) {
	nanoPerUpdate := time.Second.Nanoseconds() / framesPerSeconds

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

		fullLen := time.Duration(lag)
		for lag >= nanoPerUpdate {
			sp.Update(fullLen)
			lag -= nanoPerUpdate
		}
		g.Clear()
		sp.Draw()

		window.Redraw()
	}
}
