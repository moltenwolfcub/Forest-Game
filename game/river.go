package game

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	InteractionRange = 20
)

type River struct {
	origin   image.Point
	segments []*RiverSegment
}

func NewRiver(origin image.Point, segs ...*RiverSegment) *River {
	return &River{
		origin:   origin,
		segments: segs,
	}
}

func (r River) Overlaps(layer GameContext, other []image.Rectangle) (bool, error) {
	return DefaultHitboxOverlaps(layer, r, other)
}

func (r River) Origin(_ GameContext) (image.Point, error) {
	return r.origin, nil
}

func (r River) Size(ctx GameContext) (image.Point, error) {
	segments, err := r.GetHitbox(ctx)
	if err != nil {
		return image.Point{}, err
	}

	minX, minY, maxX, maxY := math.MaxInt, math.MaxInt, -math.MaxInt, -math.MaxInt
	for _, seg := range segments {
		minX = min(seg.Min.X, minX)
		minY = min(seg.Min.Y, minY)
		maxX = max(seg.Max.X, maxX)
		maxY = max(seg.Max.Y, maxY)
	}
	fullBounds := image.Rect(minX, minY, maxX, maxY)
	return fullBounds.Size(), nil
}

func (r River) GetHitbox(layer GameContext) ([]image.Rectangle, error) {
	hitboxes := []image.Rectangle{}
	for _, seg := range r.segments {
		rect, err := seg.GetHitbox(layer)
		if err != nil {
			return nil, err
		}

		hitboxes = append(hitboxes, rect...)
	}
	return hitboxes, nil
}

// needs to change when rendering is moved to the segments
func (r River) DrawAt(screen *ebiten.Image, pos image.Point) error {
	for _, seg := range r.segments {
		if seg.cachedTexture == nil {
			img, err := r.generateTexture(seg)
			if err != nil {
				return err
			}
			seg.cachedTexture = img
		}
		texture := seg.cachedTexture

		texture.DrawAt(screen, pos)
	}
	return nil
}

func (r *River) generateTexture(segment *RiverSegment) (*OffsetImage, error) {
	hitbox := segment.hitbox //should be getHitbox() will change when moving rendering

	img := ebiten.NewImage(hitbox.Dx(), hitbox.Dy())
	img.Fill(RiverColor)

	lineartImg, err := ApplyLineart(img, r, hitbox)
	if err != nil {
		return nil, err
	}
	img.Dispose()

	origin, err := r.Origin(Render)
	if err != nil {
		return nil, err
	}
	offset := hitbox.Min.Sub(origin)
	lineartImg.Offset = lineartImg.Offset.Add(offset)

	return lineartImg, nil
}

// func (r *River) markTextureDirty(id int) {
// 	if r.cachedTexture[id] == nil {
// 		return
// 	}
// 	r.cachedTexture[id].Image.Dispose()
// 	r.cachedTexture[id] = nil
// }

func (r River) GetZ() (int, error) {
	return -3, nil
}

type RiverSegment struct {
	// parent        *River
	hitbox        image.Rectangle
	cachedTexture *OffsetImage
}

func NewRiverSegment(rect image.Rectangle) *RiverSegment {
	return &RiverSegment{
		hitbox:        rect,
		cachedTexture: nil,
	}
}

func (r RiverSegment) Overlaps(ctx GameContext, other []image.Rectangle) (bool, error) {
	return DefaultHitboxOverlaps(ctx, r, other)
}

func (r RiverSegment) Origin(_ GameContext) (image.Point, error) {
	return r.hitbox.Min, nil
}

func (r RiverSegment) Size(ctx GameContext) (image.Point, error) {
	switch ctx {
	case Interaction:
		scaled := r.hitbox.Size().Add(image.Pt(2*InteractionRange, 2*InteractionRange))
		return scaled, nil
	default:
		return r.hitbox.Size(), nil
	}
}

func (r RiverSegment) GetHitbox(ctx GameContext) ([]image.Rectangle, error) {
	switch ctx {
	case Interaction:
		scaledHitbox := image.Rectangle{
			Min: r.hitbox.Min.Sub(image.Pt(InteractionRange, InteractionRange)),
			Max: r.hitbox.Max.Add(image.Pt(InteractionRange, InteractionRange)),
		}
		return []image.Rectangle{scaledHitbox}, nil
	default:
		return []image.Rectangle{r.hitbox}, nil
	}
}
