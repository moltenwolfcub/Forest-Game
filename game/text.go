package game

import (
	"image"
	"image/color"
	"log"

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
		log.Fatal(err)
	}

	fontFace, err = opentype.NewFace(loadedFont, &opentype.FaceOptions{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}

type TextElement struct {
	Contents string
	Pos      image.Point
}

func (t TextElement) GetPos() image.Point {
	return t.Pos
}

func (t TextElement) DrawAt(screen *ebiten.Image, pos image.Point) {
	bounds := text.BoundString(fontFace, t.Contents)
	text.Draw(screen, t.Contents, fontFace, pos.X, pos.Y+bounds.Dy(), color.Black)
}
