package game

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	lineartW    = 10
	curveRadius = float32(lineartW) / 1.25
	curveOffset = int(curveRadius - float32(lineartW)/2)

	padding = lineartW/2 + int(curveRadius*2) - lineartW
)

func pad(in image.Point, sub bool) image.Point {
	if sub {
		return in.Sub(image.Point{padding, padding})
	} else {
		return in.Add(image.Point{padding, padding})
	}
}

type OffsetImage struct {
	Image  *ebiten.Image
	Offset image.Point
}

func (o OffsetImage) Overlaps(layer GameContext, other []image.Rectangle) (bool, error) {
	return DefaultHitboxOverlaps(layer, o, other)
}

func (o OffsetImage) Origin(_ GameContext) (image.Point, error) {
	return o.Offset, nil
}

func (o OffsetImage) Size(_ GameContext) (image.Point, error) {
	return o.Image.Bounds().Size(), nil
}

func (o OffsetImage) GetHitbox(_ GameContext) ([]image.Rectangle, error) {
	return []image.Rectangle{o.Image.Bounds()}, nil
}

func (o *OffsetImage) DrawAt(screen *ebiten.Image, pos image.Point) error {
	ops := ebiten.DrawImageOptions{}
	ops.GeoM.Translate(float64(pos.X), float64(pos.Y))
	ops.GeoM.Translate(float64(o.Offset.X), float64(o.Offset.Y))

	screen.DrawImage(o.Image, &ops)
	return nil
}

/*
Takes an image and adds lineart to its sides before returning this modified image
as a an OffsetImage. The input image should be rectangular (either just a block
colour or pattern on it). More complex shapes should be passed in segment at a time.

If part of a larger image then the other segments should be given as the neighbours
so they can be used to correctly calculate where lineart should be.
*/
func ApplyLineart(blankImage *ebiten.Image, segmentOrigin image.Point, neighbours []image.Rectangle) (*OffsetImage, error) {
	// image setup
	newBounds := pad(pad(blankImage.Bounds().Size(), false), false) // 2 pads, 1 for each side of the axis
	img := ebiten.NewImage(newBounds.X, newBounds.Y)

	// original image
	ops := ebiten.DrawImageOptions{}
	ops.GeoM.Translate(float64(padding), float64(padding))
	img.DrawImage(blankImage, &ops)

	levelPos := pad(segmentOrigin, true)

	// line art
	err := drawSide(img, levelPos, neighbours, top)
	if err != nil {
		return nil, err
	}
	err = drawSide(img, levelPos, neighbours, bottom)
	if err != nil {
		return nil, err
	}
	err = drawSide(img, levelPos, neighbours, left)
	if err != nil {
		return nil, err
	}
	err = drawSide(img, levelPos, neighbours, right)
	if err != nil {
		return nil, err
	}

	return &OffsetImage{
		Image:  img,
		Offset: pad(image.Point{}, true),
	}, nil
}

func drawSide(toDrawTo *ebiten.Image, levelPos image.Point, neighbours []image.Rectangle, side lineartSide) error {
	originalSeg := image.Rectangle{
		Min: pad(toDrawTo.Bounds().Min, false),
		Max: pad(toDrawTo.Bounds().Max, true),
	}

	var current image.Point

	switch side {
	case top:
		current = originalSeg.Min
	case bottom:
		current = image.Pt(originalSeg.Min.X, originalSeg.Max.Y-1)
	case left:
		current = originalSeg.Min
	case right:
		current = image.Pt(originalSeg.Max.X-1, originalSeg.Min.Y)
	default:
		return nil
	}

	var delta image.Point
	if side.isHorizontal() {
		delta = image.Pt(1, 0)
	} else {
		delta = image.Pt(0, 1)
	}

	last := current
	first := true

	lineStart := current
	for {
		if overlaps, overlapRect := overlapsAny(current.Add(levelPos), neighbours); overlaps {
			//overlapping

			if first {
				first = false

				//jumpPast
				if side.isHorizontal() {
					current.X = overlapRect.Sub(levelPos).Max.X
				} else {
					current.Y = overlapRect.Sub(levelPos).Max.Y
				}
				lineStart = current

				drawCorner(toDrawTo, current, side, 1, -1)

				continue
			}

			if overlaps, _ := overlapsAny(last.Add(levelPos), neighbours); overlaps {
			} else {
				//just started overlapping

				//draw line
				toDrawTo.DrawImage(generateLineSegment(lineStart, current, side.isHorizontal()))

				//jumpPast
				if side.isHorizontal() {
					current.X = overlapRect.Sub(levelPos).Max.X
				} else {
					current.Y = overlapRect.Sub(levelPos).Max.Y
				}
				lineStart = current
			}
		} else if first {
			drawCorner(toDrawTo, current, side, 1, 1)
		}
		last = current
		current = current.Add(delta)
		if !current.In(originalSeg) {
			if last.Eq(lineStart) {
				break
			}

			toDrawTo.DrawImage(generateLineSegment(lineStart, last, side.isHorizontal()))

			drawCorner(toDrawTo, current.Sub(delta), side, -1, 1)

			break
		}
		first = false
	}

	return nil
}

func drawCorner(toDrawTo *ebiten.Image, current image.Point, side lineartSide, fromPoint int, reflexMod int) {

	var cx, cy float32
	if side.isHorizontal() {
		cx = float32(current.X + curveOffset*fromPoint)
	} else if side == left {
		cx = float32(current.X + curveOffset*reflexMod)
	} else if side == right {
		cx = float32(current.X - curveOffset*reflexMod)
	}
	if !side.isHorizontal() {
		cy = float32(current.Y + curveOffset*fromPoint)
	} else if side == top {
		cy = float32(current.Y + curveOffset*reflexMod)
	} else if side == bottom {
		cy = float32(current.Y - curveOffset*reflexMod)
	}
	vector.DrawFilledCircle(toDrawTo, cx, cy, curveRadius, LineartColor, false)
}

func generateLineSegment(start image.Point, end image.Point, isHorizontal bool) (*ebiten.Image, *ebiten.DrawImageOptions) {
	var lineSeg *ebiten.Image
	lineSegOps := ebiten.DrawImageOptions{}
	if isHorizontal {
		lineSeg = ebiten.NewImage(end.X-start.X-lineartW, lineartW)
		lineSegOps.GeoM.Translate(float64(lineartW/2), -float64(lineartW/2))
	} else {
		lineSeg = ebiten.NewImage(lineartW, end.Y-start.Y-lineartW)
		lineSegOps.GeoM.Translate(-float64(lineartW/2), float64(lineartW/2))
	}
	lineSeg.Fill(LineartColor)

	lineSegOps.GeoM.Translate(float64(start.X), float64(start.Y))
	return lineSeg, &lineSegOps
}

func overlapsAny(point image.Point, rects []image.Rectangle) (bool, image.Rectangle) {
	for _, rect := range rects {
		if point.In(rect) {
			return true, rect
		}
	}
	return false, image.Rectangle{}
}

type lineartSide int

const (
	left lineartSide = iota
	right
	top
	bottom
)

func (l lineartSide) isHorizontal() bool {
	switch l {
	case bottom, top:
		return true
	case left, right:
		return false
	default:
		return false
	}
}
