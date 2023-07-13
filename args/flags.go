package args

import "flag"

var (
	exampleFlag int
)

func ParseFlags() {
	flag.IntVar(&exampleFlag, "example", 5, "example flag")

	flag.Parse()
}
