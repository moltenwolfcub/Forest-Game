package game

import (
	"fmt"
	"image/color"
	"math"
)

// const transitionTime = TPGM * MinsPerHour

func GetAmbientLight(nightLight color.RGBA, dayLight color.RGBA, time Time) color.Color {
	currentDay := int(time) / TPGM / MinsPerHour / HoursPerDay
	moddedDay := currentDay % (DaysPerMonth * MonthsPerYear)

	sunrise := getSunriseTime(moddedDay)
	sunset := getSunsetTime(moddedDay)

	currentHour := math.Mod(float64(time)/TPGM/MinsPerHour, HoursPerDay)

	fmt.Println(currentHour)

	if currentHour > sunrise && currentHour < sunset {
		return dayLight
	} else {
		return nightLight
	}
}

func getSunsetTime(dayOfYear int) float64 {
	t := float64(DaysPerMonth * MonthsPerYear)
	d := float64(HoursPerDay)

	a := -d / (0.5 * t * t)
	b := d / (0.5 * t)
	c := 0.5 * d

	time := a*float64(dayOfYear*dayOfYear) + b*float64(dayOfYear) + c
	return time
}
func getSunriseTime(dayOfYear int) float64 {
	t := float64(DaysPerMonth * MonthsPerYear)
	d := float64(HoursPerDay)

	a := d / (0.5 * t * t)
	b := -d / (0.5 * t)
	c := 0.5 * d

	time := a*float64(dayOfYear*dayOfYear) + b*float64(dayOfYear) + c
	return time
}
