package game

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/moltenwolfcub/Forest-Game/errors"
)

const (
	InteractionRange = 20
)

type River struct {
	origin   image.Point
	segments []*RiverSegment
}

func NewRiver(origin image.Point, segs ...*RiverSegment) *River {
	river := &River{
		origin:   origin,
		segments: segs,
	}
	for _, seg := range segs {
		seg.parent = river
	}
	return river
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

func (r *River) DrawAt(screen *ebiten.Image, pos image.Point) error {
	for _, seg := range r.segments {
		seg.DrawAt(screen, pos)
	}
	return nil
}

func (r River) GetZ() (int, error) {
	return -3, nil
}

type RiverSegment struct {
	parent        *River
	hitbox        image.Rectangle
	cachedTexture *OffsetImage
}

func NewRiverSegment(rect image.Rectangle) *RiverSegment {
	return &RiverSegment{
		hitbox:        rect,
		parent:        nil,
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

func (r *RiverSegment) DrawAt(screen *ebiten.Image, pos image.Point) error {
	if r.cachedTexture == nil {
		img, err := r.generateTexture()
		if err != nil {
			return err
		}
		r.cachedTexture = img
	}
	texture := r.cachedTexture

	texture.DrawAt(screen, pos)
	return nil
}

func (r *RiverSegment) generateTexture() (*OffsetImage, error) {
	//use full river if it has a parent else run use self as parent for stand alone use
	var fullObj HasHitbox
	if r.parent == nil {
		fullObj = r
	} else {
		fullObj = r.parent
	}

	hitboxSet, err := r.GetHitbox(Render)
	if err != nil {
		return nil, err
	}
	if count := len(hitboxSet); count != 1 {
		return nil, errors.NewMultiHitboxRiverSegmentError(count)
	}
	hitbox := hitboxSet[0]

	img := ebiten.NewImage(hitbox.Dx(), hitbox.Dy())
	img.Fill(RiverColor)

	lineartImg, err := ApplyLineart(img, fullObj, hitbox)
	if err != nil {
		return nil, err
	}
	img.Dispose()

	origin, err := fullObj.Origin(Render)
	if err != nil {
		return nil, err
	}
	offset := hitbox.Min.Sub(origin)
	lineartImg.Offset = lineartImg.Offset.Add(offset)

	return lineartImg, nil
}

func (r *RiverSegment) markTextureDirty(id int) {
	if r.cachedTexture == nil {
		return
	}
	r.cachedTexture.Image.Dispose()
	r.cachedTexture = nil
}
