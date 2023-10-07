package game

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type River struct {
	hitbox        []image.Rectangle
	cachedTexture map[int]*OffsetImage
}

func NewRiver(hitbox ...image.Rectangle) *River {
	return &River{
		hitbox:        hitbox,
		cachedTexture: make(map[int]*OffsetImage),
	}
}

func (r River) Overlaps(layer GameContext, other []image.Rectangle) (bool, error) {
	return DefaultHitboxOverlaps(layer, r, other)
}
func (r River) Origin(layer GameContext) (image.Point, error) {
	bounds, err := r.findBounds(layer)
	return bounds.Min, err
}
func (r River) Size(layer GameContext) (image.Point, error) {
	bounds, err := r.findBounds(layer)
	return bounds.Size(), err
}
func (r River) GetHitbox(layer GameContext) ([]image.Rectangle, error) {
	switch layer {
	case Interaction:
		scaledRects := []image.Rectangle{}
		for _, rect := range r.hitbox {
			riverRect := image.Rectangle{
				Min: rect.Min.Sub(image.Point{20, 20}),
				Max: rect.Max.Add(image.Point{20, 20}),
			}
			scaledRects = append(scaledRects, riverRect)
		}
		return scaledRects, nil
	default:
		return r.hitbox, nil
	}
}

func (r River) findBounds(layer GameContext) (image.Rectangle, error) {
	minX, minY := math.MaxFloat64, math.MaxFloat64
	maxX, maxY := -math.MaxFloat64, -math.MaxFloat64

	hitbox, err := r.GetHitbox(layer)
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

func (r River) DrawAt(screen *ebiten.Image, pos image.Point) error {
	hitbox, err := r.GetHitbox(Render)
	if err != nil {
		return err
	}

	for id := range hitbox {
		texture, ok := r.cachedTexture[id]
		if !ok || texture == nil {
			err := r.generateTexture(id)
			if err != nil {
				return err
			}
			texture = r.cachedTexture[id]
		}

		texture.DrawAt(screen, pos)
	}
	return nil
}

func (r *River) markTextureDirty(id int) {
	if r.cachedTexture[id] == nil {
		return
	}
	r.cachedTexture[id].Image.Dispose()
	r.cachedTexture[id] = nil
}

func (r *River) generateTexture(id int) error {
	hitbox := r.hitbox[id]

	img := ebiten.NewImage(hitbox.Dx(), hitbox.Dy())
	img.Fill(RiverColor)

	lineartImg, err := ApplyLineart(img, r, hitbox)
	if err != nil {
		return err
	}
	img.Dispose()

	origin, err := r.Origin(Render)
	if err != nil {
		return err
	}
	offset := hitbox.Min.Sub(origin)
	lineartImg.Offset = lineartImg.Offset.Add(offset)

	r.cachedTexture[id] = lineartImg

	return nil
}

func (r River) GetZ() (int, error) {
	return -3, nil
}
