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
	Overlaps(GameContext, HasHitbox) bool
	Origin(GameContext) image.Point
	GetHitbox(GameContext) []image.Rectangle
}

func DefaultHitboxOverlaps(layer GameContext, mine HasHitbox, other HasHitbox) bool {
	for _, otherSub := range other.GetHitbox(layer) {
		for _, mySub := range mine.GetHitbox(layer) {
			if otherSub.Overlaps(mySub) {
				return true
			}
		}
	}
	return false
}

type Drawable interface {
	HasHitbox

	DrawAt(*ebiten.Image, image.Point)
}

type DepthAwareDrawable interface {
	Drawable

	GetZ() int
}

type Lightable interface {
	HasHitbox

	DrawLighting(*ebiten.Image, image.Point)
}

type Viewport struct {
	Rect image.Rectangle
}

func NewViewport() Viewport {
	return Viewport{Rect: image.Rectangle{
		Max: image.Point{
			X: WindowWidth,
			Y: WindowHeight,
		},
	}}
}

func (v Viewport) Overlaps(layer GameContext, other HasHitbox) bool {
	return DefaultHitboxOverlaps(layer, v, other)
}
func (v Viewport) Origin(layer GameContext) image.Point {
	return v.Rect.Min
}
func (v Viewport) GetHitbox(layer GameContext) []image.Rectangle {
	return []image.Rectangle{
		v.Rect,
	}
}

func (v Viewport) objectInViewport(object HasHitbox, context GameContext) bool {
	return object.Overlaps(context, v)
}

func (v Viewport) DrawToMap(mapLayer *ebiten.Image, drawable Drawable) {
	if v.objectInViewport(drawable, Render) {
		offsetPos := drawable.Origin(Render).Sub(v.Rect.Min)
		drawable.DrawAt(mapLayer, offsetPos)
	}

}
func (v Viewport) DrawToLighting(lightingLayer *ebiten.Image, lightable Lightable) {
	if v.objectInViewport(lightable, Lighting) {
		offsetPos := lightable.Origin(Lighting).Sub(v.Rect.Min)
		lightable.DrawLighting(lightingLayer, offsetPos)
	}
}

func (v Viewport) DrawToHUD(hudLayer *ebiten.Image, drawable Drawable) {
	drawable.DrawAt(hudLayer, drawable.Origin(Render))
}

func (v *Viewport) UpdatePosition(player Player) {
	v.Rect = v.Rect.Sub(v.Rect.Min).Add(image.Point{
		player.Rect.Min.X + player.Rect.Dx()/2 - v.Rect.Dx()/2,
		player.Rect.Min.Y + player.Rect.Dy()/2 - v.Rect.Dy()/2,
	})
}
