package game

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type HasPosition interface {
	GetPos() image.Point
}

type Drawable interface {
	HasPosition

	DrawAt(*ebiten.Image, image.Point)
}

type Viewport struct {
	width, height int
	offset        image.Point
}

func NewViewport() Viewport {
	return Viewport{
		width:  WindowWidth,
		height: WindowHeight,
		offset: image.Point{0, 0},
	}
}

func (v Viewport) pointInViewport(point image.Point) bool {
	if point.X < v.offset.X {
		return false
	}
	if point.X > v.offset.X+v.width {
		return false
	}
	if point.Y < v.offset.Y {
		return false
	}
	if point.Y > v.offset.Y+v.height {
		return false
	}
	return true
}

func (v Viewport) DrawToMap(mapLayer *ebiten.Image, drawable Drawable) {
	mapPos := drawable.GetPos()

	if v.pointInViewport(mapPos) {
		offsetPos := image.Point{
			mapPos.X - v.offset.X,
			mapPos.Y - v.offset.Y,
		}
		drawable.DrawAt(mapLayer, offsetPos)
	}

}
func (v Viewport) DrawToHUD(hudLayer *ebiten.Image, drawable Drawable) {
	drawable.DrawAt(hudLayer, drawable.GetPos())
}

func (v *Viewport) UpdatePosition(player Player) {
	v.offset.X = player.MapPos.X + player.PlayerWidth/2 - v.width/2
	v.offset.Y = player.MapPos.Y + player.PlayerHeight/2 - v.height/2
}
