package game

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Climbable interface {
	HasHitbox

	GetClimbSpeed() float64
}

type Incline struct {
	Collision image.Rectangle
}

func (i Incline) Hitbox(GameContext) image.Rectangle {
	return i.Collision
}

func (i Incline) DrawAt(screen *ebiten.Image, pos image.Point) {
	lineartW := 10
	lineartC := color.RGBA{0, 0, 0, 255}
	lineartOptions := ebiten.DrawImageOptions{}

	//main background colour
	img := ebiten.NewImage(i.Collision.Dx()+lineartW*2, i.Collision.Dy()+lineartW*2)
	img.Fill(color.RGBA{117, 88, 69, 255})

	//lineart
	horizontal := ebiten.NewImage(img.Bounds().Dx(), lineartW)
	vertical := ebiten.NewImage(lineartW, img.Bounds().Dy())
	horizontal.Fill(lineartC)
	vertical.Fill(lineartC)

	img.DrawImage(horizontal, nil) //top
	lineartOptions.GeoM.Reset()
	lineartOptions.GeoM.Translate(0, float64(img.Bounds().Dy()-lineartW))
	img.DrawImage(horizontal, &lineartOptions) //bottom

	img.DrawImage(vertical, nil) //left
	lineartOptions.GeoM.Reset()
	lineartOptions.GeoM.Translate(float64(img.Bounds().Dx()-lineartW), 0)
	img.DrawImage(vertical, &lineartOptions) //right

	//actual draw
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(pos.X-lineartW), float64(pos.Y-lineartW))

	screen.DrawImage(img, &options)
}

func (i Incline) GetZ() int {
	return -2
}

func (i Incline) GetClimbSpeed() float64 {
	return 0.6
}
