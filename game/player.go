package game

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	playerImage *ebiten.Image
)

const (
	playerMoveSpeed float64 = 11.5
)

func init() {
	var err error
	playerImage, _, err = ebitenutil.NewImageFromFile("assets/player.png")
	if err != nil {
		panic(err)
	}
}

type Player struct {
	Delta image.Point
	Rect  image.Rectangle
}

func NewPlayer() Player {
	width, height := playerImage.Bounds().Size().X, playerImage.Bounds().Size().Y
	return Player{
		Rect: image.Rectangle{
			Max: image.Point{width, height},
		},
	}
}

func (p Player) DrawAt(screen *ebiten.Image, pos image.Point) {
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(pos.X), float64(pos.Y))

	screen.DrawImage(playerImage, &options)
}

func (p Player) Hitbox(layer RenderLayer) image.Rectangle {
	switch layer {
	case Collision:
		baseSize := p.Rect.Size().Y / 2

		rect := image.Rectangle{
			Min: p.Rect.Max.Sub(image.Point{p.Rect.Dx(), baseSize}),
			Max: p.Rect.Max,
		}
		return rect
	default:
		return p.Rect
	}
}

func (p Player) GetZ() int {
	return 0
}

func (p *Player) Update(collidables []HasHitbox) {
	p.movePlayer(collidables)
}

func (p *Player) movePlayer(collidables []HasHitbox) {
	scalar := playerMoveSpeed

	steps := int(scalar)
	stepSize := scalar / float64(steps)

	for i := 0; i < steps; i++ {

		x := image.Point{X: int(float64(p.Delta.X) * stepSize)}
		y := image.Point{Y: int(float64(p.Delta.Y) * stepSize)}

		p.Rect = p.Rect.Add(x)
		p.checkCollisions(collidables, x)

		p.Rect = p.Rect.Add(y)
		p.checkCollisions(collidables, y)

	}
}

func (p *Player) checkCollisions(collidables []HasHitbox, direction image.Point) {
	for _, c := range collidables {
		if c.Hitbox(Collision).Overlaps(p.Hitbox(Collision)) {
			p.Rect = p.Rect.Sub(direction)
			break
		}
	}
}
