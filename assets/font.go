package assets

import (
	"encoding/json"
	"image"

	"github.com/moltenwolfcub/Forest-Game/errors"
)

func LoadFont(file string) (Font, error) {
	bytes, err := fonts.ReadFile("fonts/" + file + ".json")
	if err != nil {
		return Font{}, err
	}

	var font Font
	if err = json.Unmarshal(bytes, &font); err != nil {
		return Font{}, err
	}

	if err = font.Validate(); err != nil {
		return Font{}, err
	}
	if err = font.UpdateResources(); err != nil {
		return Font{}, err
	}

	return font, nil
}

var (
	DefaultFont Font
)

func init() {
	var err error
	DefaultFont, err = LoadFont("default")

	if err != nil {
		panic(err)
	}
}

// A monospace bitmap font
type Font struct {
	CharHeight     int        `json:"height"`
	CharWidth      int        `json:"width"`
	YShift         int        `json:"y-shift"`
	Spacing        int        `json:"spacing"`
	Rows           int        `json:"rows"`
	Cols           int        `json:"columns"`
	TexturePath    string     `json:"texture"`
	UnicodeMapping [][]string `json:"mapping"`

	Height int

	unicode2Coords map[rune]image.Point
	unknown        image.Point
}

func (f Font) Validate() error {
	if f.CharHeight <= 0 {
		return errors.NewCharacterSizeError(errors.AxisVertical, f.CharHeight)
	}
	if f.CharWidth <= 0 {
		return errors.NewCharacterSizeError(errors.AxisHorizontal, f.CharWidth)
	}

	if f.YShift > f.CharHeight {
		return errors.NewBigYShiftError(f.YShift, f.CharHeight)
	}

	if l := len(f.UnicodeMapping); l != f.Rows {
		return errors.NewIncorrectLineCountError(errors.AxisHorizontal, l, f.Rows)
	}
	for _, row := range f.UnicodeMapping {
		if l := len(row); l != f.Cols {
			return errors.NewIncorrectLineCountError(errors.AxisVertical, l, f.Cols)
		}
	}

	return nil
}

func (f *Font) UpdateResources() error {
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
					return errors.NewRepeatedCharacterDefinitionError(char)
				}

				f.unknown = image.Pt(x, y)
				continue
			}
			runeList := []rune(char)
			if len(runeList) > 1 {
				return errors.NewInvalidCharacterError(char)
			}

			currentRune := runeList[0]
			_, ok := f.unicode2Coords[currentRune]
			if ok {
				return errors.NewRepeatedCharacterDefinitionError(string(currentRune))
			}

			f.unicode2Coords[currentRune] = image.Pt(x, y)
		}
	}
	if f.unknown.X == -1 {
		return errors.NewMissingUnknownError()
	}

	return nil
}

func (f Font) GetRuneCoords(r rune) (image.Point, error) {
	if f.unicode2Coords == nil {
		return image.Point{}, errors.NewRuneFromUnloadedResourcesError()
	}

	runeCoords, ok := f.unicode2Coords[r]
	if !ok {
		return image.Pt(f.unknown.X*f.CharWidth, f.unknown.Y*f.CharHeight), nil
	}
	return image.Pt(runeCoords.X*f.CharWidth, runeCoords.Y*f.CharHeight), nil
}
