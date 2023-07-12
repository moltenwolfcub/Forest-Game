package game

import (
	"image"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	fontFace font.Face
)

func init() {
	loadedFont, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		panic(err)
	}

	fontFace, err = opentype.NewFace(loadedFont, &opentype.FaceOptions{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		panic(err)
	}
}

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
	Contents     string
	pos          image.Point
	cachedBounds image.Rectangle
	Alignment    TextAlignment
}

func (t TextElement) Overlaps(layer GameContext, other []image.Rectangle) bool {
	return DefaultHitboxOverlaps(layer, t, other)
}
func (t TextElement) Origin(GameContext) image.Point {
	return t.pos
}
func (t TextElement) Size(GameContext) image.Point {
	return t.cachedBounds.Size()
}
func (t TextElement) GetHitbox(layer GameContext) []image.Rectangle {
	return []image.Rectangle{
		t.cachedBounds.Add(t.pos),
	}
}

func (t TextElement) DrawAt(screen *ebiten.Image, pos image.Point) {
	text.Draw(screen, t.Contents, fontFace, pos.X, pos.Y, color.White)
}

var (
	screenMiddleW = WindowWidth / 2
)

func (t *TextElement) Update() {
	t.cachedBounds = text.BoundString(fontFace, t.Contents)
	switch t.Alignment {
	case TopCentre:
		t.pos.Y = int(math.Abs(float64(fontFace.Metrics().CapHeight.Ceil()))) + 10
		imgSize := t.cachedBounds.Dx()
		t.pos.X = screenMiddleW - imgSize/2
	}
}
