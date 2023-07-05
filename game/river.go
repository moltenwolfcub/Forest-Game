package game

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type River struct {
	Hitbox image.Rectangle
}

func (r River) Overlaps(layer GameContext, other HasHitbox) bool {
	return DefaultHitboxOverlaps(layer, r, other)
}
func (r River) Origin(GameContext) image.Point {
	return r.Hitbox.Min
}
func (r River) Size(GameContext) image.Point {
	return r.Hitbox.Size()
}
func (r River) GetHitbox(layer GameContext) []image.Rectangle {
	switch layer {
	case Interaction:
		riverRect := image.Rectangle{
			Min: r.Hitbox.Min.Sub(image.Point{20, 20}),
			Max: r.Hitbox.Max.Add(image.Point{20, 20}),
		}
		return []image.Rectangle{
			riverRect,
		}
	default:
		return []image.Rectangle{
			r.Hitbox,
		}
	}
}

func (r River) DrawAt(screen *ebiten.Image, pos image.Point) {
	img := ebiten.NewImage(r.Hitbox.Dx(), r.Hitbox.Dy())
	img.Fill(color.RGBA{72, 122, 173, 255})

	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(pos.X), float64(pos.Y))

	screen.DrawImage(img, &options)
}

func (r River) GetZ() int {
	return -3
}
