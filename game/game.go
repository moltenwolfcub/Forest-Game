package game

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	WindowWidth  int = 1920
	WindowHeight int = 1080
	TPS          int = 60
)

var (
	fontFace font.Face
)

type Game struct {
	time   Time
	player Player
	trees  []Tree
	view   Viewport
}

func init() {
	loadedFont, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	fontFace, err = opentype.NewFace(loadedFont, &opentype.FaceOptions{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func NewGame() Game {
	g := Game{
		player: NewPlayer(),
		trees: []Tree{
			{Pos: image.Point{960, 540}},
		},
		view: NewViewport(),
	}
	return g
}

func (g *Game) Update() error {
	g.time.Tick()
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
	for _, tree := range g.trees {
		g.view.Draw(screen, tree)
	}
	g.view.Draw(screen, g.player)

	bounds := text.BoundString(fontFace, fmt.Sprint(g.time.getTicks()))
	text.Draw(screen, fmt.Sprint(g.time.getTicks()), fontFace, WindowWidth/2-bounds.Dx()/2, 50, color.Black)
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
