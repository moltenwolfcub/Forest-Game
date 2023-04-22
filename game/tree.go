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
	ActualPos Position
}

func (p *Tree) Draw(screen *ebiten.Image) {
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(p.ActualPos.Xpos, p.ActualPos.Ypos)

	screen.DrawImage(treeImage, &options)
}

func (p *Tree) Update() {
}
