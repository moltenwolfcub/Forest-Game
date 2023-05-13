package game

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	treeImage *ebiten.Image
)

func init() {
	var err error
	treeImage, _, err = ebitenutil.NewImageFromFile("assets/tree.png")
	if err != nil {
		panic(err)
	}
}

type Tree struct {
	Rect image.Rectangle
}

func NewTree() Tree {
	width, height := treeImage.Bounds().Size().X, treeImage.Bounds().Size().Y
	return Tree{
		Rect: image.Rectangle{
			Max: image.Point{width, height},
		},
	}
}

func (t Tree) Hitbox(RenderLayer) image.Rectangle {
	return t.Rect
}

func (t Tree) DrawAt(screen *ebiten.Image, pos image.Point) {
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(pos.X), float64(pos.Y))

	screen.DrawImage(treeImage, &options)
}

func (t *Tree) Update() {
}
