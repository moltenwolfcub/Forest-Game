package game

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/moltenwolfcub/Forest-Game/assets"
)

var (
	playerImage *ebiten.Image
)

const (
	playerMoveSpeed float64 = 11.5
)

func init() {
	var err error
	playerDecoded, _, err := image.Decode(bytes.NewReader(assets.PlayerPng))
	if err != nil {
		panic(err)
	}

	playerImage = ebiten.NewImageFromImage(playerDecoded)
}

type Player struct {
	Delta            image.Point
	hitbox           image.Rectangle
	Climbing         bool
	RiverJumping     bool
	currentMoveSpeed float64
}

func NewPlayer() Player {
	width, height := playerImage.Bounds().Size().X, playerImage.Bounds().Size().Y
	return Player{
		hitbox: image.Rectangle{
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

func (p Player) Overlaps(layer GameContext, other HasHitbox) bool {
	return DefaultHitboxOverlaps(layer, p, other)
}
func (p Player) Origin(GameContext) image.Point {
	return p.hitbox.Min
}
func (p Player) Size(GameContext) image.Point {
	return p.hitbox.Size()
}
func (p Player) GetHitbox(layer GameContext) []image.Rectangle {
	switch layer {
	case Collision:
		baseSize := p.hitbox.Size().Y / 2

		playerRect := image.Rectangle{
			Min: p.hitbox.Max.Sub(image.Point{p.hitbox.Dx(), baseSize}),
			Max: p.hitbox.Max,
		}
		return []image.Rectangle{
			playerRect,
		}
	default:
		return []image.Rectangle{
			p.hitbox,
		}
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

		for _, c := range interactables {
			if p.Overlaps(Interaction, c) {
				objectToJump = c
				break
			}
		}
		if objectToJump == nil {
			return
		}

		//for now jump to top corner will need to properly re-implement at some point
		newPos := objectToJump.Origin(Collision).Sub(image.Point{p.hitbox.Dx(), p.hitbox.Dy()})

		p.hitbox = p.hitbox.Sub(p.hitbox.Min).Add(newPos)
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
			p.hitbox = p.hitbox.Sub(image.Point{
				Y: int(p.currentMoveSpeed),
			})
		} else {
			p.hitbox = p.hitbox.Add(image.Point{
				Y: int(p.currentMoveSpeed),
			})
		}
	}
}

func (p Player) findCurrentClimable(climbables []Climbable) (found Climbable) {
	for _, c := range climbables {
		if c.Overlaps(Collision, p) {
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

		p.hitbox = p.hitbox.Add(x)
		if climbingPreMove {
			p.fixCollisions(collidables, x)
		}

		if p.findCurrentClimable(climbables) != nil {
			//if currently climbing or decending Y-input should be ignored
			continue
		}
		p.hitbox = p.hitbox.Add(y)
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
		if c.Overlaps(Collision, p) {
			p.hitbox = p.hitbox.Sub(direction)
			break
		}
	}
}
