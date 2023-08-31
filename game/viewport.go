package game

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type GameContext int

const (
	Render GameContext = iota
	Lighting
	Collision
	Interaction
)

type HasHitbox interface {
	// Most objects run a direct call to DefaultHitboxOverlaps.
	// This is basically here incase of non-rectangle hitboxes
	// like circles
	Overlaps(GameContext, []image.Rectangle) (bool, error)
	Origin(GameContext) (image.Point, error)
	Size(GameContext) (image.Point, error)
	GetHitbox(GameContext) ([]image.Rectangle, error)
}

func DefaultHitboxOverlaps(layer GameContext, mine HasHitbox, other []image.Rectangle) (bool, error) {
	for _, otherSub := range other {
		myHitbox, err := mine.GetHitbox(layer)
		if err != nil {
			return false, err
		}

		for _, mySub := range myHitbox {
			if otherSub.Overlaps(mySub) {
				return true, nil
			}
		}
	}
	return false, nil
}

type Drawable interface {
	HasHitbox

	DrawAt(*ebiten.Image, image.Point) error
}

type DepthAwareDrawable interface {
	Drawable

	GetZ() (int, error)
}

type Lightable interface {
	HasHitbox

	DrawLighting(*ebiten.Image, image.Point) error
}

type Viewport struct {
	rect image.Rectangle
}

func NewViewport() Viewport {
	return Viewport{rect: image.Rectangle{
		Max: image.Point{
			X: WindowWidth,
			Y: WindowHeight,
		},
	}}
}

func (v Viewport) Overlaps(layer GameContext, other []image.Rectangle) (bool, error) {
	return DefaultHitboxOverlaps(layer, v, other)
}
func (v Viewport) Origin(layer GameContext) (image.Point, error) {
	return v.rect.Min, nil
}
func (v Viewport) Size(layer GameContext) (image.Point, error) {
	return v.rect.Size(), nil
}
func (v Viewport) GetHitbox(layer GameContext) ([]image.Rectangle, error) {
	return []image.Rectangle{v.rect}, nil
}

func (v Viewport) objectInViewport(object HasHitbox, context GameContext) (bool, error) {
	hitbox, err := v.GetHitbox(context)
	if err != nil {
		return false, err
	}

	return object.Overlaps(context, hitbox)
}

func (v Viewport) DrawToMap(mapLayer *ebiten.Image, drawable Drawable) error {
	inViewport, err := v.objectInViewport(drawable, Render)
	if err != nil {
		return err
	}
	origin, err := drawable.Origin(Render)
	if err != nil {
		return err
	}

	if inViewport {
		offsetPos := origin.Sub(v.rect.Min)
		return drawable.DrawAt(mapLayer, offsetPos)
	}
	return nil
}
func (v Viewport) DrawToLighting(lightingLayer *ebiten.Image, lightable Lightable) error {
	inViewport, err := v.objectInViewport(lightable, Lighting)
	if err != nil {
		return err
	}

	if inViewport {
		origin, err := lightable.Origin(Lighting)
		if err != nil {
			return err
		}
		offsetPos := origin.Sub(v.rect.Min)
		return lightable.DrawLighting(lightingLayer, offsetPos)
	}
	return nil
}

func (v Viewport) DrawToHUD(hudLayer *ebiten.Image, drawable Drawable) error {
	origin, err := drawable.Origin(Render)
	if err != nil {
		return err
	}

	return drawable.DrawAt(hudLayer, origin)
}

func (v *Viewport) UpdatePosition(player Player) error {
	origin, err := v.Origin(Render)
	if err != nil {
		return err
	}
	size, err := v.Size(Render)
	if err != nil {
		return err
	}

	playerOrigin, err := player.Origin(Render)
	if err != nil {
		return err
	}
	playerSize, err := player.Size(Render)
	if err != nil {
		return err
	}

	v.rect = v.rect.Sub(origin).Add(image.Point{
		playerOrigin.X + playerSize.X/2 - size.X/2,
		playerOrigin.Y + playerSize.Y/2 - size.Y/2,
	})
	return nil
}
