package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Player Player
}

func NewGame() Game {
	return Game{Player{150, 180}}
}

func (g *Game) Update() error {
	g.Player.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{34, 139, 34, 255})
	g.Player.Draw(screen)
}

func (g *Game) Layout(actualWidth, actualHeight int) (screenWidth, screenHeight int) {
	return 1920, 1080
}

func (g *Game) Run() error {
	ebiten.SetWindowSize(960, 540)
	ebiten.SetWindowTitle("Chill Forest Game")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	return ebiten.RunGame(g)
}
