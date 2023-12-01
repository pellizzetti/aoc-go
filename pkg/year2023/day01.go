package year2023

import (
	"bytes"
)

type Day01 struct{}

func (p Day01) PartA(lines []string) any {
	sum := 0
	for i := 0; i < len(lines); i++ {
		line := []byte(lines[i])
		if len(line) == 0 {
			continue
		}
		digits, _ := getDigits(line)
		sum += digits[0]*10 + digits[len(digits)-1]
	}
	return sum
}

func (p Day01) PartB(lines []string) any {
	sum := 0
	for i := 0; i < len(lines); i++ {
		line := []byte(lines[i])
		if len(line) == 0 {
			continue
		}
		_, digits := getDigits(line)
		sum += digits[0]*10 + digits[len(digits)-1]
	}
	return sum
}

func getDigits(line []byte) ([]int, []int) {
	digits := make([]int, 0)
	digitsWithLonghand := make([]int, 0)
	digitsWords := [][]byte{
		[]byte("one"),
		[]byte("two"),
		[]byte("three"),
		[]byte("four"),
		[]byte("five"),
		[]byte("six"),
		[]byte("seven"),
		[]byte("eight"),
		[]byte("nine"),
	}
	digitsMinLength := 3

	for len(line) > 0 {
		// unicode 48 == 0
		if line[0] >= 48 && line[0] < 58 {
			digit := int(line[0]) - 48
			digits = append(digits, digit)
			digitsWithLonghand = append(digitsWithLonghand, digit)
		} else {
			for i, word := range digitsWords {
				if bytes.HasPrefix(line, word) {
					digitsWithLonghand = append(digitsWithLonghand, int(i)+1)

					// check for re-use. e.g. `eightwo`
					line = line[len(word)-digitsMinLength:]
					break
				}
			}
		}
		line = line[1:]
	}
	return digits, digitsWithLonghand
}
