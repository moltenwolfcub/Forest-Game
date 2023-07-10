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

	DAYLEN = TPGM * 60 * 20

	MinsPerHour   = 60
	HoursPerDay   = 20
	DaysPerMonth  = 10
	MonthsPerYear = 8
)

type Time int

func (t *Time) Tick() {
	*t++
}

func (t Time) String() string {
	totalMins := t / TPGM
	totalHours := totalMins / MinsPerHour
	totalDays := totalHours / HoursPerDay
	totalMonths := totalDays / DaysPerMonth
	totalYears := totalMonths / MonthsPerYear

	//+1 cause humans start from 1 with these things
	mins := totalMins % MinsPerHour
	hours := totalHours % HoursPerDay
	days := totalDays%DaysPerMonth + 1
	months := totalMonths%MonthsPerYear + 1
	years := totalYears + 1

	return fmt.Sprintf("Season[unimplemented] %d/%d/%d %02d:%02d", days, months, years, hours, mins)
}

// Returns the number of minutes through the day it currently is
func (t Time) GetTimeInDay() int {
	ticks := math.Mod(float64(t), DAYLEN)
	minutes := ticks / TPGM

	return int(minutes)
}
