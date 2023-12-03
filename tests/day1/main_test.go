package day1 // if you are doing black-box testing; use 'package day1' for white-box testing

import (
	"testing"

	"github.com/burtenshaw/advent/src/day1"
)


func TestCountLine(t *testing.T) {
    tests := []struct {
        name  string
        line  string
        want  int
    }{
        {
            name: "Test 1: Numeric digits",
            line: "3hello5world7",
            want: 37,
        },
        {
            name: "Test 2: Word digits",
            line: "onehellofiveworldseven",
            want: 17,
        },
        {
            name: "Test 3: Mixed digits",
            line: "onehello5world7",
            want: 17,
        },
        {
            name: "Test 4: No digits",
            line: "hello world",
            want: 0,
        },
        {
            name: "Test 5: overlapping digits",
            line: "mbkfgktwolbvsptgsixseven1oneightzvm",
            want: 28,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got,_,_ := day1.CountLine(tt.line); got != tt.want {
                t.Errorf("countLine() = %v, want %v", got, tt.want)
            }
        })
    }
}