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
	Player *ebiten.Image = LoadPNG("textures/player.png")
	Tree   *ebiten.Image = LoadPNG("textures/tree.png")

	Icon16  *ebiten.Image = LoadPNG("textures/icon/icon16.png")
	Icon22  *ebiten.Image = LoadPNG("textures/icon/icon22.png")
	Icon24  *ebiten.Image = LoadPNG("textures/icon/icon24.png")
	Icon32  *ebiten.Image = LoadPNG("textures/icon/icon32.png")
	Icon48  *ebiten.Image = LoadPNG("textures/icon/icon48.png")
	Icon64  *ebiten.Image = LoadPNG("textures/icon/icon64.png")
	Icon128 *ebiten.Image = LoadPNG("textures/icon/icon128.png")
	Icon256 *ebiten.Image = LoadPNG("textures/icon/icon256.png")
	Icon512 *ebiten.Image = LoadPNG("textures/icon/icon512.png")

	BerriesLight1 *ebiten.Image = LoadPNG("textures/berries/light1.png")
	BerriesLight2 *ebiten.Image = LoadPNG("textures/berries/light2.png")
	BerriesLight3 *ebiten.Image = LoadPNG("textures/berries/light3.png")
	BerriesLight4 *ebiten.Image = LoadPNG("textures/berries/light4.png")
	BerriesLight5 *ebiten.Image = LoadPNG("textures/berries/light5.png")
	BerriesLight6 *ebiten.Image = LoadPNG("textures/berries/light6.png")
	BerriesLight7 *ebiten.Image = LoadPNG("textures/berries/light7.png")
	BerriesLight8 *ebiten.Image = LoadPNG("textures/berries/light8.png")

	BerriesMid1 *ebiten.Image = LoadPNG("textures/berries/mid1.png")
	BerriesMid2 *ebiten.Image = LoadPNG("textures/berries/mid2.png")
	BerriesMid3 *ebiten.Image = LoadPNG("textures/berries/mid3.png")
	BerriesMid4 *ebiten.Image = LoadPNG("textures/berries/mid4.png")
	BerriesMid5 *ebiten.Image = LoadPNG("textures/berries/mid5.png")
	BerriesMid6 *ebiten.Image = LoadPNG("textures/berries/mid6.png")
	BerriesMid7 *ebiten.Image = LoadPNG("textures/berries/mid7.png")
	BerriesMid8 *ebiten.Image = LoadPNG("textures/berries/mid8.png")

	BerriesDark1 *ebiten.Image = LoadPNG("textures/berries/dark1.png")
	BerriesDark2 *ebiten.Image = LoadPNG("textures/berries/dark2.png")
	BerriesDark3 *ebiten.Image = LoadPNG("textures/berries/dark3.png")
	BerriesDark4 *ebiten.Image = LoadPNG("textures/berries/dark4.png")
	BerriesDark5 *ebiten.Image = LoadPNG("textures/berries/dark5.png")
	BerriesDark6 *ebiten.Image = LoadPNG("textures/berries/dark6.png")
	BerriesDark7 *ebiten.Image = LoadPNG("textures/berries/dark7.png")
	BerriesDark8 *ebiten.Image = LoadPNG("textures/berries/dark8.png")
)
