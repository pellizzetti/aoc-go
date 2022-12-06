package year2022

import (
	"unicode"
)

type Day03 struct{}

type void struct{}

func (p Day03) PartA(lines []string) any {
	priority := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		containerSize := len(line) / 2
		firstContainer := line[:containerSize]
		secondContainer := line[containerSize:]
		var commonItem rune
		for _, item := range firstContainer {
			for _, otherItem := range secondContainer {
				if item == otherItem {
					commonItem = item
					break
				}
			}
		}
		priority += getItemPriority(commonItem)
	}
	return priority
}

func (p Day03) PartB(lines []string) any {
	priority := 0
	linesInGroup := 3
	var rucksacksItemList []map[rune]void
	for i, line := range lines {
		if line == "" {
			continue
		}
		position := i % linesInGroup
		var itemMap = make(map[rune]void)
		for _, item := range line {
			var v void
			itemMap[item] = v
		}
		rucksacksItemList = append(rucksacksItemList, itemMap)
		if position == 2 {
			lastLine := rucksacksItemList[position]
			for key := range lastLine {
				isCommon := true
				for _, rucksacksItems := range rucksacksItemList {
					if _, found := rucksacksItems[key]; !found {
						isCommon = false
					}
				}
				if isCommon {
					rucksacksItemList = nil
					priority += getItemPriority(key)
					break
				}
			}
		}
	}
	return priority
}

func getItemPriority(item rune) int {
	if unicode.IsUpper(item) {
		return int(item) - 38
	}
	return int(item) - 96
}
