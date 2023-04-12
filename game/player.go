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
	playerMoveSpeed float32 = 1.5
)

func init() {
	var err error
	playerImage, _, err = ebitenutil.NewImageFromFile("assets/player.png")
	if err != nil {
		log.Fatal(err)
	}
}

type Player struct {
	Dx, Dy                 int
	ActualXpos, ActualYpos float32
}

func (p *Player) Draw(screen *ebiten.Image) {
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(p.ActualXpos), float64(p.ActualYpos))

	screen.DrawImage(playerImage, &options)
}

func (p *Player) Update() {
	p.ActualXpos += float32(p.Dx) * playerMoveSpeed
	p.ActualYpos += float32(p.Dy) * playerMoveSpeed
}
