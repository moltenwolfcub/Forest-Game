package game

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// Takes an image and adds lineart to all sides of it before returning the new image.
// The image given should be a single rect out of the full object so that it can be
// computed seperately. The new offset from the original image is also returned so the
// image can be seemlessly inserted.(This is useful for keeping it in line with hitboxes)
func ApplyLineart(oldImgSeg *ebiten.Image, fullObj HasHitbox, thisSeg image.Rectangle) (*ebiten.Image, *ebiten.DrawImageOptions) {
	lineartW := 10
	lineartC := color.RGBA{0, 0, 0, 255}
	lineartOptions := ebiten.DrawImageOptions{}

	//image setup
	imgOffset := ebiten.DrawImageOptions{}
	imgOffset.GeoM.Translate(-float64(lineartW)/2, -float64(lineartW)/2)
	img := ebiten.NewImage(oldImgSeg.Bounds().Dx()+lineartW, oldImgSeg.Bounds().Dy()+lineartW)

	//original image
	oldImgOffset := ebiten.DrawImageOptions{}
	oldImgOffset.GeoM.Translate(float64(lineartW)/2, float64(lineartW)/2)
	img.DrawImage(oldImgSeg, &oldImgOffset)

	//line segments
	fullHorizontal := ebiten.NewImage(img.Bounds().Dx(), lineartW)
	fullVertical := ebiten.NewImage(lineartW, img.Bounds().Dy())
	fullHorizontal.Fill(lineartC)
	fullVertical.Fill(lineartC)

	//line art
	if !fullObj.Overlaps(Render, []image.Rectangle{image.Rect(
		thisSeg.Min.X, thisSeg.Min.Y-(lineartW/2), thisSeg.Max.X, thisSeg.Min.Y)}) {

		img.DrawImage(fullHorizontal, nil) //top
	}

	if !fullObj.Overlaps(Render, []image.Rectangle{image.Rect(
		thisSeg.Min.X, thisSeg.Max.Y+(lineartW/2), thisSeg.Max.X, thisSeg.Max.Y)}) {

		lineartOptions.GeoM.Reset()
		lineartOptions.GeoM.Translate(0, float64(img.Bounds().Dy()-lineartW))
		img.DrawImage(fullHorizontal, &lineartOptions) //bottom
	}

	if !fullObj.Overlaps(Render, []image.Rectangle{image.Rect(
		thisSeg.Min.X-(lineartW/2), thisSeg.Min.Y, thisSeg.Min.X, thisSeg.Max.Y)}) {

		img.DrawImage(fullVertical, nil) //left
	}

	if !fullObj.Overlaps(Render, []image.Rectangle{image.Rect(
		thisSeg.Max.X+(lineartW/2), thisSeg.Min.Y, thisSeg.Max.X, thisSeg.Max.Y)}) {

		lineartOptions.GeoM.Reset()
		lineartOptions.GeoM.Translate(float64(img.Bounds().Dx()-lineartW), 0)
		img.DrawImage(fullVertical, &lineartOptions) //right
	}

	return img, &imgOffset
}
