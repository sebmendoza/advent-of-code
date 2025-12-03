package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func stringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func violatesRules(arr []string, rulesmap map[string][]string) bool {
	for i, val := range arr {
		for j := i + 1; j < len(arr); j++ {
			// If arr[j] must come before arr[i], this is a violation
			if mustComeBefore, exists := rulesmap[val]; exists {
				for _, before := range mustComeBefore {
					if before == arr[j] {
						return true
					}
				}
			}
		}
	}
	return false
}

// Some rendition of bubble sort. Adapted to our own use cases.
func fix_line(arr []string, rulesmap map[string][]string) []string {
	// fmt.Println("Here is the rulesmaps: ", rulesmap)
	n := len(arr)
	result := make([]string, n)
	copy(result, arr)

	for {
		swapped := false
		for i := 0; i < n-1; i++ {
			// fmt.Println("The i and val,", i, " ", result[i])

			// Check if i+1 should come before i according to rules
			if mustComeBefore, exists := rulesmap[result[i]]; exists {
				for _, before := range mustComeBefore {
					if before == result[i+1] {
						// Swap them
						result[i], result[i+1] = result[i+1], result[i]
						// fmt.Println(arr, result)
						swapped = true
					}
				}
			}
		}
		// If no swaps needed or performed, we're done
		if !swapped {
			break
		}
	}
	return result
}

func getMidVal(arr []string) int {
	return stringToInt(arr[int(math.Floor(float64(len(arr)/2)))])
}

func main() {
	file, _ := os.Open("day5.txt")
	scanner := bufio.NewScanner(file)
	rules := []string{}
	updates := []string{}
	flag := false
	for scanner.Scan() {
		line := scanner.Text() // token (aka a line) in unicode-char
		if line == "" {
			flag = true
		} else if flag {
			updates = append(updates, line)
		} else {
			rules = append(rules, line)
		}

	}
	// fmt.Println("length of uodate", len(updates), updates[0], updates[len(updates)-1])

	rulesmap := make(map[string][]string)
	for _, val := range rules {
		key := strings.Split(val, "|")[1]
		required_before := strings.Split(val, "|")[0]
		rulesmap[key] = append(rulesmap[key], required_before)
	}
	// fmt.Printf("%T", rulesmap["helloworld"])
	valid_counts := 0
	valid_summation := 0
	invalid_count := 0
	invalid_summation := 0
	for _, val := range updates {
		isLineValid := true
		arrayed_line := strings.Split(val, ",")
		// fmt.Println(arrayed_line)
		forbidden_list := make(map[string]string)
		for _, num := range arrayed_line {
			_, exists := forbidden_list[num]

			// Break and add nothing if in forbidden list
			if exists {
				// fmt.Println("CAUGHT")
				isLineValid = false
				break
			}

			for _, newNum := range rulesmap[num] {
				forbidden_list[newNum] = newNum
			}

		}
		if isLineValid {
			// fmt.Println("VALID ONE", arrayed_line)
			valid_counts += 1
			valid_summation += getMidVal(arrayed_line)
		} else {
			// fmt.Println("InxValid Lines: ", arrayed_line)
			new_ordered_line := fix_line(arrayed_line, rulesmap)
			// fmt.Println("New Lines: ", new_ordered_line)
			invalid_count += 1
			invalid_summation += getMidVal(new_ordered_line)

		}

	}
	fmt.Println("counts:", valid_counts)      // count: 82
	fmt.Println("sum:", valid_summation)      // sum: 4957
	fmt.Println("in_counts:", invalid_count)  // count: 115
	fmt.Println("in_sum:", invalid_summation) // sum: 6951

}
