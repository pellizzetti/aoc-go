package year2023

import (
	"math"
)

type Day08 struct{}

func (p Day08) PartA(lines []string) any {
	instructions := []rune(lines[0])
	left := make(map[string]string)
	right := make(map[string]string)
	for _, line := range lines[2:] {
		if len(line) == 0 {
			continue
		}
		key := line[0:3]
		left[key] = line[7:10]
		right[key] = line[12:15]
	}

	node := "AAA"
	steps := 0
	for node != "ZZZ" {
		move := instructions[steps%len(instructions)]
		if move == 'L' {
			node = left[node]
		} else {
			node = right[node]
		}
		steps++
	}

	return steps
}

func (p Day08) PartB(lines []string) any {
	left := make(map[string]string)
	right := make(map[string]string)
	for _, line := range lines[2:] {
		if len(line) == 0 {
			continue
		}
		key := line[0:3]
		left[key] = line[7:10]
		right[key] = line[12:15]
	}

	totalSteps := int64(1)
	for key := range left {
		node := key
		if node[len(node)-1:] == "A" {
			steps := 0
			instructions := []rune(lines[0])
			for !(node[len(node)-1:] == "Z") {
				move := instructions[steps%len(instructions)]
				if move == 'L' {
					node = left[node]
				} else {
					node = right[node]
				}
				steps++
			}
			totalSteps = LCM(totalSteps, int64(steps))
		}
	}

	return totalSteps
}

func GCD(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM(a, b int64) int64 {
	return int64(math.Abs(float64(a*b)) / float64(GCD(a, b)))
}
