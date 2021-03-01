package ui

import (
	"image"
	"image/color"
	"image/draw"
	"time"

	"candy/assets"
	"candy/graphics"
	"candy/input"
	"candy/observability"
)

var _ graphics.Sprite = (*RenderEngine)(nil)

type RenderEngine struct {
	logger          *observability.Logger
	graphics        graphics.Graphics
	painter         *Painter
	rootComponent   Component
	rootConstraints Constraints
	compositeLayer  draw.Image
	batch           graphics.Batch
	updateDeps      *UpdateDeps
}

func (r *RenderEngine) Render(component Component) {
	r.rootComponent = component
	component.Init()
}

func (r *RenderEngine) Draw() {
	r.render()
	r.draw()
}

func (r *RenderEngine) Update(timeElapsed time.Duration) {
	if r.rootComponent == nil {
		return
	}
	r.rootComponent.Update(timeElapsed, Offset{}, r.updateDeps)
}

func (r *RenderEngine) HandleInput(in input.Input) {
	if r.rootComponent == nil {
		return
	}
	r.rootComponent.HandleInput(in, Offset{})
}

func (r *RenderEngine) render() {
	if r.rootComponent == nil {
		return
	}
	if !r.rootComponent.HasChanged() {
		return
	}

	r.generateLayout(r.rootComponent)

	r.compositeLayer = image.NewRGBA(image.Rectangle{
		Max: image.Point{X: r.rootConstraints.maxWidth, Y: r.rootConstraints.maxHeight},
	})

	black := color.RGBA{R: 0, G: 0, B: 0, A: 255}
	draw.Draw(r.compositeLayer, r.compositeLayer.Bounds(), &image.Uniform{C: black}, image.Point{}, draw.Src)

	r.rootComponent.Paint(r.painter, r.compositeLayer, Offset{})

	r.batch = r.graphics.StartNewBatch(r.compositeLayer)

	r.rootComponent.ResetChangeDetection()
}

func (r *RenderEngine) draw() {
	if r.compositeLayer == nil {
		return
	}
	imageBound := r.compositeLayer.Bounds()
	bound := graphics.Bound{
		X:      0,
		Y:      0,
		Width:  imageBound.Max.X,
		Height: imageBound.Max.Y,
	}
	r.batch.DrawSprite(0, 0, 0, bound, 1)
}

func (r *RenderEngine) generateLayout(component Component) {
	applyConstraints(component, r.rootConstraints)
}

func NewRenderEngine(
	rootConstraints Constraints,
	logger *observability.Logger,
	gh graphics.Graphics,
	assets *assets.Assets,
) *RenderEngine {
	return &RenderEngine{
		logger:          logger,
		graphics:        gh,
		painter:         &Painter{},
		rootConstraints: rootConstraints,
		updateDeps: &UpdateDeps{
			assets:   assets,
			fonts:    NewFonts(),
			graphics: gh,
		},
	}
}
