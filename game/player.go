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
	game             *Game
	hitbox           image.Rectangle
	currentMoveSpeed float64
}

func NewPlayer(game *Game) Player {
	width, height := assets.Player.Bounds().Size().X, assets.Player.Bounds().Size().Y
	return Player{
		game: game,
		hitbox: image.Rectangle{
			Min: image.Point{-100, -100},
			Max: image.Point{width - 100, height - 100},
		},
		currentMoveSpeed: playerMoveSpeed,
	}
}

func (p Player) DrawAt(screen *ebiten.Image, pos image.Point) error {
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(pos.X), float64(pos.Y))

	screen.DrawImage(assets.Player, &options)
	return nil
}

func (p Player) Overlaps(layer GameContext, other []image.Rectangle) (bool, error) {
	return DefaultHitboxOverlaps(layer, p, other)
}
func (p Player) Origin(layer GameContext) (image.Point, error) {
	bounds, err := p.findBounds(layer)
	return bounds.Min, err
}
func (p Player) Size(layer GameContext) (image.Point, error) {
	bounds, err := p.findBounds(layer)
	return bounds.Size(), err
}
func (p Player) GetHitbox(layer GameContext) ([]image.Rectangle, error) {
	switch layer {
	case Collision, Interaction:
		baseSize := p.hitbox.Size().Y / 2

		playerRect := image.Rectangle{
			Min: p.hitbox.Max.Sub(image.Point{p.hitbox.Dx(), baseSize}),
			Max: p.hitbox.Max,
		}
		return []image.Rectangle{playerRect}, nil
	default:
		return []image.Rectangle{p.hitbox}, nil
	}
}

func (p Player) findBounds(layer GameContext) (image.Rectangle, error) {
	minX, minY := math.MaxFloat64, math.MaxFloat64
	maxX, maxY := -math.MaxFloat64, -math.MaxFloat64

	hitbox, err := p.GetHitbox(layer)
	if err != nil {
		return image.Rectangle{}, err
	}

	for _, seg := range hitbox {
		minX = min(float64(seg.Min.X), minX)
		minY = min(float64(seg.Min.Y), minY)
		maxX = max(float64(seg.Max.X), maxX)
		maxY = max(float64(seg.Max.Y), maxY)
	}
	bounds := image.Rect(int(minX), int(minY), int(maxX), int(maxY))
	return bounds, nil
}

func (p Player) GetZ() (int, error) {
	return 0, nil
}

func (p *Player) Update() (err error) {
	p.currentMoveSpeed, err = p.calculateMovementSpeed()
	if err != nil {
		return
	}

	err = p.movePlayer()
	if err != nil {
		return
	}

	p.tryClimb()
	err = p.handleInteractions()
	if err != nil {
		return
	}

	return err
}

func (p *Player) handleInteractions() error {
	if p.game.input.IsJumping() {

		rivers := []HasHitbox{}
		for _, river := range p.game.rivers {
			rivers = append(rivers, HasHitbox(river))
		}

		newPos, found, err := p.GetSmallestJump(rivers)
		if err != nil {
			return err
		}

		if found {
			p.hitbox = p.hitbox.Sub(p.hitbox.Min).Add(newPos)

			origin, err := p.Origin(Collision)
			if err != nil {
				return err
			}

			offset := origin.Y - p.hitbox.Min.Y
			p.hitbox = p.hitbox.Sub(image.Pt(0, offset))
		}
	}
	return nil
}

func (p Player) GetSmallestJump(jumpables []HasHitbox) (image.Point, bool, error) {
	origin, err := p.Origin(Collision)
	if err != nil {
		return image.Point{}, false, err
	}
	size, err := p.Size(Collision)
	if err != nil {
		return image.Point{}, false, err
	}

	smallestJumpDist := math.MaxFloat64
	smallestJump := image.Point{}
	for _, jumpable := range jumpables {
		hitbox, err := p.GetHitbox(Collision)
		if err != nil {
			return image.Point{}, false, err
		}

		if overlaps, err := jumpable.Overlaps(Interaction, hitbox); err != nil {
			return image.Point{}, false, err

		} else if !overlaps {
			continue
		}

		jumpableHitbox, err := jumpable.GetHitbox(Interaction)
		if err != nil {
			return image.Point{}, false, err
		}

		jumpableCollideHitbox, err := jumpable.GetHitbox(Collision)
		if err != nil {
			return image.Point{}, false, err
		}

		for id, seg := range jumpableHitbox {
			if overlaps, err := p.Overlaps(Collision, []image.Rectangle{seg}); err != nil {
				return image.Point{}, false, err

			} else if !overlaps {
				continue
			}
			segHitbox := jumpableCollideHitbox[id]

			if origin.X >= segHitbox.Max.X { //right
				newPoint, err := testJump(jumpable, segHitbox,
					func(segment image.Rectangle) image.Point {
						return image.Pt(segment.Min.X, origin.Y).Sub(image.Pt(size.X, 0))
					},
					func(origin image.Point) image.Rectangle {
						return image.Rectangle{origin, origin.Add(size)}
					},
				)
				if err != nil {
					return image.Point{}, false, err
				}

				updateJumpIfSmaller(origin, newPoint, &smallestJumpDist, &smallestJump)
			} else if origin.X <= segHitbox.Min.X { //left
				newPoint, err := testJump(jumpable, segHitbox,
					func(segment image.Rectangle) image.Point {
						return image.Pt(segment.Max.X, origin.Y)
					},
					func(origin image.Point) image.Rectangle {
						return image.Rectangle{origin, origin.Add(size)}
					},
				)
				if err != nil {
					return image.Point{}, false, err
				}

				updateJumpIfSmaller(origin, newPoint, &smallestJumpDist, &smallestJump)
			}

			if origin.Y >= segHitbox.Max.Y { //bottom
				newPoint, err := testJump(jumpable, segHitbox,
					func(segment image.Rectangle) image.Point {
						return image.Pt(origin.X, segment.Min.Y).Sub(image.Pt(0, size.Y))
					},
					func(origin image.Point) image.Rectangle {
						return image.Rectangle{origin, origin.Add(size)}
					},
				)
				if err != nil {
					return image.Point{}, false, err
				}

				updateJumpIfSmaller(origin, newPoint, &smallestJumpDist, &smallestJump)
			} else if origin.Y <= segHitbox.Min.Y { //top
				newPoint, err := testJump(jumpable, segHitbox,
					func(segment image.Rectangle) image.Point {
						return image.Pt(origin.X, segment.Max.Y)
					},
					func(origin image.Point) image.Rectangle {
						return image.Rectangle{origin, origin.Add(size)}
					},
				)
				if err != nil {
					return image.Point{}, false, err
				}

				updateJumpIfSmaller(origin, newPoint, &smallestJumpDist, &smallestJump)
			}
		}
	}
	return smallestJump, smallestJumpDist != math.MaxFloat64, nil
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
func testJump(fullObj HasHitbox, jumpSeg image.Rectangle, makeJump func(image.Rectangle) image.Point, pRect func(image.Point) image.Rectangle) (image.Point, error) {
	newPoint := makeJump(jumpSeg)

	newRect := pRect(newPoint)

	overlaps, err := fullObj.Overlaps(Collision, []image.Rectangle{newRect})
	if err != nil {
		return image.Point{}, err
	}

	for overlaps {
		hitbox, err := fullObj.GetHitbox(Collision)
		if err != nil {
			return image.Point{}, err
		}
		for _, newSegTest := range hitbox {
			if !newRect.Overlaps(newSegTest) {
				continue
			}
			newPoint = makeJump(newSegTest)
			break
		}
		newRect = pRect(newPoint)
	}
	return newPoint, nil
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

func (p Player) calculateMovementSpeed() (float64, error) {
	currentClimbable, err := p.findCurrentClimbable()
	if err != nil {
		return 0, err
	}

	if currentClimbable == nil {
		return playerMoveSpeed, nil
	}
	return playerMoveSpeed * currentClimbable.GetClimbSpeed(), nil
}

func (p *Player) tryClimb() error {
	if climbable, err := p.findCurrentClimbable(); err != nil {
		return err

	} else if climbable != nil {
		if p.game.input.IsClimbing() {
			p.hitbox = p.hitbox.Sub(image.Point{
				Y: int(p.currentMoveSpeed),
			})
		} else {
			p.hitbox = p.hitbox.Add(image.Point{
				Y: int(p.currentMoveSpeed),
			})
		}
	}
	return nil
}

func (p Player) findCurrentClimbable() (found Climbable, err error) {

	climbables := []Climbable{}
	for _, incline := range p.game.inclines {
		climbables = append(climbables, Climbable(incline))
	}

	hitbox, err := p.GetHitbox(Collision)
	if err != nil {
		return nil, err
	}

	for _, c := range climbables {
		if overlaps, err := c.Overlaps(Collision, hitbox); err != nil {
			return nil, err

		} else if overlaps {
			found = c
			break
		}
	}
	return
}

func (p *Player) movePlayer() error {
	scalar := p.currentMoveSpeed

	steps := int(scalar)
	stepSize := scalar / float64(steps)

	x := image.Point{X: int(float64(-p.game.input.LeftImpulse()) * stepSize)}
	y := image.Point{Y: int(float64(-p.game.input.ForwardsImpulse()) * stepSize)}

	preMoveClimbable, err := p.findCurrentClimbable()
	if err != nil {
		return err
	}
	climbingPreMove := preMoveClimbable != nil

	for i := 0; i < steps; i++ {

		p.hitbox = p.hitbox.Add(x)
		if !climbingPreMove {
			err := p.fixCollisions(x)
			if err != nil {
				return err
			}
		}

		if climbable, err := p.findCurrentClimbable(); err != nil {
			return err

		} else if climbable != nil {
			//if currently climbing or decending Y-input should be ignored
			continue
		}
		p.hitbox = p.hitbox.Add(y)

		currentClimbable, err := p.findCurrentClimbable()
		if err != nil {
			return err
		}

		if p.game.input.IsClimbing() && currentClimbable != nil && y.Y <= 0 {
			//if hitting the bottom of a climbable while trying to climb don't fix collisions
			continue
		}
		if !p.game.input.IsClimbing() && currentClimbable != nil && y.Y >= 0 {
			//if hitting the top of a climbable and not climbing don't fix collisions
			continue
		}
		err = p.fixCollisions(y)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Player) fixCollisions(direction image.Point) error {
	collidables := []HasHitbox{}
	for _, incline := range p.game.inclines {
		collidables = append(collidables, HasHitbox(incline))
	}
	for _, river := range p.game.rivers {
		collidables = append(collidables, HasHitbox(river))
	}

	hitbox, err := p.GetHitbox(Collision)
	if err != nil {
		return err
	}

	for _, c := range collidables {
		if overlaps, err := c.Overlaps(Collision, hitbox); err != nil {
			return err

		} else if overlaps {
			p.hitbox = p.hitbox.Sub(direction)
			break
		}
	}
	return nil
}
