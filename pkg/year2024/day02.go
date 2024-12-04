package year2024

import (
	"strconv"
	"strings"
)

type Day02 struct{}

func getSafeReports(lines []string, allowDampener bool) (safeReports int) {
	for _, line := range lines {
		if len(line) == 0 {
			break
		}
		parts := strings.Fields(line)
		levels := make([]int, len(parts))
		for i, part := range parts {
			levels[i], _ = strconv.Atoi(part)
		}

		changeType := 0
		isSafe := true
		dampenerUsed := false
		for i := 1; i < len(levels); i++ {
			diff := levels[i] - levels[i-1]

			if diff == 0 || diff > 3 || diff < -3 {
				if !allowDampener || dampenerUsed {
					isSafe = false
					break
				}
				dampenerUsed = true
				continue
			}

			if diff > 0 {
				if changeType == -1 {
					if !allowDampener || dampenerUsed {
						isSafe = false
						break
					}
					dampenerUsed = true
				}
				changeType = 1
			} else {
				if changeType == 1 {
					if !allowDampener || dampenerUsed {
						isSafe = false
						break
					}
					dampenerUsed = true
				}
				changeType = -1
			}
		}
		if isSafe {
			safeReports++
		}
	}

	return
}

func (p Day02) PartA(lines []string) any {
	return getSafeReports(lines, false)
}

func (p Day02) PartB(lines []string) any {
	return getSafeReports(lines, true)
}
