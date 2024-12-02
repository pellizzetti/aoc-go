package year2023

type Day09 struct{}

func (p Day09) PartA(lines []string) any {
	total := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		numbers := getNumbers(line)
		allDiffs := make([][]int, 0)
		getDiffs(numbers, &allDiffs)
		allDiffs = append([][]int{numbers}, allDiffs...)

		total += getExtrapolated(allDiffs, false)
	}
	return total
}

func (p Day09) PartB(lines []string) any {
	total := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		numbers := getNumbers(line)
		allDiffs := make([][]int, 0)
		getDiffs(numbers, &allDiffs)
		allDiffs = append([][]int{numbers}, allDiffs...)

		total += getExtrapolated(allDiffs, true)
	}
	return total
}

func getDiffs(numbers []int, allDiffs *[][]int) {
	sumDiff := 0
	diffs := make([]int, 0)
	for i := 0; i < len(numbers)-1; i++ {
		diff := numbers[i+1] - numbers[i]
		sumDiff += diff
		diffs = append(diffs, diff)
	}
	*allDiffs = append(*allDiffs, diffs)
	if sumDiff == 0 {
		return
	}

	getDiffs(diffs, allDiffs)
}

func getExtrapolated(diffs [][]int, backwards bool) (extrapolated int) {
	for i := len(diffs) - 1; i >= 0; i-- {
		index := len(diffs[i]) - 1
		mult := 1
		if backwards {
			index = 0
			mult = -1
		}
		elem := diffs[i][index]
		if i == len(diffs)-1 {
			last := diffs[i][index]
			extrapolated = last + elem
		}
		extrapolated = (mult * extrapolated) + elem
	}

	return
}
