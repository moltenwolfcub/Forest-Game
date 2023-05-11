package game

import (
	"image"
	"log"

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
		log.Fatal(err)
	}
}

type Tree struct {
	Pos image.Point
}

func (t Tree) GetPos() image.Point {
	return t.Pos
}

func (t Tree) DrawAt(screen *ebiten.Image, pos image.Point) {
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(pos.X), float64(pos.Y))

	screen.DrawImage(treeImage, &options)
}

func (t *Tree) Update() {
}
