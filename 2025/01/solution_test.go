package template

import (
	_ "embed"
	"math"
	"strconv"
	"strings"
	"testing"
)

//go:embed input.txt
var input string

//go:embed sample.txt
var sample string

func TestSolution(t *testing.T) {
	type test struct {
		name            string
		input           string
		expectedPartOne int
		expectedPartTwo int
	}

	tests := []test{
		{
			name:            "with sample",
			input:           sample,
			expectedPartOne: 3,
			expectedPartTwo: 6,
		}, {
			name:            "with large input",
			input:           input,
			expectedPartOne: 0,
			expectedPartTwo: 0,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			offsets := parseInput(tst.input)
			if tst.expectedPartOne != -1 {
				if got := partOne(offsets); got != tst.expectedPartOne {
					t.Errorf("partOne() = %v, want %v", got, tst.expectedPartOne)
				}
			}

			if tst.expectedPartTwo != -1 {
				if got := partTwo(offsets); got != tst.expectedPartTwo {
					t.Errorf("partTwo() = %v, want %v", got, tst.expectedPartTwo)
				}
			}
		})

	}

}

func parseInput(input string) []int {
	lines := strings.Split(input, "\n")
	offsets := make([]int, len(lines))
	for i, line := range lines {
		num, _ := strconv.Atoi(line[1:])
		if strings.HasPrefix(line, "L") {
			offsets[i] = num * -1
		} else {
			offsets[i] = num
		}
	}
	return offsets
}

func partOne(offsets []int) int {
	nextDigit := 50
	countZero := 0
	for _, offset := range offsets {
		nextDigit += offset
		nextDigit = ((nextDigit % 100) + 100) % 100
		if nextDigit == 0 {
			countZero++
		}
	}
	return countZero
}

func partTwo(offsets []int) int {
	pos := 50
	countZero := 0
	for _, offset := range offsets {
		next := pos + offset
		if next <= 0 && pos != 0 {
			countZero++
		}
		countZero += int(math.Abs(float64(next))) / 100
		pos = ((next % 100) + 100) % 100
	}
	return countZero
}
