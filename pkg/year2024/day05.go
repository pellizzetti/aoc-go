package year2024

import (
	"strconv"
	"strings"
)

type Day05 struct{}

func reorderPages(updates []string, orderingRules map[string]map[string]struct{}) ([]string, bool) {
	swapped := false
	for i := 0; i < len(updates); i++ {
		for j := i + 1; j < len(updates); j++ {
			if _, found := orderingRules[updates[j]][updates[i]]; found {
				updates[i], updates[j] = updates[j], updates[i]
				swapped = true
			}
		}
	}
	return updates, swapped
}

func parseOrderingRules(lines []string, orderingRules map[string]map[string]struct{}) int {
	i := 0
	for {
		if len(lines[i]) == 0 {
			break
		}
		pages := strings.Split(lines[i], "|")
		prev, next := pages[0], pages[1]
		if _, found := orderingRules[prev]; !found {
			orderingRules[prev] = make(map[string]struct{})
		}
		orderingRules[prev][next] = struct{}{}
		i++
	}
	return i
}

func (p Day05) PartA(lines []string) any {
	orderingRules := make(map[string]map[string]struct{})
	sum := 0
	i := parseOrderingRules(lines, orderingRules)
	for i := i + 1; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			break
		}
		updates, reordered := reorderPages(strings.Split(lines[i], ","), orderingRules)
		if !reordered {
			middleValue, _ := strconv.Atoi(updates[len(updates)/2])
			sum += middleValue
		}
	}
	return sum
}

func (p Day05) PartB(lines []string) any {
	orderingRules := make(map[string]map[string]struct{})
	sum := 0
	i := parseOrderingRules(lines, orderingRules)
	for i := i + 1; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			break
		}
		updates, reordered := reorderPages(strings.Split(lines[i], ","), orderingRules)
		if reordered {
			middleValue, _ := strconv.Atoi(updates[len(updates)/2])
			sum += middleValue
		}
	}
	return sum
}
