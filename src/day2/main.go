package day2

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/burtenshaw/advent/src/utils"
)

func parseGame(line string) (int, bool) {
    re := regexp.MustCompile(`Game (\d+): (.+)`)
    matches := re.FindStringSubmatch(line)
    if matches == nil {
        return 0, false
    }

    gameId, _ := strconv.Atoi(matches[1])
    sets := strings.Split(matches[2], ";")

    for _, set := range sets {
        red, green, blue := 0, 0, 0
        cubes := strings.Split(set, ",")
        for _, cube := range cubes {
            parts := strings.Fields(strings.TrimSpace(cube))
            count, _ := strconv.Atoi(parts[0])
            color := parts[1]

            switch color {
            case "red":
                red += count
            case "green":
                green += count
            case "blue":
                blue += count
            }
        }

        if red > 12 || green > 13 || blue > 14 {
            return 0, false
        }
    }

    return gameId, true
}

func parseGamePowers(line string) (int, int, int, int, bool) {
    re := regexp.MustCompile(`Game (\d+): (.+)`)
    matches := re.FindStringSubmatch(line)
    if matches == nil {
        return 0, 0, 0, 0, false
    }

    gameId, _ := strconv.Atoi(matches[1])
    sets := strings.Split(matches[2], ";")

    maxRed, maxGreen, maxBlue := 0, 0, 0
    for _, set := range sets {
        red, green, blue := 0, 0, 0
        cubes := strings.Split(set, ",")
        for _, cube := range cubes {
            parts := strings.Fields(strings.TrimSpace(cube))
            count, _ := strconv.Atoi(parts[0])
            color := parts[1]

            switch color {
            case "red":
                red += count
            case "green":
                green += count
            case "blue":
                blue += count
            }
        }

        if red > maxRed {
            maxRed = red
        }
        if green > maxGreen {
            maxGreen = green
        }
        if blue > maxBlue {
            maxBlue = blue
        }
    }

    return gameId, maxRed, maxGreen, maxBlue, true
}

func Run(inputPath string) {
	input := utils.Reader(inputPath)

    sum := 0
	sumOfPowers := 0

    for _, line := range strings.Split(input, "\n") {
        gameId, possible := parseGame(line)
        if possible {
            sum += gameId
        }
		gameId, maxRed, maxGreen, maxBlue, valid := parseGamePowers(line)
		if valid {
			power := maxRed * maxGreen * maxBlue
			sumOfPowers += power
		}

    }
	fmt.Println("Day 2")
    fmt.Println("Sum of IDs of possible games:", sum)
	fmt.Println("Sum of powers of possible games:", sumOfPowers)
}
