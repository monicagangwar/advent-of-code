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
			expectedPartOne: 55312,
			expectedPartTwo: 65601038650482,
		}, {
			name:            "with large input",
			input:           input,
			expectedPartOne: 0,
			expectedPartTwo: 0,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			numArr := parseInput(tst.input)
			if tst.expectedPartOne != -1 {
				if got := partOne(numArr, 25); got != tst.expectedPartOne {
					t.Errorf("partOne() = %v, want %v", got, tst.expectedPartOne)
				}
			}

			if tst.expectedPartTwo != -1 {
				if got := partOne(numArr, 75); got != tst.expectedPartTwo {
					t.Errorf("partTwo() = %v, want %v", got, tst.expectedPartTwo)
				}
			}
		})

	}

}

func performOP(num int) (int, int) {

	if num == 0 {
		return 1, -1
	}

	numStr := fmt.Sprintf("%d", num)

	if len(numStr)%2 == 0 { // even
		newNumLeftStr := numStr[:len(numStr)/2]
		newNumRightStr := numStr[len(numStr)/2:]
		newNumLeft, _ := strconv.Atoi(newNumLeftStr)
		newNumRight, _ := strconv.Atoi(newNumRightStr)

		return newNumLeft, newNumRight
	}
	return num * 2024, -1
}

func parseInput(input string) []int {
	nums := strings.Split(input, " ")
	numArr := make([]int, 0)

	for _, num := range nums {
		numInt, _ := strconv.Atoi(num)
		numArr = append(numArr, numInt)
	}
	return numArr
}

func partOne(numArr []int, blinks int) int {

	numFreq := make(map[int]int)

	for _, num := range numArr {
		if _, found := numFreq[num]; found {
			numFreq[num]++
		} else {
			numFreq[num] = 1
		}

	}

	for i := 1; i <= blinks; i++ {
		//fmt.Println(numFreq)
		newFreq := make(map[int]int)
		for num, freq := range numFreq {

			num1, num2 := performOP(num)
			if _, found := newFreq[num1]; found {
				newFreq[num1] += freq
			} else {
				newFreq[num1] = freq
			}

			if num2 != -1 {
				if _, found := newFreq[num2]; found {
					newFreq[num2] += freq
				} else {
					newFreq[num2] = freq
				}
			}
		}

		numFreq = newFreq
	}

	countStones := 0
	for _, freq := range numFreq {
		countStones += freq
	}

	return countStones

}
