package game

import "image"

type Lamp struct {
	Rect image.Rectangle
}

func NewLamp() Lamp {
	radius := 70
	lamp := Lamp{
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

func (l Lamp) Radius() int {
	return 70
}
