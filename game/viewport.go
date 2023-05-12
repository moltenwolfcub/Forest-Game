package game

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type HasPosition interface {
	GetPos() image.Rectangle
}

type Drawable interface {
	HasPosition

	DrawAt(*ebiten.Image, image.Point)
}

type Lightable interface {
	HasPosition

	GetLight() Light
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
	mapPos := drawable.GetPos()

	if v.objectInViewport(mapPos) {
		offsetPos := mapPos.Min.Sub(v.Rect.Min)
		drawable.DrawAt(mapLayer, offsetPos)
	}

}
func (v Viewport) DrawToLighting(lightingLayer *ebiten.Image, lightable Lightable) {
	mapPos := lightable.GetPos()
	light := lightable.GetLight()

	if v.objectInViewport(mapPos) || true {
		offsetPos := mapPos.Min.Sub(v.Rect.Min)

		img := ebiten.NewImage(light.Radius*2, light.Radius*2)
		vector.DrawFilledCircle(img, float32(light.Radius), float32(light.Radius), float32(light.Radius), color.RGBA{250, 129, 40, 255}, false)

		options := ebiten.DrawImageOptions{}
		options.GeoM.Translate(float64(offsetPos.X), float64(offsetPos.Y))

		lightingLayer.DrawImage(img, &options)
	}
}

func (v Viewport) DrawToHUD(hudLayer *ebiten.Image, drawable Drawable) {
	drawable.DrawAt(hudLayer, drawable.GetPos().Min)
}

func (v *Viewport) UpdatePosition(player Player) {
	v.Rect = v.Rect.Sub(v.Rect.Min).Add(image.Point{
		player.Rect.Min.X + player.Rect.Dx()/2 - v.Rect.Dx()/2,
		player.Rect.Min.Y + player.Rect.Dy()/2 - v.Rect.Dy()/2,
	})
}
