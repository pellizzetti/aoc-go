package year2024

type Day04 struct{}

func countWordOccurrences(grid [][]rune, word string) int {
	rows := len(grid)
	cols := len(grid[0])
	wordLength := len(word)
	directions := [][]int{
		{0, 1}, {0, -1}, // Horizontal
		{1, 0}, {-1, 0}, // Vertical
		{1, 1}, {1, -1}, // Diagonal down
		{-1, 1}, {-1, -1}, // Diagonal up
	}

	isValidPosition := func(r, c int) bool {
		return r >= 0 && r < rows && c >= 0 && c < cols
	}

	checkWord := func(r, c, directionRow, directionCol int) bool {
		for i := 0; i < wordLength; i++ {
			newRow, newCol := r+i*directionRow, c+i*directionCol
			if !isValidPosition(newRow, newCol) || grid[newRow][newCol] != rune(word[i]) {
				return false
			}
		}
		return true
	}

	count := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			for _, d := range directions {
				if checkWord(r, c, d[0], d[1]) {
					count++
				}
			}
		}
	}
	return count
}

func countMASOccurrences(grid [][]rune) int {
	rows := len(grid)
	cols := len(grid[0])

	// Deltas for diagonal directions
	deltas := [][][]int{
		{{-1, -1}, {1, 1}},
		{{-1, 1}, {1, -1}},
	}

	isValidPosition := func(r, c int) bool {
		return r >= 0 && r < rows && c >= 0 && c < cols
	}

	checkMAS := func(r, c int, deltas [][][]int) bool {
		for _, direction := range deltas {
			newRowUp, newColUp := r+direction[0][0], c+direction[0][1]
			newRowDown, newColDown := r+direction[1][0], c+direction[1][1]
			if !isValidPosition(newRowUp, newColUp) || !isValidPosition(newRowDown, newColDown) {
				return false
			}

			isMAS := grid[newRowUp][newColUp] == 'M' && grid[newRowDown][newColDown] == 'S'
			isSAM := grid[newRowUp][newColUp] == 'S' && grid[newRowDown][newColDown] == 'M'
			if !isMAS && !isSAM {
				return false
			}
		}
		return true
	}

	count := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 'A' && checkMAS(r, c, deltas) {
				count++
			}
		}
	}
	return count
}

func getGrid(lines []string) [][]rune {
	var grid [][]rune
	for _, line := range lines {
		if len(line) == 0 {
			break
		}
		grid = append(grid, []rune(line))
	}
	return grid
}

func (p Day04) PartA(lines []string) any {
	grid := getGrid(lines)
	return countWordOccurrences(grid, "XMAS")
}

func (p Day04) PartB(lines []string) any {
	grid := getGrid(lines)
	return countMASOccurrences(grid)
}
