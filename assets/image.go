package assets

import (
	"bytes"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

func LoadPNG(file string) (*ebiten.Image, error) {

	embeddedImage, err := textures.ReadFile("textures/" + file + ".png")
	if err != nil {
		return nil, err
	}

	imageDecoded, _, err := image.Decode(bytes.NewReader(embeddedImage))
	if err != nil {
		return nil, err
	}
	return ebiten.NewImageFromImage(imageDecoded), nil
}

func MustLoadPNG(file string) *ebiten.Image {
	img, err := LoadPNG(file)
	if err != nil {
		panic("Failed to load PNG: " + err.Error())
	}
	return img
}

var (
	Player *ebiten.Image = MustLoadPNG("entity/player")
	Tree   *ebiten.Image = MustLoadPNG("object/tree")

	Icon16  *ebiten.Image = MustLoadPNG("icon/icon16")
	Icon22  *ebiten.Image = MustLoadPNG("icon/icon22")
	Icon24  *ebiten.Image = MustLoadPNG("icon/icon24")
	Icon32  *ebiten.Image = MustLoadPNG("icon/icon32")
	Icon48  *ebiten.Image = MustLoadPNG("icon/icon48")
	Icon64  *ebiten.Image = MustLoadPNG("icon/icon64")
	Icon128 *ebiten.Image = MustLoadPNG("icon/icon128")
	Icon256 *ebiten.Image = MustLoadPNG("icon/icon256")
	Icon512 *ebiten.Image = MustLoadPNG("icon/icon512")

	Berries   TextureCache = NewTextureCache()
	Mushrooms TextureCache = NewTextureCache()

	Fonts TextureCache = NewTextureCache()
)

type TextureCache struct {
	cache map[string]*ebiten.Image
}

func NewTextureCache() TextureCache {
	return TextureCache{
		cache: make(map[string]*ebiten.Image),
	}
}

func (t *TextureCache) GetTexture(id string) *ebiten.Image {
	img, ok := t.cache[id]
	if ok {
		return img
	}
	img = MustLoadPNG(id)
	t.cache[id] = img
	return img
}

func (t TextureCache) String() string {
	str := "["
	length := len(t.cache)
	i := 1

	for k := range t.cache {
		str += k
		if i < length {
			str += ", "
		}
		i++
	}
	return str + "]"
}
