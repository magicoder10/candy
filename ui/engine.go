package ui

import (
    "image"
    "image/color"
    "image/draw"

    "candy/graphics"
    "candy/observability"
)

type RenderEngine struct {
    logger          *observability.Logger
    graphics        graphics.Graphics
    painter         *painter
    rootConstraints Constraints
    compositeLayer  draw.Image
    hasUpdate       bool
    batch           graphics.Batch
}

func (r *RenderEngine) Render(component Component) {
    if !r.hasUpdate {
        return
    }

    r.generateLayout(component)

    r.compositeLayer = image.NewRGBA(image.Rectangle{
        Max: image.Point{X: r.rootConstraints.maxWidth, Y: r.rootConstraints.maxHeight},
    })
    transparent := color.RGBA{R: 0, G: 0, B: 0, A: 0}
    draw.Draw(r.compositeLayer, r.compositeLayer.Bounds(), &image.Uniform{C: transparent}, image.Point{}, draw.Src)

    component.paint(r.painter, r.compositeLayer, offset{})

    r.batch = r.graphics.StartNewBatch(r.compositeLayer)
    r.hasUpdate = false
}

func (r RenderEngine) Draw() {
    imageBound := r.compositeLayer.Bounds()
    bound := graphics.Bound{
        X:      0,
        Y:      0,
        Width:  imageBound.Max.X,
        Height: imageBound.Max.Y,
    }
    r.batch.DrawSprite(0, 0, 0, bound, 1)
    r.batch.RenderBatch()
}

func (r RenderEngine) generateLayout(component Component) {
    applyConstraints(component, r.rootConstraints)
}

func NewRenderEngine(
    logger *observability.Logger,
    g graphics.Graphics,
    rootConstraints Constraints,
) RenderEngine {
    return RenderEngine{
        logger:          logger,
        graphics:        g,
        painter:         &painter{},
        rootConstraints: rootConstraints,
        hasUpdate:       true,
    }
}
