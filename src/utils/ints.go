package utils

import (
	"fmt"
	"strconv"
	"strings"
)


func MustParseInt64(s string) int64 {
	val, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("unable to parse string to int64: %s", s))
	}
	return val
}

func MustStringToInt64Slice(input string) []int64 {
	fields := strings.Fields(input)
	vals := make([]int64, len(fields))
	for i, field := range fields {
		vals[i] = MustParseInt64(field)
	}
	return vals
}


// takes a string which is a list of integers separated by spaces and returns the corresponding array of ints
func ParseIntList(input string) (output []int) {
	maybeInts := strings.Split(input, " ")
	for _, maybeInt := range maybeInts {
		if intToAdd, err := strconv.Atoi(maybeInt); err == nil {
			output = append(output, intToAdd)
		}
	}
	return
}

func Max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func Min(a, b int) int {	
	if a < b {
		return a
	}
	return b
}