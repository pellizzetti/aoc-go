package year2022

import (
	"regexp"
	"sort"
	"strconv"
	"unicode"
)

type Day05 struct{}

func (p Day05) PartA(lines []string) any {
	crateStacks, startOfInstructions := getCrateStacks(lines)

	for _, line := range lines[startOfInstructions:] {
		if line == "" {
			break
		}
		amount, from, to := getInstructions(line)

		fromStack := crateStacks[from]
		toStack := crateStacks[to]
		for i := 0; i < amount; i++ {
			n := len(fromStack) - 1
			toStack = append(toStack, fromStack[n])
			fromStack = fromStack[:n]
		}
		crateStacks[from] = fromStack
		crateStacks[to] = toStack
	}

	topCrates := getTopCrates(crateStacks)
	return topCrates
}

func (p Day05) PartB(lines []string) any {
	crateStacks, startOfInstructions := getCrateStacks(lines)

	for _, line := range lines[startOfInstructions:] {
		if line == "" {
			break
		}
		amount, from, to := getInstructions(line)

		fromStack := crateStacks[from]
		fromStackLen := len(fromStack)
		crateStacks[to] = append(crateStacks[to], fromStack[fromStackLen-amount:]...)
		crateStacks[from] = fromStack[:fromStackLen-amount]
	}

	topCrates := getTopCrates(crateStacks)
	return topCrates
}

func getCrateStacks(lines []string) (map[int][]rune, int) {
	var crateStacks = make(map[int][]rune)
	startOfInstructions := 0
	for i, line := range lines {
		if unicode.IsDigit(rune(line[1])) {
			startOfInstructions = i + 2
			break
		}
		j := 1
		for k := 1; k < len(line); k += 4 {
			if crate := rune(line[k]); !unicode.IsSpace(crate) {
				crateStacks[j] = append(crateStacks[j], crate)
			}
			j++
		}

	}
	for _, k := range crateStacks {
		sort.Slice(k, func(i, j int) bool {
			return i > j
		})
	}

	return crateStacks, startOfInstructions
}

func getInstructions(line string) (int, int, int) {
	re := regexp.MustCompile(`(?P<move>\d*)\sfrom\s(?P<from>\d*)\sto\s(?P<to>\d*)`)
	m := re.FindStringSubmatch(line)
	move, _ := strconv.Atoi(m[re.SubexpIndex("move")])
	from, _ := strconv.Atoi(m[re.SubexpIndex("from")])
	to, _ := strconv.Atoi(m[re.SubexpIndex("to")])
	return move, from, to
}

func getTopCrates(crateStacks map[int][]rune) string {
	keys := make([]int, 0)
	for k := range crateStacks {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	var topCrates string
	for _, k := range keys {
		stack := crateStacks[k]
		topCrates += string(stack[len(stack)-1])
	}
	return topCrates
}
