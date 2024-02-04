package args

import (
	"flag"
)

var (
	TimeRateFlag int
	ProfilerFlag bool
)

func ParseFlags() {
	flag.IntVar(&TimeRateFlag, "time-rate", 1, "A multiplier for the rate that time should progress. 1 is normal time rate")
	flag.BoolVar(&ProfilerFlag, "profile", false, "Runs a profiler on the cpu performance")

	flag.Parse()
}
