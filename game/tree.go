package game

import (
	"bytes"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/moltenwolfcub/Forest-Game/assets"
)

var (
	treeImage *ebiten.Image
)

func init() {
	var err error
	treeDecoded, _, err := image.Decode(bytes.NewReader(assets.TreePng))
	if err != nil {
		panic(err)
	}

	treeImage = ebiten.NewImageFromImage(treeDecoded)
}

type Tree struct {
	hitbox image.Rectangle
}

func NewTree() Tree {
	width, height := treeImage.Bounds().Size().X, treeImage.Bounds().Size().Y
	return Tree{
		hitbox: image.Rectangle{
			Max: image.Point{width, height},
		},
	}
}

func (t Tree) Overlaps(layer GameContext, other HasHitbox) bool {
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

	screen.DrawImage(treeImage, &options)
}

func (t Tree) GetZ() int {
	return -1
}

func (t *Tree) Update() {
}
