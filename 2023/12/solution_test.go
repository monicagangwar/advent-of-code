package template

import (
	_ "embed"
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
			expectedPartOne: 21,
			expectedPartTwo: 0,
		}, {
			name:            "with large input",
			input:           input,
			expectedPartOne: 0,
			expectedPartTwo: -1,
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

			modifiedReports := expandReports(reports)

			if tst.expectedPartTwo != -1 {
				if got := partOne(modifiedReports); got != tst.expectedPartTwo {
					t.Errorf("partTwo() = %v, want %v", got, tst.expectedPartTwo)
				}
			}
		})

	}

}

type Report struct {
	arrangement       []byte
	arrangementGroups []int
}

func parseInput(input string) []Report {
	lines := strings.Split(input, "\n")
	arrangementRegex := regexp.MustCompile(`(?P<agmt>[\?\.#]*) (?P<nums>[\d,]*)`)
	reports := make([]Report, 0)
	for _, line := range lines {
		matches := arrangementRegex.FindStringSubmatch(line)
		agmt := matches[1]
		nums := convertStrToIntSlice(matches[2])
		reports = append(reports, Report{
			arrangement:       []byte(agmt),
			arrangementGroups: nums,
		})
	}
	return reports
}

func convertStrToIntSlice(str string) []int {
	numsStr := strings.Split(str, ",")
	numArr := make([]int, len(numsStr))
	for i, numStr := range numsStr {
		num, _ := strconv.Atoi(numStr)
		numArr[i] = num
	}
	return numArr
}

func expandReports(reports []Report) []Report {
	for i, report := range reports {
		tmpAgmtGrp := report.arrangementGroups
		tmpAgmt := report.arrangement
		agmtArr := make([]string, 5)
		afmtGrpArr := make([]int, 0)
		for j := 0; j < 5; j++ {
			agmtArr[j] = string(tmpAgmt)
			afmtGrpArr = append(afmtGrpArr, tmpAgmtGrp...)
		}

		reports[i].arrangementGroups = afmtGrpArr
		reports[i].arrangement = []byte(strings.Join(agmtArr, "?"))

	}
	return reports
}

func isValid(arrangement []byte, arrangementGroups []int) bool {
	groups := make([]int, 0)
	count := 0
	for i := 0; i < len(arrangement); i++ {
		if arrangement[i] == '#' {
			count++
		} else if arrangement[i] == '.' && count > 0 {
			groups = append(groups, count)
			count = 0
		}
	}
	if count > 0 {
		groups = append(groups, count)
	}

	if len(groups) != len(arrangementGroups) {
		return false
	}
	for i := 0; i < len(groups); i++ {
		if groups[i] != arrangementGroups[i] {
			return false
		}
	}

	return true
}

func getValidAgmtCount(arrangement []byte, arrangementGroups []int, index int, memoized map[string]int) int {
	if count, found := memoized[string(arrangement)]; found {
		return count
	}

	if index == len(arrangement) {
		if isValid(arrangement, arrangementGroups) {
			return 1
		}
		return 0
	}

	if arrangement[index] == '?' {
		newArrangement1 := copyArrangement(arrangement)
		newArrangement1[index] = '.'
		count1 := getValidAgmtCount(newArrangement1, arrangementGroups, index+1, memoized)

		memoized[string(newArrangement1)] = count1

		newArrangement2 := copyArrangement(arrangement)
		newArrangement2[index] = '#'
		count2 := getValidAgmtCount(newArrangement2, arrangementGroups, index+1, memoized)

		memoized[string(newArrangement2)] = count2

		return count1 + count2
	}

	count := getValidAgmtCount(arrangement, arrangementGroups, index+1, memoized)
	memoized[string(arrangement)] = count
	return count
}

func copyArrangement(arrangement []byte) []byte {
	newArrangement := make([]byte, len(arrangement))
	copy(newArrangement, arrangement)
	return newArrangement
}

func partOne(reports []Report) int {
	validArrangements := 0
	memoized := make(map[string]int)
	for _, report := range reports {
		validArrangements += getValidAgmtCount(report.arrangement, report.arrangementGroups, 0, memoized)
	}
	return validArrangements
}
