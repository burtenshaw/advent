package main

import (
	"fmt"
	"os"

	"github.com/burtenshaw/advent/src/day1" // Replace with your actual module path
	"github.com/burtenshaw/advent/src/day2" // Replace with your actual module path
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Please specify a day to run (e.g., 'go run main.go 1')")
        return
    }

    day := os.Args[1]

    switch day {
    case "1":
        day1.Run() // Ensure there's a Run function in day1 package
    case "2":
        day2.Run() // Ensure there's a Run function in day2 package
    default:
        fmt.Println("Day not implemented")
    }
}
