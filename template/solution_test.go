package template

import (
	_ "embed"
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
			expectedPartOne: 0,
			expectedPartTwo: -1,
		}, {
			name:            "with large input",
			input:           input,
			expectedPartOne: 0,
			expectedPartTwo: -1,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			if tst.expectedPartOne != -1 {
				if got := partOne(parseInput(tst.input)); got != tst.expectedPartOne {
					t.Errorf("partOne() = %v, want %v", got, tst.expectedPartOne)
				}
			}

			if tst.expectedPartTwo != -1 {
				if got := partTwo(parseInput(tst.input)); got != tst.expectedPartTwo {
					t.Errorf("partTwo() = %v, want %v", got, tst.expectedPartTwo)
				}
			}
		})

	}

}

func parseInput(input string) string {
	return ""
}

func partOne(input string) int {
	return 0
}

func partTwo(input string) int {
	return 0
}
