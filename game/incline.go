package game

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Climbable interface {
	HasHitbox

	GetClimbSpeed() float64
}

type Incline struct {
	Hitbox image.Rectangle
}

func (i Incline) Overlaps(layer GameContext, other HasHitbox) bool {
	return DefaultHitboxOverlaps(layer, i, other)
}
func (i Incline) Origin(GameContext) image.Point {
	return i.Hitbox.Min
}
func (i Incline) Size(GameContext) image.Point {
	return i.Hitbox.Size()
}
func (i Incline) GetHitbox(layer GameContext) []image.Rectangle {
	return []image.Rectangle{
		i.Hitbox,
	}
}

func (i Incline) DrawAt(screen *ebiten.Image, pos image.Point) {
	img := ebiten.NewImage(i.Hitbox.Dx(), i.Hitbox.Dy())
	img.Fill(color.RGBA{117, 88, 69, 255})

	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(pos.X), float64(pos.Y))

	screen.DrawImage(img, &options)
}

func (i Incline) GetZ() int {
	return -2
}

func (i Incline) GetClimbSpeed() float64 {
	return 0.6
}
