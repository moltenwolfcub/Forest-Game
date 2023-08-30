package assets

import (
	"encoding/json"
	"fmt"
	"image"
)

func LoadFont(file string) Font {
	bytes, err := fonts.ReadFile("fonts/" + file + ".json")
	if err != nil {
		panic(err)
	}

	var font Font
	if err = json.Unmarshal(bytes, &font); err != nil {
		panic(err)
	}

	font.Validate()
	font.UpdateResources()

	return font
}

var (
	DefaultFont Font = LoadFont("default")
)

// A monospace bitmap font
type Font struct {
	CharHeight     int        `json:"height"`
	CharWidth      int        `json:"width"`
	YShift         int        `json:"y-shift"`
	Rows           int        `json:"rows"`
	Cols           int        `json:"columns"`
	TexturePath    string     `json:"texture"`
	UnicodeMapping [][]string `json:"mapping"`

	Height int

	unicode2Coords map[rune]image.Point
	unknown        image.Point
}

func (f Font) Validate() {
	if l := len(f.UnicodeMapping); l != f.Rows {
		panic(fmt.Sprintf("Font has the wrong number of rows defined in the mapping than expected. Found %d, Expected %d.", l, f.Rows))
	}
	for i, row := range f.UnicodeMapping {
		if l := len(row); l != f.Cols {
			panic(fmt.Sprintf("Font has the wrong number of columns on row %d defined in the mapping than expected. Found %d, Expected %d.", i, l, f.Cols))
		}
	}

	if f.YShift > f.CharHeight {
		panic(fmt.Sprintf("Font's Y-Shift %d is lower than the character height %d.", f.YShift, f.CharHeight))
	}

	if f.CharHeight <= 0 {
		panic(fmt.Sprintf("Font needs a Character height of 1 or more. Not %d.", f.CharHeight))
	}
	if f.CharWidth <= 0 {
		panic(fmt.Sprintf("Font needs a Character width of 1 or more. Not %d.", f.CharWidth))
	}
}

func (f *Font) UpdateResources() {
	f.Height = f.CharHeight - f.YShift

	f.unicode2Coords = make(map[rune]image.Point)
	f.unknown = image.Pt(-1, -1)

	for y, rows := range f.UnicodeMapping {
		for x, char := range rows {
			if char == "\u0000" {
				continue
			}

			if char == "unknown" {
				if f.unknown.X != -1 {
					panic("Can't have multiple unknowns defined.")
				}

				f.unknown = image.Pt(x, y)
				continue
			}
			runeList := []rune(char)
			if len(runeList) > 1 {
				panic(fmt.Sprintf("%s isn't a valid unicode character or 'unknown'.", char))
			}

			currentRune := runeList[0]
			_, ok := f.unicode2Coords[currentRune]
			if ok {
				panic(fmt.Sprintf("Can't have multiple definitions of %v in a font.", currentRune))
			}

			f.unicode2Coords[currentRune] = image.Pt(x, y)
		}
	}
	if f.unknown.X == -1 {
		panic("Font requires an 'unknown' glyph for missing characters.")
	}

}

func (f Font) GetRuneCoords(r rune) image.Point {
	if f.unicode2Coords == nil {
		panic("Can't get a rune from a font before calling `UpdateResources` atleast once.")
	}

	runeCoords, ok := f.unicode2Coords[r]
	if !ok {
		return image.Pt(f.unknown.X*f.CharWidth, f.unknown.Y*f.CharHeight)
	}
	return image.Pt(runeCoords.X*f.CharWidth, runeCoords.Y*f.CharHeight)
}
