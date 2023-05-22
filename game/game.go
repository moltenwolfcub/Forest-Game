package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	WindowWidth  = 1920
	WindowHeight = 1080
	TPS          = 60
)

type Game struct {
	time Time
	view Viewport

	renderer Renderer

	timeHud TextElement
	player  Player
	trees   []Tree
	lamp    Lamp
}

func NewGame() Game {
	startTime := 5

	g := Game{
		player: NewPlayer(),
		trees: []Tree{
			NewTree(),
		},
		lamp:     NewLamp(),
		view:     NewViewport(),
		renderer: NewRenderer(),
		time:     Time(TPGM * 60 * startTime),
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
	mapElements := []Drawable{
		g.player,
	}
	for _, tree := range g.trees {
		mapElements = append(mapElements, Drawable(tree))
	}

	lights := []Lightable{
		g.lamp,
	}

	hudElements := []Drawable{
		g.timeHud,
	}

	screen.DrawImage(g.renderer.Render(g.view, g.time, mapElements, lights, hudElements), nil)
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
