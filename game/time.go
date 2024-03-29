package game

import (
	"fmt"
	"log/slog"

	"github.com/moltenwolfcub/Forest-Game/args"
	"github.com/moltenwolfcub/Forest-Game/errors"
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

	MinsPerHour   = 60
	HoursPerDay   = 20
	DaysPerMonth  = 10
	MonthsPerYear = 8

	hoursPerMonth = HoursPerDay * DaysPerMonth
	ticksPerDay   = TPGM * MinsPerHour * HoursPerDay

	SolsticeMonthsOffset = 1
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
// Any value provided outside of the range 1-8 will return an error
func GetSeason(monthInYear int) (Season, error) {
	switch monthInYear {
	case 1, 2:
		return Spring, nil
	case 3, 4:
		return Summer, nil
	case 5, 6:
		return Autumn, nil
	case 7, 8:
		return Winter, nil
	default:
		return 0, errors.NewSeasonOutOfBoundsError(monthInYear)
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

	rawSeason, err := GetSeason(int(months))
	var seasonStr string
	if err != nil {
		slog.Warn(err.Error())
		seasonStr = "SeasonError"
	} else {
		seasonStr = rawSeason.String()
	}

	season := cases.Title(language.AmericanEnglish).String(seasonStr)

	return fmt.Sprintf("%s %d/%d/%d %02d:%02d", season, days, months, years, hours, mins)
}

func (t Time) ThroughMonth() float64 {
	return float64(t.HoursThroughMonth()) / float64(hoursPerMonth)
}
func (t Time) ThroughDay() float64 {
	return float64(t.TicksThroughDay()) / float64(ticksPerDay)
}

func (t Time) Hours() int {
	return int(float64(t) / TPGM / MinsPerHour)
}
func (t Time) Days() int {
	return int(float64(t.Hours()) / HoursPerDay)
}
func (t Time) Months() int {
	return int(float64(t.Days()) / DaysPerMonth)
}

func (t Time) TicksThroughDay() int {
	return int(t) % ticksPerDay
}
func (t Time) HoursThroughMonth() int {
	return t.Hours() % (hoursPerMonth)
}
func (t Time) DaysThroughYear() int {
	return t.Days() % (DaysPerMonth * MonthsPerYear)
}
func (t Time) MonthsThroughYear() int {
	return t.Months()%MonthsPerYear + 1
}
