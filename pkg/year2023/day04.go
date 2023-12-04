package year2023

import (
	"strconv"
	"strings"
)

type Day04 struct{}

func (p Day04) PartA(lines []string) any {
	totalPoints := 0
	for i := 0; i < len(lines)-1; i++ {
		totalPoints += getCardPoints(lines[i])
	}

	return totalPoints
}

func (p Day04) PartB(lines []string) any {
	totalScratchcards := 0
	scratchcards := make([]int, len(lines)-1)
	for i := 0; i < len(lines)-1; i++ {
		totalScratchcards += getCardCopies(lines[i], i, &scratchcards)
	}

	return totalScratchcards
}

func getCardPoints(card string) (totalPoints int) {
	parts := strings.Split(card, "|")
	winningNumbers := getNumbers(parts[0])
	cardNumbers := getNumbers(parts[1])
	for i := 0; i < len(cardNumbers); i++ {
		for j := 0; j < len(winningNumbers); j++ {
			if cardNumbers[i] == winningNumbers[j] {
				if totalPoints *= 2; totalPoints == 0 {
					totalPoints += 1
				}
			}
		}
	}

	return
}

func getCardCopies(card string, cardIndex int, scratchcards *[]int) (totalCopies int) {
	parts := strings.Split(card, "|")
	winningNumbers := getNumbers(parts[0])
	cardNumbers := getNumbers(parts[1])
	copyIndex := cardIndex + 1
	(*scratchcards)[cardIndex]++
	for i := 0; i < len(cardNumbers); i++ {
		for j := 0; j < len(winningNumbers); j++ {
			if cardNumbers[i] == winningNumbers[j] {
				(*scratchcards)[copyIndex] += 1 * (*scratchcards)[cardIndex]
				copyIndex++
			}
		}
	}
	totalCopies = (*scratchcards)[cardIndex]

	return
}

func getNumbers(part string) []int {
	fields := strings.Fields(part)
	numbers := make([]int, len(fields))
	for i, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			continue
		}
		numbers[i] = num
	}

	return numbers
}
