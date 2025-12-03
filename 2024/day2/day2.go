package day2

import (
	"2024/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func isNotValid(l, r int, isIncreasing bool) bool {
	return (l < r) != isIncreasing || int(math.Abs(float64(r-l))) < 1 || int(math.Abs(float64(r-l))) > 3
}

func isLineSafe(line []string) bool {
	var isIncreasing bool = utils.StringToInt(line[0]) < utils.StringToInt(line[1])
	for j := 1; j < len(line); j++ {
		var l int = utils.StringToInt(line[j-1])
		var r int = utils.StringToInt(line[j])
		if isNotValid(l, r, isIncreasing) {
			return false
		}
	}
	return true
}

func checkAllVariations(line []string) bool {
	if isLineSafe(line) {
		return true
	} else {
		fmt.Println("Original Line", line)
		for i := 0; i < len(line); i++ {
			print(i)
			// IMPORTANT!!! NEED TO MAKE DEEP COPY
			tempLine := make([]string, len(line))
			copy(tempLine, line)
			alteredLine := remove(tempLine, i)
			// fmt.Println(alteredLine)
			if isLineSafe(alteredLine) {
				return true
			}
		}
	}
	return false
}

func Solve() {
	file, err := os.Open("day2.txt")
	if err != nil {
	}
	defer func() {
		if err = file.Close(); err != nil {

		}
	}()

	scanner := bufio.NewScanner(file)

	var count int = 0
	for scanner.Scan() {
		line := scanner.Text() // token (aka a line) in unicode-char
		lineAsList := strings.Fields(line)

		if checkAllVariations(lineAsList) {
			count++
		}
	}

	fmt.Println("Count: ", count)

}
