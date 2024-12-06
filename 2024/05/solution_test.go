package template

import (
	_ "embed"
	"fmt"
	"regexp"
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
			expectedPartOne: 143,
			expectedPartTwo: 123,
		}, {
			name:            "with large input",
			input:           input,
			expectedPartOne: 0,
			expectedPartTwo: 0,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			sortOrder, updatePagesList := parseInput(tst.input)
			if tst.expectedPartOne != -1 {
				if got := partOne(sortOrder, updatePagesList); got != tst.expectedPartOne {
					t.Errorf("partOne() = %v, want %v", got, tst.expectedPartOne)
				}
			}

			if tst.expectedPartTwo != -1 {
				if got := partTwo(sortOrder, updatePagesList); got != tst.expectedPartTwo {
					t.Errorf("partTwo() = %v, want %v", got, tst.expectedPartTwo)
				}
			}
		})

	}

}

func parseInput(input string) (map[string]struct{}, [][]int) {
	lines := strings.Split(input, "\n")
	sortOrderRegEx := regexp.MustCompile(`^\d+|\d+$`)
	updateRegex := regexp.MustCompile(`^(\d+,)+\d+$`)
	sortOrder := make(map[string]struct{}, 0)
	updatePages := make([][]int, 0)
	for _, line := range lines {
		matches := sortOrderRegEx.FindAllString(line, -1)
		if len(matches) == 2 {
			sortOrder[line] = struct{}{}
		}

		matches = updateRegex.FindAllString(line, -1)
		if len(matches) > 0 {
			nums := make([]int, 0)
			for _, num := range strings.Split(matches[0], ",") {
				n, _ := strconv.Atoi(num)
				nums = append(nums, n)
			}
			updatePages = append(updatePages, nums)
		}

	}

	return sortOrder, updatePages
}

func partOne(sortOrder map[string]struct{}, updatePagesList [][]int) int {
	sumMidPageNo := 0
	for _, updatePages := range updatePagesList {
		validList := true
		for i := 0; i < len(updatePages); i++ {
			for j := i + 1; j < len(updatePages); j++ {
				if _, found := sortOrder[fmt.Sprintf("%d|%d", updatePages[j], updatePages[i])]; found {
					validList = false
					break
				}
			}
		}

		if validList {
			sumMidPageNo += updatePages[len(updatePages)/2]
		}

	}
	return sumMidPageNo
}

func partTwo(sortOrder map[string]struct{}, updatePagesList [][]int) int {
	sumMidPageNo := 0
	for _, updatePages := range updatePagesList {
		validList := true
		for i := 0; i < len(updatePages); i++ {
			for j := i + 1; j < len(updatePages); j++ {
				if _, found := sortOrder[fmt.Sprintf("%d|%d", updatePages[j], updatePages[i])]; found {
					validList = false
					tmp := updatePages[i]
					updatePages[i] = updatePages[j]
					updatePages[j] = tmp
				}
			}
		}
		if !validList {
			sumMidPageNo += updatePages[len(updatePages)/2]
		}

	}
	return sumMidPageNo
}
