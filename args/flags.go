package args

import "flag"

var (
	exampleFlag  int
	TimeRateFlag int
)

func ParseFlags() {
	flag.IntVar(&exampleFlag, "example", 5, "example flag")
	flag.IntVar(&TimeRateFlag, "time-rate", 1, "A multiplier for the rate that time should progress. 1 is normal time rate")

	flag.Parse()
}
