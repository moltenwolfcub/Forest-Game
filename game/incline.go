package game

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Incline struct {
	Collision image.Rectangle
}

func (i Incline) Hitbox(RenderLayer) image.Rectangle {
	return i.Collision
}

func (i Incline) DrawAt(screen *ebiten.Image, pos image.Point) {
	img := ebiten.NewImage(i.Collision.Dx(), i.Collision.Dy())
	img.Fill(color.RGBA{185, 124, 0, 255})

	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(pos.X), float64(pos.Y))

	screen.DrawImage(img, &options)
}

func (t Incline) GetZ() int {
	return -100
}
