package year2023

import (
	"regexp"
	"strconv"
)

type Day03 struct{}

func (p Day03) PartA(lines []string) any {
	sumPartNumbers := 0
	for i := 0; i < len(lines)-1; i++ {
		line := lines[i]
		if i > 0 {
			sumPartNumbers += getAdjacentsToSymbol(line, lines[i-1])
		}
		sumPartNumbers += getAdjacentsToSymbol(line, line)
		sumPartNumbers += getAdjacentsToSymbol(line, lines[i+1])
	}
	return sumPartNumbers
}

func (p Day03) PartB(lines []string) any {
	sumGearRatio := 0
	for i := 1; i < len(lines)-2; i++ {
		line := lines[i]
		prev := lines[i-1]
		next := lines[i+1]
		sumGearRatio += getGearsRatio(line, []string{prev, line, next})
	}
	return sumGearRatio
}

func getAdjacentsToGear(line string, gearIndex int) (adjacentsNumbers []int) {
	numberRegex := regexp.MustCompile(`\d{1,}`)
	numbersIndices := numberRegex.FindAllStringIndex(line, -1)
	for _, indexPair := range numbersIndices {
		startIndex := indexPair[0]
		endIndex := indexPair[1]
		number, _ := strconv.Atoi(line[startIndex:endIndex])
		if startIndex-1 <= gearIndex && gearIndex <= endIndex {
			adjacentsNumbers = append(adjacentsNumbers, number)
		}
	}
	return
}

func getGearsRatio(gearLine string, lines []string) (totalGearRatio int) {
	gearRegex := regexp.MustCompile(`[*]`)
	gearsIndices := gearRegex.FindAllStringIndex(gearLine, -1)
	if gearsIndices == nil {
		return
	}
	for _, gearIndexPair := range gearsIndices {
		adjacentsToGear := make([]int, 0)
		for _, line := range lines {
			adjacentsToGear = append(adjacentsToGear, getAdjacentsToGear(line, gearIndexPair[0])...)
		}
		if len(adjacentsToGear) != 2 {
			continue
		}
		totalGearRatio += adjacentsToGear[0] * adjacentsToGear[1]
	}
	return
}

func getAdjacentsToSymbol(numbersLine, symbolsLine string) (sumNumbers int) {
	symbols := make([]int, 0)
	symbolRegex := regexp.MustCompile(`[^\d.\n]+`)
	symbolsIndices := symbolRegex.FindAllStringIndex(symbolsLine, -1)
	for _, indexPair := range symbolsIndices {
		symbols = append(symbols, indexPair[0])
	}

	numberRegex := regexp.MustCompile(`\d{1,}`)
	numbersIndices := numberRegex.FindAllStringIndex(numbersLine, -1)
	for _, indexPair := range numbersIndices {
		number, _ := strconv.Atoi(numbersLine[indexPair[0]:indexPair[1]])
		for _, symbolIndex := range symbols {
			if indexPair[0]-1 <= symbolIndex && symbolIndex <= indexPair[1] {
				sumNumbers += number
				break
			}
		}
	}
	return
}
