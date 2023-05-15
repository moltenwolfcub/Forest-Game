package game

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Lamp struct {
	Rect image.Rectangle
}

func NewLamp() Lamp {
	radius := 70
	lamp := Lamp{
		Rect: image.Rectangle{
			Max: image.Point{radius * 2, radius * 2},
		},
	}
	lamp.Rect = lamp.Rect.Add(image.Point{50, 40})
	return lamp
}

func (l Lamp) Hitbox(layer RenderLayer) image.Rectangle {
	return l.Rect
}

func (l Lamp) DrawLighting(lightingLayer *ebiten.Image, pos image.Point) {
	radius := 70

	img := ebiten.NewImage(radius*2, radius*2)

	vector.DrawFilledCircle(img, float32(radius), float32(radius), float32(radius), color.Opaque, false)

	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(pos.X), float64(pos.Y))

	lightingLayer.DrawImage(img, &options)
	img.Dispose()

}
