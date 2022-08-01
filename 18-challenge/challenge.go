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

	snailfishNumbers := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		snailfishNumbers = append(snailfishNumbers, line)
	}

	output := performAddition(snailfishNumbers)

	fmt.Printf("\nadded       : %s", output)

	//fmt.Println(getMagnitude(output))
}

func performAddition(numbers []string) string {
	output := ""
	for _, number := range numbers {
		if output == "" {
			output = number
		} else {
			output = fmt.Sprintf("[%s,%s]", output, number)
			output = reduce(output)
		}
		//fmt.Printf("\nadded       : %s", output)
	}
	return output
}

func reduce(number string) string {
	numberRepresentationList := make([]interface{}, 0)
	for _, ch := range number {
		if !(ch == '[' || ch == ']' || ch == ',') {
			num, _ := strconv.ParseInt(string(ch), 10, 64)
			numberRepresentationList = append(numberRepresentationList, num)
		} else {
			numberRepresentationList = append(numberRepresentationList, string(ch))
		}
	}
	fmt.Printf("orginal     : %v", numberRepresentationList)
	idx := 0
	level := 0
	for {
		if idx == len(numberRepresentationList)-1 {
			break
		}

		if numberRepresentationList[idx] == "[" {
			level++
		} else if numberRepresentationList[idx] == "]" {
			level--
		}

		if level == 5 {
			curIdx := idx
			idx = 0
			level = 0

			left := numberRepresentationList[curIdx+1].(int64)
			right := numberRepresentationList[curIdx+3].(int64)

			for i := curIdx - 1; i >= 0; i-- {
				num, ok := numberRepresentationList[i].(int64)
				if ok {
					numberRepresentationList[i] = num + left
					break
				}
			}

			for i := curIdx + 4; i < len(numberRepresentationList); i++ {
				num, ok := numberRepresentationList[i].(int64)
				if ok {
					numberRepresentationList[i] = num + right
					break
				}
			}

			newNumberRepresentationList := make([]interface{}, 0)
			newNumberRepresentationList = append(newNumberRepresentationList, numberRepresentationList[:curIdx]...)
			newNumberRepresentationList = append(newNumberRepresentationList, int64(0))
			newNumberRepresentationList = append(newNumberRepresentationList, numberRepresentationList[curIdx+5:]...)

			numberRepresentationList = newNumberRepresentationList

			fmt.Printf("\nexploded    : %v", numberRepresentationList)
			continue

		} else {
			num, ok := numberRepresentationList[idx].(int64)
			if ok && num >= 10 {
				curIdx := idx
				idx = 0
				level = 0

				newNumberRepresentationList := make([]interface{}, 0)
				newNumberRepresentationList = append(newNumberRepresentationList, numberRepresentationList[:curIdx]...)
				newNumberRepresentationList = append(newNumberRepresentationList, "[", num/2, ",", (num/2)+(num%2), "]")
				newNumberRepresentationList = append(newNumberRepresentationList, numberRepresentationList[curIdx+1:]...)
				numberRepresentationList = newNumberRepresentationList

				fmt.Printf("\nsplit       : %v", numberRepresentationList)

				continue

			}
		}

		idx++
	}

	return convertToString(numberRepresentationList)
}

func convertToString(numberList []interface{}) string {
	output := ""
	for _, ch := range numberList {
		num, ok := ch.(int64)
		if ok {
			output = fmt.Sprintf("%s%d", output, num)
		} else {
			output = fmt.Sprintf("%s%s", output, ch.(string))
		}
	}
	return output
}

func getMagnitude(number string) int {
	return 0
}
