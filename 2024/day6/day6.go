package main

import (
	"bufio"
	"fmt"
	"os"
)

func printGridWithHighlight(grid []string, row, col int) {
	// ANSI escape codes for colors
	red := "\033[31m"
	reset := "\033[0m"

	for i, line := range grid {
		for j, char := range line {
			if i == row && j == col {
				// Print character in red
				fmt.Print(red, string(char), reset)
			} else {
				// Print character normally
				fmt.Print(string(char))
			}
		}
		fmt.Println()
	}
	fmt.Println("\n")
}

func part1_loop(bp []string, start_pos [2]int) map[[2]int]struct{} {
	x, y := start_pos[0], start_pos[1]
	direction := "UP"
	count := 0
	visited := make(map[[2]int]struct{})
	for {
		if _, ok := visited[[2]int{x, y}]; !ok {
			count++
			visited[[2]int{x, y}] = struct{}{}
		}
		// fmt.Println("The curr pos: ", x, y)
		// printGridWithHighlight(bp, y, x)

		anticipated := getNextStep(direction, x, y)
		anticipated_x, anticipated_y := anticipated[0], anticipated[1]
		nextIsOutOfBounds := anticipated_x < 0 || anticipated_x >= len(bp[0]) || anticipated_y < 0 || anticipated_y >= len(bp)
		if nextIsOutOfBounds {
			break
		}
		isObstacle := bp[anticipated_y][anticipated_x] == 35
		if isObstacle {
			direction = switchDirection(direction)
			new_anticipated := getNextStep(direction, x, y)
			new_anticipated_x, new_anticipated_y := new_anticipated[0], new_anticipated[1]
			anotherObstacle := bp[new_anticipated_y][new_anticipated_x] == 35
			if anotherObstacle {
				direction = switchDirection(direction)
			}

		}
		nextStep := getNextStep(direction, x, y)
		x, y = nextStep[0], nextStep[1]
	}
	return visited

}

func checkCandidateLoop(bp []string, start_pos [2]int, insertObstacle [2]int) int {
	the_copy := make([]string, len(bp))
	copy(the_copy, bp)
	row := []rune(the_copy[insertObstacle[1]])
	row[insertObstacle[0]] = '#'
	the_copy[insertObstacle[1]] = string(row)
	// printGridWithHighlight(the_copy, insertObstacle[1], insertObstacle[0])
	// fmt.Print(insertObstacle)
	path := make(map[PathState]struct{})

	x, y := start_pos[0], start_pos[1]
	direction := "UP"
	for {
		if _, ok := path[PathState{x, y, direction}]; ok {
			return 1
		} else {
			// fmt.Println("This was not in the path", PathState{x, y, direction})
			// printGridWithHighlight(the_copy, y, x)
			// fmt.Println("Check Path: ", path)
			// fmt.Println("----------")

			// jsonBytes, _ := json.MarshalIndent(path, "", "    ")
			// fmt.Println(string(jsonBytes))

			path[PathState{x, y, direction}] = struct{}{}
		}
		anticipated := getNextStep(direction, x, y)
		anticipated_x, anticipated_y := anticipated[0], anticipated[1]
		nextIsOutOfBounds := anticipated_x < 0 || anticipated_x >= len(the_copy[0]) || anticipated_y < 0 || anticipated_y >= len(the_copy)
		if nextIsOutOfBounds {
			break
		}
		isObstacle := the_copy[anticipated_y][anticipated_x] == 35
		if isObstacle {
			direction = switchDirection(direction)
		}

		nextStep := getNextStep(direction, x, y)
		x, y = nextStep[0], nextStep[1]
	}
	return 0
}

func main() {
	file, _ := os.Open("day6.txt")
	scanner := bufio.NewScanner(file)
	start_pos := [2]int{} // { x, y}
	blueprint := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		blueprint = append(blueprint, line)
	}

	for i := range blueprint {
		for j := range blueprint[i] {
			if blueprint[i][j] == 94 {
				start_pos[0] = j
				start_pos[1] = i
			}
		}

	}
	// // Now Onto Part 2
	candidates := part1_loop(blueprint, start_pos)
	fmt.Println("All Valid Paths: ", len(candidates))
	// delete(candidates, [2]int{start_pos[0], start_pos[1]})
	total := 0
	// fmt.Println(candidates)
	for k := range candidates {
		total += checkCandidateLoop(blueprint, start_pos, [2]int{k[0], k[1]})
	}
	fmt.Println("Total: ", total) // 1721
}

func getNextStep(dir string, x int, y int) [2]int {
	switch dir {
	case "UP":
		return [2]int{x, y - 1}
	case "RIGHT":
		return [2]int{x + 1, y}
	case "DOWN":
		return [2]int{x, y + 1}
	case "LEFT":
		return [2]int{x - 1, y}
	default:
		panic("invalid direction")
	}
}

func switchDirection(d string) string {
	switch d {
	case "UP":
		return "RIGHT"
	case "RIGHT":
		return "DOWN"
	case "DOWN":
		return "LEFT"
	case "LEFT":
		return "UP"
	default:
		panic("invalid direction")
	}
}

type PathState struct {
	x, y int
	d    string
}

// 	fmt.Println("Part 2 count: ", valid_obstacle_placements)
// }

// DEAD CODE
// from main
// WIDTH := len(blueprint[0])
// HEIGHT := len(blueprint)
// visited := make(map[Coord]struct{})
// path_travelled := travel([2]int{start_pos[0], start_pos[1]}, blueprint, "UP", WIDTH, HEIGHT, visited)
// fmt.Println("Amount Travelled: ", path_travelled)
// // fmt.Println(visited)

// func travel(coords [2]int, bp []string, dir string, WIDTH int, HEIGHT int, visited map[Coord]struct{}) int {
// 	x, y := coords[0], coords[1]
// 	// printGridWithHighlight(bp, y, x)
// 	anticipated := getNextStep(dir, x, y)
// 	anticipated_x, anticipated_y := anticipated[0], anticipated[1]
// 	if anticipated_x < 0 || anticipated_x >= WIDTH || anticipated_y < 0 || anticipated_y >= HEIGHT {
// 		return 1
// 	}

// 	isObstacle := bp[anticipated_y][anticipated_x] == 35
// 	if isObstacle {
// 		dir = switchDirection(dir)
// 	}

// 	if _, exists := visited[Coord{x, y}]; exists {
// 		return travel(getNextStep(dir, x, y), bp, dir, WIDTH, HEIGHT, visited)
// 	} else {
// 		visited[Coord{x, y}] = struct{}{}
// 		return 1 + travel(getNextStep(dir, x, y), bp, dir, WIDTH, HEIGHT, visited)
// 	}

// }
