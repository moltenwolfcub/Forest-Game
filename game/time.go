package game

//  Minute
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

type Time struct {
	ticksElapsed uint
}

func (t *Time) Tick() {
	t.ticksElapsed++
}

func (t Time) getTicks() uint {
	return t.ticksElapsed
}
