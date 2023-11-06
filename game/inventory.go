package game

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
