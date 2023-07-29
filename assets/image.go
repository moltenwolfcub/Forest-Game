package assets

import (
	"bytes"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

func LoadPNG(file string) *ebiten.Image {

	embeddedImage, err := textures.ReadFile("textures/" + file + ".png")
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
	Player *ebiten.Image = LoadPNG("player")
	Tree   *ebiten.Image = LoadPNG("tree")

	Icon16  *ebiten.Image = LoadPNG("icon/icon16")
	Icon22  *ebiten.Image = LoadPNG("icon/icon22")
	Icon24  *ebiten.Image = LoadPNG("icon/icon24")
	Icon32  *ebiten.Image = LoadPNG("icon/icon32")
	Icon48  *ebiten.Image = LoadPNG("icon/icon48")
	Icon64  *ebiten.Image = LoadPNG("icon/icon64")
	Icon128 *ebiten.Image = LoadPNG("icon/icon128")
	Icon256 *ebiten.Image = LoadPNG("icon/icon256")
	Icon512 *ebiten.Image = LoadPNG("icon/icon512")

	BerriesLight1 *ebiten.Image = LoadPNG("berries/light1")
	BerriesLight2 *ebiten.Image = LoadPNG("berries/light2")
	BerriesLight3 *ebiten.Image = LoadPNG("berries/light3")
	BerriesLight4 *ebiten.Image = LoadPNG("berries/light4")
	BerriesLight5 *ebiten.Image = LoadPNG("berries/light5")
	BerriesLight6 *ebiten.Image = LoadPNG("berries/light6")
	BerriesLight7 *ebiten.Image = LoadPNG("berries/light7")
	BerriesLight8 *ebiten.Image = LoadPNG("berries/light8")

	BerriesMid1 *ebiten.Image = LoadPNG("berries/mid1")
	BerriesMid2 *ebiten.Image = LoadPNG("berries/mid2")
	BerriesMid3 *ebiten.Image = LoadPNG("berries/mid3")
	BerriesMid4 *ebiten.Image = LoadPNG("berries/mid4")
	BerriesMid5 *ebiten.Image = LoadPNG("berries/mid5")
	BerriesMid6 *ebiten.Image = LoadPNG("berries/mid6")
	BerriesMid7 *ebiten.Image = LoadPNG("berries/mid7")
	BerriesMid8 *ebiten.Image = LoadPNG("berries/mid8")

	BerriesDark1 *ebiten.Image = LoadPNG("berries/dark1")
	BerriesDark2 *ebiten.Image = LoadPNG("berries/dark2")
	BerriesDark3 *ebiten.Image = LoadPNG("berries/dark3")
	BerriesDark4 *ebiten.Image = LoadPNG("berries/dark4")
	BerriesDark5 *ebiten.Image = LoadPNG("berries/dark5")
	BerriesDark6 *ebiten.Image = LoadPNG("berries/dark6")
	BerriesDark7 *ebiten.Image = LoadPNG("berries/dark7")
	BerriesDark8 *ebiten.Image = LoadPNG("berries/dark8")
)
