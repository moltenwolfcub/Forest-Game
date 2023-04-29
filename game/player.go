package game

import (
	"image"
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
	Delta                     image.Point
	MapPos                    image.Point
	PlayerWidth, PlayerHeight int
}

func NewPlayer() Player {
	width, height := playerImage.Bounds().Size().X, playerImage.Bounds().Size().Y
	return Player{image.Point{}, image.Point{}, width, height}
}

func (p Player) DrawAt(screen *ebiten.Image, pos image.Point) {
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(pos.X), float64(pos.Y))

	screen.DrawImage(playerImage, &options)
}

func (p Player) GetMapPos() image.Point {
	return p.MapPos
}

func (p *Player) Update() {
	p.MapPos.X += int(float64(p.Delta.X) * playerMoveSpeed)
	p.MapPos.Y += int(float64(p.Delta.Y) * playerMoveSpeed)
}
