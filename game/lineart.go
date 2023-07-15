package game

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	lineartW = 10
)

// Takes an image and adds lineart to all sides of it before returning the new image.
// The image given should be a single rect out of the full object so that it can be
// computed seperately. The new offset from the original image is also returned so the
// image can be seemlessly inserted.(This is useful for keeping it in line with hitboxes)
func ApplyLineart(oldImgSeg *ebiten.Image, fullObj HasHitbox, thisSeg image.Rectangle) (*ebiten.Image, *ebiten.DrawImageOptions) {

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
	drawSide(img, thisSeg, fullObj, fullHorizontal, top)
	drawSide(img, thisSeg, fullObj, fullHorizontal, bottom)
	drawSide(img, thisSeg, fullObj, fullVertical, left)
	drawSide(img, thisSeg, fullObj, fullVertical, right)

	return img, &imgOffset
}

func drawSide(toDrawTo *ebiten.Image, thisSeg image.Rectangle, fullObj HasHitbox, edge image.Rectangle, side lineartSide) {
	start := getPointOnBounds(side, startSide, thisSeg, fullObj.GetHitbox(Render), edge)
	end := getPointOnBounds(side, endSide, thisSeg, fullObj.GetHitbox(Render), edge)
	diff := int(end - start)
	if diff < lineartW {
		return
	}
	imgSize := side.swapAxis(image.Pt(lineartW, diff+lineartW))
	partialSide := ebiten.NewImage(imgSize.X, imgSize.Y)
	partialSide.Fill(LineartColor)

	lineartOptions := ebiten.DrawImageOptions{}
	offset := side.getAxisPoint(image.Pt(int(start)-thisSeg.Min.X, int(start)-thisSeg.Min.Y))
	lineartOptions.GeoM.Translate(float64(offset.X), float64(offset.Y))
	offset = side.getOffset(toDrawTo.Bounds())
	lineartOptions.GeoM.Translate(float64(offset.X), float64(offset.Y))

	toDrawTo.DrawImage(partialSide, &lineartOptions)
	partialSide.Dispose()
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

func (l lineartSide) getOffset(rect image.Rectangle) image.Point {
	switch l {
	case bottom:
		return image.Pt(0, rect.Dy()-lineartW)
	case right:
		return image.Pt(rect.Dx()-lineartW, 0)
	default:
		return image.Point{}
	}
}

func (l lineartSide) getAxis(point image.Point) float64 {
	if l == left || l == right {
		return float64(point.Y)
	} else if l == top || l == bottom {
		return float64(point.X)
	} else {
		return 0
	}
}
func (l lineartSide) getAxisPoint(point image.Point) image.Point {
	if l == left || l == right {
		return image.Pt(0, point.Y)
	} else if l == top || l == bottom {
		return image.Pt(point.X, 0)
	} else {
		return image.Point{}
	}
}

// returns the point if `l` is horizontal if it's vertical then
// x and y get flipped in the point. Returns point of 0, 0 if
// there is a problem
func (l lineartSide) swapAxis(point image.Point) image.Point {
	if l == left || l == right {
		return point
	} else if l == top || l == bottom {
		return image.Pt(point.Y, point.X)
	} else {
		return image.Point{}
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
