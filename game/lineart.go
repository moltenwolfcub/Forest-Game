package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// Takes an image and adds lineart to all sides of it before returning the new image.
// The new offset from the original image is also returned so the image can be seemlessly
// inserted.(This is useful for keeping it in line with hitboxes)
func ApplyLineart(oldImg *ebiten.Image) (*ebiten.Image, *ebiten.DrawImageOptions) {
	lineartW := 10
	lineartC := color.RGBA{0, 0, 0, 255}
	lineartOptions := ebiten.DrawImageOptions{}

	//image setup
	imgOffset := ebiten.DrawImageOptions{}
	imgOffset.GeoM.Translate(-float64(lineartW)/2, -float64(lineartW)/2)
	img := ebiten.NewImage(oldImg.Bounds().Dx()+lineartW, oldImg.Bounds().Dy()+lineartW)

	//original image
	oldImgOffset := ebiten.DrawImageOptions{}
	oldImgOffset.GeoM.Translate(float64(lineartW)/2, float64(lineartW)/2)
	img.DrawImage(oldImg, &oldImgOffset)

	//line segments
	horizontal := ebiten.NewImage(img.Bounds().Dx(), lineartW)
	vertical := ebiten.NewImage(lineartW, img.Bounds().Dy())
	horizontal.Fill(lineartC)
	vertical.Fill(lineartC)

	//line art
	img.DrawImage(horizontal, nil) //top

	lineartOptions.GeoM.Reset()
	lineartOptions.GeoM.Translate(0, float64(img.Bounds().Dy()-lineartW))
	img.DrawImage(horizontal, &lineartOptions) //bottom

	img.DrawImage(vertical, nil) //left

	lineartOptions.GeoM.Reset()
	lineartOptions.GeoM.Translate(float64(img.Bounds().Dx()-lineartW), 0)
	img.DrawImage(vertical, &lineartOptions) //right

	return img, &imgOffset
}
