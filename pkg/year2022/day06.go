package year2022

type Day06 struct{}

func (p Day06) PartA(lines []string) any {
	return getMarker(lines[0], 4)
}

func (p Day06) PartB(lines []string) any {
	return getMarker(lines[0], 14)
}

func getMarker(dataStream string, threshold int) int {
	var lastFour []rune
	for i, c := range dataStream {
		lastFour = append(lastFour, c)
		lastFourLen := len(lastFour)
		if lastFourLen >= threshold {
			if lastFourLen == threshold+1 {
				lastFour = lastFour[1:]
			}

			unique := true
			for j := 0; j < len(lastFour); j++ {
				for k := 0; k < len(lastFour); k++ {
					if j == k {
						continue
					}
					if lastFour[j] == lastFour[k] {
						unique = false
						break
					}
				}
			}
			if unique {
				return i + 1
			}
		}
	}
	return -1
}
