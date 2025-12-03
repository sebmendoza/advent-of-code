package main

import (
	"fmt"
	"os"
	"regexp"
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

func main() {
	// From PERPLEXITY
	content, err := os.ReadFile("day3.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Convert []byte to string
	fileContent := string(content)

	// fmt.Println(fileContent)
	r, _ := regexp.Compile(`mul\(\d+,\d+\)|do\(\)|don\'t\(\)`)
	fmt.Println(r.FindAllString(fileContent, -1))

	var total int
	enabled := true
	for _, value := range r.FindAllString(fileContent, -1) {
		fmt.Println(value)
		if value == "do()" {
			enabled = true
			fmt.Println("Enabled")
		} else if value == "don't()" {
			enabled = false
			fmt.Println("Disabled ")
		} else {
			if enabled {
				newr, _ := regexp.Compile(`\d+,\d+`)
				nums := strings.Split(newr.FindAllString(value, -1)[0], ",")
				first_num := stringToInt(nums[0])
				second_num := stringToInt(nums[1])
				total += (first_num * second_num)
			}
		}
		// fmt.Println(newr.FindAllString(value, -1))
		// fmt.Println(nums)
		// fmt.Printf("%T", newr.FindAllString(value, -1)[0][1])
	}
	fmt.Println(total)

}

// Part 1:  173731097
