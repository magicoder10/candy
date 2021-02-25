package ui

import (
	"image"
	"image/draw"
	"time"
)

type ImageProps struct {
	ImagePath string
}

var _ Component = (*Image)(nil)

type Image struct {
	SharedComponent
	props         ImageProps
	prevImagePath string
	image         image.Image
}

func (i *Image) Paint(painter *Painter, destLayer draw.Image, offset Offset) {
	if i.image == nil {
		return
	}
	contentLayer := image.NewRGBA(image.Rectangle{
		Max: image.Point{
			X: i.SharedComponent.size.width,
			Y: i.SharedComponent.size.height,
		},
	})
	painter.drawImage(i.image, i.image.Bounds(), contentLayer, image.Point{
		X: 0,
		Y: 0,
	})
	painter.drawImage(contentLayer, contentLayer.Bounds(), destLayer, image.Point{
		X: offset.x,
		Y: offset.y,
	})
}

func (i Image) ComputeLeafSize(_ Constraints) Size {
	if i.image == nil {
		return Size{}
	}
	imageBound := i.image.Bounds()
	width := imageBound.Max.X - imageBound.Min.X
	if i.style.Width != nil {
		width = *i.style.Width
	}
	height := imageBound.Max.Y - imageBound.Min.Y
	if i.style.Height != nil {
		height = *i.style.Height
	}
	return Size{
		width:  width,
		height: height,
	}
}

func (i *Image) Update(timeElapsed time.Duration, deps *UpdateDeps) {
	i.SharedComponent.Update(timeElapsed, deps)

	if i.props.ImagePath != i.prevImagePath {
		if len(i.props.ImagePath) > 0 {
			i.image = deps.assets.GetImage(i.props.ImagePath)
		}
		i.hasChanged = true
		i.prevImagePath = i.props.ImagePath
	}

	if i.style.hasChanged {
		i.hasChanged = true
	}
}

func NewImage(props *ImageProps, style *Style) *Image {
	if props == nil {
		props = &ImageProps{}
	}
	if style == nil {
		style = &Style{}
	}
	return &Image{
		props: *props,
		SharedComponent: SharedComponent{
			name:  "Image",
			style: style,
		},
	}
}
