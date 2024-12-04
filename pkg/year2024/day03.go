package year2024

import (
	"strconv"
)

type Day03 struct{}

func parseMuls(lines []string, useInstructions bool) (result int) {
	do := true
	for _, line := range lines {
		if len(line) == 0 {
			break
		}
		for i := 0; i <= len(line); i++ {
			if useInstructions && i+7 <= len(line) && line[i:i+7] == "don't()" {
				do = false
			}
			if useInstructions && i+4 <= len(line) && line[i:i+4] == "do()" {
				do = true
			}
			if do && i+4 <= len(line) && line[i:i+4] == "mul(" {
				c := i + 4
				var firstValue, secondValue int
				isCorrectFormat := true
				for {
					_, err := strconv.Atoi(string(line[c]))
					if err != nil {
						if line[c] != ',' {
							isCorrectFormat = false
						}
						firstValue, _ = strconv.Atoi(line[i+4 : c])
						c++
						break
					}
					c++
				}
				if isCorrectFormat {
					cc := c
					for {
						_, err := strconv.Atoi(string(line[c]))
						if err != nil {
							if line[c] != ')' {
								isCorrectFormat = false
							}
							secondValue, _ = strconv.Atoi(line[cc:c])
							break
						}
						c++
					}
				}

				if isCorrectFormat {
					result += firstValue * secondValue
				}
			}
		}
	}
	return
}

func (p Day03) PartA(lines []string) any {
	return parseMuls(lines, false)
}

func (p Day03) PartB(lines []string) any {
	return parseMuls(lines, true)
}
