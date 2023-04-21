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
	ActualXpos, ActualYpos float32
}

func (p *Tree) Draw(screen *ebiten.Image) {
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(p.ActualXpos), float64(p.ActualYpos))

	screen.DrawImage(treeImage, &options)
}

func (p *Tree) Update() {
}
