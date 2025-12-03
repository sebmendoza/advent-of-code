package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkDir(fir byte, sec byte, thi byte) int {
	if fir == 'M' && sec == 'A' && thi == 'S' {
		return 1
	}
	return 0

}

func part1_word_search(ws []string) {
	hashmap := make(map[string][][]int)
	hashmap["right"] = [][]int{{1, 0}, {2, 0}, {3, 0}}
	hashmap["left"] = [][]int{{-1, 0}, {-2, 0}, {-3, 0}}
	hashmap["up"] = [][]int{{0, -1}, {0, -2}, {0, -3}}
	hashmap["down"] = [][]int{{0, 1}, {0, 2}, {0, 3}}
	hashmap["d-up-right"] = [][]int{{1, -1}, {2, -2}, {3, -3}}
	hashmap["d-up-left"] = [][]int{{-1, -1}, {-2, -2}, {-3, -3}}
	hashmap["d-down-right"] = [][]int{{1, 1}, {2, 2}, {3, 3}}
	hashmap["d-down-left"] = [][]int{{-1, 1}, {-2, 2}, {-3, 3}}

	LENGTH_ROW := len(ws[0])
	LENGTH_WS := len(ws)
	count := 0

	for i := range ws {
		for j := range ws[i] {
			if ws[i][j] == 'X' {
				invalid_r := j+3 >= LENGTH_ROW
				invalid_l := j-3 < 0
				invalid_up := i-3 < 0
				invalid_d := i+3 >= LENGTH_WS
				for key, val := range hashmap {
					switch {
					case key == "right" && invalid_r:
						continue
					case key == "left" && invalid_l:
						continue
					case key == "up" && invalid_up:
						continue
					case key == "down" && invalid_d:
						continue
					case key == "d-up-right" && (invalid_up || invalid_r):
						continue
					case key == "d-up-left" && (invalid_up || invalid_l):
						continue
					case key == "d-down-right" && (invalid_d || invalid_r):
						continue
					case key == "d-down-left" && (invalid_d || invalid_l):
						continue
					default:
						f := ws[i+val[0][1]][j+val[0][0]]
						s := ws[i+val[1][1]][j+val[1][0]]
						t := ws[i+val[2][1]][j+val[2][0]]
						count += checkDir(f, s, t)
					}
				}
			}
		}
	}
	fmt.Println("Count: ", count) // Part 1: Count:  2454
}

func main() {
	file, _ := os.Open("day4.txt")
	scanner := bufio.NewScanner(file)
	ws := []string{}
	for scanner.Scan() {
		line := scanner.Text() // token (aka a line) in unicode-char
		ws = append(ws, line)
	}

	// part1_word_search(ws)
	part2_X_MAS(ws)
}
func part2_X_MAS(ws []string) {
	// focus on the A!
	count := 0
	// X_patterns := [][]int{{-1, -1}, {1, -1}, {-1, 1}, {1, 1}} //  topl, topr, downl, downl { horizontal, vertical }
	for i := range ws {
		for j := range ws[i] {
			if ws[i][j] == 'A' {
				if i <= 0 || i+1 >= len(ws) || j <= 0 || j+1 >= len(ws[i]) { // invalid
					continue
				}

				topl := ws[i-1][j-1]
				topr := ws[i-1][j+1]
				downl := ws[i+1][j-1]
				downr := ws[i+1][j+1]

				// configs: X on top, X on bottom, X on rigt, X on left
				if topl == 'M' && topr == 'M' && downl == 'S' && downr == 'S' {
					count++
				} else if topl == 'M' && downl == 'M' && topr == 'S' && downr == 'S' {
					count++
				} else if downl == 'M' && downr == 'M' && topr == 'S' && topl == 'S' {
					count++
				} else if downr == 'M' && topr == 'M' && topl == 'S' && downl == 'S' {
					count++
				}
			}
		}
	}
	fmt.Println("Count:  ", count)
}

// the way to get aorund the problem of double counting is to only check right and left. Never check up or down and we'll never double count
