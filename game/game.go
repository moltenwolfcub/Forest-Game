package game

import (
	"image"

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
	keys     Keybinds

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
		keys:     NewKeybinds(),
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
	var delta = image.Point{0, 0}

	if g.keys.Forwards.Triggered() {
		delta.Y -= 1
	}
	if g.keys.Backwards.Triggered() {
		delta.Y += 1
	}
	if g.keys.Left.Triggered() {
		delta.X -= 1
	}
	if g.keys.Right.Triggered() {
		delta.X += 1
	}
	g.player.Delta = delta
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
