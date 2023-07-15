package game

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Climbable interface {
	HasHitbox

	GetClimbSpeed() float64
}

type Incline struct {
	hitbox image.Rectangle
}

func (i Incline) Overlaps(layer GameContext, other []image.Rectangle) bool {
	return DefaultHitboxOverlaps(layer, i, other)
}
func (i Incline) Origin(GameContext) image.Point {
	return i.hitbox.Min
}
func (i Incline) Size(GameContext) image.Point {
	return i.hitbox.Size()
}
func (i Incline) GetHitbox(layer GameContext) []image.Rectangle {
	return []image.Rectangle{
		i.hitbox,
	}
}

func (i Incline) DrawAt(screen *ebiten.Image, pos image.Point) {
	img := ebiten.NewImage(i.hitbox.Dx(), i.hitbox.Dy())
	img.Fill(InclineColor)

	lineartImg, drawOps := ApplyLineart(img, i, i.hitbox)
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
