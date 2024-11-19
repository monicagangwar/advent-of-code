package _4

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
		expectedPartOne int64
		expectedPartTwo int64
	}

	tests := []test{
		{
			name:            "with sample",
			input:           sample,
			expectedPartOne: 288,
			expectedPartTwo: 71503,
		}, {
			name:            "with large input",
			input:           input,
			expectedPartOne: ***REMOVED***,
			expectedPartTwo: ***REMOVED***,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			times, distances := parseInput(tst.input)
			singleTime, singleDistance := parseInputAsSingleNum(tst.input)
			if tst.expectedPartOne != -1 {
				if got := getNumberOfChances(times, distances); got != tst.expectedPartOne {
					t.Errorf("partOne() = %v, want %v", got, tst.expectedPartOne)
				}
			}

			if tst.expectedPartTwo != -1 {
				if got := getNumberOfChances(singleTime, singleDistance); got != tst.expectedPartTwo {
					t.Errorf("partTwo() = %v, want %v", got, tst.expectedPartTwo)
				}
			}
		})

	}

}

func parseInput(input string) ([]int64, []int64) {
	lines := strings.Split(input, "\n")

	times := convertStrListToInt(strings.Replace(lines[0], "Time:", "", 1))
	distances := convertStrListToInt(strings.Replace(lines[1], "Distances:", "", 1))

	return times, distances
}

func parseInputAsSingleNum(input string) ([]int64, []int64) {
	lines := strings.Split(input, "\n")

	times := convertStrListToInt(strings.Replace(strings.Replace(lines[0], "Time:", "", 1), " ", "", -1))
	distances := convertStrListToInt(strings.Replace(strings.Replace(lines[1], "Distance:", "", 1), " ", "", -1))

	return times, distances
}

func convertStrListToInt(strList string) []int64 {
	nums := make([]int64, 0)
	for _, numStr := range strings.Split(strList, " ") {
		num, err := strconv.ParseInt(strings.TrimSpace(numStr), 10, 64)
		if err == nil {
			nums = append(nums, num)
		}

	}
	return nums
}

func getNumberOfChances(times, distances []int64) int64 {
	numWaysToBreakRecordTotal := int64(1)

	for i := 0; i < len(distances); i++ {
		numWaysToBreakRecord := int64(0)
		for speed := int64(1); speed < times[i]; speed++ {
			timeRemaining := times[i] - speed
			if speed*timeRemaining > distances[i] {
				numWaysToBreakRecord++
			}
		}
		numWaysToBreakRecordTotal *= numWaysToBreakRecord
	}
	return numWaysToBreakRecordTotal
}
