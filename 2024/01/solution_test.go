package _1

import (
	_ "embed"
	"math"
	"sort"
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
			expectedPartOne: 11,
			expectedPartTwo: 31,
		}, {
			name:            "with large input",
			input:           input,
			expectedPartOne: 0,
			expectedPartTwo: 0,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			leftNums, rightNums := parseInput(tst.input)
			if tst.expectedPartOne != -1 {
				if got := partOne(leftNums, rightNums); got != tst.expectedPartOne {
					t.Errorf("partOne() = %v, want %v", got, tst.expectedPartOne)
				}
			}

			if tst.expectedPartTwo != -1 {
				if got := partTwo(leftNums, rightNums); got != tst.expectedPartTwo {
					t.Errorf("partTwo() = %v, want %v", got, tst.expectedPartTwo)
				}
			}
		})

	}

}

func parseInput(input string) ([]int, []int) {
	lines := strings.Split(input, "\n")
	leftNums := make([]int, len(lines))
	rightNums := make([]int, len(lines))
	for i, line := range lines {
		nums := strings.Split(line, "   ")
		//fmt.Println(nums)
		leftNum, _ := strconv.Atoi(strings.TrimSpace(nums[0]))
		rightNum, _ := strconv.Atoi(strings.TrimSpace(nums[1]))
		leftNums[i] = leftNum
		rightNums[i] = rightNum
	}
	return leftNums, rightNums
}

func partOne(leftNums, rightNums []int) int {
	sort.Slice(leftNums, func(i, j int) bool {
		return leftNums[i] < leftNums[j]
	})

	sort.Slice(rightNums, func(i, j int) bool {
		return rightNums[i] < rightNums[j]
	})

	//fmt.Println(leftNums)
	//fmt.Println(rightNums)

	dist := 0
	for i := 0; i < len(leftNums); i++ {

		dist += int(math.Abs(float64(leftNums[i] - rightNums[i])))
	}

	return dist
}

func partTwo(leftNums, rightNums []int) int {
	numMap := make(map[int]int)
	for _, num := range leftNums {
		numMap[num] = 0
	}
	for _, num := range rightNums {
		if _, found := numMap[num]; found {
			numMap[num]++
		}
	}

	similarityScore := 0

	for _, num := range leftNums {
		numFreq := numMap[num]
		similarityScore += num * numFreq
	}

	return similarityScore
}
