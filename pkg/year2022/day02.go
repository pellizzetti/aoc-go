package year2022

type Day02 struct{}

const (
	Rock = iota + 1
	Paper
	Scissors
)

const (
	Defeat = iota
	Draw   = iota * 3
	Victory
)

var (
	handShapeMap = map[rune]int{
		'A': Rock,
		'B': Paper,
		'C': Scissors,
		'X': Rock,
		'Y': Paper,
		'Z': Scissors,
	}
	guessedStrategyMapp = map[int]map[int]int{
		Rock: {
			Scissors: Defeat,
			Paper:    Victory,
		},
		Paper: {
			Rock:     Defeat,
			Scissors: Victory,
		},
		Scissors: {
			Rock:  Victory,
			Paper: Defeat,
		},
	}

	roundOutcomeMap = map[rune]int{
		'X': Defeat,
		'Y': Draw,
		'Z': Victory,
	}
	explainedStrategyMap = map[int]map[int]int{
		Rock: {
			Defeat:  Scissors,
			Victory: Paper,
		},
		Paper: {
			Defeat:  Rock,
			Victory: Scissors,
		},
		Scissors: {
			Defeat:  Paper,
			Victory: Rock,
		},
	}
)

func (p Day02) PartA(lines []string) any {
	score := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		opponentHandShape := handShapeMap[[]rune(line)[0]]
		myHandShape := handShapeMap[[]rune(line)[2]]
		roundOutcome := guessedStrategyMapp[opponentHandShape][myHandShape]
		if opponentHandShape == myHandShape {
			score += Draw + myHandShape
		} else {
			score += roundOutcome + myHandShape
		}
	}
	return score
}

func (p Day02) PartB(lines []string) any {
	score := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		opponentHandShape := handShapeMap[[]rune(line)[0]]
		roundOutcome := roundOutcomeMap[[]rune(line)[2]]
		strategyHandShape := explainedStrategyMap[opponentHandShape]
		if roundOutcome == Draw {
			score += Draw + opponentHandShape
		} else if roundOutcome == Victory {
			score += Victory + strategyHandShape[Victory]
		} else {
			score += Defeat + strategyHandShape[Defeat]
		}
	}
	return score
}
