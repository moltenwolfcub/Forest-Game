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

	cachedTexture *OffsetImage
}

func NewIncline(hitbox image.Rectangle) *Incline {
	return &Incline{
		hitbox: hitbox,
	}
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

func (i *Incline) DrawAt(screen *ebiten.Image, pos image.Point) error {
	if i.cachedTexture == nil {
		err := i.generateTexture()
		if err != nil {
			return err
		}
	}

	return i.cachedTexture.DrawAt(screen, pos)
}

func (i *Incline) markTextureDirty() {
	if i.cachedTexture == nil {
		return
	}
	i.cachedTexture.Image.Dispose()
	i.cachedTexture = nil
}

func (i *Incline) generateTexture() error {
	img := ebiten.NewImage(i.hitbox.Dx(), i.hitbox.Dy())
	img.Fill(InclineColor)

	lineartImg, err := ApplyLineart(img, i, i.hitbox)
	if err != nil {
		return err
	}
	img.Dispose()
	i.cachedTexture = lineartImg

	return nil
}

func (i Incline) GetZ() (int, error) {
	return -2, nil
}

func (i Incline) GetClimbSpeed() float64 {
	return 0.6
}
