package game

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/moltenwolfcub/Forest-Game/assets"
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

	inclines []Incline
	rivers   []River
	trees    []Tree
	berries  []Berry
}

func NewGame() Game {
	startTime := 10

	g := Game{
		player:   NewPlayer(),
		view:     NewViewport(),
		renderer: NewRenderer(),
		time:     Time(TPGM * 60 * startTime),
		keys:     NewKeybinds(),

		trees:   []Tree{},
		berries: []Berry{{}},
		inclines: []Incline{
			{NewBasicTerrainElement(0, 0, 1024, 256)},
			{NewBasicTerrainElement(1024, -128, 448, 256)},
			{NewBasicTerrainElement(1472, -256, 320, 256)},
		},
		rivers: []River{
			{hitbox: []image.Rectangle{
				NewBasicTerrainElement(0, 448, 1024, 256),
				NewBasicTerrainElement(768, 576, 768, 256),
				NewBasicTerrainElement(1280, 768, 448, 256),
			}},
		},
	}
	g.timeHud = TextElement{
		Contents:  g.time.String(),
		Alignment: TopCentre,
	}
	return g
}
func NewBasicTerrainElement(x int, y int, dx int, dy int) (returnVal image.Rectangle) {
	returnVal = image.Rectangle{
		image.Point{x, y},
		image.Point{x + dx, y + dy},
	}
	return
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
	g.timeHud.Update()
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
	for _, berry := range g.berries {
		mapElements = append(mapElements, DepthAwareDrawable(berry))
	}
	for _, incline := range g.inclines {
		mapElements = append(mapElements, DepthAwareDrawable(incline))
	}
	for _, river := range g.rivers {
		mapElements = append(mapElements, DepthAwareDrawable(river))
	}

	lights := []Lightable{}

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
	ebiten.SetWindowIcon([]image.Image{
		assets.Icon16, assets.Icon22, assets.Icon24,
		assets.Icon32, assets.Icon48, assets.Icon64,
		assets.Icon128, assets.Icon256, assets.Icon512,
	})
	ebiten.SetTPS(TPS)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	return ebiten.RunGame(g)
}
