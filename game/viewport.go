package game

import "github.com/hajimehoshi/ebiten/v2"

type Drawable interface {
	GetMapPos() Position

	DrawAt(*ebiten.Image, Position)
}

type Viewport struct {
	width, height int
	offset        Position
}

func NewViewport() Viewport {
	return Viewport{
		width:  WindowWidth,
		height: WindowHeight,
		offset: Position{0, 0},
	}
}

func (v *Viewport) pointInViewport(point Position) bool {
	if point.Xpos < v.offset.Xpos {
		return false
	}
	if point.Xpos > v.offset.Xpos+float64(v.width) {
		return false
	}
	if point.Ypos < v.offset.Ypos {
		return false
	}
	if point.Ypos > v.offset.Ypos+float64(v.height) {
		return false
	}
	return true
}

func (v *Viewport) Draw(screen *ebiten.Image, drawable Drawable) {
	mapPos := drawable.GetMapPos()

	if v.pointInViewport(mapPos) {
		offsetPos := Position{
			mapPos.Xpos - v.offset.Xpos,
			mapPos.Ypos - v.offset.Ypos,
		}
		drawable.DrawAt(screen, offsetPos)
	}

}
