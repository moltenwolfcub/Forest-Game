package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type InputHandler struct {
	forwards  ebiten.Key
	backwards ebiten.Key
	left      ebiten.Key
	right     ebiten.Key
	jump      ebiten.Key
	climb     ebiten.Key
	inventory ebiten.Key
}

func NewInputHandler() InputHandler {
	return InputHandler{
		forwards:  ebiten.KeyW,
		backwards: ebiten.KeyS,
		left:      ebiten.KeyA,
		right:     ebiten.KeyD,
		jump:      ebiten.KeySpace,
		climb:     ebiten.KeySpace,
		inventory: ebiten.KeyE,
	}
}

func (p InputHandler) ForwardsImpulse() float64 {
	return calculateImpulse(ebiten.IsKeyPressed(p.forwards), ebiten.IsKeyPressed(p.backwards))
}
func (p InputHandler) LeftImpulse() float64 {
	return calculateImpulse(ebiten.IsKeyPressed(p.left), ebiten.IsKeyPressed(p.right))
}
func (p InputHandler) IsJumping() bool {
	return inpututil.IsKeyJustPressed(p.jump)
}
func (p InputHandler) IsClimbing() bool {
	return ebiten.IsKeyPressed(p.jump)
}
func (p InputHandler) IsTogglingInventory() bool {
	return inpututil.IsKeyJustPressed(p.inventory)
}

func calculateImpulse(positive bool, negative bool) float64 {
	if positive == negative {
		return 0
	}
	if positive {
		return 1
	}
	return -1
}
