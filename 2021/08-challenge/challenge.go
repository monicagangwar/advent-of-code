package main

import (
	"bufio"
	"fmt"
	"math"
	"runtime"
	"strings"

	"github.com/monicagangwar/advent-of-code/input"
)

func main() {
	one()
	two()
}

func isUnique(digitStr string) bool {
	digitStore := make([]bool, 7)
	for idx := 0; idx < 7; idx++ {
		digitStore[idx] = false
	}
	for _, digit := range digitStr {
		if digitStore[digit-'a'] == true {
			return false
		}
		digitStore[digit-'a'] = true
	}
	return true
}

func one() {

	_, currentFilePath, _, _ := runtime.Caller(0)
	file := input.GetFileMarker(currentFilePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	counter := 0

	for scanner.Scan() {
		line := scanner.Text()
		outputDigits := strings.Split(strings.Trim(strings.Split(line, "|")[1], " "), " ")
		for _, digit := range outputDigits {
			lenDigit := len(digit)
			if (lenDigit == 2 || lenDigit == 3 || lenDigit == 4 || lenDigit == 7) && isUnique(digit) {
				counter += 1
			}
		}
	}

	fmt.Printf("%d\n", counter)
}

func diff(string1 string, string2 string) string {

	digits := make([]int, 7)
	for idx := 0; idx < 7; idx++ {
		digits[idx] = 0
	}
	for _, char := range string1 {
		digits[char-'a'] += 1
	}

	for _, char := range string2 {
		digits[char-'a'] += 1
	}

	diffStr := ""

	for idx, count := range digits {
		if count == 1 {
			diffStr = fmt.Sprintf("%s%s", diffStr, string(idx+'a'))
		}
	}

	//fmt.Printf("%s - %s = %s\n", string1, string2, diffStr)

	return diffStr
}

func getNumDigit(digitSum int, digitLen int) int {

	switch digitSum {
	case 21:
		if digitLen == 6 {
			return 0
		} else {
			return 5
		}
	case 5:
		return 1
	case 19:
		return 2
	case 17:
		return 3
	case 18:
		return 4
	case 26:
		return 6
	case 6:
		return 7
	case 28:
		return 8
	case 23:
		return 9
	}
	return -1
}

func two() {
	_, currentFilePath, _, _ := runtime.Caller(0)
	file := input.GetFileMarker(currentFilePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := float64(0)

	for scanner.Scan() {
		line := scanner.Text()
		inputDigits := strings.Split(strings.TrimSpace(strings.Split(line, "|")[0]), " ")
		outputDigits := strings.Split(strings.TrimSpace(strings.Split(line, "|")[1]), " ")

		digits := make([]string, 7)

		digits1 := ""
		digits7 := ""
		digits4 := ""
		digits8 := ""

		digitCount := make([]int, 7)

		for idx := 0; idx < 7; idx++ {
			digitCount[idx] = 0
		}

		//fmt.Printf("%+v\n", inputDigits)

		for _, digit := range inputDigits {
			if len(digit) == 2 {
				digits1 = digit
			} else if len(digit) == 4 {
				digits4 = digit
			} else if len(digit) == 3 {
				digits7 = digit
			} else if len(digit) == 7 {
				digits8 = digit
			}

			for _, char := range digit {
				digitCount[char-'a'] += 1
			}
		}

		for idx, count := range digitCount {
			if count == 9 {
				digits[2] = string(idx + 'a')
			}
			if count == 4 {
				digits[4] = string(idx + 'a')
			}
			if count == 6 {
				digits[5] = string(idx + 'a')
			}
		}

		digits[0] = diff(digits7, digits1)
		digits[1] = diff(digits1, digits[2])
		digits[3] = diff(digits8, fmt.Sprintf("%s%s%s", digits4, digits[0], digits[4]))
		digits[6] = diff(digits4, fmt.Sprintf("%s%s", digits1, digits[5]))

		//fmt.Printf("%+v\n", digits)

		transposedDigits := make([]int, 7)

		for idx, char := range digits {
			transposedDigits[char[0]-'a'] = idx + 1
		}

		num := float64(0)

		for idx, digit := range outputDigits {
			digitSum := 0
			for _, char := range digit {
				digitSum += transposedDigits[char-'a']
			}
			num += float64(getNumDigit(digitSum, len(digit))) * math.Pow10(3-idx)
			//fmt.Printf("%s %d, %d = %d\n", digit, digitSum, len(digit), getNumDigit(digitSum, len(digit)))
		}

		//fmt.Printf("%+v", digits)
		//fmt.Printf("%+v", transposedDigits)
		//
		//fmt.Printf("%s = %f\n", outputDigits, num)

		sum += num

	}

	fmt.Printf("%f\n", sum)
}
