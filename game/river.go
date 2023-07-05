package game

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type River struct {
	hitbox image.Rectangle
}

func (r River) Overlaps(layer GameContext, other HasHitbox) bool {
	return DefaultHitboxOverlaps(layer, r, other)
}
func (r River) Origin(GameContext) image.Point {
	return r.hitbox.Min
}
func (r River) Size(GameContext) image.Point {
	return r.hitbox.Size()
}
func (r River) GetHitbox(layer GameContext) []image.Rectangle {
	switch layer {
	case Interaction:
		riverRect := image.Rectangle{
			Min: r.hitbox.Min.Sub(image.Point{20, 20}),
			Max: r.hitbox.Max.Add(image.Point{20, 20}),
		}
		return []image.Rectangle{
			riverRect,
		}
	default:
		return []image.Rectangle{
			r.hitbox,
		}
	}
}

func (r River) DrawAt(screen *ebiten.Image, pos image.Point) {
	img := ebiten.NewImage(r.hitbox.Dx(), r.hitbox.Dy())
	img.Fill(color.RGBA{72, 122, 173, 255})

	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(pos.X), float64(pos.Y))

	screen.DrawImage(img, &options)
}

func (r River) GetZ() int {
	return -3
}
