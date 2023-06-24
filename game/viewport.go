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
	Hitbox(GameContext) image.Rectangle
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

func (v Viewport) objectInViewport(object image.Rectangle) bool {
	return v.Rect.Overlaps(object)
}

func (v Viewport) DrawToMap(mapLayer *ebiten.Image, drawable Drawable) {
	mapPos := drawable.Hitbox(Render)

	if v.objectInViewport(mapPos) {
		offsetPos := mapPos.Min.Sub(v.Rect.Min)
		drawable.DrawAt(mapLayer, offsetPos)
	}

}
func (v Viewport) DrawToLighting(lightingLayer *ebiten.Image, lightable Lightable) {
	mapPos := lightable.Hitbox(Lighting)

	if v.objectInViewport(mapPos) {
		offsetPos := mapPos.Min.Sub(v.Rect.Min)
		lightable.DrawLighting(lightingLayer, offsetPos)
	}
}

func (v Viewport) DrawToHUD(hudLayer *ebiten.Image, drawable Drawable) {
	drawable.DrawAt(hudLayer, drawable.Hitbox(Render).Min)
}

func (v *Viewport) UpdatePosition(player Player) {
	v.Rect = v.Rect.Sub(v.Rect.Min).Add(image.Point{
		player.Rect.Min.X + player.Rect.Dx()/2 - v.Rect.Dx()/2,
		player.Rect.Min.Y + player.Rect.Dy()/2 - v.Rect.Dy()/2,
	})
}
