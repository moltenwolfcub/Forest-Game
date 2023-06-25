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
	lamp    Lamp

	inclines []Incline
	rivers   []River
	trees    []Tree
}

func NewGame() Game {
	startTime := 5

	g := Game{
		player:   NewPlayer(),
		lamp:     NewLamp(),
		view:     NewViewport(),
		renderer: NewRenderer(),
		time:     Time(TPGM * 60 * startTime),
		keys:     NewKeybinds(),

		trees: []Tree{
			NewTree(),
		},
		inclines: []Incline{
			{Collision: image.Rect(400, -150, 800, 300)},
		},
		rivers: []River{
			{Collision: image.Rect(500, 500, 1500, 700)},
			{Collision: image.Rect(1500, 550, 2000, 750)},
		},
	}
	g.timeHud = TextElement{
		Contents: g.time.String(),
	}
	return g
}

func (g *Game) Update() error {
	collideables := []HasHitbox{}
	for _, incline := range g.inclines {
		collideables = append(collideables, HasHitbox(incline))
	}
	for _, river := range g.rivers {
		collideables = append(collideables, HasHitbox(river))
	}

	climbables := []Climbable{}
	for _, incline := range g.inclines {
		climbables = append(climbables, Climbable(incline))
	}

	rivers := []HasHitbox{}
	for _, river := range g.rivers {
		rivers = append(rivers, HasHitbox(river))
	}

	g.time.Tick()
	g.timeHud.Contents = g.time.String()
	g.HandleInput()
	g.player.Update(collideables, climbables, rivers)
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

	g.player.Climbing = g.keys.Climb.Triggered()
	g.player.RiverJumping = g.keys.RiverJump.Triggered()
}

func (g Game) Draw(screen *ebiten.Image) {
	mapElements := []DepthAwareDrawable{
		g.player,
	}
	for _, tree := range g.trees {
		mapElements = append(mapElements, DepthAwareDrawable(tree))
	}
	for _, incline := range g.inclines {
		mapElements = append(mapElements, DepthAwareDrawable(incline))
	}
	for _, river := range g.rivers {
		mapElements = append(mapElements, DepthAwareDrawable(river))
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
