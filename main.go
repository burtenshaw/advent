package main

import (
	"fmt"
	"os"

	"github.com/burtenshaw/advent/src/day1"
	"github.com/burtenshaw/advent/src/day2"
	"github.com/burtenshaw/advent/src/day3"
	"github.com/burtenshaw/advent/src/day4"
	"github.com/burtenshaw/advent/src/day5"
	"github.com/burtenshaw/advent/src/day6"
	"github.com/burtenshaw/advent/src/day7"
	"github.com/burtenshaw/advent/src/day8"
	"github.com/burtenshaw/advent/src/day9"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please specify a day to run (e.g., 'go run main.go 1')")
		return
	}

	day := os.Args[1]

	inputData := "data/day" + day + "/input.txt"

	switch day {
	case "1":
		day1.Run(inputData)
	case "2":
		day2.Run(inputData)
	case "3":
		day3.Run(inputData)
	case "4":
		day4.Run(inputData)
	case "5":
		day5.Run(inputData)
	case "6":
		day6.Run(inputData)
	case "7":
		day7.Run(inputData)
	case "8":
		day8.Run(inputData)
	case "9":
		day9.Run(inputData)
	default:
		fmt.Println("Day not implemented")
	}
}
