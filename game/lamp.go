package game

import "image"

type Lamp struct {
	Rect  image.Rectangle
	light Light
}

func NewLamp() Lamp {
	radius := 70
	lamp := Lamp{
		light: NewLight(radius),
		Rect: image.Rectangle{
			Max: image.Point{radius * 2, radius * 2},
		},
	}
	lamp.Rect = lamp.Rect.Add(image.Point{50, 40})
	return lamp
}

func (l Lamp) GetPos() image.Rectangle {
	return l.Rect
}

func (l Lamp) GetLight() Light {
	return l.light
}
