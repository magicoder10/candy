package ui

import (
	"image"
	"image/color"
	"image/draw"
	"time"

	"candy/assets"
	"candy/input"
	"candy/observability"
	"candy/ui/graphics"
)

var _ graphics.Sprite = (*RenderEngine)(nil)

type RenderEngine struct {
	logger          *observability.Logger
	canvas          graphics.Canvas
	painter         *Painter
	rootComponent   Component
	rootConstraints Constraints
	compositeLayer  draw.Image
	bound           graphics.Bound
	updateDeps      *UpdateDeps
}

func (r *RenderEngine) Render(component Component) {
	r.rootComponent = component
	component.Init()
}

func (r *RenderEngine) Update(timeElapsed time.Duration) {
	if r.rootComponent == nil {
		return
	}
	r.rootComponent.Update(timeElapsed, Offset{}, r.updateDeps)

	if !r.rootComponent.HasChanged() {
		return
	}
	r.generateLayout(r.rootComponent)

	black := color.RGBA{R: 0, G: 0, B: 0, A: 255}
	draw.Draw(r.compositeLayer, r.compositeLayer.Bounds(), &image.Uniform{C: black}, image.Point{}, draw.Src)
	r.rootComponent.Paint(r.painter, r.compositeLayer, Offset{})
	r.rootComponent.ResetChangeDetection()
	r.canvas.OverrideContent(r.compositeLayer)
}

func (r *RenderEngine) HandleInput(in input.Input) {
	if r.rootComponent == nil {
		return
	}
	r.rootComponent.HandleInput(in, Offset{})
}

func (r *RenderEngine) generateLayout(component Component) {
	applyConstraints(component, r.rootConstraints)
}

func NewRenderEngine(
	rootConstraints Constraints,
	logger *observability.Logger,
	assets *assets.Assets,
	gh graphics.Graphics,
	canvas graphics.Canvas,
) *RenderEngine {
	compositeLayer := image.NewRGBA(image.Rectangle{
		Max: image.Point{X: rootConstraints.maxWidth, Y: rootConstraints.maxHeight},
	})
	imageBound := compositeLayer.Bounds()
	bound := graphics.Bound{
		X:      0,
		Y:      0,
		Width:  imageBound.Max.X,
		Height: imageBound.Max.Y,
	}
	return &RenderEngine{
		logger:          logger,
		canvas:          canvas,
		painter:         &Painter{},
		rootConstraints: rootConstraints,
		compositeLayer:  compositeLayer,
		bound:           bound,
		updateDeps: &UpdateDeps{
			assets:   assets,
			fonts:    NewFonts(),
			graphics: gh,
		},
	}
}
