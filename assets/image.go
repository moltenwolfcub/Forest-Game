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

	BerriesLight1 *ebiten.Image = LoadPNG(lightBerries1Bytes)
	BerriesLight2 *ebiten.Image = LoadPNG(lightBerries2Bytes)
	BerriesLight3 *ebiten.Image = LoadPNG(lightBerries3Bytes)
	BerriesLight4 *ebiten.Image = LoadPNG(lightBerries4Bytes)
	BerriesLight5 *ebiten.Image = LoadPNG(lightBerries5Bytes)
	BerriesLight6 *ebiten.Image = LoadPNG(lightBerries6Bytes)
	BerriesLight7 *ebiten.Image = LoadPNG(lightBerries7Bytes)
	BerriesLight8 *ebiten.Image = LoadPNG(lightBerries8Bytes)

	BerriesMid1 *ebiten.Image = LoadPNG(midBerries1Bytes)
	BerriesMid2 *ebiten.Image = LoadPNG(midBerries2Bytes)
	BerriesMid3 *ebiten.Image = LoadPNG(midBerries3Bytes)
	BerriesMid4 *ebiten.Image = LoadPNG(midBerries4Bytes)
	BerriesMid5 *ebiten.Image = LoadPNG(midBerries5Bytes)
	BerriesMid6 *ebiten.Image = LoadPNG(midBerries6Bytes)
	BerriesMid7 *ebiten.Image = LoadPNG(midBerries7Bytes)
	BerriesMid8 *ebiten.Image = LoadPNG(midBerries8Bytes)

	BerriesDark1 *ebiten.Image = LoadPNG(darkBerries1Bytes)
	BerriesDark2 *ebiten.Image = LoadPNG(darkBerries2Bytes)
	BerriesDark3 *ebiten.Image = LoadPNG(darkBerries3Bytes)
	BerriesDark4 *ebiten.Image = LoadPNG(darkBerries4Bytes)
	BerriesDark5 *ebiten.Image = LoadPNG(darkBerries5Bytes)
	BerriesDark6 *ebiten.Image = LoadPNG(darkBerries6Bytes)
	BerriesDark7 *ebiten.Image = LoadPNG(darkBerries7Bytes)
	BerriesDark8 *ebiten.Image = LoadPNG(darkBerries8Bytes)
)
