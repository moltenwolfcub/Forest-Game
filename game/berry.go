package game

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/moltenwolfcub/Forest-Game/assets"
)

type berryPhase uint8

func (b berryPhase) GetTexture() *ebiten.Image {
	switch b {
	case 1:
		return assets.Berries1
	case 2:
		return assets.Berries2
	case 3:
		return assets.Berries3
	case 4:
		return assets.Berries4
	case 5:
		return assets.Berries5
	case 6:
		return assets.Berries6
	case 7:
		return assets.Berries7
	case 8:
		return assets.Berries8
	default:
		panic("not a valid berry phase")
	}
}

type Berry struct {
	phase berryPhase
	pos   image.Point
}

func NewBerry() Berry {
	return Berry{
		phase: 1,
	}
}

func (b Berry) Overlaps(layer GameContext, other []image.Rectangle) bool {
	return DefaultHitboxOverlaps(layer, b, other)
}

func (b Berry) Origin(GameContext) image.Point {
	return b.pos
}

func (b Berry) Size(GameContext) image.Point {
	return b.phase.GetTexture().Bounds().Size()
}

func (b Berry) GetHitbox(layer GameContext) []image.Rectangle {
	width := b.Size(layer).X
	height := b.Size(layer).Y
	offsetRect := b.phase.GetTexture().Bounds().Add(b.pos).Sub(image.Pt(width/2, height))
	return []image.Rectangle{offsetRect}
}

func (b Berry) DrawAt(screen *ebiten.Image, pos image.Point) {
	width := b.Size(Render).X
	height := b.Size(Render).Y

	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(pos.X), float64(pos.Y))
	options.GeoM.Translate(-float64(width/2), -float64(height))

	screen.DrawImage(b.phase.GetTexture(), &options)
}

func (b Berry) GetZ() int {
	return 1
}
