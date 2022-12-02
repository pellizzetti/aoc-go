package year2022

import (
	"sort"
	"strconv"
)

type Day01 struct{}

func (p Day01) PartA(lines []string) any {
	elves := getElves(lines)
	maxCalories := elves[len(elves)-1]
	return maxCalories
}

func (p Day01) PartB(lines []string) any {
	elves := getElves(lines)
	topElves := elves[len(elves)-3:]
	return topElves[0] + topElves[1] + topElves[2]
}

func getElves(lines []string) []int {
	var elves []int
	caloriesCount := 0
	for _, line := range lines {
		if line == "" {
			elves = append(elves, caloriesCount)
			caloriesCount = 0
			continue
		}
		calories, _ := strconv.Atoi(line)
		caloriesCount += calories
	}
	sort.Ints(elves)
	return elves
}
