package game

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/moltenwolfcub/Forest-Game/assets"
)

type InventoryOverlay struct {
	game   *Game
	isOpen bool
}

func NewInventoryOverlay(game *Game) InventoryOverlay {
	return InventoryOverlay{
		game: game,
	}
}

func (i *InventoryOverlay) Update() {
	if i.game.input.IsTogglingInventory() {
		i.Toggle()
	}
}

func (i *InventoryOverlay) Toggle() {
	i.isOpen = !i.isOpen
}

func (i InventoryOverlay) IsFocused() bool {
	return !i.isOpen
}

func (i InventoryOverlay) Overlaps(layer GameContext, other []image.Rectangle) (bool, error) {
	return DefaultHitboxOverlaps(layer, i, other)
}

func (i InventoryOverlay) Origin(layer GameContext) (image.Point, error) {
	s, err := i.Size(layer)
	if err != nil {
		return image.Point{}, err
	}

	x := (WindowWidth - s.X) / 2
	y := (WindowHeight - s.Y) / 2
	return image.Pt(x, y), nil
}

func (i InventoryOverlay) Size(_ GameContext) (image.Point, error) {
	return image.Pt(960, 960), nil
}

func (i InventoryOverlay) GetHitbox(layer GameContext) ([]image.Rectangle, error) {
	origin, err := i.Origin(layer)
	if err != nil {
		return nil, err
	}
	size, err := i.Size(layer)
	if err != nil {
		return nil, err
	}

	r := image.Rectangle{
		Min: origin,
		Max: origin.Add(size),
	}

	return []image.Rectangle{r}, nil
}

func (i InventoryOverlay) DrawAt(screen *ebiten.Image, pos image.Point) error {
	if i.IsFocused() {
		return nil
	}

	options := ebiten.DrawImageOptions{}

	targetSize, _ := i.Size(Render)
	sf := float64(targetSize.Y) / float64(assets.Inventory.Bounds().Size().Y)
	options.GeoM.Scale(float64(sf), float64(sf))

	options.GeoM.Translate(float64(pos.X), float64(pos.Y))

	screen.DrawImage(assets.Inventory, &options)
	return nil
}
