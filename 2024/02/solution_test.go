package _2

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
			expectedPartOne: 2,
			expectedPartTwo: 4,
		}, {
			name:            "with large input",
			input:           input,
			expectedPartOne: 0,
			expectedPartTwo: 0,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			reports := parseInput(tst.input)
			if tst.expectedPartOne != -1 {
				if got := partOne(reports); got != tst.expectedPartOne {
					t.Errorf("partOne() = %v, want %v", got, tst.expectedPartOne)
				}
			}

			if tst.expectedPartTwo != -1 {
				if got := partTwo(reports); got != tst.expectedPartTwo {
					t.Errorf("partTwo() = %v, want %v", got, tst.expectedPartTwo)
				}
			}
		})

	}

}

func parseInput(input string) [][]int {
	lines := strings.Split(input, "\n")
	reports := make([][]int, len(lines))
	for i, line := range lines {
		nums := strings.Split(line, " ")
		levels := make([]int, len(nums))
		for j, num := range nums {
			levels[j], _ = strconv.Atoi(num)
		}
		reports[i] = levels
	}
	return reports
}

func partOne(reports [][]int) int {
	countSafeReports := 0
	for _, levels := range reports {
		if checkIfReportSafe(levels) {
			countSafeReports++
		}
	}
	return countSafeReports
}

func checkIfReportSafe(levels []int) bool {
	direction := 0
	for i := 0; i < len(levels)-1; i++ {
		diff := levels[i+1] - levels[i]
		if !checkIfSafe(&direction, diff) {
			return false
		}
	}
	return true
}

func checkIfSafe(direction *int, diff int) bool {
	absDiff := int(math.Abs(float64(diff)))
	if absDiff == 0 || absDiff > 3 {
		return false
	}

	if diff > 0 {
		diff = 1
	} else {
		diff = -1
	}

	if *direction == 0 {
		*direction = diff
	} else if *direction != diff {
		return false
	}
	return true
}

func partTwo(reports [][]int) int {
	countSafeReports := 0

	for _, levels := range reports {
		if checkIfReportSafe(levels) {
			countSafeReports++
		} else {
			for i := 0; i < len(levels); i++ {
				newLevels := make([]int, 0)
				newLevels = append(newLevels, levels[:i]...)
				newLevels = append(newLevels, levels[i+1:]...)
				if checkIfReportSafe(newLevels) {
					countSafeReports++
					break
				}
			}
		}
	}
	return countSafeReports
}
