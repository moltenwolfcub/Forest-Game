package game

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/moltenwolfcub/Forest-Game/assets"
)

type Tree struct {
	hitbox image.Rectangle
}

func NewTree() Tree {
	width, height := assets.Tree.Bounds().Size().X, assets.Tree.Bounds().Size().Y
	return Tree{
		hitbox: image.Rectangle{
			Max: image.Point{width, height},
		},
	}
}

func (t Tree) Overlaps(layer GameContext, other []image.Rectangle) (bool, error) {
	return DefaultHitboxOverlaps(layer, t, other)
}
func (t Tree) Origin(GameContext) (image.Point, error) {
	return t.hitbox.Min, nil
}
func (t Tree) Size(GameContext) (image.Point, error) {
	return t.hitbox.Size(), nil
}
func (t Tree) GetHitbox(layer GameContext) ([]image.Rectangle, error) {
	return []image.Rectangle{t.hitbox}, nil
}

func (t Tree) DrawAt(screen *ebiten.Image, pos image.Point) error {
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(pos.X), float64(pos.Y))

	screen.DrawImage(assets.Tree, &options)
	return nil
}

func (t Tree) GetZ() (int, error) {
	return -1, nil
}
