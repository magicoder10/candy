package ui

import (
	"fmt"
	"io/ioutil"

	"github.com/adrg/sysfont"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

type Fonts struct {
	fontFinder *sysfont.Finder
	families   map[string]*fontFamily
}

func (f *Fonts) getFamily(familyName string) *fontFamily {
	f.loadFont(familyName)
	matchedFont := f.fontFinder.Match(familyName)
	return f.families[matchedFont.Family]
}

func (f *Fonts) loadFont(familyName string) {
	matchedFont := f.fontFinder.Match(familyName)
	if _, ok := f.families[matchedFont.Family]; ok {
		return
	}
	ff, err := newFontFamily(f.fontFinder, matchedFont.Family)
	if err != nil {
		return
	}
	f.families[matchedFont.Family] = &ff
}

func NewFonts() *Fonts {
	finder := sysfont.NewFinder(nil)
	return &Fonts{
		fontFinder: finder,
		families:   make(map[string]*fontFamily, 0),
	}
}

type fontFamily struct {
	dpi        int
	fontFinder *sysfont.Finder
	fontFamily string
	fonts      map[sysfont.Font]*truetype.Font
}

type fontStyle struct {
	weight string
	italic bool
}

func (f fontFamily) face(style fontStyle, fontSize int) (*truetype.Font, *fontFace, error) {
	ft, err := f.getFont(style)
	if err != nil {
		return nil, nil, err
	}
	options := truetype.Options{Size: float64(fontSize), DPI: float64(f.dpi)}
	face := truetype.NewFace(ft, &options)
	return ft, &fontFace{fontFace: face}, nil
}

func (f fontFamily) getFont(style fontStyle) (*truetype.Font, error) {
	matchedFont := f.fontFinder.Match(f.getQuery(style))
	if matchedFont == nil {
		return nil, fmt.Errorf("style not found:%s", style.toString())
	}
	return f.fonts[*matchedFont], nil
}

func (f fontFamily) getQuery(style fontStyle) string {
	return f.fontFamily + style.toString()
}

func findFontFamily(fontFinder *sysfont.Finder, family string) (string, []*sysfont.Font) {
	fonts := make([]*sysfont.Font, 0)
	for _, fontInstalled := range fontFinder.List() {
		if family == fontInstalled.Family {
			fonts = append(fonts, fontInstalled)
		}
	}
	return family, fonts
}

func loadFonts(fonts []*sysfont.Font) (map[sysfont.Font]*truetype.Font, error) {
	parsedFonts := make(map[sysfont.Font]*truetype.Font, 0)
	for _, f := range fonts {
		fontData, err := ioutil.ReadFile(f.Filename)
		if err != nil {
			return nil, err
		}
		parsedFont, err := truetype.Parse(fontData)
		if err != nil {
			return nil, err
		}
		parsedFonts[*f] = parsedFont
	}
	return parsedFonts, nil
}

func newFontFamily(fontFinder *sysfont.Finder, family string) (fontFamily, error) {
	familyName, fontsInFamily := findFontFamily(fontFinder, family)
	if len(fontsInFamily) < 1 {
		return fontFamily{}, fmt.Errorf("font family not found: %s", family)
	}

	fonts, err := loadFonts(fontsInFamily)
	if err != nil {
		return fontFamily{}, err
	}
	return fontFamily{
		dpi:        72,
		fontFamily: familyName,
		fontFinder: fontFinder,
		fonts:      fonts,
	}, nil
}

func (f fontStyle) toString() string {
	str := f.weight
	if f.italic {
		str += "Italic"
	}
	return str
}

type fontFace struct {
	fontFace font.Face
}

func (f fontFace) getRuneSize(rune rune) (Size, error) {
	_, advance, ok := f.fontFace.GlyphBounds(rune)
	if !ok {
		return Size{}, fmt.Errorf("rune not found:%c", rune)
	}
	metrics := f.fontFace.Metrics()
	return Size{width: advance.Round(), height: metrics.Height.Round()}, nil
}

func (f fontFace) getKern(firstRune rune, secondRune rune) int {
	return f.fontFace.Kern(firstRune, secondRune).Round()
}
