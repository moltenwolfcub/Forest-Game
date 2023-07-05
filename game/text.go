package game

import (
	"image"
	"image/color"

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

type TextElement struct {
	Contents     string
	Pos          image.Point
	cachedBounds image.Rectangle
}

func (t TextElement) Overlaps(layer GameContext, other HasHitbox) bool {
	return DefaultHitboxOverlaps(layer, t, other)
}
func (t TextElement) Origin(GameContext) image.Point {
	return t.Pos
}
func (t TextElement) Size(GameContext) image.Point {
	return t.cachedBounds.Size()
}
func (t TextElement) GetHitbox(layer GameContext) []image.Rectangle {
	return []image.Rectangle{
		t.cachedBounds.Add(t.Pos),
	}
}

func (t TextElement) DrawAt(screen *ebiten.Image, pos image.Point) {
	text.Draw(screen, t.Contents, fontFace, pos.X, pos.Y+(t.cachedBounds.Dy()), color.Black)
}

func (t *TextElement) Update() {
	t.cachedBounds = text.BoundString(fontFace, t.Contents)
}
