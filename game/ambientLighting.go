package game

import (
	"image/color"
	"math"
)

// the number of hours sunrise/sunset takes
const transitionTime = 1.0

func GetAmbientLight(nightLight color.RGBA, dayLight color.RGBA, time Time) color.Color {
	currentDay := int(math.Floor(float64(time) / TPGM / MinsPerHour / HoursPerDay))
	moddedDay := currentDay % (DaysPerMonth * MonthsPerYear)

	sunriseStart, sunriseEnd := getSunriseTime(moddedDay)
	sunsetStart, sunsetEnd := getSunsetTime(moddedDay)

	currentHour := math.Mod(float64(time)/TPGM/MinsPerHour, HoursPerDay)

	if currentHour > sunriseEnd && currentHour < sunsetStart {
		return dayLight
	} else if currentHour < sunriseStart || currentHour > sunsetEnd {
		return nightLight
	} else {
		return color.RGBA{255, 0, 0, 0}
	}
}

func getSunsetTime(dayOfYear int) (start float64, end float64) {
	t := float64(DaysPerMonth * MonthsPerYear)
	d := float64(HoursPerDay)

	a := -d / (0.5 * t * t)
	b := d / (0.5 * t)
	c := 0.5 * d

	time := a*float64(dayOfYear*dayOfYear) + b*float64(dayOfYear) + c
	return time - transitionTime/2, time + transitionTime/2
}
func getSunriseTime(dayOfYear int) (start float64, end float64) {
	t := float64(DaysPerMonth * MonthsPerYear)
	d := float64(HoursPerDay)

	a := d / (0.5 * t * t)
	b := -d / (0.5 * t)
	c := 0.5 * d

	time := a*float64(dayOfYear*dayOfYear) + b*float64(dayOfYear) + c
	return time - transitionTime/2, time + transitionTime/2
}
