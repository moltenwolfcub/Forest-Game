package game

import (
	"image"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Lamp struct {
	Rect image.Rectangle
}

func NewLamp() Lamp {
	radius := 100
	lamp := Lamp{
		Rect: image.Rectangle{
			Max: image.Point{radius * 2, radius * 2},
		},
	}
	lamp.Rect = lamp.Rect.Add(image.Point{50, 40})
	return lamp
}

func (l Lamp) Hitbox(layer GameContext) image.Rectangle {
	return l.Rect
}

func (l Lamp) DrawLighting(lightingLayer *ebiten.Image, pos image.Point) {
	radius := 100
	diameter := radius * 2

	img := ebiten.NewImage(diameter, diameter)

	// not very efficient (prolly gonna pre-generate it eventually so it doesn't cause lag with bigger light sources)
	for y := 0; y < diameter; y++ {
		for x := 0; x < diameter; x++ {
			strength := lightingStrength(diameter, x+1, y+1)

			// current way of simulating a 'screen' blend operation with alpha
			img.Set(x, y, color.RGBA{strength, strength, strength, strength})
		}
	}

	// prolly need to find a way to do it in blend options to remove dark rim
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(pos.X), float64(pos.Y))

	lightingLayer.DrawImage(img, &options)
	img.Dispose()
}

func lightingStrength(size int, x int, y int) uint8 {
	centre := size / 2
	maxStrength := 255

	// find distance from centre
	dx := float64(x - centre)
	dy := float64(y - centre)
	dist := math.Hypot(dx, dy)

	// map it to value between 0 and 1
	mappedDist := dist / float64(size/2)
	if mappedDist >= 1 {
		return 0
	}
	if mappedDist == 0 {
		return uint8(maxStrength)
	}
	// invert it to make light in the middle rather than at the edge
	invDist := 1 - mappedDist

	// multiply it back up from 0-1 up to the desired strength
	strength := invDist * float64(maxStrength)

	return uint8(strength)
}
