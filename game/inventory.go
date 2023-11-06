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
	return assets.Inventory.Bounds().Size(), nil
}

func (i InventoryOverlay) GetHitbox(_ GameContext) ([]image.Rectangle, error) {
	return []image.Rectangle{assets.Inventory.Bounds()}, nil
}

func (i InventoryOverlay) DrawAt(screen *ebiten.Image, pos image.Point) error {
	if i.IsFocused() {
		return nil
	}

	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(pos.X), float64(pos.Y))

	screen.DrawImage(assets.Inventory, &options)
	return nil
}
