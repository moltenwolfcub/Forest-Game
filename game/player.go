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
	Dx, Dy                    int
	MapPos                    Position
	PlayerWidth, PlayerHeight int
}

func NewPlayer() Player {
	width, height := playerImage.Bounds().Size().X, playerImage.Bounds().Size().Y
	return Player{0, 0, Position{}, width, height}
}

func (p Player) DrawAt(screen *ebiten.Image, pos Position) {
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(pos.Xpos), float64(pos.Ypos))

	screen.DrawImage(playerImage, &options)
}

func (p Player) GetMapPos() Position {
	return p.MapPos
}

func (p *Player) Update() {
	p.MapPos.Xpos += int(float64(p.Dx) * playerMoveSpeed)
	p.MapPos.Ypos += int(float64(p.Dy) * playerMoveSpeed)
}
