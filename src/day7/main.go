package day7

import (
	"fmt"
	"slices"
	"strings"

	"github.com/burtenshaw/advent/src/utils"
)

type Hand struct {
	Cards string
	Bid   int
}

func Run(inputPath string) {
	input := utils.ReaderSplit(inputPath)

	hands := []Hand{}
	for _, s := range input {
		h := Hand{}
		fmt.Sscanf(s, "%s %d", &h.Cards, &h.Bid)
		hands = append(hands, h)
	}

	winnings := func(jokers bool) (w int) {
		slices.SortFunc(hands, func(a, b Hand) int {
			return compareHands(a.Cards, b.Cards, jokers)
		})
		for i, h := range hands {
			w += (i + 1) * h.Bid
		}
		return
	}

	fmt.Println("Day 7 - Part 1")
	fmt.Println(winnings(false))
	fmt.Println("Day 7 - Part 2")
	fmt.Println(winnings(true))
}

func compareHands(a, b string, jokers bool) int {
    j, r := "J", "TAJBQCKDAE"
    if jokers {
        j, r = "23456789TQKA", "TAJ0QCKDAE"
    }

    handType := func(cards string) string {
        maxCount := 0
        for _, j := range strings.Split(j, "") {
            replaced, count := strings.ReplaceAll(cards, "J", j), 0
            for _, s := range replaced {
                count += strings.Count(replaced, string(s))
            }
            maxCount = utils.Max(maxCount, count)
        }
        return map[int]string{5: "0", 7: "1", 9: "2", 11: "3", 13: "4", 17: "5", 25: "6"}[maxCount]
    }

    replaceWithOrder := func(cards string, order string) string {
        return strings.NewReplacer(strings.Split(order, "")...).Replace(cards)
    }

    return strings.Compare(
        handType(a)+replaceWithOrder(a, r),
        handType(b)+replaceWithOrder(b, r),
    )
}

