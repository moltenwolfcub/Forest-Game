package assets

import (
	"bytes"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

func LoadPNG(embeddedImage []byte) *ebiten.Image {
	imageDecoded, _, err := image.Decode(bytes.NewReader(embeddedImage))
	if err != nil {
		panic(err)
	}
	return ebiten.NewImageFromImage(imageDecoded)
}
