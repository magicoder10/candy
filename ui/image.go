package ui

import (
	"image"
	"image/draw"

	"candy/assets"
)

var _ Component = (*Image)(nil)

type Image struct {
	path string
	sharedComponent
	ass   *assets.Assets
	image image.Image
}

func (i *Image) paint(painter *painter, destLayer draw.Image, offset offset) {
	contentLayer := image.NewRGBA(image.Rectangle{
		Max: image.Point{
			X: i.sharedComponent.size.width,
			Y: i.sharedComponent.size.height,
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

func (i Image) computeLeafSize() size {
	if i.image == nil {
		return size{}
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
	return size{
		width:  width,
		height: height,
	}
}

type ImageBuilder struct {
	imagePath string
	image     image.Image
	componentBuilder
}

func (i *ImageBuilder) ImagePath(ass *assets.Assets, path string) *ImageBuilder {
	i.imagePath = path
	i.image = ass.GetImage(path)
	return i
}

func (i ImageBuilder) Build() *Image {
	if i.style == nil {
		i.style = &Style{}
	}
	if i.layout == nil {
		i.layout = BoxLayout{}
	}
	if i.image == nil {
		i.image = image.NewRGBA(image.Rectangle{})
	}
	return &Image{
		path: i.imagePath,
		sharedComponent: sharedComponent{
			layout: i.layout,
			style:  *i.style,
		},
		image: i.image,
	}
}

func NewImageBuilder() *ImageBuilder {
	return &ImageBuilder{}
}
