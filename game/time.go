package game

import (
	"fmt"
	"math"

	"github.com/moltenwolfcub/Forest-Game/args"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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

type Season int

const (
	Spring Season = iota
	Summer
	Autumn
	Winter
)

// Returns a Season from a month in the year.
// The parsed month should be 1-indexed so 0 in invalid.
//
// Any value provided outside of the range 1-8 will panic
func GetSeason(monthInYear int) Season {
	switch monthInYear {
	case 1, 2:
		return Spring
	case 3, 4:
		return Summer
	case 5, 6:
		return Autumn
	case 7, 8:
		return Winter
	default:
		panic(fmt.Sprintf("Can't figure out season from month: %d. Season.GetSeason only accepts values in the range 1-8 inclusive.", monthInYear))
	}
}

func (s Season) String() string {
	switch s {
	case Spring:
		return "spring"
	case Summer:
		return "summer"
	case Autumn:
		return "autumn"
	case Winter:
		return "winter"
	default:
		return "error"
	}
}

type Time int

func (t *Time) Tick() {
	*t += Time(args.TimeRateFlag)
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
	years := totalYears

	season := cases.Title(language.AmericanEnglish).String(GetSeason(int(months)).String())

	return fmt.Sprintf("%s %d/%d/%d %02d:%02d", season, days, months, years, hours, mins)
}

// Returns the number of minutes through the day it currently is
func (t Time) GetTimeInDay() int {
	ticks := math.Mod(float64(t), DAYLEN)
	minutes := ticks / TPGM

	return int(minutes)
}
