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
	Overlaps(GameContext, []image.Rectangle) bool
	Origin(GameContext) image.Point
	Size(GameContext) image.Point
	GetHitbox(GameContext) []image.Rectangle
}

func DefaultHitboxOverlaps(layer GameContext, mine HasHitbox, other []image.Rectangle) bool {
	for _, otherSub := range other {
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

func (v Viewport) Overlaps(layer GameContext, other []image.Rectangle) bool {
	return DefaultHitboxOverlaps(layer, v, other)
}
func (v Viewport) Origin(layer GameContext) image.Point {
	return v.rect.Min
}
func (v Viewport) Size(layer GameContext) image.Point {
	return v.rect.Size()
}
func (v Viewport) GetHitbox(layer GameContext) []image.Rectangle {
	return []image.Rectangle{
		v.rect,
	}
}

func (v Viewport) objectInViewport(object HasHitbox, context GameContext) bool {
	return object.Overlaps(context, v.GetHitbox(context))
}

func (v Viewport) DrawToMap(mapLayer *ebiten.Image, drawable Drawable) {
	if v.objectInViewport(drawable, Render) {
		offsetPos := drawable.Origin(Render).Sub(v.rect.Min)
		drawable.DrawAt(mapLayer, offsetPos)
	}

}
func (v Viewport) DrawToLighting(lightingLayer *ebiten.Image, lightable Lightable) {
	if v.objectInViewport(lightable, Lighting) {
		offsetPos := lightable.Origin(Lighting).Sub(v.rect.Min)
		lightable.DrawLighting(lightingLayer, offsetPos)
	}
}

func (v Viewport) DrawToHUD(hudLayer *ebiten.Image, drawable Drawable) {
	drawable.DrawAt(hudLayer, drawable.Origin(Render))
}

func (v *Viewport) UpdatePosition(player Player) {
	v.rect = v.rect.Sub(v.Origin(Render)).Add(image.Point{
		player.Origin(Render).X + player.Size(Render).X/2 - v.Size(Render).X/2,
		player.Origin(Render).Y + player.Size(Render).Y/2 - v.Size(Render).Y/2,
	})
}
