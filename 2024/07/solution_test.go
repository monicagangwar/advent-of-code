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
		expectedPartTwo int64
	}

	tests := []test{
		{
			name:            "with sample",
			input:           sample,
			expectedPartOne: 3749,
			expectedPartTwo: 11387,
		}, {
			name:            "with large input",
			input:           input,
			expectedPartOne: 0,
			expectedPartTwo: 0,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			equations := parseInput(tst.input)
			if tst.expectedPartOne != -1 {
				if got := partOne(equations); got != tst.expectedPartOne {
					t.Errorf("partOne() = %v, want %v", got, tst.expectedPartOne)
				}
			}

			if tst.expectedPartTwo != -1 {
				if got := partTwo(equations); got != tst.expectedPartTwo {
					t.Errorf("partTwo() = %v, want %v", got, tst.expectedPartTwo)
				}
			}
		})

	}

}

type equation struct {
	answer int
	nums   []int
}

func parseInput(input string) []equation {
	lines := strings.Split(input, "\n")
	eqRegex := regexp.MustCompile(`(?P<answer>\d+): (?P<nums>((\d+ )+(\d+)))`)
	equations := make([]equation, 0)
	for _, line := range lines {
		matches := eqRegex.FindStringSubmatch(line)
		answer, _ := strconv.Atoi(matches[1])
		nums := convertArrToStr(matches[2])
		equations = append(equations, equation{answer: answer, nums: nums})
	}
	return equations
}

func convertArrToStr(arrStr string) []int {
	nums := make([]int, 0)
	for _, numStr := range strings.Split(arrStr, " ") {
		num, _ := strconv.Atoi(numStr)
		nums = append(nums, num)
	}
	return nums
}

func partOne(equations []equation) int {
	validEqAnsSum := 0
	for _, eq := range equations {
		if isValid(eq.nums, 1, eq.nums[0], eq.answer) {
			validEqAnsSum += eq.answer
		}
	}
	return validEqAnsSum
}

func isValid(nums []int, index int, ans int, target int) bool {
	if len(nums) == index {
		return ans == target
	}

	return isValid(nums, index+1, ans+nums[index], target) || isValid(nums, index+1, ans*nums[index], target)
}

func isValidPartTwo(nums []int, index int, ans int, target int) bool {
	if len(nums) == index {
		return ans == target
	}

	return isValidPartTwo(nums, index+1, ans+nums[index], target) ||
		isValidPartTwo(nums, index+1, ans*nums[index], target) ||
		isValidPartTwo(nums, index+1, opConcat(ans, nums[index]), target)
}

func opConcat(a, b int) int {
	digitsB := []int{}
	for b > 0 {
		digitsB = append(digitsB, b%10)
		b /= 10
	}
	for i := len(digitsB) - 1; i >= 0; i-- {
		a = a*10 + digitsB[i]
	}
	return a
}

func partTwo(equations []equation) int64 {
	validEqAnsSum := int64(0)
	for _, eq := range equations {
		if isValidPartTwo(eq.nums, 1, eq.nums[0], eq.answer) {
			validEqAnsSum += int64(eq.answer)
		}
	}
	return validEqAnsSum
}
