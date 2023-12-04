package day4

import (
	"testing"

	"github.com/burtenshaw/advent/src/day4"
)

func TestNewScratchcard(t *testing.T) {
	card := day4.NewScratchcard("41 48 83 86 17 | 83 86  6 31 17  9 48 53")
	if len(card.WinningNumbers) != 5 || len(card.YourNumbers) != 8 {
		t.Errorf("NewScratchcard did not parse the card correctly")
	}
}

func TestCountMatches(t *testing.T) {
	card := day4.NewScratchcard("41 48 83 86 17 | 83 86  6 31 17  9 48 53")
	matches := card.CountMatches()
	if matches != 4 {
		t.Errorf("Expected 4 matches, got %d", matches)
	}
}

func TestCalculateTotalPoints(t *testing.T) {
	cards := []day4.Scratchcard{
		day4.NewScratchcard("41 48 83 86 17 | 83 86  6 31 17  9 48 53"),
		day4.NewScratchcard("13 32 20 16 61 | 61 30 68 82 17 32 24 19"),
	}
	totalPoints := day4.CalculateTotalPoints(cards)
	if totalPoints != 10 {
		t.Errorf("Expected 10 total points, got %d", totalPoints)
	}
}
