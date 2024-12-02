package year2023

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Day05 struct{}

type Rule struct {
	Destination int64
	Source      int64
	Range       int64
}

type Range struct {
	From int64
	To   int64
}

var test = []string{
	"seeds: 79 14 55 13",
	"",
	"seed-to-soil map:",
	"50 98 2",
	"52 50 48",
	"",
	"soil-to-fertilizer map:",
	"0 15 37",
	"37 52 2",
	"39 0 15",
	"",
	"fertilizer-to-water map:",
	"49 53 8",
	"0 11 42",
	"42 0 7",
	"57 7 4",
	"",
	"water-to-light map:",
	"88 18 7",
	"18 25 70",
	"",
	"light-to-temperature map:",
	"45 77 23",
	"81 45 19",
	"68 64 13",
	"",
	"temperature-to-humidity map:",
	"0 69 1",
	"1 0 69",
	"",
	"humidity-to-location map:",
	"60 56 37",
	"56 93 4",
	"",
}

func (p Day05) PartA(lines []string) any {
	return 0
	seeds := parseSeeds(strings.TrimPrefix(lines[0], "seeds: "))
	maps := parseMaps(lines[1:])
	for _, m := range maps {
		fmt.Printf("%v\n", m)
	}

	var lowestLocation int64 = math.MaxInt64
	for _, seed := range seeds {
		curr := seed
		fmt.Printf("Seed %d\n", seed)
		for i, m := range maps {
			switch i {
			case 0:
				fmt.Println("Soil")
			case 1:
				fmt.Println("Fertilizer")
			case 2:
				fmt.Println("Water")
			case 3:
				fmt.Println("Light")
			case 4:
				fmt.Println("Temperature")
			case 5:
				fmt.Println("Humidity")
			case 6:
				fmt.Println("Location")
			}
			fmt.Printf("%v\n", m)
			for _, rule := range m {
				fmt.Printf("rule %v\n", rule)
				fmt.Printf("%d >= %d (%t)\n", curr, rule.Source, curr >= rule.Source)
				fmt.Printf("%d <= %d (%t)\n", curr, rule.Source+rule.Range-1, curr <= rule.Source+rule.Range-1)
				if curr >= rule.Source && curr <= rule.Source+rule.Range-1 {
					// Apply the rule
					offset := curr - rule.Source

					curr = rule.Destination + offset
					fmt.Printf("%d = %d + %d\n", curr, rule.Destination, offset)

					// Move to the next map
					break
				}
			}
		}

		if curr < lowestLocation {
			lowestLocation = curr
		}
	}

	// Print the result
	return lowestLocation
}

func (p Day05) PartB(lines []string) any {
	lines = test
	seeds := strings.TrimPrefix(lines[0], "seeds: ")
	seedsList := make([]int64, 0)
	for _, s := range strings.Fields(seeds) {
		val, _ := strconv.ParseInt(s, 10, 64)
		seedsList = append(seedsList, val)
	}

	fmt.Printf("%v\n", seedsList)
	println(len(seedsList))

	seedRanges := make([]Range, 0)
	for i := 0; i < len(seedsList); i += 2 {
		seed, seedRange := seedsList[i], seedsList[i+1]
		seedRanges = append(seedRanges, Range{From: seed, To: seed + seedRange})
	}

	fmt.Printf("%v\n", seedRanges)
	println(len(seedRanges))

	maps := make([][]Rule, 0)
	// mapBlocks := strings.Split(mapsStr, "\n\n")
	// fmt.Printf("lines %v\n", lines[1:])

	for _, block := range lines[2:] {
		fmt.Printf("block %s\n", block)

		lines := strings.Split(block, "\n")
		fmt.Printf("lines %v\n", lines)
		ruleList := make([]Rule, 0)
		for _, line := range lines {
			nums := strings.Fields(line)
			dest, _ := strconv.ParseInt(nums[0], 10, 64)
			src, _ := strconv.ParseInt(nums[1], 10, 64)
			rng, _ := strconv.ParseInt(nums[2], 10, 64)
			ruleList = append(ruleList, Rule{Destination: dest, Source: src, Range: rng})
		}
		fmt.Printf("%v\n", ruleList)
		sort.Slice(ruleList, func(i, j int) bool {
			return ruleList[i].Source < ruleList[j].Source
		})
		maps = append(maps, ruleList)
	}

	currRanges := make([]Range, len(seedRanges))
	copy(currRanges, seedRanges)

	for _, m := range maps {
		newRanges := make([]Range, 0)

		for _, r := range currRanges {
			curr := r

			for _, rule := range m {
				offset := rule.Destination - rule.Source
				ruleApplies := curr.From <= curr.To && curr.From <= rule.Source+rule.Range && curr.To >= rule.Source

				if ruleApplies {
					if curr.From < rule.Source {
						newRanges = append(newRanges, Range{From: curr.From, To: rule.Source - 1})
						curr.From = rule.Source
						if curr.To < rule.Source+rule.Range {
							newRanges = append(newRanges, Range{From: curr.From + offset, To: curr.To + offset})
							curr.From = curr.To + 1
						} else {
							newRanges = append(newRanges, Range{From: curr.From + offset, To: rule.Source + rule.Range - 1 + offset})
							curr.From = rule.Source + rule.Range
						}
					} else if curr.To < rule.Source+rule.Range {
						newRanges = append(newRanges, Range{From: curr.From + offset, To: curr.To + offset})
						curr.From = curr.To + 1
					} else {
						newRanges = append(newRanges, Range{From: curr.From + offset, To: rule.Source + rule.Range - 1 + offset})
						curr.From = rule.Source + rule.Range
					}
				}
			}
			if curr.From <= curr.To {
				newRanges = append(newRanges, curr)
			}
		}
		currRanges = newRanges
	}

	lowestLocation := currRanges[0].From
	for _, r := range currRanges {
		if r.From < lowestLocation {
			lowestLocation = r.From
		}
	}

	return lowestLocation
}

func parseSeeds(line string) []int64 {
	seeds := make([]int64, 0)
	for _, s := range strings.Fields(line) {
		seed, err := strconv.ParseInt(s, 10, 64)
		if err == nil {
			seeds = append(seeds, seed)
		}
	}
	return seeds
}

func parseMaps(lines []string) [][]Rule {
	maps := make([][]Rule, 0)
	mapRules := make([]Rule, 0)
	for _, line := range lines {
		rulesLine := strings.Split(line, " ")
		if len(rulesLine) != 3 {
			fmt.Printf("%v (%d) mapRules = %v\n", rulesLine, len(rulesLine), mapRules)
			if len(mapRules) > 0 {
				maps = append(maps, mapRules)
				mapRules = make([]Rule, 0)
			}
			continue
		}

		rulesNumbers := make([]int64, 3)
		for i := 0; i < 3; i++ {
			rulesNumbers[i], _ = strconv.ParseInt(rulesLine[i], 10, 64)
		}
		mapRules = append(mapRules, Rule{
			Destination: rulesNumbers[0],
			Source:      rulesNumbers[1],
			Range:       rulesNumbers[2],
		})
	}
	return maps
}
