package game

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	WindowWidth  = 1920
	WindowHeight = 1080
	TPS          = 60
)

type Game struct {
	time     Time
	timeHud  TextElement
	player   Player
	trees    []Tree
	view     Viewport
	mapLayer *ebiten.Image
	hudLayer *ebiten.Image
}

func NewGame() Game {
	g := Game{
		player: NewPlayer(),
		trees: []Tree{
			{Pos: image.Point{960, 540}},
		},
		view:     NewViewport(),
		mapLayer: ebiten.NewImage(WindowWidth, WindowHeight),
		hudLayer: ebiten.NewImage(WindowWidth, WindowHeight),
	}
	g.timeHud = TextElement{
		Contents: g.time.String(),
	}
	return g
}

func (g *Game) Update() error {
	g.time.Tick()
	g.timeHud.Contents = g.time.String()
	g.HandleInput()
	g.player.Update()
	g.view.UpdatePosition(g.player)
	return nil
}

func (g *Game) HandleInput() {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		if ebiten.IsKeyPressed(ebiten.KeyS) {
			g.player.Delta.Y = 0
		} else {
			g.player.Delta.Y = -1
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.player.Delta.Y = 1
	} else {
		g.player.Delta.Y = 0
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		if ebiten.IsKeyPressed(ebiten.KeyA) {
			g.player.Delta.X = 0
		} else {
			g.player.Delta.X = 1
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.player.Delta.X = -1
	} else {
		g.player.Delta.X = 0
	}
}

func (g Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{34, 139, 34, 255})

	g.mapLayer.Clear()
	g.hudLayer.Clear()

	for _, tree := range g.trees {
		g.view.DrawToMap(g.mapLayer, tree)
	}
	g.view.DrawToMap(g.mapLayer, g.player)
	g.view.DrawToHUD(g.hudLayer, g.timeHud)

	screen.DrawImage(g.mapLayer, &ebiten.DrawImageOptions{})
	screen.DrawImage(g.hudLayer, &ebiten.DrawImageOptions{})
}

func (g Game) Layout(actualWidth, actualHeight int) (screenWidth, screenHeight int) {
	return WindowWidth, WindowHeight
}

func (g *Game) Run() error {
	ebiten.SetWindowSize(960, 540)
	ebiten.SetWindowTitle("Chill Forest Game")
	ebiten.SetTPS(TPS)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	return ebiten.RunGame(g)
}
