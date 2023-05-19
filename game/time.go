package game

import (
	"fmt"
	"math"
)

//  Minute
// TPGM ticks
//irl - second

//  Hour
//60 minutes
//irl - minute

//  Day
//20 hours
//irl - 20 minutes

//  Month
//10 days
//irl - 3hr 20mins

//  Season
//2 months
//irl - 6hr 40mins

//  Year
//4 Seasons
//irl - 26hr 40mins

const (
	TPGM = TPS
)

type Time int

func (t *Time) Tick() {
	*t++
}

// func (t Time) Minutes() int {
// 	return int((t / TPGM) % 60)
// }
// func (t Time) Hours() int {
// 	return (t.Minutes() / 60) % 20
// }
// func (t Time) Days() int {
// 	return (t.Hours() / 20) % 10
// }
// func (t Time) Months() int {
// 	return (t.Days() / 10) % 2
// }
// func (t Time) Seasons() int {
// 	return (t.Months() / 2) % 4
// }
// func (t Time) Years() int {
// 	return t.Seasons() / 4
// }

func (t Time) String() string {
	minutes := (t / TPGM) % 60
	hours := (t / TPGM / 60) % 20

	return fmt.Sprintf("%02d:%02d", hours, minutes)
}

func (t Time) GetLighting() uint8 {
	// at 00:00 minimum light
	// at 10:00 maximum light
	var dayLen float64 = TPGM * 60 * 20
	var halfDayLen float64 = dayLen / 2
	var minLight float64 = 48
	var maxLight float64 = 255

	var current float64 = math.Mod(float64(t), dayLen)

	colorPerTick := (maxLight - minLight) / (halfDayLen)
	mappedLight := minLight + colorPerTick*float64(current)
	if mappedLight > maxLight {
		diff := mappedLight - maxLight
		mappedLight = maxLight - diff
	}

	return uint8(mappedLight)
}
