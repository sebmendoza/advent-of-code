package utils

import (
	"fmt"
	"time"
)

// Solution represents a daily solution
type Solution interface {
	Part1() (string, error)
	Part2() (string, error)
}

// Runner handles running and displaying solutions
type Runner struct {
	Solutions map[int]Solution
}

// NewRunner returns a map of each day and respective solution
func NewRunner() *Runner {
	return &Runner{
		Solutions: make(map[int]Solution),
	}
}

// Add will add a day-solution pair
func (r *Runner) Add(day int, solution Solution) {
	r.Solutions[day] = solution
}

// Run will ha
func (r *Runner) Run(day int) {
	solution, exists := r.Solutions[day]
	if !exists {
		fmt.Printf("❌ Day %d not implemented yet\n", day)
		return
	}

	fmt.Printf("\n🎄 Day %d 🎄\n", day)
	fmt.Println("================")

	start := time.Now()
	part1, err := solution.Part1()
	duration1 := time.Since(start)
	if err != nil {
		fmt.Printf("Part 1 error: %v\n", err)
	} else {
		fmt.Printf("Part 1: %s (%v)\n", part1, duration1)
	}

	start = time.Now()
	part2, err := solution.Part2()
	duration2 := time.Since(start)
	if err != nil {
		fmt.Printf("Part 2 error: %v\n", err)
	} else {
		fmt.Printf("Part 2: %s (%v)\n", part2, duration2)
	}
	fmt.Println()
}
