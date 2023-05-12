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
	Delta image.Point
	Rect  image.Rectangle
}

func NewPlayer() Player {
	width, height := playerImage.Bounds().Size().X, playerImage.Bounds().Size().Y
	return Player{
		Rect: image.Rectangle{
			Max: image.Point{width, height},
		},
	}
}

func (p Player) DrawAt(screen *ebiten.Image, pos image.Point) {
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(pos.X), float64(pos.Y))

	screen.DrawImage(playerImage, &options)
}

func (p Player) Hitbox(RenderLayer) image.Rectangle {
	return p.Rect
}

func (p *Player) Update() {
	p.Rect = p.Rect.Add(image.Point{
		X: int(float64(p.Delta.X) * playerMoveSpeed),
		Y: int(float64(p.Delta.Y) * playerMoveSpeed),
	})
}
