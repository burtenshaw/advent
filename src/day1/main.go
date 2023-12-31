package day1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/burtenshaw/advent/src/utils"
)

// mapWordToDigit maps spelled-out numbers to digits.
func mapWordToDigit(word string) string {
	numberMap := map[string]string{
		"one": "1", "two": "2", "three": "3", "four": "4",
		"five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9",
	}
	if digit, exists := numberMap[word]; exists {
		return digit
	}
	return word
}

// CountLine calculates the calibration value from a line of text.
func CountLine(line string) int {
	words := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	firstDigit := ""
	lastDigit := ""

	// Extract the first and last digit/number-word from the line
	for i := 0; i < len(line); i++ {
		// Check for spelled-out numbers
		for _, word := range words {
			if strings.HasPrefix(line[i:], word) {
				digit := mapWordToDigit(word)
				if firstDigit == "" {
					firstDigit = digit
					// Continue searching for the last digit
				}
				lastDigit = digit
				// Fuck life
				// i += len(word) - 1 // Skip the length of the word
				break
			}

		}

		// Check for digits
		if char := line[i]; char >= '0' && char <= '9' {
			if firstDigit == "" {
				firstDigit = string(char)
			}
			lastDigit = string(char)
		}
	}

	// Handle case where no valid digit/number-word is found
	if firstDigit == "" || lastDigit == "" {
		return 0
	}

	combined := firstDigit + lastDigit
	number, err := strconv.Atoi(combined)
	if err != nil {
		fmt.Println("Error converting to number:", combined)
		return 0
	}

	return number
}

func Run(inputPath string) {

	input := utils.Reader(inputPath)
	totalSum := 0

	for _, line := range strings.Split(input, "\n") {
		number := CountLine(line) // Modify CountLine to return digits too
		totalSum += number
	}

	fmt.Println("Total sum:", totalSum)
}
