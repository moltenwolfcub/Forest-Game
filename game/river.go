package game

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type River struct {
	Collision image.Rectangle
}

func (r River) Hitbox(GameContext) image.Rectangle {
	return r.Collision
}

func (r River) DrawAt(screen *ebiten.Image, pos image.Point) {
	img := ebiten.NewImage(r.Collision.Dx(), r.Collision.Dy())
	img.Fill(color.RGBA{72, 122, 173, 255})

	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(pos.X), float64(pos.Y))

	screen.DrawImage(img, &options)
}

func (r River) GetZ() int {
	return -3
}
