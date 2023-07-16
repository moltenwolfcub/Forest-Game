package game

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Berry struct {
}

func (b Berry) Overlaps(layer GameContext, other []image.Rectangle) bool {
	return DefaultHitboxOverlaps(layer, b, other)
}

func (b Berry) Origin(GameContext) image.Point {
	return image.Pt(128, 256)
}

func (b Berry) Size(GameContext) image.Point {
	return image.Pt(256, 256)
}

func (b Berry) GetHitbox(GameContext) []image.Rectangle {
	return []image.Rectangle{image.Rect(0, 0, 256, 256)}
}

func (b Berry) DrawAt(screen *ebiten.Image, pos image.Point) {
	img := ebiten.NewImage(256, 256)
	img.Fill(color.RGBA{255, 0, 0, 255})

	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(pos.X-128), float64(pos.Y-256))

	screen.DrawImage(img, &options)
	img.Dispose()
}

func (b Berry) GetZ() int {
	return 1
}
