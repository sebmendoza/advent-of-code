package main

import (
	"2024/day1"
	"2024/utils"
	"flag"
)

func main() {

	day := flag.Int("day", 1, "day to run")
	flag.Parse()

	runner := utils.NewRunner()

	// Register solutions
	runner.Add(1, day1.New())
	// Add more days as you implement them

	if *day == 0 {
		// Run all implemented days
		for d := 1; d <= 25; d++ {
			runner.Run(d)
		}
	} else {
		runner.Run(*day)
	}

}
