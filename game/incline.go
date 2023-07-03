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
	Collision image.Rectangle
}

func (i Incline) Hitbox(GameContext) image.Rectangle {
	return i.Collision
}

func (i Incline) DrawAt(screen *ebiten.Image, pos image.Point) {
	img := ebiten.NewImage(i.Collision.Dx(), i.Collision.Dy())
	img.Fill(color.RGBA{117, 88, 69, 255})

	lineartImg, drawOps := ApplyLineart(img)
	drawOps.GeoM.Translate(float64(pos.X), float64(pos.Y))

	screen.DrawImage(lineartImg, drawOps)
	img.Dispose()
	lineartImg.Dispose()
}

func (i Incline) GetZ() int {
	return -2
}

func (i Incline) GetClimbSpeed() float64 {
	return 0.6
}
