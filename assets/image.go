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

	Berries BerryCache = NewBerryCache()
)

type BerryCache struct {
	cache map[string]*ebiten.Image
}

func NewBerryCache() BerryCache {
	return BerryCache{
		cache: make(map[string]*ebiten.Image),
	}
}

func (b *BerryCache) GetTexture(id string) *ebiten.Image {
	img, ok := b.cache[id]
	if ok {
		return img
	}
	img = LoadPNG(id)
	b.cache[id] = img
	return img
}

func (b BerryCache) String() string {
	str := "["
	length := len(b.cache)
	i := 1

	for k := range b.cache {
		str += k
		if i < length {
			str += ", "
		}
		i++
	}
	return str + "]"
}
