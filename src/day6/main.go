package day6

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/burtenshaw/advent/src/utils"
)

type Race struct {
	Time     int
	Distance int
}

func Run(inputPath string) {
	fmt.Println("Day 6 - Part 1")
	input := utils.Reader(inputPath)
	races := parseInput(input)

	result := 1
	for _, race := range races {
		ways := race.GetWaysToWin()
		result *= ways
	}
	fmt.Println("Total ways to win:", result)
	fmt.Println("Day 6 - Part 2")
	input = utils.Reader(inputPath)
	race := parseSingleRaceInput(input)
	ways := race.GetWaysToWinSingleRace()
	fmt.Println("Total ways to win:", ways)
}


func (r Race) GetWaysToWin() int {
	ways := 0
	for holdTime := 0; holdTime < r.Time; holdTime++ {
		speed := holdTime
		travelTime := r.Time - holdTime
		distance := speed * travelTime
		if distance > r.Distance {
			ways++
		}
	}
	return ways
}

func parseInput(input string) []Race {
	timesAndDistancesString := strings.Split(input, "\n")
	// parse out time
	timeString := strings.Split(timesAndDistancesString[0], ":")[1]
	times := utils.ParseIntList(timeString)
	// parse out distance
	distanceString := strings.Split(timesAndDistancesString[1], ":")[1]
	distances := utils.ParseIntList(distanceString)
	races := []Race{}
	for i, time := range times {
		races = append(races, Race{Time: time, Distance: distances[i]})
	}
	return races
}

func (r Race) GetWaysToWinSingleRace() int {
	ways := 0
	for holdTime := 0; holdTime < r.Time; holdTime++ {
		speed := holdTime
		travelTime := r.Time - holdTime
		distance := speed * travelTime
		if distance > r.Distance {
			ways++
		}
	}
	return ways
}

func parseSingleRaceInput(input string) Race {
	timesAndDistancesString := strings.Split(input, "\n")
	// parse out time
	timeString := strings.Split(timesAndDistancesString[0], ":")[1]
	timeString = strings.ReplaceAll(timeString, " ", "")
	time, _ := strconv.Atoi(timeString)
	// parse out distance
	distanceString := strings.Split(timesAndDistancesString[1], ":")[1]
	distanceString = strings.ReplaceAll(distanceString, " ", "")
	distance, _ := strconv.Atoi(distanceString)
	return Race{Time: time, Distance: distance}
}

