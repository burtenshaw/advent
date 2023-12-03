package day3

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/burtenshaw/advent/src/utils"
)

func Run(input string) {
	fmt.Println("Day 3")
	fmt.Println("first part:")
	inputText := utils.Reader(input)
	schematic := NewSchematic(inputText)
	schematicNumbers, gearCandidates := schematic.GetSchematicParts()
	partNumbersSum := sumPartNumbers(schematicNumbers)
	fmt.Println(partNumbersSum)
	fmt.Println("second part:")
	fmt.Printf("%d\n", gearCandidates.SumAllGearRatios())
}

type SchematicNumber struct {
	Value           int
	AdjacentSymbols map[string]bool
}

type SchematicDimensions struct {
	Width  int
	Length int
}

type Schematic struct {
	Contents   string
	Dimensions SchematicDimensions
}

func NewSchematic(input string) (schematic Schematic) {
	schematic.Contents = input
	for i, c := range input {
		if c == rune('\n') {
			schematic.Dimensions.Width = i
			schematic.Dimensions.Length = len(input) / schematic.Dimensions.Width
			break
		}
	}
	return
}

func (s Schematic) GetSymbol(x, y int) (rune, error) {
    if x >= s.Dimensions.Width || y >= s.Dimensions.Length || x < 0 || y < 0 {
        return 0, errors.New("coordinates out of bounds")
    }
    symbolIndex := x + s.Dimensions.Width*y + y
    return rune(s.Contents[symbolIndex]), nil
}

func NewSchematicNumber(value int) (s SchematicNumber) {
	s.Value = value
	s.AdjacentSymbols = make(map[string]bool)
	return
}


// GEAR CANDIDATES

type Coordinates struct {
	x int
	y int
}

// gear candidate is a list of potential part number values, which will later be used to calculate ratio
type GearCandidate []int

type GearCandidates map[Coordinates]GearCandidate

func (g GearCandidate) IsGear() bool {
	return len(g) == 2
}

func (g GearCandidate) GetRatio() (ratio int) {
	ratio = 1

	for _, value := range g {
		ratio = ratio * value
	}

	return
}

func (g GearCandidates) SumAllGearRatios() (sumRatios int) {
	for _, gearCandidate := range g {
		if gearCandidate.IsGear() {
			sumRatios += gearCandidate.GetRatio()
		}
	}
	return
}

func (s Schematic) GetSchematicParts() ([]SchematicNumber, GearCandidates) {
    numbers := []SchematicNumber{}
    gearCandidates := GearCandidates{}

    x, y := 0, 0
    incrementCoordinates := func(c rune) {
        switch c {
        case rune('\n'):
            x, y = 0, y+1
        default:
            x++
        }
    }

    currentValue := 0
    numberLength := 0

    addSymbolIfPossible := func(number *SchematicNumber, cx, cy int) {
        if symbol, err := s.GetSymbol(cx, cy); err == nil {
            number.AdjacentSymbols[string(symbol)] = true
            if symbol == rune('*') {
                gearCandidates[Coordinates{x: cx, y: cy}] = append(gearCandidates[Coordinates{x: cx, y: cy}], currentValue)
            }
        }
    }

    newSchematicNumberWithAdjacentSymbols := func() SchematicNumber {
        number := NewSchematicNumber(currentValue)
        addSymbolIfPossible(&number, x, y)                // symbol in front of number
        addSymbolIfPossible(&number, x-numberLength-1, y) // symbol behind number
        for i := x - numberLength - 1; i <= x; i++ {
            addSymbolIfPossible(&number, i, y-1) // symbols in the row on top of number
            addSymbolIfPossible(&number, i, y+1) // symbols in the row under the number
        }
        return number
    }

    for _, char := range s.Contents {
        if digit, err := strconv.Atoi(string(char)); err == nil {
            currentValue = currentValue*10 + digit
            numberLength++
        } else if currentValue > 0 {
            currentNumber := newSchematicNumberWithAdjacentSymbols()
            numbers = append(numbers, currentNumber)
            currentValue = 0
            numberLength = 0
        }
        incrementCoordinates(char)
    }
    return numbers, gearCandidates
}

func (n SchematicNumber) IsPartNumber() bool {
	for symbol := range n.AdjacentSymbols {
		if symbol != "." {
			return true
		}
	}
	return false
}

func sumPartNumbers(schematicNumbers []SchematicNumber) (partNumbersSum int) {
	for _, schematicNumber := range schematicNumbers {
		if schematicNumber.IsPartNumber() {
			partNumbersSum += schematicNumber.Value
		}
	}
	return
}

