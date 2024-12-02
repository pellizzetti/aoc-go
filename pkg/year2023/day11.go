package year2023

import "strings"

type Day11 struct{}

func (p Day11) PartA(lines []string) any {
	// (9-1)+(9-2)+(9-3)+(9-4)+(9-5)+(9-6)+(9-7)+(9-8) = 36
	var galaxyMap = []string{
		"....#........",
		".............",
		".........#...",
	}

	return galaxyMap
}

func (p Day11) PartB(lines []string) any {
	return "implement_me"
}

func aaa(lines []string) {

	for _, line := range lines {
		hasGalaxy := false
		if strings.IndexAny(line, "#") >= 0 {

		}

		if !hasGalaxy {

		}
	}
}
