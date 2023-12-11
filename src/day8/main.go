package day8

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/burtenshaw/advent/src/utils"
)

type Node struct {
	Left  string
	Right string
}

func Run(inputPath string) {
	lines := utils.ReaderSplit(inputPath)
	network, instructions, startNodes := buildNetwork(lines)
	fmt.Println("Day 8")
	// FOR PART 1
	steps := 0
	startNode := "AAA"
	steps = findSteps(network, startNode, instructions, false)
	fmt.Println("Part 1", steps)
	// FOR PART 2
	stepsTaken := []int{}
	for _, currentNode := range startNodes {
		steps = findSteps(network, currentNode, instructions, true)
		stepsTaken = append(stepsTaken, steps)
	}
	steps = lcmN(stepsTaken)
	fmt.Println("Part 2", steps)
}

func buildNetwork(lines []string) (map[string]Node, []string, []string) {
	var instructions []string
	network := make(map[string]Node)
	startNodes := []string{}
	for i, line := range lines {
		if i == 0 {
			instructions = strings.Split(line, "")
			continue
		} else if line != "" {
			matches := reMatch("(\\w+)", line)
			n, l, r := getNodeVal(matches)
			network[n] = Node{l, r}
			if strings.HasSuffix(n, "A") {
				startNodes = append(startNodes, n)
			}
		}
	}
	return network, instructions, startNodes
}

func getNodeVal(matches []string) (string, string, string) {
	return matches[0], matches[1], matches[2]
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func lcmN(numbers []int) int {
	if len(numbers) == 2 {
		return lcm(numbers[0], numbers[1])
	}
	return lcm(numbers[0], lcmN(numbers[1:]))
}

func reMatch(pattern string, s string) []string {
	r, _ := regexp.Compile(pattern)
	return r.FindAllString(s, -1)
}

func findSteps(network map[string]Node, startNode string, instructions []string, isPart2 bool) int {
	steps := 0
	for {
		currentNode := startNode
		for _, nav := range instructions {
			steps++
			node := network[currentNode]
			if nav == "L" {
				currentNode = node.Left
			} else {
				currentNode = node.Right
			}
			if currentNode == "ZZZ" && !isPart2 {
				return steps
			}
			if strings.HasSuffix(currentNode, "Z") && isPart2 {
				return steps
			}
		}
		startNode = currentNode
	}
}
