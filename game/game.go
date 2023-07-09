package game

import (
	"bytes"
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
}

func NewGame() Game {
	startTime := 10

	g := Game{
		player:   NewPlayer(),
		view:     NewViewport(),
		renderer: NewRenderer(),
		time:     Time(TPGM * 60 * startTime),
		keys:     NewKeybinds(),

		trees: []Tree{
			// NewTree(),
		},
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
		Contents: g.time.String(),
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
	ebiten.SetWindowIcon([]image.Image{icon16, icon22, icon24, icon32, icon48, icon64, icon128, icon256, icon512})
	ebiten.SetTPS(TPS)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	return ebiten.RunGame(g)
}

var (
	icon16  *ebiten.Image
	icon22  *ebiten.Image
	icon24  *ebiten.Image
	icon32  *ebiten.Image
	icon48  *ebiten.Image
	icon64  *ebiten.Image
	icon128 *ebiten.Image
	icon256 *ebiten.Image
	icon512 *ebiten.Image
)

func init() {
	iconDecoded, _, err := image.Decode(bytes.NewReader(assets.Icon16))
	if err != nil {
		panic(err)
	}
	icon16 = ebiten.NewImageFromImage(iconDecoded)

	iconDecoded, _, err = image.Decode(bytes.NewReader(assets.Icon22))
	if err != nil {
		panic(err)
	}
	icon22 = ebiten.NewImageFromImage(iconDecoded)

	iconDecoded, _, err = image.Decode(bytes.NewReader(assets.Icon24))
	if err != nil {
		panic(err)
	}
	icon24 = ebiten.NewImageFromImage(iconDecoded)

	iconDecoded, _, err = image.Decode(bytes.NewReader(assets.Icon32))
	if err != nil {
		panic(err)
	}
	icon32 = ebiten.NewImageFromImage(iconDecoded)

	iconDecoded, _, err = image.Decode(bytes.NewReader(assets.Icon48))
	if err != nil {
		panic(err)
	}
	icon48 = ebiten.NewImageFromImage(iconDecoded)

	iconDecoded, _, err = image.Decode(bytes.NewReader(assets.Icon64))
	if err != nil {
		panic(err)
	}
	icon64 = ebiten.NewImageFromImage(iconDecoded)

	iconDecoded, _, err = image.Decode(bytes.NewReader(assets.Icon128))
	if err != nil {
		panic(err)
	}
	icon128 = ebiten.NewImageFromImage(iconDecoded)

	iconDecoded, _, err = image.Decode(bytes.NewReader(assets.Icon256))
	if err != nil {
		panic(err)
	}
	icon256 = ebiten.NewImageFromImage(iconDecoded)

	iconDecoded, _, err = image.Decode(bytes.NewReader(assets.Icon512))
	if err != nil {
		panic(err)
	}
	icon512 = ebiten.NewImageFromImage(iconDecoded)
}
