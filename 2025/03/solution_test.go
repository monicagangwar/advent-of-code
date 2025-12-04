package template

import (
	_ "embed"
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
			expectedPartOne: 357,
			expectedPartTwo: 3121910778619,
		}, {
			name:            "with large input",
			input:           input,
			expectedPartOne: 0,
			expectedPartTwo: 0,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			lines := parseInput(tst.input)
			if tst.expectedPartOne != -1 {
				if got := partOne(lines); got != tst.expectedPartOne {
					t.Errorf("partOne() = %v, want %v", got, tst.expectedPartOne)
				}
			}

			if tst.expectedPartTwo != -1 {
				if got := partTwo(lines); got != tst.expectedPartTwo {
					t.Errorf("partTwo() = %v, want %v", got, tst.expectedPartTwo)
				}
			}
		})

	}

}

func parseInput(input string) []string {
	return strings.Split(input, "\n")

}

func partOne(lines []string) int {
	sumMaxNum := 0
	for _, line := range lines {
		maxNum := 0
		for i := 0; i < len(line); i++ {
			numI, _ := strconv.Atoi(string(line[i]))
			for j := i + 1; j < len(line); j++ {
				numJ, _ := strconv.Atoi(string(line[j]))
				num := numI*10 + numJ
				if num > maxNum {
					maxNum = num
				}
			}
		}
		//fmt.Printf("\n %s: %d", line, maxNum)
		sumMaxNum += maxNum

	}

	return sumMaxNum
}

func partTwo(lines []string) int {
	sumMaxNum := 0
	for _, line := range lines {
		maxNums := make([]int, 0)
		numSelectedIdx := -1
		for i := 0; i < 12; i++ {
			maxNumIdx := 0
			maxNum := 0
			digitsRequired := 12 - i - 1
			for j := numSelectedIdx + 1; j < len(line); j++ {
				num, _ := strconv.Atoi(string(line[j]))
				digitsRemaining := len(line) - 1 - j
				if num > maxNum && digitsRemaining >= digitsRequired {
					maxNum = num
					maxNumIdx = j
				}
			}
			maxNums = append(maxNums, maxNum)
			numSelectedIdx = maxNumIdx
		}
		maxNumInt := 0
		for i := 0; i < len(maxNums); i++ {
			if maxNumInt == 0 {
				maxNumInt = maxNums[i]
			} else {
				maxNumInt = maxNumInt*10 + maxNums[i]
			}
		}
		sumMaxNum += maxNumInt
		//fmt.Printf("\n%s => %d\n", line, maxNumInt)
	}
	return sumMaxNum
}
