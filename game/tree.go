package game

import (
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
	Pos Position
}

func (t Tree) GetMapPos() Position {
	return t.Pos
}

func (t Tree) DrawAt(screen *ebiten.Image, pos Position) {
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(pos.Xpos, pos.Ypos)

	screen.DrawImage(treeImage, &options)
}

func (t *Tree) Update() {
}
