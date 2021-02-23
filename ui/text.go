package ui

import (
	"image"
	"image/draw"
	"strings"
	"time"

	"github.com/golang/freetype/truetype"
)

type TextProps struct {
	Text string
}

var _ Component = (*Text)(nil)

type Text struct {
	SharedComponent
	props        TextProps
	font         *truetype.Font
	fontFace     *fontFace
	lines        []string
	shouldUpdate bool
}

func (t *Text) ComputeLeafSize(constraints Constraints) Size {
	if len(t.props.Text) == 0 || t.fontFace == nil {
		return Size{width: 0, height: 0}
	}

	t.breakLines(constraints)

	bottomOffset := t.style.FontStyle.Size / 5 * 2
	height := len(t.lines)*t.style.FontStyle.LineHeight + bottomOffset
	return Size{width: constraints.maxWidth, height: height}
}

func (t *Text) Update(timeElapsed time.Duration) {
	if !t.shouldUpdate {
		return
	}
	t.shouldUpdate = false

	family, err := newFontFamily(t.style.FontStyle.Family)
	if err != nil {
		return
	}
	font, face, err := family.face(fontStyle{
		weight: t.style.FontStyle.Weight,
		italic: t.style.FontStyle.Italic,
	}, t.style.FontStyle.Size)
	if err != nil {
		return
	}
	t.font = font
	t.fontFace = &face
}

func (t *Text) breakLines(constraints Constraints) {
	var prevRune *rune
	drawEnd := 0

	runes := []rune(strings.ReplaceAll(t.props.Text, "\n", ""))
	line := make([]rune, 0)
	for _, currRune := range runes {
		runeSize, err := t.fontFace.getRuneSize(currRune)
		if err != nil {
			return
		}
		nextDrawEnd := drawEnd
		if prevRune != nil {
			// draw at the end of the previous rune
			nextDrawEnd -= t.fontFace.getKern(*prevRune, currRune)
		}

		// draw the current rune
		nextDrawEnd += runeSize.width
		if nextDrawEnd > constraints.maxWidth {
			// curr char should be placed on the next row
			t.lines = append(t.lines, string(line))
			line = make([]rune, 0)
			drawEnd = runeSize.width
		} else {
			drawEnd = nextDrawEnd
		}
		line = append(line, currRune)
		prevRune = &currRune
	}
	if len(line) > 0 {
		t.lines = append(t.lines, string(line))
	}
}

func (t Text) Paint(painter *Painter, destLayer draw.Image, offset Offset) {
	if len(t.lines) == 0 {
		return
	}
	contentLayer := image.NewRGBA(image.Rectangle{
		Max: image.Point{
			X: t.SharedComponent.size.width,
			Y: t.SharedComponent.size.height,
		},
	})

	for index, line := range t.lines {
		y := offset.y + index*t.style.FontStyle.LineHeight
		destPoint := image.Point{X: offset.x, Y: y}
		painter.drawString(
			contentLayer, destPoint,
			t.font, line,
			t.style.FontStyle.Size, t.style.FontStyle.Color,
		)

	}

	painter.drawImage(contentLayer, contentLayer.Bounds(), destLayer, image.Point{
		X: offset.x,
		Y: offset.y,
	})
}

func NewText(props *TextProps, style *Style) *Text {
	if props == nil {
		props = &TextProps{}
	}
	if style == nil {
		style = &Style{}
	}
	return &Text{
		shouldUpdate: true,
		SharedComponent: SharedComponent{
			style: *style,
		},
		props: *props,
	}
}
