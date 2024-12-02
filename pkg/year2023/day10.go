package year2023

import (
	"fmt"
	"image"
	"slices"
	"strings"
)

type Day10 struct{}

func (p Day10) PartA(lines []string) any {
	// lines = []string{
	// 	"..F7.",
	// 	".FJ|.",
	// 	"SJ.L7",
	// 	"|F--J",
	// 	"LJ...",
	// }

	grid, start := map[image.Point]rune{}, image.Point{}
	for y, s := range lines {
		for x, r := range s {
			grid[image.Point{x, y}] = r
			if r == 'S' {
				start = image.Point{x, y}
			}
		}
	}

	grid[start] = map[[4]bool]rune{
		{true, false, true, false}: '|', {false, true, false, true}: '-',
		{true, true, false, false}: 'L', {true, false, false, true}: 'J',
		{false, false, true, true}: '7', {false, true, true, false}: 'F',
	}[[4]bool{
		strings.ContainsRune("7F|", grid[start.Add(image.Point{0, -1})]),
		strings.ContainsRune("-7J", grid[start.Add(image.Point{1, 0})]),
		strings.ContainsRune("JL|", grid[start.Add(image.Point{0, 1})]),
		strings.ContainsRune("-FL", grid[start.Add(image.Point{-1, 0})]),
	}]

	path, area := []image.Point{start}, 0
	for p, n := start, start; ; path = append(path, n) {
		p, n = n, start
		for _, d := range map[rune][]image.Point{
			'|': {{0, -1}, {0, 1}}, '-': {{1, 0}, {-1, 0}}, 'L': {{0, -1}, {1, 0}},
			'J': {{0, -1}, {-1, 0}}, '7': {{0, 1}, {-1, 0}}, 'F': {{0, 1}, {1, 0}},
		}[grid[p]] {
			if !slices.Contains(path, p.Add(d)) {
				n = p.Add(d)
			}
		}
		area += p.X*n.Y - p.Y*n.X
		if n == start {
			break
		}
	}

	fmt.Println((slices.Max([]int{area})-len(path))/2 + 1)

	return len(path) / 2
}

func (p Day10) PartB(lines []string) any {
	// lines = []string{}

	history, count := findCount(lines)
	numberOfInsideElements, visual := findArea(history, lines)
	fmt.Println("answer for part 1: ", count)
	fmt.Println("answer for part 2: ", numberOfInsideElements)

	for i := 0; i < len(lines); i++ {
		if _, ok := visual[i]; ok {
			fmt.Println(visual[i].v, visual[i].count)
		}
	}
	// total := 0
	// for _, line := range lines {
	// 	if len(line) == 0 {
	// 		continue
	// 	}
	// 	total++
	// }
	return 0
}

func findS(input []string) (int, int) {
	for i, line := range input {
		for j, char := range line {
			if char == 'S' {
				return i, j
			}
		}
	}
	return 0, 0
}

type position struct {
	x int
	y int
}

func findCount(input []string) ([]position, int) {
	history := []position{}

	x, y := findS(input)
	history = append(history, position{
		x: x,
		y: y,
	})

	x, y = getStartingPosition(x, y, input)

	for input[x][y] != 'S' {
		lastP := history[len(history)-1]
		x1, y1 := findNext(x, y, input, lastP.x, lastP.y)
		history = append(history, position{
			x: x,
			y: y,
		})
		x, y = x1, y1
	}
	return history, len(history) / 2
}

func getStartingPosition(x, y int, input []string) (int, int) {

	if y != (len(input[x])-1) && (input[x][y+1] == '-' || input[x][y+1] == 'J' || input[x][y+1] == '7') {
		y++
		return x, y
	}

	if x != (len(input)-1) && (input[x+1][y] == '|' || input[x+1][y] == 'J' || input[x+1][y] == 'L') {
		x++
		return x, y
	}

	if y != 0 && (input[x][y-1] == 'F' || input[x][y-1] == '-' || input[x][y-1] == 'L') {
		y--
		return x, y
	}

	if x != 0 && (input[x-1][y] == '|' || input[x-1][y] == 'L' || input[x-1][y] == '7') {
		x--
		return x, y
	}
	return x, y
}

func findNext(x, y int, input []string, lastx, lasty int) (int, int) {
	if input[x][y] == '-' {
		if y != 0 && lasty == (y-1) {
			y++
			return x, y
		}

		if y != (len(input[x])-1) && lasty == (y+1) {
			y--
			return x, y
		}
	}

	if input[x][y] == 'J' {
		if x != 0 && y != 0 && lastx == (x-1) {
			y--
			return x, y
		}

		if x != 0 && y != 0 && lasty == (y-1) {
			x--
			return x, y
		}
	}

	if input[x][y] == '|' {
		if x != 0 && x != len(input)-1 && lastx == (x-1) {
			x++
			return x, y
		}

		if x != 0 && x != len(input)-1 && lastx == (x+1) {
			x--
			return x, y
		}
	}

	if input[x][y] == 'L' {
		if x != 0 && y != len(input[x])-1 && lastx == (x-1) {
			y++
			return x, y
		}

		if x != 0 && y != len(input[x])-1 && lasty == (y+1) {
			x--
			return x, y
		}
	}

	if input[x][y] == '7' {
		if y != 0 && x != (len(input)-1) && lasty == (y-1) {
			x++
			return x, y
		}

		if x != (len(input)-1) && y != 0 && lastx == (x+1) {
			y--
			return x, y
		}
	}

	if input[x][y] == 'F' {
		if x != (len(input)-1) && y != (len(input[x])-1) && lasty == (y+1) {
			x++
			return x, y
		}

		if x != (len(input)-1) && y != (len(input[x])-1) && lastx == (x+1) {
			y++
			return x, y
		}
	}

	return x, y
}

type zoo struct {
	v     string
	count int
}

func findArea(path []position, input []string) (int, map[int]zoo) {
	replaceWith := map[string]string{}
	replaceWith["J"] = "┘"
	replaceWith["L"] = "└"
	replaceWith["7"] = "┐"
	replaceWith["F"] = "┌"
	replaceWith["|"] = "│"
	replaceWith["-"] = "─"

	mapPosition := make(map[int][]int)
	resultMap := make(map[int]zoo)
	sum := 0

	for _, p := range path {
		mapPosition[p.x] = append(mapPosition[p.x], p.y)
	}

	for k, v := range mapPosition {
		a := strings.Split(input[k], "")
		for _, j := range v {
			if a[j] != "S" {
				a[j] = replaceWith[a[j]]
			} else {
				a[j] = replaceWith[string(replaceSWith(path[0], path[1], path[len(path)-1], input))]
			}
		}
		// clean edges
		left, right := false, false
		for i, j := range a {
			if i == (len(a) - 1 - i) {
				break
			}
			r := a[len(a)-i-1]
			if j == "┘" || j == "└" || j == "┐" || j == "┌" || j == "│" || j == "─" {
				left = true
			}
			if r == "┘" || r == "└" || r == "┐" || r == "┌" || r == "│" || r == "─" {
				right = true
			}

			if !left {
				a[i] = " "
			}
			if !right {
				a[len(a)-1-i] = " "
			}
		}
		count := 0
		isInside := 0
		last := "-"
		// this is the crux of the solution:
		// we count the number of | encountered in pipe path when traversing each  cleaned input line
		// if it is odd, that means the character encountered is inside the pipe path, else outside
		// corner cases:
		// 		- (F is followed by a J) or (L is followed by a 7) then its equivalent to a single |
		// 		- '-' does not affect count
		// 		- J, L, 7, F do not affect count individually, unless 1st case is encountered
		for i, char := range a {
			if char == "│" {
				isInside++
				continue
			}
			if char == "─" {
				continue
			}

			if last == "-" && (char == "┘" || char == "└" || char == "┐" || char == "┌") {
				if char == "┘" || char == "└" || char == "┐" || char == "┌" {
					last = char
					continue
				}
			} else if last != "-" && (char == "┘" || char == "└" || char == "┐" || char == "┌") {
				if last == "└" && char == "┐" {
					isInside++
				}

				if last == "┌" && char == "┘" {
					isInside++
				}
				last = "-"
				continue
			}

			if isInside%2 == 0 {
				a[i] = " "
			}

			if isInside%2 != 0 {
				a[i] = "\033[0;32m█\033[0m"
				count++
			}
		}
		resultMap[k] = zoo{
			v:     strings.Join(a, ""),
			count: count,
		}
		sum += count
	}

	return sum, resultMap
}

func replaceSWith(s, first, last position, input []string) byte {

	for _, char := range "-|JL7F" {
		duplicateInput := input
		a := strings.Split(duplicateInput[s.x], "")
		a[s.y] = string(char)
		duplicateInput[s.x] = strings.Join(a, "")
		x, y := findNext(s.x, s.y, duplicateInput, last.x, last.y)
		if x == first.x && y == first.y {
			return byte(char)
		}
	}

	return 'S'
}
