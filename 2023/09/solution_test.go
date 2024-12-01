package _8

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
			expectedPartOne: 114,
			expectedPartTwo: 2,
		}, {
			name:            "with large input",
			input:           input,
			expectedPartOne: 0,
			expectedPartTwo: 0,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			nums := parseInput(tst.input)
			if tst.expectedPartOne != -1 {
				if got := partOne(nums); got != tst.expectedPartOne {
					t.Errorf("partOne() = %v, want %v", got, tst.expectedPartOne)
				}
			}

			if tst.expectedPartTwo != -1 {
				if got := partTwo(nums); got != tst.expectedPartTwo {
					t.Errorf("partTwo() = %v, want %v", got, tst.expectedPartTwo)
				}
			}
		})

	}

}

func partOne(nums [][]int) int {
	result := 0
	for _, row := range nums {
		_, nextNum := findNextAndFirstNum(row)
		result += nextNum
	}
	return result
}

func partTwo(nums [][]int) int {
	result := 0
	for _, row := range nums {
		firstNum, _ := findNextAndFirstNum(row)
		result += firstNum
	}
	return result
}

func findNextAndFirstNum(row []int) (int, int) {
	diffRows := make([][]int, 0)
	diffRows = append(diffRows, row)
	diffRowIdx := 0
	for {
		newDiffRow := make([]int, 0)
		nonZeroFound := false
		for i := 0; i < len(diffRows[diffRowIdx])-1; i++ {
			diff := diffRows[diffRowIdx][i+1] - diffRows[diffRowIdx][i]
			if diff != 0 {
				nonZeroFound = true
			}
			newDiffRow = append(newDiffRow, diff)
		}
		if !nonZeroFound {
			break
		}
		diffRows = append(diffRows, newDiffRow)
		diffRowIdx++
	}
	nextNum := 0
	firstNum := 0
	for diffRowIdx := len(diffRows) - 1; diffRowIdx >= 0; diffRowIdx-- {
		curRow := diffRows[diffRowIdx]
		nextNum += curRow[len(curRow)-1]
		if firstNum == 0 {
			firstNum = curRow[0]
		} else {
			firstNum = curRow[0] - firstNum
		}
	}

	//fmt.Printf("diffRows: %v, newNum: %d firstNum: %d\n", diffRows, nextNum, firstNum)
	return firstNum, nextNum
}

func parseInput(input string) [][]int {
	lines := strings.Split(input, "\n")

	nums := make([][]int, len(lines))

	for i, line := range lines {
		rowNumsStr := strings.Split(line, " ")
		row := make([]int, len(rowNumsStr))
		for j, numStr := range rowNumsStr {
			num, _ := strconv.Atoi(numStr)
			row[j] = num
		}
		nums[i] = row
	}
	return nums
}
