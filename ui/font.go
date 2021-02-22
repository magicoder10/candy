package ui

import (
	"fmt"
	"io/ioutil"

	"github.com/adrg/sysfont"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

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

func (f fontFamily) face(style fontStyle, fontSize int) (*truetype.Font, fontFace, error) {
	ft, err := f.getFont(style)
	if err != nil {
		return nil, fontFace{}, err
	}
	options := truetype.Options{Size: float64(fontSize), DPI: float64(f.dpi)}
	face := truetype.NewFace(ft, &options)
	return ft, fontFace{fontFace: face}, nil
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

func findFontFamily(fontFinder *sysfont.Finder, fontFamily string) (string, []*sysfont.Font) {
	matchedFont := fontFinder.Match(fontFamily)
	if matchedFont == nil {
		return "", []*sysfont.Font{}
	}
	fonts := make([]*sysfont.Font, 0)
	for _, fontInstalled := range fontFinder.List() {
		if matchedFont.Family == fontInstalled.Family {
			fonts = append(fonts, fontInstalled)
		}
	}
	return matchedFont.Family, fonts
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

func newFontFamily(fontFamilyName string) (fontFamily, error) {
	finder := sysfont.NewFinder(nil)
	familyName, fontsInFamily := findFontFamily(finder, fontFamilyName)
	if len(fontsInFamily) < 1 {
		return fontFamily{}, fmt.Errorf("font family not found: %s", fontFamilyName)
	}

	fonts, err := loadFonts(fontsInFamily)
	if err != nil {
		return fontFamily{}, err
	}
	return fontFamily{
		dpi:        72,
		fontFamily: familyName,
		fontFinder: finder,
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
