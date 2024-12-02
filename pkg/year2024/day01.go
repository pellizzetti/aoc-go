package year2024

import (
	"sort"
	"strconv"
	"strings"
)

type Day01 struct{}

func (p Day01) PartA(lines []string) any {
	leftList, rightList := getLists(lines)
	sort.Ints(leftList)
	sort.Ints(rightList)

	var totalDistance int
	for i, leftValue := range leftList {
		v := rightList[i] - leftValue
		if v < 0 {
			v *= -1
		}
		totalDistance += v
	}
	return totalDistance
}

func (p Day01) PartB(lines []string) any {
	var leftList, rightList = getLists(lines)
	var similarityScore int
	for _, leftValue := range leftList {
		var matches int
		for _, rightValue := range rightList {
			if leftValue == rightValue {
				matches++
			}
		}
		similarityScore += leftValue * matches
	}

	return similarityScore
}

func getLists(lines []string) (leftList []int, rightList []int) {
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		split := strings.Split(line, "   ")

		leftValue, _ := strconv.Atoi(split[0])
		leftList = append(leftList, leftValue)

		rightValue, _ := strconv.Atoi(split[1])
		rightList = append(rightList, rightValue)
	}
	return
}
