package game

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type River struct {
	hitbox []image.Rectangle
}

func (r River) Overlaps(layer GameContext, other []image.Rectangle) bool {
	return DefaultHitboxOverlaps(layer, r, other)
}
func (r River) Origin(layer GameContext) image.Point {
	bounds := r.findBounds(layer)
	return bounds.Min
}
func (r River) Size(layer GameContext) image.Point {
	bounds := r.findBounds(layer)
	return bounds.Size()
}
func (r River) GetHitbox(layer GameContext) []image.Rectangle {
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
		return scaledRects
	default:
		return r.hitbox
	}
}

func (r River) findBounds(layer GameContext) image.Rectangle {
	minX, minY := math.MaxFloat64, math.MaxFloat64
	maxX, maxY := -math.MaxFloat64, -math.MaxFloat64
	for _, seg := range r.GetHitbox(layer) {
		minX = min(float64(seg.Min.X), minX)
		minY = min(float64(seg.Min.Y), minY)
		maxX = max(float64(seg.Max.X), maxX)
		maxY = max(float64(seg.Max.Y), maxY)
	}
	bounds := image.Rect(int(minX), int(minY), int(maxX), int(maxY))
	return bounds
}

func (r River) DrawAt(screen *ebiten.Image, pos image.Point) {
	for _, rect := range r.GetHitbox(Render) {
		rectImg := ebiten.NewImage(rect.Dx(), rect.Dy())
		rectImg.Fill(RiverColor)

		lineartImg, ops := ApplyLineart(rectImg, r, rect)
		ops.GeoM.Translate(float64(pos.X), float64(pos.Y))
		ops.GeoM.Translate(float64(rect.Min.X), float64(rect.Min.Y))
		origin := r.Origin(Render)
		ops.GeoM.Translate(-float64(origin.X), -float64(origin.Y))
		screen.DrawImage(lineartImg, ops)

		rectImg.Dispose()
		lineartImg.Dispose()
	}
}

func (r River) GetZ() int {
	return -3
}
