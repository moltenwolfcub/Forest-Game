package game

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	playerImage *ebiten.Image
)

const (
	playerMoveSpeed float64 = 11.5
)

func init() {
	var err error
	playerImage, _, err = ebitenutil.NewImageFromFile("assets/player.png")
	if err != nil {
		log.Fatal(err)
	}
}

type Player struct {
	Dx, Dy    int
	ActualPos Position
}

func (p *Player) Draw(screen *ebiten.Image) {
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(p.ActualPos.Xpos, p.ActualPos.Ypos)

	screen.DrawImage(playerImage, &options)
}

func (p *Player) Update() {
	p.ActualPos.Xpos += float64(p.Dx) * playerMoveSpeed
	p.ActualPos.Ypos += float64(p.Dy) * playerMoveSpeed
}
