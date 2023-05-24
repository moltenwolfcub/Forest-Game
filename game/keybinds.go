package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type KeyPressType int

const (
	IsPressed KeyPressType = iota
	JustPressed
	JustReleased
)

type Keybind struct {
	currentKey ebiten.Key
	pressType  KeyPressType
}

func NewKeybind(key ebiten.Key, pressType KeyPressType) Keybind {
	return Keybind{
		currentKey: key,
		pressType:  pressType,
	}
}

func (k Keybind) Triggered() (activated bool) {
	switch k.pressType {
	case IsPressed:
		activated = ebiten.IsKeyPressed(k.currentKey)
	case JustPressed:
		activated = inpututil.IsKeyJustPressed(k.currentKey)
	case JustReleased:
		activated = inpututil.IsKeyJustReleased(k.currentKey)
	}
	return
}

type Keybinds struct {
	Forwards  Keybind
	Backwards Keybind
	Left      Keybind
	Right     Keybind
}

func NewKeybinds() Keybinds {
	bindings := Keybinds{}

	bindings.Forwards = NewKeybind(ebiten.KeyW, IsPressed)
	bindings.Backwards = NewKeybind(ebiten.KeyS, IsPressed)
	bindings.Left = NewKeybind(ebiten.KeyA, IsPressed)
	bindings.Right = NewKeybind(ebiten.KeyD, IsPressed)

	return bindings
}
