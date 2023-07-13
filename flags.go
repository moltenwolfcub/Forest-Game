package main

import "flag"

var (
	exampleFlag int
)

func parseFlags() {
	flag.IntVar(&exampleFlag, "example", 5, "example flag")

	flag.Parse()
}
