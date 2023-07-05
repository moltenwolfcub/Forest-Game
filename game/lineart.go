package game

import (
	"image"
	"image/color"
	"math"

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
	} else {
		var start float64 = float64(thisSeg.Min.X)
		for _, seg := range fullObj.GetHitbox(Render) {
			if seg == thisSeg || !seg.Overlaps(fullHorizontal.Bounds().Add(thisSeg.Min)) {
				continue
			}
			if seg.Max.X < thisSeg.Max.X {
				start = math.Max(start, float64(seg.Max.X))
			}
		}
		var end float64 = float64(thisSeg.Max.X)
		for _, seg := range fullObj.GetHitbox(Render) {
			if seg == thisSeg || !seg.Overlaps(fullHorizontal.Bounds().Add(thisSeg.Min)) {
				continue
			}
			if seg.Min.X > thisSeg.Min.X {
				end = math.Min(end, float64(seg.Min.X))
			}
		}
		diff := int(end - start)
		if diff >= lineartW {
			partialHorizontal := ebiten.NewImage(diff+lineartW, lineartW)
			partialHorizontal.Fill(lineartC)

			lineartOptions.GeoM.Reset()
			lineartOptions.GeoM.Translate(start-float64(thisSeg.Min.X), 0)

			img.DrawImage(partialHorizontal, &lineartOptions)

			partialHorizontal.Dispose()
		}
	}

	if !fullObj.Overlaps(Render, []image.Rectangle{image.Rect(
		thisSeg.Min.X, thisSeg.Max.Y+(lineartW/2), thisSeg.Max.X, thisSeg.Max.Y)}) {

		lineartOptions.GeoM.Reset()
		lineartOptions.GeoM.Translate(0, float64(img.Bounds().Dy()-lineartW))
		img.DrawImage(fullHorizontal, &lineartOptions) //bottom
	} else {
		var start float64 = float64(thisSeg.Min.X)
		for _, seg := range fullObj.GetHitbox(Render) {
			if seg == thisSeg || !seg.Overlaps(fullHorizontal.Bounds().Add(thisSeg.Min.Add(image.Pt(0, img.Bounds().Dy()-lineartW)))) {
				continue
			}
			if seg.Max.X < thisSeg.Max.X {
				start = math.Max(start, float64(seg.Max.X))
			}
		}
		var end float64 = float64(thisSeg.Max.X)
		for _, seg := range fullObj.GetHitbox(Render) {
			if seg == thisSeg || !seg.Overlaps(fullHorizontal.Bounds().Add(thisSeg.Min.Add(image.Pt(0, img.Bounds().Dy()-lineartW)))) {
				continue
			}
			if seg.Min.X > thisSeg.Min.X {
				end = math.Min(end, float64(seg.Min.X))
			}
		}
		diff := int(end - start)
		if diff >= lineartW {
			partialHorizontal := ebiten.NewImage(diff+lineartW, lineartW)
			partialHorizontal.Fill(lineartC)

			lineartOptions.GeoM.Reset()
			lineartOptions.GeoM.Translate(0, float64(img.Bounds().Dy()-lineartW))
			lineartOptions.GeoM.Translate(start-float64(thisSeg.Min.X), 0)

			img.DrawImage(partialHorizontal, &lineartOptions)

			partialHorizontal.Dispose()
		}
	}

	if !fullObj.Overlaps(Render, []image.Rectangle{image.Rect(
		thisSeg.Min.X-(lineartW/2), thisSeg.Min.Y, thisSeg.Min.X, thisSeg.Max.Y)}) {

		img.DrawImage(fullVertical, nil) //left
	} else {
		var start float64 = float64(thisSeg.Min.Y)
		for _, seg := range fullObj.GetHitbox(Render) {
			if seg == thisSeg || !seg.Overlaps(fullVertical.Bounds().Add(thisSeg.Min)) {
				continue
			}
			if seg.Max.Y < thisSeg.Max.Y {
				start = math.Max(start, float64(seg.Max.Y))
			}
		}
		var end float64 = float64(thisSeg.Max.Y)
		for _, seg := range fullObj.GetHitbox(Render) {
			if seg == thisSeg || !seg.Overlaps(fullVertical.Bounds().Add(thisSeg.Min)) {
				continue
			}
			if seg.Min.Y > thisSeg.Min.Y {
				end = math.Min(end, float64(seg.Min.Y))
			}
		}
		diff := int(end - start)
		if diff >= lineartW {
			partialVertical := ebiten.NewImage(lineartW, diff+lineartW)
			partialVertical.Fill(lineartC)

			lineartOptions.GeoM.Reset()
			lineartOptions.GeoM.Translate(0, start-float64(thisSeg.Min.Y))

			img.DrawImage(partialVertical, &lineartOptions)

			partialVertical.Dispose()
		}
	}

	if !fullObj.Overlaps(Render, []image.Rectangle{image.Rect(
		thisSeg.Max.X+(lineartW/2), thisSeg.Min.Y, thisSeg.Max.X, thisSeg.Max.Y)}) {

		lineartOptions.GeoM.Reset()
		lineartOptions.GeoM.Translate(float64(img.Bounds().Dx()-lineartW), 0)
		img.DrawImage(fullVertical, &lineartOptions) //right
	} else {
		var start float64 = float64(thisSeg.Min.Y)
		for _, seg := range fullObj.GetHitbox(Render) {
			if seg == thisSeg || !seg.Overlaps(fullVertical.Bounds().Add(thisSeg.Min.Add(image.Pt(img.Bounds().Dx()-lineartW, 0)))) {
				continue
			}
			if seg.Max.Y < thisSeg.Max.Y {
				start = math.Max(start, float64(seg.Max.Y))
			}
		}
		var end float64 = float64(thisSeg.Max.Y)
		for _, seg := range fullObj.GetHitbox(Render) {
			if seg == thisSeg || !seg.Overlaps(fullVertical.Bounds().Add(thisSeg.Min.Add(image.Pt(img.Bounds().Dx()-lineartW, 0)))) {
				continue
			}
			if seg.Min.Y > thisSeg.Min.Y {
				end = math.Min(end, float64(seg.Min.Y))
			}
		}
		diff := int(end - start)
		if diff >= lineartW {
			partialVertical := ebiten.NewImage(lineartW, diff+lineartW)
			partialVertical.Fill(lineartC)

			lineartOptions.GeoM.Reset()
			lineartOptions.GeoM.Translate(0, start-float64(thisSeg.Min.Y))
			lineartOptions.GeoM.Translate(float64(img.Bounds().Dx()-lineartW), 0)

			img.DrawImage(partialVertical, &lineartOptions)

			partialVertical.Dispose()
		}
	}

	return img, &imgOffset
}

//this function needs a serious refactor it's long and repetitive
