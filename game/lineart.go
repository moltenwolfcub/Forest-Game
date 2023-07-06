package game

import (
	"image"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	lineartW = 10
	lineartC = color.RGBA{0, 0, 0, 255}
)

// Takes an image and adds lineart to all sides of it before returning the new image.
// The image given should be a single rect out of the full object so that it can be
// computed seperately. The new offset from the original image is also returned so the
// image can be seemlessly inserted.(This is useful for keeping it in line with hitboxes)
func ApplyLineart(oldImgSeg *ebiten.Image, fullObj HasHitbox, thisSeg image.Rectangle) (*ebiten.Image, *ebiten.DrawImageOptions) {
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
	fullHorizontal := image.Rect(0, 0, img.Bounds().Dx(), lineartW)
	fullVertical := image.Rect(0, 0, lineartW, img.Bounds().Dy())

	//line art
	start := getPointOnBounds(top, startSide, thisSeg, fullObj.GetHitbox(Render), fullHorizontal)
	end := getPointOnBounds(top, endSide, thisSeg, fullObj.GetHitbox(Render), fullHorizontal)
	diff := int(end - start)
	if diff >= lineartW {
		partialHorizontal := ebiten.NewImage(diff+lineartW, lineartW)
		partialHorizontal.Fill(lineartC)

		lineartOptions.GeoM.Reset()
		lineartOptions.GeoM.Translate(start-float64(thisSeg.Min.X), 0)

		img.DrawImage(partialHorizontal, &lineartOptions)

		partialHorizontal.Dispose()
	}

	start = getPointOnBounds(bottom, startSide, thisSeg, fullObj.GetHitbox(Render), fullHorizontal)
	end = getPointOnBounds(bottom, endSide, thisSeg, fullObj.GetHitbox(Render), fullHorizontal)
	diff = int(end - start)
	if diff >= lineartW {
		partialHorizontal := ebiten.NewImage(diff+lineartW, lineartW)
		partialHorizontal.Fill(lineartC)

		lineartOptions.GeoM.Reset()
		lineartOptions.GeoM.Translate(0, float64(img.Bounds().Dy()-lineartW))
		lineartOptions.GeoM.Translate(start-float64(thisSeg.Min.X), 0)

		img.DrawImage(partialHorizontal, &lineartOptions)

		partialHorizontal.Dispose()
	}

	start = getPointOnBounds(left, startSide, thisSeg, fullObj.GetHitbox(Render), fullVertical)
	end = getPointOnBounds(left, endSide, thisSeg, fullObj.GetHitbox(Render), fullVertical)
	diff = int(end - start)
	if diff >= lineartW {
		partialVertical := ebiten.NewImage(lineartW, diff+lineartW)
		partialVertical.Fill(lineartC)

		lineartOptions.GeoM.Reset()
		lineartOptions.GeoM.Translate(0, start-float64(thisSeg.Min.Y))

		img.DrawImage(partialVertical, &lineartOptions)

		partialVertical.Dispose()
	}

	start = getPointOnBounds(right, startSide, thisSeg, fullObj.GetHitbox(Render), fullVertical)
	end = getPointOnBounds(right, endSide, thisSeg, fullObj.GetHitbox(Render), fullVertical)
	diff = int(end - start)
	if diff >= lineartW {
		partialVertical := ebiten.NewImage(lineartW, diff+lineartW)
		partialVertical.Fill(lineartC)

		lineartOptions.GeoM.Reset()
		lineartOptions.GeoM.Translate(0, start-float64(thisSeg.Min.Y))
		lineartOptions.GeoM.Translate(float64(img.Bounds().Dx()-lineartW), 0)

		img.DrawImage(partialVertical, &lineartOptions)

		partialVertical.Dispose()
	}

	return img, &imgOffset
}

func getPointOnBounds(side lineartSide, end lineartEnd, thisSeg image.Rectangle, occluders []image.Rectangle, edge image.Rectangle) (point float64) {
	point = side.getAxis(end.getCheckStart(thisSeg))

	offset := thisSeg.Min
	switch side {
	case bottom:
		offset = offset.Add(image.Pt(0, thisSeg.Dy()-lineartW))
	case right:
		offset = offset.Add(image.Pt(thisSeg.Dx()-lineartW, 0))
	}
	offsetEdge := edge.Add(offset)

	for _, seg := range occluders {
		if seg == thisSeg || !seg.Overlaps(offsetEdge) {
			continue
		}
		if end.checkOnRightSide(side.getAxis(end.getCheckEnd(thisSeg)), side.getAxis(end.getCheckEnd(seg))) {
			point = end.minMax(point, side.getAxis(end.getCheckEnd(seg)))
		}
	}
	return
}

type lineartSide int

const (
	left lineartSide = iota
	right
	top
	bottom
)

func (l lineartSide) getAxis(point image.Point) float64 {
	if l == left || l == right {
		return float64(point.Y)
	} else if l == top || l == bottom {
		return float64(point.X)
	} else {
		return 0
	}
}

type lineartEnd int

const (
	startSide lineartEnd = iota
	endSide
)

func (l lineartEnd) getCheckStart(rect image.Rectangle) (point image.Point) {
	switch l {
	case startSide:
		point = rect.Min
	case endSide:
		point = rect.Max
	}
	return
}
func (l lineartEnd) getCheckEnd(rect image.Rectangle) (point image.Point) {
	switch l {
	case startSide:
		point = rect.Max
	case endSide:
		point = rect.Min
	}
	return
}
func (l lineartEnd) minMax(pos float64, potentialNew float64) (newPos float64) {
	switch l {
	case startSide:
		newPos = math.Max(pos, potentialNew)
	case endSide:
		newPos = math.Min(pos, potentialNew)
	}
	return
}
func (l lineartEnd) checkOnRightSide(this float64, other float64) bool {
	switch l {
	case startSide:
		return other < this
	case endSide:
		return other > this
	default:
		return false
	}
}
