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

func (t Tree) Overlaps(layer GameContext, other []image.Rectangle) bool {
	return DefaultHitboxOverlaps(layer, t, other)
}
func (t Tree) Origin(GameContext) image.Point {
	return t.hitbox.Min
}
func (t Tree) Size(GameContext) image.Point {
	return t.hitbox.Size()
}
func (t Tree) GetHitbox(layer GameContext) []image.Rectangle {
	return []image.Rectangle{
		t.hitbox,
	}
}

func (t Tree) DrawAt(screen *ebiten.Image, pos image.Point) {
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(pos.X), float64(pos.Y))

	screen.DrawImage(assets.Tree, &options)
}

func (t Tree) GetZ() int {
	return -1
}

func (t *Tree) Update() {
}
