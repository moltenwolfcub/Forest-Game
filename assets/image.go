package assets

import (
	"bytes"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

func LoadPNG(embeddedImage []byte) *ebiten.Image {
	imageDecoded, _, err := image.Decode(bytes.NewReader(embeddedImage))
	if err != nil {
		panic(err)
	}
	return ebiten.NewImageFromImage(imageDecoded)
}

var (
	Player *ebiten.Image = LoadPNG(playerBytes)
	Tree   *ebiten.Image = LoadPNG(treeBytes)

	Icon16  *ebiten.Image = LoadPNG(icon16Bytes)
	Icon22  *ebiten.Image = LoadPNG(icon22Bytes)
	Icon24  *ebiten.Image = LoadPNG(icon24Bytes)
	Icon32  *ebiten.Image = LoadPNG(icon32Bytes)
	Icon48  *ebiten.Image = LoadPNG(icon48Bytes)
	Icon64  *ebiten.Image = LoadPNG(icon64Bytes)
	Icon128 *ebiten.Image = LoadPNG(icon128Bytes)
	Icon256 *ebiten.Image = LoadPNG(icon256Bytes)
	Icon512 *ebiten.Image = LoadPNG(icon512Bytes)

	Berries1 *ebiten.Image = LoadPNG(berries1Bytes)
	Berries2 *ebiten.Image = LoadPNG(berries2Bytes)
	Berries3 *ebiten.Image = LoadPNG(berries3Bytes)
	Berries4 *ebiten.Image = LoadPNG(berries4Bytes)
	Berries5 *ebiten.Image = LoadPNG(berries5Bytes)
	Berries6 *ebiten.Image = LoadPNG(berries6Bytes)
	Berries7 *ebiten.Image = LoadPNG(berries7Bytes)
	Berries8 *ebiten.Image = LoadPNG(berries8Bytes)
)
