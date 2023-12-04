package day4

import (
	"fmt"
	"strings"

	"github.com/burtenshaw/advent/src/utils"
)

func Run(inputPath string) {
	inputData := utils.Reader(inputPath)
	cardStrings := strings.Split(inputData, "\n")

	var cards []Scratchcard
	for _, cardString := range cardStrings {
		cards = append(cards, NewScratchcard(cardString))
	}

	totalPoints := CalculateTotalPoints(cards)
	totalCards := ProcessCards(cards)

	fmt.Println("Day 4")
	fmt.Println("Part 1")
	fmt.Printf("Total points: %d\n", totalPoints)
	fmt.Println("Part 2")
	fmt.Printf("Total scratchcards: %d\n", totalCards)
}

type Scratchcard struct {
	WinningNumbers []string
	YourNumbers    []string
}

func NewScratchcard(card string) Scratchcard {
	split := strings.Split(card, "|")
	return Scratchcard{
		WinningNumbers: strings.Fields(split[0]),
		YourNumbers:    strings.Fields(split[1]),
	}
}

func (s Scratchcard) CountMatches() int {
	matches := 0
	for _, wn := range s.WinningNumbers {
		for _, yn := range s.YourNumbers {
			if wn == yn {
				matches++
			}
		}
	}
	return matches
}

func CalculateTotalPoints(cards []Scratchcard) int {
	totalPoints := 0
	for _, card := range cards {
		matches := card.CountMatches()
		if matches > 0 {
			totalPoints += 1 << (matches - 1)
		}
	}
	return totalPoints
}

func ProcessCards(cards []Scratchcard) int {
	totalCards := len(cards)
	cardsToProcess := make([]int, len(cards))

	// Initialize with representing the original cards
	for i := range cardsToProcess {
		cardsToProcess[i] = 1
	}
	for i := 0; i < len(cards); i++ {
		for cardsToProcess[i] > 0 {
			matches := cards[i].CountMatches()
			totalCards += matches * cardsToProcess[i]
			for j := 1; j <= matches; j++ {
				if i+j < len(cards) {
					cardsToProcess[i+j] += cardsToProcess[i]
				}
			}
			cardsToProcess[i] = 0
		}
	}
	return totalCards
}
