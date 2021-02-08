package ui

import (
    "image"
    "image/draw"
)

type painter struct {
}

func (painter) drawImage(src image.Image, srcRect image.Rectangle, dest draw.Image, destPoint image.Point) {
    width := srcRect.Max.X - srcRect.Min.X
    height := srcRect.Max.Y - srcRect.Min.Y

    destRect := image.Rectangle{
        Min: destPoint,
        Max: image.Point{
            X: destPoint.X + width,
            Y: destPoint.Y + height,
        },
    }
    draw.Draw(dest, destRect, src, srcRect.Min, draw.Over)
}
