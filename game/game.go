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
	nextFrame *ebiten.Image

	time Time
	view Viewport

	renderer Renderer
	input    InputHandler

	timeHud TextElement
	player  Player

	inclines []*Incline
	rivers   []*River
	trees    []*Tree
	berries  []*Berry
	pumpkins []*Pumpkin
}

func NewGame() (*Game, error) {
	startTime := 10

	g := &Game{
		view:  NewViewport(),
		input: NewInputHandler(),
		time:  Time(TPGM * 60 * startTime),

		trees: []*Tree{},
		inclines: []*Incline{
			NewIncline(NewBasicTerrainElement(0, 0, 1024, 256)),
			NewIncline(NewBasicTerrainElement(1024, -128, 448, 256)),
			NewIncline(NewBasicTerrainElement(1472, -256, 320, 256)),
		},
	}
	g.player = NewPlayer(g)
	g.renderer = NewRenderer(g)

	river, _ := NewRiver(
		image.Pt(0, 0),
		NewRiverSegment(NewBasicTerrainElement(0, 448, 1024, 256)),
		NewRiverSegment(NewBasicTerrainElement(768, 576, 768, 256)),
		NewRiverSegment(NewBasicTerrainElement(1280, 768, 448, 256)),
	)
	g.rivers = append(g.rivers, river)

	berry, err := NewBerry(g, image.Pt(256, -128))
	if err != nil {
		return nil, err
	}
	g.berries = []*Berry{berry}

	pumpkin, err := NewPumpkin(g, image.Pt(-64, -128))
	if err != nil {
		return nil, err
	}
	g.pumpkins = []*Pumpkin{pumpkin}

	g.timeHud = NewTextElement(g.time.String(), TopCentre, assets.DefaultFont, 24)
	return g, nil
}
func NewBasicTerrainElement(x int, y int, dx int, dy int) (returnVal image.Rectangle) {
	returnVal = image.Rectangle{
		image.Point{x, y},
		image.Point{x + dx, y + dy},
	}
	return
}

func (g *Game) Update() (err error) {

	g.time.Tick()
	g.timeHud.Contents = g.time.String()
	g.timeHud.Update()
	for i := range g.berries {
		err = g.berries[i].Update()
		if err != nil {
			return err
		}
	}
	for i := range g.pumpkins {
		err = g.pumpkins[i].Update()
		if err != nil {
			return err
		}
	}

	err = g.player.Update()
	if err != nil {
		return err
	}
	err = g.view.UpdatePosition(g.player)
	if err != nil {
		return err
	}

	frame, err := g.GenerateFrame()
	if err != nil {
		return err
	}
	g.nextFrame = frame

	return nil
}

func (g Game) GenerateFrame() (*ebiten.Image, error) {
	mapElements := []DepthAwareDrawable{
		g.player,
	}
	for _, tree := range g.trees {
		mapElements = append(mapElements, DepthAwareDrawable(tree))
	}
	for _, berry := range g.berries {
		mapElements = append(mapElements, DepthAwareDrawable(berry))
	}
	for _, pumpkin := range g.pumpkins {
		mapElements = append(mapElements, DepthAwareDrawable(pumpkin))
	}
	for _, incline := range g.inclines {
		mapElements = append(mapElements, DepthAwareDrawable(incline))
	}
	for _, river := range g.rivers {
		mapElements = append(mapElements, DepthAwareDrawable(river))
	}

	lights := []Lightable{}

	hudElements := []Drawable{
		&g.timeHud,
	}

	image, err := g.renderer.Render(mapElements, lights, hudElements)
	if err != nil {
		return nil, err
	}
	return image, nil
}

func (g Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.nextFrame, nil)
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
