package year2022

import (
	"strconv"
	"strings"
)

type Day04 struct{}

func (p Day04) PartA(lines []string) any {
	overlappingAssignments := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		x, y, x2, y2 := getSectionsFromLine(line)
		if (x2 >= x && y2 <= y) || (x >= x2 && y <= y2) {
			overlappingAssignments += 1
		}
	}
	return overlappingAssignments
}

func (p Day04) PartB(lines []string) any {
	overlappingAssignments := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		x, y, x2, y2 := getSectionsFromLine(line)
		if (x2 <= x && x2 >= y) || (x <= y2 && y >= x2) {
			overlappingAssignments += 1
		}
	}
	return overlappingAssignments
}

func getSectionsFromLine(line string) (int, int, int, int) {
	elvesPair := strings.Split(line, ",")
	firstSections := strings.Split(elvesPair[0], "-")
	secondSections := strings.Split(elvesPair[1], "-")
	x, _ := strconv.Atoi(firstSections[0])
	y, _ := strconv.Atoi(firstSections[1])
	x2, _ := strconv.Atoi(secondSections[0])
	y2, _ := strconv.Atoi(secondSections[1])
	return x, y, x2, y2
}
