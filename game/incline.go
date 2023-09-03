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

func (i Incline) Overlaps(layer GameContext, other []image.Rectangle) (bool, error) {
	return DefaultHitboxOverlaps(layer, i, other)
}
func (i Incline) Origin(GameContext) (image.Point, error) {
	return i.hitbox.Min, nil
}
func (i Incline) Size(GameContext) (image.Point, error) {
	return i.hitbox.Size(), nil
}
func (i Incline) GetHitbox(layer GameContext) ([]image.Rectangle, error) {
	return []image.Rectangle{i.hitbox}, nil
}

func (i Incline) DrawAt(screen *ebiten.Image, pos image.Point) error {
	img := ebiten.NewImage(i.hitbox.Dx(), i.hitbox.Dy())
	img.Fill(InclineColor)

	lineartImg, drawOps, err := ApplyLineart(img, i, i.hitbox)
	if err != nil {
		return err
	}

	drawOps.GeoM.Translate(float64(pos.X), float64(pos.Y))

	screen.DrawImage(lineartImg, drawOps)
	img.Dispose()
	lineartImg.Dispose()
	return nil
}

func (i Incline) GetZ() (int, error) {
	return -2, nil
}

func (i Incline) GetClimbSpeed() float64 {
	return 0.6
}
