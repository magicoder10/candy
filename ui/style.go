package ui

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"reflect"

	"candy/ui/ptr"

	"github.com/golang/freetype/truetype"
)

type Style struct {
	Width      *int
	Height     *int
	LayoutType *LayoutType
	FontStyle  *FontStyle
	Padding    *EdgeSpacing
	Margin     *EdgeSpacing
	Alignment  *Alignment
	Background *Background
	hasChanged bool
}

func (s Style) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return ""
	}
	return fmt.Sprintf(string(buf))
}

type StatefulStyle struct {
	Styles map[State]*Style
}

func (s *StatefulStyle) ComputeStyle(states map[State]struct{}) *Style {
	// TODO
	style := copyStyle(s.Styles[NormalState], true)
	value := reflect.ValueOf(style).Elem()
	for _, state := range statePriority {
		if _, ok := states[state]; !ok {
			continue
		}
		stateStyle, ok := s.Styles[state]
		if !ok {
			continue
		}
		stateValue := reflect.ValueOf(stateStyle).Elem()
		styleType := stateValue.Type()

		for i := 0; i < styleType.NumField(); i++ {
			valueField := value.Field(i)
			if !valueField.CanSet() {
				continue
			}
			stateValueField := stateValue.Field(i)
			if stateValueField.IsNil() {
				continue
			}
			valueField.Set(stateValueField)
		}
	}
	return style
}

func (s *StatefulStyle) ResetChangeDetection() {
	for _, style := range s.Styles {
		style.ResetChangeDetection()
	}
}

func (s *StatefulStyle) Update(deps *UpdateDeps) {
	for _, style := range s.Styles {
		style.Update(deps)
	}
}

func (s *StatefulStyle) HasChanged() bool {
	for _, style := range s.Styles {
		if style.hasChanged {
			return true
		}
	}
	return false
}

func (s *StatefulStyle) Init() {
	for _, style := range s.Styles {
		style.hasChanged = true
	}
}

func NewStatefulStyle() *StatefulStyle {
	return &StatefulStyle{Styles: map[State]*Style{
		NormalState: {},
	}}
}

func NewStatefulStyleWithLayout(layoutType LayoutType) *StatefulStyle {
	return &StatefulStyle{Styles: map[State]*Style{
		NormalState: {
			LayoutType: LayoutTypePtr(layoutType),
		},
	}}
}

func (s *Style) Update(deps *UpdateDeps) {
	if s.Background != nil {
		s.Background.Update(deps.assets)
		if s.Background.hasChanged {
			s.hasChanged = true
		}
	}
	if s.FontStyle == nil {
		s.FontStyle = &FontStyle{}
	}
	s.FontStyle.Update(deps.fonts)
	if s.FontStyle.hasChanged {
		s.hasChanged = true
	}
}

func (s *Style) ResetChangeDetection() {
	if s.Background != nil {
		s.Background.ResetChangeDetection()
	}
	if s.FontStyle != nil {
		s.FontStyle.ResetChangeDetection()
	}
	s.hasChanged = false
}

func (s Style) GetFontStyle() FontStyle {
	if s.FontStyle == nil {
		return FontStyle{}
	}
	return *s.FontStyle
}

func (s Style) GetWidth() int {
	if s.Width == nil {
		return 0
	} else {
		return *s.Width
	}
}

func (s Style) GetHeight() int {
	if s.Height == nil {
		return 0
	} else {
		return *s.Height
	}
}

func (s Style) GetPadding() EdgeSpacing {
	if s.Padding == nil {
		return EdgeSpacing{}
	} else {
		return *s.Padding
	}
}

func (s Style) GetMargin() EdgeSpacing {
	if s.Margin == nil {
		return EdgeSpacing{}
	} else {
		return *s.Margin
	}
}

func (s Style) GetAlignment() Alignment {
	if s.Alignment == nil {
		return Alignment{}
	}
	return *s.Alignment
}

type FontStyle struct {
	Family     *string
	Weight     *string
	Italic     *bool
	LineHeight *int
	Color      *Color
	Size       *int

	prevFamily    string
	prevWeight    string
	prevItalic    bool
	preLineHeight int
	prevColor     Color
	prevSize      int
	hasChanged    bool

	family   *fontFamily
	font     *truetype.Font
	fontFace *fontFace
}

func (f *FontStyle) Update(fonts *Fonts) {
	if f.Family == nil {
		f.Family = ptr.String(defaultFontFamily)
	}
	if f.Size == nil {
		f.Size = ptr.Int(defaultFontFontSize)
	}
	if f.LineHeight == nil {
		f.LineHeight = ptr.Int(defaultFontFontSize)
	}
	if f.Color == nil {
		f.Color = &Color{
			Red:   255,
			Green: 255,
			Blue:  255,
			Alpha: 255,
		}
	}
	if f.Weight == nil {
		f.Weight = ptr.String(defaultFontWeight)
	}
	if f.Italic == nil {
		f.Italic = ptr.Bool(false)
	}

	if *f.Family != f.prevFamily {
		f.family = fonts.getFamily(*f.Family)
		f.hasChanged = true
		f.prevFamily = *f.Family
	}

	if f.hasChanged || *f.Weight != f.prevWeight ||
		*f.Italic != f.prevItalic ||
		*f.Size != f.prevSize {

		font, face, err := f.family.face(fontStyle{
			weight: *f.Weight,
			italic: *f.Italic,
		}, *f.Size)
		if err != nil {
			return
		}
		f.font, f.fontFace = font, face
		f.hasChanged = true

		f.prevSize = *f.Size
		f.prevWeight = *f.Weight
		f.prevItalic = *f.Italic
	}
}

func (f *FontStyle) ResetChangeDetection() {
	f.hasChanged = false
}

var _ color.Color = (*Color)(nil)

type Color struct {
	Red   uint8
	Green uint8
	Blue  uint8
	Alpha uint8
}

func (c Color) RGBA() (r, g, b, a uint32) {
	return highBits(c.Red), highBits(c.Green), highBits(c.Blue), highBits(c.Alpha)
}

func highBits(num uint8) uint32 {
	red := uint32(num)
	red |= red << 8
	return red
}

func (c Color) toUniform() *image.Uniform {
	return image.NewUniform(c)
}

func copyStatefulStyle(src *StatefulStyle, includeMargin bool) *StatefulStyle {
	newStatefulStyle := NewStatefulStyle()
	for state, style := range src.Styles {
		newStatefulStyle.Styles[state] = copyStyle(style, includeMargin)
	}
	return newStatefulStyle
}

func copyStyle(src *Style, includeMargin bool) *Style {
	target := Style{}
	target.FontStyle = src.FontStyle
	target.LayoutType = src.LayoutType
	target.Padding = src.Padding
	target.Alignment = src.Alignment
	target.Background = src.Background
	target.Width = src.Width
	target.Height = src.Height
	if includeMargin {
		target.Margin = src.Margin
	}
	return &target
}
