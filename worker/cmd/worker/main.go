package main

import (
	"flag"
)

func main() {
	var (
		avgLoad  = flag.Uint("avg-load", 10, "Targeted average load in %")
		memUsage = flag.Uint("mem-usage", 512, "Memory allocation target in MB")
	)
	flag.Parse()

	for {
		// TODO
		_ = avgLoad
		_ = memUsage
	}
}
