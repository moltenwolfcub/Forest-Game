package game

import (
	"image/color"
	"math"
)

// the number of hours sunrise/sunset takes
const transitionTime = 1.0

func GetAmbientLight(time Time) color.Color {
	currentDay := int(math.Floor(float64(time) / TPGM / MinsPerHour / HoursPerDay))
	moddedDay := currentDay % (DaysPerMonth * MonthsPerYear)

	sunriseStart, sunriseEnd := getSunriseTime(moddedDay)
	sunsetStart, sunsetEnd := getSunsetTime(moddedDay)

	currentHour := math.Mod(float64(time)/TPGM/MinsPerHour, HoursPerDay)

	if currentHour > sunriseEnd && currentHour < sunsetStart {
		return DayAmbientLightColor
	} else if currentHour < sunriseStart || currentHour > sunsetEnd {
		return NightAmbientLightColor
	} else {
		return getPartialLighting(currentHour, sunriseStart, sunsetStart)
	}
}

// Gets the ambient lighting for when it's not day or night time
//
// It finds the lighting between day and night light for sunrise
// and sunset and returns the correct one based upon which one is
// currently happening. If both are happening due to it being close
// to a solstice it returns the darker of the 2 colors.
func getPartialLighting(hour float64, riseTime float64, setTime float64) color.Color {
	risePercent := (hour - riseTime) / transitionTime
	riseColor := getColorBetween(NightAmbientLightColor, DayAmbientLightColor, risePercent)

	setPercent := (hour - setTime) / transitionTime
	setColor := getColorBetween(DayAmbientLightColor, NightAmbientLightColor, setPercent)

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

// Returns the darkest of 2 colors
func getDarkest(a color.RGBA, b color.RGBA) color.RGBA {
	aSum := int(a.R) + int(a.G) + int(a.B) + int(a.A)
	bSum := int(b.R) + int(b.G) + int(b.B) + int(b.A)

	if aSum <= bSum {
		return a
	} else {
		return b
	}
}

// Returns a color between the start and end color based on the percentage.
// Percentage represents how far between the 2 colors is the desired output.
//
// This currently uses linear interpolation between the 2 points but could be
// changed to a different kind of transition possibly parsed as a parameter.
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

// Gets the time that the sun should set on a given day of the year
//
// This is based upon a parabolic curve betweeen midday at winter solstice
// and very late evening at summer solstice.
//
// It describes the function for the parabola and parses the day
// of the year into the function.
//
// First returned value is the starting time of the sunset an second the ending time.
// Currently their difference (length of sunset) is `transitionTime` but will
// eventually become more dynamic based upon time of year.
func getSunsetTime(dayOfYear int) (start float64, end float64) {
	t := float64(DaysPerMonth * MonthsPerYear)
	d := float64(HoursPerDay)

	offsetDay := math.Mod(float64(dayOfYear)+(float64(SolsticeMonthsOffset)/MonthsPerYear)*t, t)

	a := -d / (0.5 * t * t)
	b := d / (0.5 * t)
	c := 0.5 * d

	time := a*float64(offsetDay*offsetDay) + b*float64(offsetDay) + c
	return time - transitionTime/2, time + transitionTime/2
}

// Gets the time that the sun should rise on a given day of the year
//
// This is based upon a parabolic curve betweeen midday at winter solstice
// and very early morning at summer solstice.
//
// It describes the function for the parabola and parses the day
// of the year into the function.
//
// First returned value is the starting time of the sunrise an second the ending
// time. Currently their difference (length of sunrise) is `transitionTime` but
// will eventually become more dynamic based upon time of year.
func getSunriseTime(dayOfYear int) (start float64, end float64) {
	t := float64(DaysPerMonth * MonthsPerYear)
	d := float64(HoursPerDay)

	offsetDay := math.Mod(float64(dayOfYear)+(float64(SolsticeMonthsOffset)/MonthsPerYear)*t, t)

	a := d / (0.5 * t * t)
	b := -d / (0.5 * t)
	c := 0.5 * d

	time := a*float64(offsetDay*offsetDay) + b*float64(offsetDay) + c
	return time - transitionTime/2, time + transitionTime/2
}
