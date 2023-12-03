package utils

import (
	"os"
	"strings"
)

func Reader(input string) string {
	content, err := os.ReadFile(input)

	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(string(content))
}