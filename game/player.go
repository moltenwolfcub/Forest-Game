package game

import (
	_ "embed"
	"image"
	_ "image/png"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/moltenwolfcub/Forest-Game/assets"
)

const (
	playerMoveSpeed float64 = 11.5
)

type Player struct {
	Delta            image.Point
	hitbox           image.Rectangle
	Climbing         bool
	RiverJumping     bool
	currentMoveSpeed float64
}

func NewPlayer() Player {
	width, height := assets.Player.Bounds().Size().X, assets.Player.Bounds().Size().Y
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

	screen.DrawImage(assets.Player, &options)
}

func (p Player) Overlaps(layer GameContext, other []image.Rectangle) bool {
	return DefaultHitboxOverlaps(layer, p, other)
}
func (p Player) Origin(layer GameContext) image.Point {
	bounds := p.findBounds(layer)
	return bounds.Min
}
func (p Player) Size(layer GameContext) image.Point {
	bounds := p.findBounds(layer)
	return bounds.Size()
}
func (p Player) GetHitbox(layer GameContext) []image.Rectangle {
	switch layer {
	case Collision, Interaction:
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

func (p Player) findBounds(layer GameContext) image.Rectangle {
	minX, minY := math.MaxFloat64, math.MaxFloat64
	maxX, maxY := -math.MaxFloat64, -math.MaxFloat64
	for _, seg := range p.GetHitbox(layer) {
		minX = min(float64(seg.Min.X), minX)
		minY = min(float64(seg.Min.Y), minY)
		maxX = max(float64(seg.Max.X), maxX)
		maxY = max(float64(seg.Max.Y), maxY)
	}
	bounds := image.Rect(int(minX), int(minY), int(maxX), int(maxY))
	return bounds
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

		newPos, found := p.GetSmallestJump(interactables)

		if found {
			p.hitbox = p.hitbox.Sub(p.hitbox.Min).Add(newPos)
			offset := p.Origin(Collision).Y - p.hitbox.Min.Y
			p.hitbox = p.hitbox.Sub(image.Pt(0, offset))
		}
	}
}

func (p Player) GetSmallestJump(jumpables []HasHitbox) (point image.Point, found bool) {
	origin := p.Origin(Collision)
	size := p.Size(Collision)

	smallestJumpDist := math.MaxFloat64
	smallestJump := image.Point{}
	for _, jumpable := range jumpables {
		if !jumpable.Overlaps(Interaction, p.GetHitbox(Collision)) {
			continue
		}

		for id, seg := range jumpable.GetHitbox(Interaction) {
			if !p.Overlaps(Collision, []image.Rectangle{seg}) {
				continue
			}
			segHitbox := jumpable.GetHitbox(Collision)[id]

			if origin.X >= segHitbox.Max.X { //right
				newPoint := testJump(jumpable, segHitbox,
					func(segment image.Rectangle) image.Point {
						return image.Pt(segment.Min.X, origin.Y).Sub(image.Pt(size.X, 0))
					},
					func(origin image.Point) image.Rectangle {
						return image.Rectangle{origin, origin.Add(size)}
					},
				)

				updateJumpIfSmaller(origin, newPoint, &smallestJumpDist, &smallestJump)
			} else if origin.X <= segHitbox.Min.X { //left
				newPoint := testJump(jumpable, segHitbox,
					func(segment image.Rectangle) image.Point {
						return image.Pt(segment.Max.X, origin.Y)
					},
					func(origin image.Point) image.Rectangle {
						return image.Rectangle{origin, origin.Add(size)}
					},
				)

				updateJumpIfSmaller(origin, newPoint, &smallestJumpDist, &smallestJump)
			}

			if origin.Y >= segHitbox.Max.Y { //bottom
				newPoint := testJump(jumpable, segHitbox,
					func(segment image.Rectangle) image.Point {
						return image.Pt(origin.X, segment.Min.Y).Sub(image.Pt(0, size.Y))
					},
					func(origin image.Point) image.Rectangle {
						return image.Rectangle{origin, origin.Add(size)}
					},
				)

				updateJumpIfSmaller(origin, newPoint, &smallestJumpDist, &smallestJump)
			} else if origin.Y <= segHitbox.Min.Y { //top
				newPoint := testJump(jumpable, segHitbox,
					func(segment image.Rectangle) image.Point {
						return image.Pt(origin.X, segment.Max.Y)
					},
					func(origin image.Point) image.Rectangle {
						return image.Rectangle{origin, origin.Add(size)}
					},
				)

				updateJumpIfSmaller(origin, newPoint, &smallestJumpDist, &smallestJump)
			}
		}
	}
	return smallestJump, smallestJumpDist != math.MaxFloat64
}

// Finds the location the player would jump to taking into account multiple
// segments that might need to be jumped. If the player would land in another
// segment it continues to test further jumps on that new segment.
//
// Once that new land location is found it gets returned.
//
// makeJump returns the origin of the new player hitbox after jumping the given
// segment. This function handles how the jump should be made (what direction)
//
// pRect generates a version of the player's hitbox to test for collisions after
// each jump without actually moving the player yet. It should just return a hitbox
// of the player's size with it's origin at the provided point.
func testJump(fullObj HasHitbox, jumpSeg image.Rectangle, makeJump func(image.Rectangle) image.Point, pRect func(image.Point) image.Rectangle) image.Point {
	newPoint := makeJump(jumpSeg)

	newRect := pRect(newPoint)
	for fullObj.Overlaps(Collision, []image.Rectangle{newRect}) {
		for _, newSegTest := range fullObj.GetHitbox(Collision) {
			if !newRect.Overlaps(newSegTest) {
				continue
			}
			newPoint = makeJump(newSegTest)
			break
		}
		newRect = pRect(newPoint)
	}
	return newPoint
}

// Calculates jump distance between points before and new and updates the
// pointers with the new distance and point respectively if the new distance
// is smaller than the previous smallest.
func updateJumpIfSmaller(before image.Point, new image.Point, dist *float64, point *image.Point) {
	delta := math.Hypot(float64(before.X-new.X), float64(before.Y-new.Y))
	if delta < *dist {
		*point = new
		*dist = delta
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
		if c.Overlaps(Collision, p.GetHitbox(Collision)) {
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
		if c.Overlaps(Collision, p.GetHitbox(Collision)) {
			p.hitbox = p.hitbox.Sub(direction)
			break
		}
	}
}
