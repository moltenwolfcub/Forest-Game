package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	WindowWidth  = 1920
	WindowHeight = 1080
	TPS          = 60
)

type Game struct {
	time Time
	view Viewport

	layeredImage  *ebiten.Image
	bgLayer       *ebiten.Image
	mapLayer      *ebiten.Image
	lightingLayer *ebiten.Image
	hudLayer      *ebiten.Image

	debugLighting bool

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
		lamp:          NewLamp(),
		view:          NewViewport(),
		layeredImage:  ebiten.NewImage(WindowWidth, WindowHeight),
		bgLayer:       ebiten.NewImage(WindowWidth, WindowHeight),
		mapLayer:      ebiten.NewImage(WindowWidth, WindowHeight),
		hudLayer:      ebiten.NewImage(WindowWidth, WindowHeight),
		lightingLayer: ebiten.NewImage(WindowWidth, WindowHeight),
		time:          Time(TPGM * 60 * startTime),
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

	if inpututil.IsKeyJustPressed(ebiten.KeyL) {
		g.debugLighting = !g.debugLighting
	}
}

func (g Game) ambientLight(min float64, max float64) color.Color {
	colorPerTick := (max - min) / float64(DAYLEN/2)
	mappedLight := min + colorPerTick*float64(g.time.GetTimeInDay())
	if mappedLight > max {
		diff := mappedLight - max
		mappedLight = max - diff
	}
	intLight := uint8(mappedLight)

	return color.RGBA{intLight, intLight, intLight, 255}
}

func (g Game) Draw(screen *ebiten.Image) {
	g.layeredImage.Clear()
	g.bgLayer.Clear()
	g.mapLayer.Clear()
	g.lightingLayer.Clear()
	g.hudLayer.Clear()

	g.bgLayer.Fill(color.RGBA{34, 139, 34, 255})
	g.lightingLayer.Fill(g.ambientLight(48, 255))

	for _, tree := range g.trees {
		g.view.DrawToMap(g.mapLayer, tree)
	}
	g.view.DrawToMap(g.mapLayer, g.player)
	g.view.DrawToHUD(g.hudLayer, g.timeHud)
	g.view.DrawToLighting(g.lightingLayer, g.lamp)

	g.layeredImage.DrawImage(g.bgLayer, nil)
	g.layeredImage.DrawImage(g.mapLayer, nil)

	options := ebiten.DrawImageOptions{}

	if !g.debugLighting {
		options.Blend.BlendOperationRGB = ebiten.BlendOperationAdd
		options.Blend.BlendFactorSourceRGB = ebiten.BlendFactorDestinationColor
		options.Blend.BlendFactorDestinationRGB = ebiten.BlendFactorZero
	}

	g.layeredImage.DrawImage(g.lightingLayer, &options)
	g.layeredImage.DrawImage(g.hudLayer, nil)

	screen.DrawImage(g.layeredImage, nil)
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
