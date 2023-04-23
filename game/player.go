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
	Dx, Dy int
	MapPos Position
}

func (p Player) DrawAt(screen *ebiten.Image, pos Position) {
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(pos.Xpos, pos.Ypos)

	screen.DrawImage(playerImage, &options)
}

func (p Player) GetMapPos() Position {
	return p.MapPos
}

func (p *Player) Update() {
	p.MapPos.Xpos += float64(p.Dx) * playerMoveSpeed
	p.MapPos.Ypos += float64(p.Dy) * playerMoveSpeed
}
