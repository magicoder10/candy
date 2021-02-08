package main

import (
	"time"

	"candy/assets"
	"candy/graphics"
	"candy/input"
	"candy/ui"
)

var _ graphics.Sprite = (*app)(nil)

type app struct {
	renderEngine *ui.RenderEngine
	component    ui.Component
}

func (a app) Draw() {
	a.renderEngine.Render(a.component)
	a.renderEngine.Draw()
}

func (a app) Update(timeElapsed time.Duration) {
	a.component.Update(timeElapsed)
}

func (a app) HandleInput(in input.Input) {
	a.component.HandleInput(in)
}

func newApp(ass *assets.Assets, renderEngine *ui.RenderEngine) *app {
	return &app{
		renderEngine: renderEngine,
		component: ui.NewBoxBuilder().
			Children([]ui.Component{
				ui.NewImageBuilder().
					ImagePath(ass, "test/image3.png").
					Build(),
				ui.NewImageBuilder().
					ImagePath(ass, "test/image1.jpg").
					Build(),
				ui.NewImageBuilder().
					ImagePath(ass, "test/image2.jpg").
					Build(),
			}).
			Build(),
	}
}
