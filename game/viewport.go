package game

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type RenderLayer int

const (
	Map RenderLayer = iota
	Lighting
	GUI
)

type HasHitbox interface {
	Hitbox(RenderLayer) image.Rectangle
}

type Drawable interface {
	HasHitbox

	DrawAt(*ebiten.Image, image.Point)
}

type Lightable interface {
	HasHitbox

	Radius() int
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
	mapPos := drawable.Hitbox(Map)

	if v.objectInViewport(mapPos) {
		offsetPos := mapPos.Min.Sub(v.Rect.Min)
		drawable.DrawAt(mapLayer, offsetPos)
	}

}
func (v Viewport) DrawToLighting(lightingLayer *ebiten.Image, lightable Lightable) {
	mapPos := lightable.Hitbox(Lighting)
	radius := lightable.Radius()

	if v.objectInViewport(mapPos) || true {
		offsetPos := mapPos.Min.Sub(v.Rect.Min)

		img := ebiten.NewImage(radius*2, radius*2)
		vector.DrawFilledCircle(img, float32(radius), float32(radius), float32(radius), color.Opaque, false)

		options := ebiten.DrawImageOptions{}
		options.GeoM.Translate(float64(offsetPos.X), float64(offsetPos.Y))

		lightingLayer.DrawImage(img, &options)
		img.Dispose()
	}
}

func (v Viewport) DrawToHUD(hudLayer *ebiten.Image, drawable Drawable) {
	drawable.DrawAt(hudLayer, drawable.Hitbox(GUI).Min)
}

func (v *Viewport) UpdatePosition(player Player) {
	v.Rect = v.Rect.Sub(v.Rect.Min).Add(image.Point{
		player.Rect.Min.X + player.Rect.Dx()/2 - v.Rect.Dx()/2,
		player.Rect.Min.Y + player.Rect.Dy()/2 - v.Rect.Dy()/2,
	})
}
