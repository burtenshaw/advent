package main

import (
	"fmt"
	"os"

	"github.com/burtenshaw/advent/src/day1"
	"github.com/burtenshaw/advent/src/day2"
	"github.com/burtenshaw/advent/src/day3"
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
        day1.Run(inputData) // Ensure there's a Run function in day1 package
    case "2":
        day2.Run(inputData) // Ensure there's a Run function in day2 package
    case "3":
        day3.Run(inputData) // Ensure there's a Run function in day3 package
    default:
        fmt.Println("Day not implemented")
    }
}
