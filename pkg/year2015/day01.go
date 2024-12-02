package year2015

type Day01 struct{}

func (p Day01) PartA(lines []string) any {
	instructions := lines[0]
	floor := 0
	for i := 0; i < len(instructions); i++ {
		instruction := instructions[i]
		if instruction == '(' {
			floor += 1
		} else {
			floor -= 1
		}
	}
	return floor
}

func (p Day01) PartB(lines []string) any {
	instructions := lines[0]
	floor := 0
	for i := 0; i < len(instructions); i++ {
		instruction := instructions[i]
		if floor == -1 {
			return i
		}
		if instruction == '(' {
			floor += 1
		} else {
			floor -= 1
		}
	}
	return 0
}
