package _1

import (
	_ "embed"
	"strings"
	"testing"
)

//go:embed input.txt
var input string

func TestSolution(t *testing.T) {
	type test struct {
		input           string
		expectedPartOne int32
		expectedPartTwo int32
	}

	tests := []test{
		{
			input:           "1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet",
			expectedPartOne: 142,
			expectedPartTwo: -1,
		}, {
			input:           "two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen",
			expectedPartOne: -1,
			expectedPartTwo: 281,
		},
		{
			input:           input,
			expectedPartOne: ***REMOVED***,
			expectedPartTwo: ***REMOVED***,
		},
	}

	for _, tst := range tests {
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

	}

}

func partOne(input string) int32 {
	lines := strings.Split(input, "\n")
	calibrationSum := int32(0)

	for _, line := range lines {
		var firstDigit, lastDigit int32
		for _, char := range line {
			if char >= '0' && char <= '9' {
				firstDigit = char - '0'
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if line[i] >= '0' && line[i] <= '9' {
				lastDigit = int32(line[i]) - '0'
				break
			}
		}

		num := (firstDigit * 10) + lastDigit
		calibrationSum += num
	}

	return calibrationSum
}

func partTwo(input string) int32 {
	lines := strings.Split(input, "\n")

	numKeywords := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	calibrationSum := int32(0)

	for _, line := range lines {
		var firstDigit, lastDigit int32

		minIdx, maxIdx := 999, -1

		for keywordDigit, keyword := range numKeywords {
			if idx := strings.Index(line, keyword); idx != -1 && idx < minIdx {
				minIdx = idx
				firstDigit = int32(keywordDigit + 1)
			}

			if idx := strings.LastIndex(line, keyword); idx != -1 && idx > maxIdx {
				maxIdx = idx
				lastDigit = int32(keywordDigit + 1)
			}
		}

		for i := 0; i < minIdx && i < len(line); i++ {
			if line[i] >= '0' && line[i] <= '9' {
				firstDigit = int32(line[i] - '0')
				break
			}
		}

		for i := len(line) - 1; i > maxIdx && i >= 0; i-- {
			if line[i] >= '0' && line[i] <= '9' {
				lastDigit = int32(line[i] - '0')
				break
			}
		}

		num := (firstDigit * 10) + lastDigit
		//fmt.Printf("line: %s, num: %d\n", line, num)
		calibrationSum += num
	}
	return calibrationSum
}
