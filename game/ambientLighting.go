package game

import (
	"image/color"
	"math"
)

// the number of hours sunrise/sunset takes
const transitionTime = 1.0

var (
	nightLight = color.RGBA{115, 100, 135, 0}
	dayLight   = color.RGBA{255, 255, 255, 0}
)

func GetAmbientLight(time Time) color.Color {
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
		return getPartialLighting(currentHour, sunriseStart, sunsetStart)
	}
}

func getPartialLighting(hour float64, riseTime float64, setTime float64) color.Color {
	risePercent := (hour - riseTime) / transitionTime
	riseColor := getColorBetween(nightLight, dayLight, risePercent)

	setPercent := (hour - setTime) / transitionTime
	setColor := getColorBetween(dayLight, nightLight, setPercent)

	rising := hour >= riseTime && hour <= riseTime+transitionTime
	setting := hour >= setTime && hour <= setTime+transitionTime

	if rising && !setting {
		return riseColor
	} else if setting && !rising {
		return setColor
	} else {
		return getDarkest(riseColor, setColor)
	}
}

func getDarkest(a color.RGBA, b color.RGBA) color.RGBA {
	aSum := a.R + a.G + a.B + a.A
	bSum := b.R + b.G + b.B + b.A

	if aSum <= bSum {
		return a
	} else {
		return b
	}
}

func getColorBetween(start color.RGBA, end color.RGBA, percent float64) color.RGBA {
	sr, sg, sb, sa := float64(start.R), float64(start.G), float64(start.B), float64(start.A)
	er, eg, eb, ea := float64(end.R), float64(end.G), float64(end.B), float64(end.A)

	dr := er - sr
	dg := eg - sg
	db := eb - sb
	da := ea - sa

	r := uint8(percent*dr + sr)
	g := uint8(percent*dg + sg)
	b := uint8(percent*db + sb)
	a := uint8(percent*da + sa)

	return color.RGBA{r, g, b, a}
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
