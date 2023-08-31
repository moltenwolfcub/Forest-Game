package game

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/moltenwolfcub/Forest-Game/assets"
)

type TextAlignment int

const (
	// Position doesn't get automatically controlled
	// and needs to be done separately
	Free TextAlignment = iota
	// Automatically adjust the position of the text
	// to be centred at the top of the screen
	TopCentre
)

type TextElement struct {
	Contents  string
	Alignment TextAlignment
	Font      assets.Font
	Scale     int
	pos       image.Point

	runeCache map[rune]*ebiten.Image
}

func NewTextElement(contents string, alignment TextAlignment, font assets.Font, scale int) TextElement {
	return TextElement{
		Contents:  contents,
		Alignment: alignment,
		Font:      font,
		Scale:     scale,
		runeCache: make(map[rune]*ebiten.Image),
	}
}

func (t TextElement) Overlaps(layer GameContext, other []image.Rectangle) (bool, error) {
	return DefaultHitboxOverlaps(layer, t, other)
}
func (t TextElement) Origin(GameContext) (image.Point, error) {
	return t.pos, nil
}
func (t TextElement) Size(GameContext) (image.Point, error) {
	return image.Pt(int(float64(len(t.Contents)*t.Font.CharWidth)*t.getScalar()), int(float64(t.Font.CharHeight)*t.getScalar())), nil
}
func (t TextElement) GetHitbox(layer GameContext) ([]image.Rectangle, error) {
	origin, err := t.Origin(layer)
	if err != nil {
		return nil, err
	}
	size, err := t.Size(layer)
	if err != nil {
		return nil, err
	}

	return []image.Rectangle{
		{
			Min: origin,
			Max: origin.Add(size),
		},
	}, nil
}

func (t TextElement) getScalar() float64 {
	return float64(t.Scale) / float64(t.Font.CharHeight)
}

func (t *TextElement) DrawAt(screen *ebiten.Image, pos image.Point) error {
	glyphs := []*ebiten.Image{}
	for _, c := range t.Contents {
		cachedGlyph, ok := t.runeCache[c]
		if ok {
			glyphs = append(glyphs, cachedGlyph)
			continue
		}

		coords, err := t.Font.GetRuneCoords(c)
		if err != nil {
			panic(err)
		}

		rect := image.Rectangle{
			Min: coords,
			Max: coords.Add(image.Pt(t.Font.CharWidth, t.Font.CharHeight)),
		}

		glyph := assets.Fonts.GetTexture(t.Font.TexturePath).SubImage(rect).(*ebiten.Image)
		t.runeCache[c] = glyph

		glyphs = append(glyphs, glyph)
	}

	scalar := t.getScalar()

	for i, glyph := range glyphs {
		options := ebiten.DrawImageOptions{}
		options.GeoM.Scale(float64(scalar), float64(scalar))

		options.GeoM.Translate(float64(pos.X), float64(pos.Y))
		options.GeoM.Translate(float64(i)*(float64(t.Font.CharWidth)*scalar+float64(t.Font.Spacing)), float64(-t.Font.CharHeight)*scalar+float64(t.Font.YShift))

		screen.DrawImage(glyph, &options)
	}
	return nil
}

const (
	screenMiddleW = WindowWidth / 2
)

func (t *TextElement) Update() {
	switch t.Alignment {
	case TopCentre:
		t.pos.Y = int(float64(t.Font.Height)*t.getScalar() + 10)
		imgSize := float64(len(t.Contents)) * float64(t.Font.CharWidth) * t.getScalar()
		t.pos.X = screenMiddleW - int(imgSize/2)
	}
}
