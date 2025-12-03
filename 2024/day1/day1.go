package day1

import (
	"2024/utils"
	"bufio"
	"fmt"
	"math"
	"sort"
	"strings"
)

type Solution struct {
	input string
}

func New() *Solution {
	return &Solution{
		input: "day1/day1.txt",
	}
}

func (s *Solution) Part1() (string, error) {
	first_ids, second_ids, err := pre_calculations(s)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", calculateDistances(first_ids, second_ids)), nil // Part f1: 2264607
}

func calculateDistances(slice1 []int, slice2 []int) int {
	var total_dist int = 0
	for i := 0; i < len(slice1); i++ {
		total_dist += int(math.Abs(float64(slice1[i] - slice2[i]))) // Abs only takes float64, not ints
	}

	return total_dist
}

func (s *Solution) Part2() (string, error) {
	first_ids, second_ids, _ := pre_calculations(s)
	return fmt.Sprintf("%d", getSimilarityScore(first_ids, second_ids)), nil //Part 2: 19457120
}

func getSimilarityScore(slice1 []int, slice2 []int) int {
	// How many times does each number in left list appear in the other?
	// Let's firgure out how to loop through the list only printing the numbers in left list.
	total_length := 0
	for i := 0; i < len(slice1); i++ {
		if i == 0 {
			total_length += 1
			continue
		}
		if slice2[i-1] == slice2[i] {
			continue
		}
		total_length += 1
	}

	// There are no duplicates in the left! Weird

	// -------------------------------------------
	// Count appearances in slice2
	sim_score := 0
	for i := 0; i < len(slice1); i++ {
		total_duplicates := 0
		num := slice1[i]
		for j := 0; j < len(slice2); j++ {
			if num == slice2[j] {
				total_duplicates += 1
			}
		}
		sim_score += total_duplicates * num
	}
	return sim_score
}

func pre_calculations(s *Solution) ([]int, []int, error) {
	file, err := utils.OpenFile(s.input)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var first_ids []int
	var second_ids []int
	for scanner.Scan() {
		line := scanner.Text() // token (aka a line) in unicode-char
		fields := strings.Fields(line)
		first_ids = append(first_ids, utils.StringToInt(fields[0]))
		second_ids = append(second_ids, utils.StringToInt(fields[1]))
	}

	sort.Ints(first_ids)
	sort.Ints(second_ids)

	return first_ids, second_ids, nil
}
