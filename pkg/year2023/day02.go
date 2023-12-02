package year2023

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Day02 struct{}

func (p Day02) PartA(lines []string) any {
	targetCounts := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	totalPossibleGames := 0
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if len(line) == 0 {
			continue
		}
		gameID, cubeCounts := getCubesCounts(line)
		if isValidGame(cubeCounts, targetCounts) {
			totalPossibleGames += gameID
		}
	}
	return totalPossibleGames
}

func (p Day02) PartB(lines []string) any {
	sumPowerOfSet := 0
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if len(line) == 0 {
			continue
		}
		_, cubeCounts := getCubesCounts(line)
		powerOfSet := getPowerOfSet(cubeCounts)
		sumPowerOfSet += powerOfSet
	}
	return sumPowerOfSet
}

func getPowerOfSet(cubeCounts []map[string]int) int {
	counts := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	for i := 0; i < len(cubeCounts); i++ {
		for color, count := range cubeCounts[i] {
			if count > counts[color] {
				counts[color] = count
			}
		}
	}
	return counts["red"] * counts["blue"] * counts["green"]
}

func isValidGame(cubeCounts []map[string]int, targetCounts map[string]int) bool {
	for i := 0; i < len(cubeCounts); i++ {
		for color, count := range cubeCounts[i] {
			targetCount, exists := targetCounts[color]
			if !exists || count > targetCount {
				fmt.Printf("Fail: %s == %d\n", color, count)
				return false
			}
		}
	}
	return true
}

func getCubesCounts(line string) (int, []map[string]int) {
	gameMatch := regexp.MustCompile(`Game (\d+): (.+?);`)
	match := gameMatch.FindStringSubmatch(line)
	gameID, _ := strconv.Atoi(match[1])

	colorCounts := make([]map[string]int, 0)
	gameResults := strings.Split(line, ";")
	for i, gameResult := range gameResults {
		colorMatches := regexp.MustCompile(`(\d+) (\w+)`)
		matches := colorMatches.FindAllStringSubmatch(gameResult, -1)
		colorCount := make(map[string]int)
		for _, match := range matches {
			count, _ := strconv.Atoi(match[1])
			color := match[2]
			colorCount[color] += count
			if len(colorCounts) == i {
				colorCounts = append(colorCounts, colorCount)
			} else {
				colorCounts[i][color] = count
			}
		}
	}
	return gameID, colorCounts
}
