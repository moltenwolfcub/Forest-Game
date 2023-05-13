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
	Contents string
	Pos      image.Point
}

func (t TextElement) Hitbox(RenderLayer) image.Rectangle {
	return text.BoundString(fontFace, t.Contents).Add(t.Pos)
}

func (t TextElement) DrawAt(screen *ebiten.Image, pos image.Point) {
	bounds := text.BoundString(fontFace, t.Contents)
	text.Draw(screen, t.Contents, fontFace, pos.X, pos.Y+(2*bounds.Dy()), color.Black)
}
