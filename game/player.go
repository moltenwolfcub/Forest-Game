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
	Delta            image.Point
	Rect             image.Rectangle
	Climbing         bool
	currentMoveSpeed float64
}

func NewPlayer() Player {
	width, height := playerImage.Bounds().Size().X, playerImage.Bounds().Size().Y
	return Player{
		Rect: image.Rectangle{
			Max: image.Point{width, height},
		},
		currentMoveSpeed: playerMoveSpeed,
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

func (p *Player) Update(collidables []HasHitbox, climbables []Climbable) {
	currentClimeable := p.findCurrentClimable(climbables)
	p.currentMoveSpeed = p.calculateMovementSpeed(currentClimeable)

	p.movePlayer(collidables, currentClimeable)
	p.tryClimb(currentClimeable)
}

func (p Player) calculateMovementSpeed(currentClimable Climbable) (speed float64) {
	if currentClimable == nil {
		return playerMoveSpeed
	}
	return playerMoveSpeed * currentClimable.GetClimbSpeed()
}

func (p *Player) tryClimb(currentClimable Climbable) {
	if p.Climbing && currentClimable != nil {
		p.Rect = p.Rect.Sub(image.Point{
			Y: int(p.currentMoveSpeed),
		})
	}
}

func (p Player) findCurrentClimable(climbables []Climbable) (found Climbable) {
	rect := p.Hitbox(Collision)

	for _, c := range climbables {
		if rect.Overlaps(c.Hitbox(Collision)) {
			found = c
			break
		}
	}
	return
}

func (p *Player) movePlayer(collidables []HasHitbox, currentClimable Climbable) {
	scalar := p.currentMoveSpeed

	steps := int(scalar)
	stepSize := scalar / float64(steps)

	for i := 0; i < steps; i++ {

		x := image.Point{X: int(float64(p.Delta.X) * stepSize)}
		y := image.Point{Y: int(float64(p.Delta.Y) * stepSize)}

		p.Rect = p.Rect.Add(x)
		if currentClimable == nil {
			p.fixCollisions(collidables, x)
		}

		p.Rect = p.Rect.Add(y)
		p.fixCollisions(collidables, y)

	}
}

func (p *Player) fixCollisions(collidables []HasHitbox, direction image.Point) {
	for _, c := range collidables {
		if c.Hitbox(Collision).Overlaps(p.Hitbox(Collision)) {
			p.Rect = p.Rect.Sub(direction)
			break
		}
	}
}
