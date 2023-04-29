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

func (v Viewport) pointInViewport(point Position) bool {
	if point.Xpos < v.offset.Xpos {
		return false
	}
	if point.Xpos > v.offset.Xpos+v.width {
		return false
	}
	if point.Ypos < v.offset.Ypos {
		return false
	}
	if point.Ypos > v.offset.Ypos+v.height {
		return false
	}
	return true
}

func (v Viewport) Draw(screen *ebiten.Image, drawable Drawable) {
	mapPos := drawable.GetMapPos()

	if v.pointInViewport(mapPos) {
		offsetPos := Position{
			mapPos.Xpos - v.offset.Xpos,
			mapPos.Ypos - v.offset.Ypos,
		}
		drawable.DrawAt(screen, offsetPos)
	}

}

func (v *Viewport) UpdatePosition(player Player) {
	v.offset.Xpos = player.MapPos.Xpos + player.PlayerWidth/2 - v.width/2
	v.offset.Ypos = player.MapPos.Ypos + player.PlayerHeight/2 - v.height/2
}
