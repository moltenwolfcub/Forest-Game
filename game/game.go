package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	player Player
	trees  []Tree
}

func NewGame() Game {
	return Game{Player{}, []Tree{{Position{960, 540}}}}
}

func (g *Game) Update() error {
	g.HandleInput()
	g.player.Update()
	return nil
}

func (g *Game) HandleInput() {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		if ebiten.IsKeyPressed(ebiten.KeyS) {
			g.player.Dy = 0
		} else {
			g.player.Dy = -1
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.player.Dy = 1
	} else {
		g.player.Dy = 0
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		if ebiten.IsKeyPressed(ebiten.KeyA) {
			g.player.Dx = 0
		} else {
			g.player.Dx = 1
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.player.Dx = -1
	} else {
		g.player.Dx = 0
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{34, 139, 34, 255})
	g.player.Draw(screen)
	for _, tree := range g.trees {
		tree.Draw(screen)
	}
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
