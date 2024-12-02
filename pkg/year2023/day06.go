package year2023

import (
	"strconv"
	"strings"
)

type Day06 struct{}

type Race struct {
	Time            int
	Distance        int
	PossibleRecords int
}

var testa = []string{
	"Time:      7  15   30",
	"Distance:  9  40  200",
	"",
}

const boatSpeed = 1

func (p Day06) PartA(lines []string) any {
	races := getRaces(lines)
	for i, race := range races {
		for holding := 0; holding < race.Time; holding++ {
			totalDistance := holding * boatSpeed * (race.Time - holding)
			if totalDistance > race.Distance {
				races[i].PossibleRecords += 1
			}
		}
	}
	records := 1
	for _, race := range races {
		println(race.PossibleRecords)
		records *= race.PossibleRecords
	}
	return records
}

func (p Day06) PartB(lines []string) any {
	race := getRace(lines)
	for holding := 0; holding < race.Time; holding++ {
		totalDistance := holding * boatSpeed * (race.Time - holding)
		if totalDistance > race.Distance {
			race.PossibleRecords += 1
		}
	}
	return race.PossibleRecords
}

func getRaces(lines []string) []Race {
	races := make([]Race, 0)
	times := strings.Fields(lines[0])
	for _, time := range times {
		t, err := strconv.Atoi(time)
		if err != nil {
			continue
		}
		races = append(races, Race{Time: t})
	}
	distances := strings.Fields(lines[1])
	raceDistances := make([]int, 0)
	for _, distance := range distances {
		d, err := strconv.Atoi(distance)
		if err != nil {
			continue
		}
		raceDistances = append(raceDistances, d)
	}

	for i := range races {
		races[i].Distance = raceDistances[i]
		races[i].PossibleRecords = 0
	}

	return races
}

func getRace(lines []string) Race {
	timeStr := strings.TrimPrefix(lines[0], "Time:")
	distanceStr := strings.TrimPrefix(lines[1], "Distance:")
	time, _ := strconv.Atoi(strings.ReplaceAll(timeStr, " ", ""))
	distance, _ := strconv.Atoi(strings.ReplaceAll(distanceStr, " ", ""))
	return Race{
		Time:     time,
		Distance: distance,
	}
}
