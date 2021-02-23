package ui

import (
	"image"
	"image/color"
	"image/draw"
	"time"

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
	hasUpdate       bool
	batch           graphics.Batch
}

func (r *RenderEngine) Render(component Component) {
	r.rootComponent = component
}

func (r *RenderEngine) Draw() {
	r.render()
	r.draw()
}

func (r *RenderEngine) Update(timeElapsed time.Duration) {
	if r.rootComponent == nil {
		return
	}
	r.rootComponent.Update(timeElapsed)
}

func (r *RenderEngine) HandleInput(in input.Input) {
	if r.rootComponent == nil {
		return
	}
	r.rootComponent.HandleInput(in)
}

func (r *RenderEngine) render() {
	if !r.hasUpdate {
		return
	}

	r.generateLayout(r.rootComponent)

	r.compositeLayer = image.NewRGBA(image.Rectangle{
		Max: image.Point{X: r.rootConstraints.maxWidth, Y: r.rootConstraints.maxHeight},
	})
	transparent := color.RGBA{R: 0, G: 0, B: 0, A: 0}
	draw.Draw(r.compositeLayer, r.compositeLayer.Bounds(), &image.Uniform{C: transparent}, image.Point{}, draw.Src)

	r.rootComponent.Paint(r.painter, r.compositeLayer, Offset{})

	r.batch = r.graphics.StartNewBatch(r.compositeLayer)
	r.hasUpdate = false
}

func (r *RenderEngine) draw() {
	imageBound := r.compositeLayer.Bounds()
	bound := graphics.Bound{
		X:      0,
		Y:      0,
		Width:  imageBound.Max.X,
		Height: imageBound.Max.Y,
	}
	r.batch.DrawSprite(0, 0, 0, bound, 1)
}

func (r RenderEngine) generateLayout(component Component) {
	applyConstraints(component, r.rootConstraints)
}

func NewRenderEngine(
	logger *observability.Logger,
	g graphics.Graphics,
	rootConstraints Constraints,
) *RenderEngine {
	return &RenderEngine{
		logger:          logger,
		graphics:        g,
		painter:         &Painter{},
		rootConstraints: rootConstraints,
		hasUpdate:       true,
	}
}
