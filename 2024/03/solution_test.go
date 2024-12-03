package _3

import (
	_ "embed"
	"regexp"
	"strconv"
	"testing"
)

//go:embed input.txt
var input string

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
			input:           "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
			expectedPartOne: 161,
			expectedPartTwo: 161,
		}, {
			name:            "with sample 2",
			input:           "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
			expectedPartOne: 161,
			expectedPartTwo: 48,
		},
		{
			name:            "with large input",
			input:           input,
			expectedPartOne: 0,
			expectedPartTwo: 0,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			if tst.expectedPartOne != -1 {
				if got := partOne(tst.input); got != tst.expectedPartOne {
					t.Errorf("partOne() = %v, want %v", got, tst.expectedPartOne)
				}
			}

			if tst.expectedPartTwo != -1 {
				if got := partTwo(tst.input); got != tst.expectedPartTwo {
					t.Errorf("partTwo() = %v, want %v", got, tst.expectedPartTwo)
				}
			}
		})

	}

}

func partOne(input string) int {
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	matches := re.FindAllString(input, -1)
	ans := 0
	for _, match := range matches {
		nums := regexp.MustCompile(`\d+`).FindAllString(match, -1)
		a, _ := strconv.Atoi(nums[0])
		b, _ := strconv.Atoi(nums[1])
		ans += a * b
	}
	return ans
}

func partTwo(input string) int {
	re := regexp.MustCompile(`do\(\)|don't\(\)|mul\(\d+,\d+\)`)
	matches := re.FindAllString(input, -1)
	ans := 0
	enabled := true
	for _, match := range matches {
		if match == "do()" {
			enabled = true

		} else if match == "don't()" {
			enabled = false
		} else if enabled {
			nums := regexp.MustCompile(`\d+`).FindAllString(match, -1)
			a, _ := strconv.Atoi(nums[0])
			b, _ := strconv.Atoi(nums[1])
			ans += a * b
		}
	}
	return ans
}
