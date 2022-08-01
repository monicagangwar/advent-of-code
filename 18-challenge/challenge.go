package main

import (
	"bufio"
	"fmt"
	"runtime"
	"strconv"

	"github.com/monicagangwar/advent-of-code-2021/input"
)

func main() {
	_, currentFilePath, _, _ := runtime.Caller(0)
	file := input.GetFileMarker(currentFilePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	snailfishNumbers := make([][]interface{}, 0)

	for scanner.Scan() {
		line := scanner.Text()
		newNum := make([]interface{}, 0)
		for _, ch := range line {
			if ch == '[' || ch == ']' || ch == ',' {
				newNum = append(newNum, string(ch))
			} else {
				num, _ := strconv.ParseInt(string(ch), 10, 64)
				newNum = append(newNum, num)
			}
		}
		snailfishNumbers = append(snailfishNumbers, newNum)
	}

	partOne(snailfishNumbers)
	partTwo(snailfishNumbers)

}

func partOne(snailfishNumbers [][]interface{}) {
	output := performAdditionList(snailfishNumbers)

	fmt.Printf("\nadded       : %v", output)

	fmt.Printf("\nmagnitude   : %v", getMagnitude(output))
}

func partTwo(snailfishNumbers [][]interface{}) {
	largestMagnitude := int64(0)
	for i := 0; i < len(snailfishNumbers); i++ {
		for j := i + 1; j < len(snailfishNumbers); j++ {
			output1 := performAddition(snailfishNumbers[i], snailfishNumbers[j])
			magnitude1 := getMagnitude(output1)

			if magnitude1 > largestMagnitude {
				largestMagnitude = magnitude1
			}

			output2 := performAddition(snailfishNumbers[j], snailfishNumbers[i])
			magnitude2 := getMagnitude(output2)

			if magnitude2 > largestMagnitude {
				largestMagnitude = magnitude2
			}

		}
	}

	fmt.Printf("\nlargest magnitude   : %d", largestMagnitude)
}

func performAddition(num1 []interface{}, num2 []interface{}) []interface{} {
	addedNum := make([]interface{}, 0)
	addedNum = append(addedNum, "[")
	addedNum = append(addedNum, num1...)
	addedNum = append(addedNum, ",")
	addedNum = append(addedNum, num2...)
	addedNum = append(addedNum, "]")
	return reduce(addedNum)
}

func performAdditionList(numbers [][]interface{}) []interface{} {
	output := make([]interface{}, 0)
	for _, number := range numbers {
		if len(output) == 0 {
			output = number
		} else {
			output = performAddition(output, number)
		}
	}
	return output
}

func tryExplode(number []interface{}) ([]interface{}, bool) {
	level := 0
	for idx, ch := range number {
		if ch == "[" {
			level++
		} else if ch == "]" {
			level--
		}

		if level == 5 {
			left := number[idx+1].(int64)
			right := number[idx+3].(int64)
			for i := idx - 1; i >= 0; i-- {
				num, ok := number[i].(int64)
				if ok {
					number[i] = num + left
					break
				}
			}

			for i := idx + 4; i < len(number); i++ {
				num, ok := number[i].(int64)
				if ok {
					number[i] = num + right
					break
				}
			}

			reducedNumber := make([]interface{}, 0)
			reducedNumber = append(reducedNumber, number[:idx]...)
			reducedNumber = append(reducedNumber, int64(0))
			reducedNumber = append(reducedNumber, number[idx+5:]...)
			return reducedNumber, true

		}
	}

	return number, false
}

func trySplit(number []interface{}) ([]interface{}, bool) {
	for idx, ch := range number {
		num, ok := ch.(int64)
		if ok && num >= 10 {
			reducedNumber := make([]interface{}, 0)
			reducedNumber = append(reducedNumber, number[:idx]...)
			reducedNumber = append(reducedNumber, "[", num/2, ",", (num/2)+(num%2), "]")
			reducedNumber = append(reducedNumber, number[idx+1:]...)
			return reducedNumber, true
		}
	}
	return number, false
}

func reduce(number []interface{}) []interface{} {
	//fmt.Printf("orginal     : %v", number)
	changed := false
	for {
		number, changed = tryExplode(number)
		if changed {
			//fmt.Printf("\nexploded    : %v", number)
			continue
		}
		number, changed = trySplit(number)
		if !changed {
			return number
		}
		//else {
		//	fmt.Printf("\nsplit       : %v", number)
		//}
	}
}

func getMagnitude(number []interface{}) int64 {
	stack := make([]interface{}, 0)
	for _, ch := range number {
		if ch == "]" {
			right := stack[len(stack)-1].(int64)
			left := stack[len(stack)-2].(int64)

			stack = stack[:len(stack)-3]

			stack = append(stack, (left*3)+(right*2))

		} else if ch != "," {
			stack = append(stack, ch)
		}
	}
	return stack[0].(int64)
}
