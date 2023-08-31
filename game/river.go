package game

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type River struct {
	hitbox []image.Rectangle
}

func (r River) Overlaps(layer GameContext, other []image.Rectangle) (bool, error) {
	return DefaultHitboxOverlaps(layer, r, other)
}
func (r River) Origin(layer GameContext) (image.Point, error) {
	bounds, err := r.findBounds(layer)
	return bounds.Min, err
}
func (r River) Size(layer GameContext) (image.Point, error) {
	bounds, err := r.findBounds(layer)
	return bounds.Size(), err
}
func (r River) GetHitbox(layer GameContext) ([]image.Rectangle, error) {
	switch layer {
	case Interaction:
		scaledRects := []image.Rectangle{}
		for _, rect := range r.hitbox {
			riverRect := image.Rectangle{
				Min: rect.Min.Sub(image.Point{20, 20}),
				Max: rect.Max.Add(image.Point{20, 20}),
			}
			scaledRects = append(scaledRects, riverRect)
		}
		return scaledRects, nil
	default:
		return r.hitbox, nil
	}
}

func (r River) findBounds(layer GameContext) (image.Rectangle, error) {
	minX, minY := math.MaxFloat64, math.MaxFloat64
	maxX, maxY := -math.MaxFloat64, -math.MaxFloat64

	hitbox, err := r.GetHitbox(layer)
	if err != nil {
		return image.Rectangle{}, err
	}

	for _, seg := range hitbox {
		minX = min(float64(seg.Min.X), minX)
		minY = min(float64(seg.Min.Y), minY)
		maxX = max(float64(seg.Max.X), maxX)
		maxY = max(float64(seg.Max.Y), maxY)
	}
	bounds := image.Rect(int(minX), int(minY), int(maxX), int(maxY))
	return bounds, nil
}

func (r River) DrawAt(screen *ebiten.Image, pos image.Point) error {
	hitbox, err := r.GetHitbox(Render)
	if err != nil {
		return err
	}

	for _, rect := range hitbox {
		rectImg := ebiten.NewImage(rect.Dx(), rect.Dy())
		rectImg.Fill(RiverColor)

		lineartImg, ops, err := ApplyLineart(rectImg, r, rect)
		if err != nil {
			return err
		}
		origin, err := r.Origin(Render)
		if err != nil {
			return err
		}

		ops.GeoM.Translate(float64(pos.X), float64(pos.Y))
		ops.GeoM.Translate(float64(rect.Min.X), float64(rect.Min.Y))
		ops.GeoM.Translate(-float64(origin.X), -float64(origin.Y))
		screen.DrawImage(lineartImg, ops)

		rectImg.Dispose()
		lineartImg.Dispose()
	}
	return nil
}

func (r River) GetZ() (int, error) {
	return -3, nil
}
