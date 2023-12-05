package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Reader(input string) string {
	content, err := os.ReadFile(input)

	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(string(content))
}

func ReaderSplit(filePath string) []string {
	fileContent := Reader(filePath)
	return strings.Split(fileContent, "\n")
}

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