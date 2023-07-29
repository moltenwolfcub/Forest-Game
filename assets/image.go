package assets

import (
	"bytes"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

func LoadPNG(file string) *ebiten.Image {

	embeddedImage, err := textures.ReadFile(file)
	if err != nil {
		panic(err)
	}

	imageDecoded, _, err := image.Decode(bytes.NewReader(embeddedImage))
	if err != nil {
		panic(err)
	}
	return ebiten.NewImageFromImage(imageDecoded)
}

var (
	Player *ebiten.Image = LoadPNG("player.png")
	Tree   *ebiten.Image = LoadPNG("tree.png")

	Icon16  *ebiten.Image = LoadPNG("icon/icon16.png")
	Icon22  *ebiten.Image = LoadPNG("icon/icon22.png")
	Icon24  *ebiten.Image = LoadPNG("icon/icon24.png")
	Icon32  *ebiten.Image = LoadPNG("icon/icon32.png")
	Icon48  *ebiten.Image = LoadPNG("icon/icon48.png")
	Icon64  *ebiten.Image = LoadPNG("icon/icon64.png")
	Icon128 *ebiten.Image = LoadPNG("icon/icon128.png")
	Icon256 *ebiten.Image = LoadPNG("icon/icon256.png")
	Icon512 *ebiten.Image = LoadPNG("icon/icon512.png")

	BerriesLight1 *ebiten.Image = LoadPNG("berries/light1.png")
	BerriesLight2 *ebiten.Image = LoadPNG("berries/light2.png")
	BerriesLight3 *ebiten.Image = LoadPNG("berries/light3.png")
	BerriesLight4 *ebiten.Image = LoadPNG("berries/light4.png")
	BerriesLight5 *ebiten.Image = LoadPNG("berries/light5.png")
	BerriesLight6 *ebiten.Image = LoadPNG("berries/light6.png")
	BerriesLight7 *ebiten.Image = LoadPNG("berries/light7.png")
	BerriesLight8 *ebiten.Image = LoadPNG("berries/light8.png")

	BerriesMid1 *ebiten.Image = LoadPNG("berries/mid1.png")
	BerriesMid2 *ebiten.Image = LoadPNG("berries/mid2.png")
	BerriesMid3 *ebiten.Image = LoadPNG("berries/mid3.png")
	BerriesMid4 *ebiten.Image = LoadPNG("berries/mid4.png")
	BerriesMid5 *ebiten.Image = LoadPNG("berries/mid5.png")
	BerriesMid6 *ebiten.Image = LoadPNG("berries/mid6.png")
	BerriesMid7 *ebiten.Image = LoadPNG("berries/mid7.png")
	BerriesMid8 *ebiten.Image = LoadPNG("berries/mid8.png")

	BerriesDark1 *ebiten.Image = LoadPNG("berries/dark1.png")
	BerriesDark2 *ebiten.Image = LoadPNG("berries/dark2.png")
	BerriesDark3 *ebiten.Image = LoadPNG("berries/dark3.png")
	BerriesDark4 *ebiten.Image = LoadPNG("berries/dark4.png")
	BerriesDark5 *ebiten.Image = LoadPNG("berries/dark5.png")
	BerriesDark6 *ebiten.Image = LoadPNG("berries/dark6.png")
	BerriesDark7 *ebiten.Image = LoadPNG("berries/dark7.png")
	BerriesDark8 *ebiten.Image = LoadPNG("berries/dark8.png")
)
