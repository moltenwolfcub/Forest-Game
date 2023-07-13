package game

import "image/color"

func GetAmbientLight(min color.RGBA, max color.RGBA, time Time) color.Color {

	colorPerTick := float64(max.R-min.R) / float64(DAYLEN/2)
	mappedLight := float64(min.R) + colorPerTick*float64(time.GetTimeInDay()*TPGM)
	if mappedLight > float64(max.R) {
		diff := mappedLight - float64(max.R)
		mappedLight = float64(max.R) - diff
	}
	redLight := uint8(mappedLight)

	colorPerTick = float64(max.G-min.G) / float64(DAYLEN/2)
	mappedLight = float64(min.G) + colorPerTick*float64(time.GetTimeInDay()*TPGM)
	if mappedLight > float64(max.G) {
		diff := mappedLight - float64(max.G)
		mappedLight = float64(max.G) - diff
	}
	greenLight := uint8(mappedLight)

	colorPerTick = float64(max.B-min.B) / float64(DAYLEN/2)
	mappedLight = float64(min.B) + colorPerTick*float64(time.GetTimeInDay()*TPGM)
	if mappedLight > float64(max.B) {
		diff := mappedLight - float64(max.B)
		mappedLight = float64(max.B) - diff
	}
	blueLight := uint8(mappedLight)

	return color.RGBA{redLight, greenLight, blueLight, 255}
}
