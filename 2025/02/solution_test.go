package template

import (
	_ "embed"
	"fmt"
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
			expectedPartOne: 1227775554,
			expectedPartTwo: 4174379265,
		}, {
			name:            "with large input",
			input:           input,
			expectedPartOne: 0,
			expectedPartTwo: 0,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			productIDList := parseInput(tst.input)
			if tst.expectedPartOne != -1 {
				if got := partOne(productIDList); got != tst.expectedPartOne {
					t.Errorf("partOne() = %v, want %v", got, tst.expectedPartOne)
				}
			}

			if tst.expectedPartTwo != -1 {
				if got := partTwo(productIDList); got != tst.expectedPartTwo {
					t.Errorf("partTwo() = %v, want %v", got, tst.expectedPartTwo)
				}
			}
		})

	}

}

func parseInput(input string) [][2]int {
	lines := strings.Split(input, ",")
	productIDList := make([][2]int, len(lines))
	for i, line := range lines {
		productIDsStr := strings.Split(line, "-")
		productIDList[i][0], _ = strconv.Atoi(productIDsStr[0])
		productIDList[i][1], _ = strconv.Atoi(productIDsStr[1])
	}
	return productIDList
}

func partOne(productIDList [][2]int) int {
	sumInvalidIds := 0
	for _, productIDs := range productIDList {
		for i := productIDs[0]; i <= productIDs[1]; i++ {
			if hasOnlyRepeats(fmt.Sprintf("%d", i)) {
				sumInvalidIds += i
			}
		}
	}
	return sumInvalidIds
}

func hasOnlyRepeats(num string) bool {
	numLen := len(num)
	if numLen%2 != 0 {
		return false
	}
	numLeft := num[:numLen/2]
	numRight := num[numLen/2:]
	return numLeft == numRight
}

func partTwo(productIDList [][2]int) int {
	sumInvalidIds := 0
	for _, productIDs := range productIDList {
		for i := productIDs[0]; i <= productIDs[1]; i++ {
			if hasRepeatedPattern(fmt.Sprintf("%d", i), 1) {
				sumInvalidIds += i
			}
		}
	}
	return sumInvalidIds
}

func hasRepeatedPattern(num string, requiredRepeats int) bool {
	for i := 1; i < len(num); i++ {
		repeatedCount := 0
		pattern := num[:i]
		for j := i; j < len(num); j += i {
			if j+len(pattern) > len(num) {
				repeatedCount = 0
				break
			}
			subNum := num[j : j+len(pattern)]
			if subNum == pattern {
				repeatedCount++
			} else {
				repeatedCount = 0
				break
			}
		}
		if repeatedCount >= requiredRepeats {
			//fmt.Printf("\n num: %s, pattern: %s", num, pattern)
			return true
		}
	}
	return false
}
