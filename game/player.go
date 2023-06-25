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
	RiverJumping     bool
	currentMoveSpeed float64
}

func NewPlayer() Player {
	width, height := playerImage.Bounds().Size().X, playerImage.Bounds().Size().Y
	return Player{
		Rect: image.Rectangle{
			Min: image.Point{-100, -100},
			Max: image.Point{width - 100, height - 100},
		},
		currentMoveSpeed: playerMoveSpeed,
	}
}

func (p Player) DrawAt(screen *ebiten.Image, pos image.Point) {
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(pos.X), float64(pos.Y))

	screen.DrawImage(playerImage, &options)
}

func (p Player) Hitbox(layer GameContext) image.Rectangle {
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

func (p *Player) Update(collidables []HasHitbox, climbables []Climbable, rivers []HasHitbox) {
	currentClimeable := p.findCurrentClimable(climbables)
	p.currentMoveSpeed = p.calculateMovementSpeed(currentClimeable)

	p.movePlayer(collidables, climbables)
	p.tryClimb(currentClimeable)
	p.handleInteractions(rivers)
}

func (p *Player) handleInteractions(interactables []HasHitbox) {
	if p.RiverJumping {
		var objectToJump HasHitbox = nil
		playerHitbox := p.Hitbox(Collision)

		for _, c := range interactables {
			if playerHitbox.Overlaps(c.Hitbox(Interaction)) {
				objectToJump = c
				break
			}
		}
		if objectToJump == nil {
			return
		}
		objectHitbox := objectToJump.Hitbox(Collision)

		newPos := image.Point{}

		if playerHitbox.Max.Y <= objectHitbox.Min.Y {
			newPos.X = playerHitbox.Min.X
			newPos.Y = objectHitbox.Max.Y
		} else if playerHitbox.Min.Y >= objectHitbox.Max.Y {
			newPos.X = playerHitbox.Min.X
			newPos.Y = objectHitbox.Min.Y - playerHitbox.Dy()
		} else if playerHitbox.Max.X <= objectHitbox.Min.X {
			newPos.X = objectHitbox.Max.X
			newPos.Y = playerHitbox.Min.Y
		} else if playerHitbox.Min.X >= objectHitbox.Max.X {
			newPos.X = objectHitbox.Min.X - playerHitbox.Dx()
			newPos.Y = playerHitbox.Min.Y
		}

		p.Rect = p.Rect.Sub(playerHitbox.Min).Add(newPos)
	}
}

func (p Player) calculateMovementSpeed(currentClimable Climbable) (speed float64) {
	if currentClimable == nil {
		return playerMoveSpeed
	}
	return playerMoveSpeed * currentClimable.GetClimbSpeed()
}

func (p *Player) tryClimb(currentClimable Climbable) {
	if currentClimable != nil {
		if p.Climbing {
			p.Rect = p.Rect.Sub(image.Point{
				Y: int(p.currentMoveSpeed),
			})
		} else {
			p.Rect = p.Rect.Add(image.Point{
				Y: int(p.currentMoveSpeed),
			})
		}
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

func (p *Player) movePlayer(collidables []HasHitbox, climbables []Climbable) {
	scalar := p.currentMoveSpeed

	steps := int(scalar)
	stepSize := scalar / float64(steps)

	x := image.Point{X: int(float64(p.Delta.X) * stepSize)}
	y := image.Point{Y: int(float64(p.Delta.Y) * stepSize)}

	climbingPreMove := p.findCurrentClimable(climbables) == nil

	for i := 0; i < steps; i++ {

		p.Rect = p.Rect.Add(x)
		if climbingPreMove {
			p.fixCollisions(collidables, x)
		}

		if p.findCurrentClimable(climbables) != nil {
			//if currently climbing or decending Y-input should be ignored
			continue
		}
		p.Rect = p.Rect.Add(y)
		if p.Climbing && p.findCurrentClimable(climbables) != nil && y.Y <= 0 {
			//if hitting the bottom of a climbable while trying to climb don't fix collisions
			continue
		}
		if !p.Climbing && p.findCurrentClimable(climbables) != nil && y.Y >= 0 {
			//if hitting the top of a climbable and not climbing don't fix collisions
			continue
		}
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
