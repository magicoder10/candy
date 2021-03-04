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
	if i.hasChanged {
		i.initContentLayer()
		painter.drawImage(i.image, i.image.Bounds(), i.contentLayer, image.Point{
			X: 0,
			Y: 0,
		})
	}

	painter.drawImage(i.contentLayer, i.contentLayer.Bounds(), destLayer, image.Point{
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

	style := i.getStyle()
	if style.Width != nil {
		width = *style.Width
	}
	height := imageBound.Max.Y - imageBound.Min.Y
	if style.Height != nil {
		height = *style.Height
	}
	return Size{
		width:  width,
		height: height,
	}
}

func (i *Image) Update(timeElapsed time.Duration, screenOffset Offset, deps *UpdateDeps) {
	i.SharedComponent.Update(timeElapsed, screenOffset, deps)

	if i.props.ImagePath != i.prevImagePath {
		if len(i.props.ImagePath) > 0 {
			i.image = deps.assets.GetImage(i.props.ImagePath)
		}
		i.hasChanged = true
		i.prevImagePath = i.props.ImagePath
	}

	if i.StatefulStyle.HasChanged() {
		i.hasChanged = true
	}
}

func NewImage(props *ImageProps, statefulStyle *StatefulStyle) *Image {
	if props == nil {
		props = &ImageProps{}
	}
	if statefulStyle == nil {
		statefulStyle = NewStatefulStyle()
	}
	return &Image{
		props: *props,
		SharedComponent: SharedComponent{
			Name:          "Image",
			States:        map[State]struct{}{},
			StatefulStyle: statefulStyle,
		},
	}
}
