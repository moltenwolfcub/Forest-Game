package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	WindowWidth  int = 1920
	WindowHeight int = 1080
)

type Game struct {
	player Player
	trees  []Tree
	view   Viewport
}

func NewGame() Game {
	g := Game{
		player: Player{},
		trees: []Tree{
			{Pos: Position{960, 540}},
		},
		view: NewViewport(),
	}
	return g
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

func (g Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{34, 139, 34, 255})
	g.view.Draw(screen, g.player)
	for _, tree := range g.trees {
		g.view.Draw(screen, tree)
	}
}

func (g Game) Layout(actualWidth, actualHeight int) (screenWidth, screenHeight int) {
	return WindowWidth, WindowHeight
}

func (g *Game) Run() error {
	ebiten.SetWindowSize(960, 540)
	ebiten.SetWindowTitle("Chill Forest Game")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	return ebiten.RunGame(g)
}
