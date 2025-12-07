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
			expectedPartOne: 4277556,
			expectedPartTwo: 3263827,
		}, {
			name:            "with large input",
			input:           input,
			expectedPartOne: 0,
			expectedPartTwo: 0,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			if tst.expectedPartOne != -1 {
				if got := partOne(parseInput(tst.input)); got != tst.expectedPartOne {
					t.Errorf("partOne() = %v, want %v", got, tst.expectedPartOne)
				}
			}

			if tst.expectedPartTwo != -1 {
				if got := partOne(parseInputPartTwo(tst.input)); got != tst.expectedPartTwo {
					t.Errorf("partTwo() = %v, want %v", got, tst.expectedPartTwo)
				}
			}
		})

	}

}

func parseInput(input string) [][]int {
	lines := strings.Split(input, "\n")
	nums := make([][]int, len(lines))
	maxI := len(lines)
	maxJ := 0
	for i, line := range lines {
		line = strings.ReplaceAll(line, "  ", " ")
		numStr := strings.Split(line, " ")
		numList := make([]int, 0)
		for _, str := range numStr {
			str = strings.TrimSpace(str)
			if str == "" {
				continue
			}
			num, err := strconv.Atoi(str)
			if err != nil {
				if str == "+" {
					num = 0
				}
				if str == "*" {
					num = 1
				}
			}
			numList = append(numList, num)
		}
		nums[i] = numList
		if len(nums[i]) > maxJ {
			maxJ = len(nums[i])
		}
	}

	numsTranspose := make([][]int, 0)
	for j := 0; j < maxJ; j++ {
		numsTransposeRow := make([]int, 0)
		for i := 0; i < maxI; i++ {
			if j < len(nums[i]) {
				numsTransposeRow = append(numsTransposeRow, nums[i][j])
			}
		}
		numsTranspose = append(numsTranspose, numsTransposeRow)
	}
	return numsTranspose
}
func parseInputPartTwo(input string) [][]int {
	lines := make([][]byte, 0)
	maxLineLen := 0
	for _, line := range strings.Split(input, "\n") {
		lines = append(lines, []byte(line))
		if len(line) > maxLineLen {
			maxLineLen = len(line)
		}
	}
	nums := make([][]int, 0)
	numArr := make([]int, 0)
	for j := 0; j < maxLineLen; j++ {
		num := 0
		numFound := false
		for i := 0; i < len(lines)-1; i++ {
			if j >= len(lines[i]) {
				continue
			}
			numVal, err := strconv.Atoi(string(lines[i][j]))
			if err != nil {
				continue
			}
			numFound = true
			num = num*10 + numVal
		}
		if numFound {
			numArr = append(numArr, num)
		} else if len(numArr) > 0 {
			nums = append(nums, numArr)
			numArr = make([]int, 0)
		}
	}
	if len(numArr) > 0 {
		nums = append(nums, numArr)
	}
	idx := 0
	for _, char := range lines[len(lines)-1] {
		if char == '*' {
			nums[idx] = append(nums[idx], 1)
			idx++
		} else if char == '+' {
			nums[idx] = append(nums[idx], 0)
			idx++
		}
	}
	return nums
}

func partOne(nums [][]int) int {
	total := 0

	for i := 0; i < len(nums); i++ {
		operator := nums[i][len(nums[i])-1]
		result := 0
		if operator == 1 {
			result = 1
		}
		for j := 0; j < len(nums[i])-1; j++ {
			if operator == 1 {
				result *= nums[i][j]
			} else {
				result += nums[i][j]
			}
		}
		total += result
	}
	return total
}
