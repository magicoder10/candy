package ui

import (
	"image"
	"image/draw"
	"math"
	"strings"
	"time"
)

const defaultFontFamily = "Arial"
const defaultFontWeight = "Regular"
const defaultFontFontSize = 16

type TextProps struct {
	Text string
}

var _ Component = (*Text)(nil)

type Text struct {
	SharedComponent
	props    TextProps
	prevText string
	lines    []string
}

func (t *Text) ComputeLeafSize(constraints Constraints) Size {
	if !t.hasChanged {
		return t.size
	}

	lineWidth := t.breakIntoLines(constraints)
	bottomOffset := *t.Style.FontStyle.Size / 5 * 2
	height := len(t.lines)*(*t.Style.FontStyle.LineHeight) + bottomOffset
	return Size{width: lineWidth, height: height}
}

func (t *Text) Update(_ time.Duration, _ Offset, deps *UpdateDeps) {
	if t.prevText != t.props.Text {
		t.hasChanged = true
		t.prevText = t.props.Text
	}

	t.Style.Update(deps)
	if t.Style.hasChanged {
		t.hasChanged = true
	}
}

func (t *Text) breakIntoLines(constraints Constraints) int {
	var prevRune *rune
	maxLineWidth := 0

	drawEnd := 0

	fontFace := t.Style.FontStyle.fontFace

	runes := []rune(strings.ReplaceAll(t.props.Text, "\n", ""))
	line := make([]rune, 0)
	for _, currRune := range runes {
		runeSize, err := fontFace.getRuneSize(currRune)
		if err != nil {
			return 0
		}
		nextDrawEnd := drawEnd
		if prevRune != nil {
			// draw at the end of the previous rune
			nextDrawEnd -= fontFace.getKern(*prevRune, currRune)
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
		maxLineWidth = int(math.Max(float64(maxLineWidth), float64(drawEnd)))

		line = append(line, currRune)
		prevRune = &currRune
	}
	if len(line) > 0 {
		t.lines = append(t.lines, string(line))
	}
	return maxLineWidth
}

func (t Text) Paint(painter *Painter, destLayer draw.Image, offset Offset) {
	if !t.hasChanged {
		return
	}

	contentLayer := image.NewRGBA(image.Rectangle{
		Max: image.Point{
			X: t.SharedComponent.size.width,
			Y: t.SharedComponent.size.height,
		},
	})

	for index, line := range t.lines {
		y := index * (*t.Style.FontStyle.LineHeight)
		destPoint := image.Point{X: 0, Y: y}
		painter.drawString(
			contentLayer, destPoint,
			t.Style.FontStyle.font, line,
			*t.Style.FontStyle.Size, *t.Style.FontStyle.Color,
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
		SharedComponent: SharedComponent{
			Name:  "Text",
			Style: style,
		},
		props: *props,
	}
}
