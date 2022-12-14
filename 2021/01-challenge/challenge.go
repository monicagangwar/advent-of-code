package main

import (
	"bufio"
	"fmt"
	"runtime"
	"strconv"

	"github.com/monicagangwar/advent-of-code/input"
)

func main() {
	one()
	two()
}

func one() {
	_, currentFilePath, _, _ := runtime.Caller(0)
	file := input.GetFileMarker(currentFilePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	prevNum := -1
	increase := 0

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		if prevNum == -1 {
			prevNum = num
		}

		if num > prevNum {
			increase += 1
		}

		prevNum = num
	}

	fmt.Printf("%d\n", increase)

}

func two() {
	_, currentFilePath, _, _ := runtime.Caller(0)
	file := input.GetFileMarker(currentFilePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	prevNum1 := -1
	prevNum2 := -1
	prevNum3 := -1

	increase := 0

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		if prevNum1 == -1 {
			prevNum1 = num
		} else if prevNum2 == -1 {
			prevNum2 = num
		} else if prevNum3 == -1 {
			prevNum3 = num
		} else {
			if num > prevNum1 {
				increase += 1
			}

			prevNum1 = prevNum2
			prevNum2 = prevNum3
			prevNum3 = num
		}
	}

	fmt.Printf("%d\n", increase)

}
