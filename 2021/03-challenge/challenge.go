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

func getNum(count0 []int, count1 []int, numType string) float64 {
	num := float64(0)
	length := len(count0) - 1

	for idx := length; idx >= 0; idx-- {
		bit := 0
		if count0[idx] > count1[idx] {
			if numType == "gamma" {
				bit = 0
			} else {
				bit = 1
			}
		} else {
			if numType == "gamma" {
				bit = 1
			} else {
				bit = 0
			}
		}

		if bit == 1 {
			num = num + math.Pow(float64(2), float64(length-idx))
		}
	}

	fmt.Printf("%s : %f\n", numType, num)
	return num
}

func one() {
	_, currentFilePath, _, _ := runtime.Caller(0)
	file := input.GetFileMarker(currentFilePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	count0 := make([]int, 0)
	count1 := make([]int, 0)

	for scanner.Scan() {
		num := scanner.Text()
		if len(count0) == 0 {
			for _, char := range num {
				if char == '0' {
					count0 = append(count0, 1)
					count1 = append(count1, 0)
				} else {
					count0 = append(count0, 0)
					count1 = append(count1, 1)
				}
			}
		} else {
			for idx, char := range num {
				if char == '0' {
					count0[idx] += 1
				} else {
					count1[idx] += 1
				}
			}
		}
	}

	fmt.Printf("%f\n", getNum(count0, count1, "gamma")*getNum(count0, count1, "epsillon"))
}

func binaryToDecimal(num string) float64 {
	decimalNum := float64(0)
	length := len(num) - 1
	for idx := length; idx >= 0; idx-- {
		if num[idx] == '1' {
			decimalNum = decimalNum + math.Pow(float64(2), float64(length-idx))
		}
	}
	return decimalNum
}

func getNumTwo(nums []string, bitPos int, numType string) float64 {
	if len(nums) == 1 {
		decimalNum := binaryToDecimal(nums[0])
		fmt.Printf("%s: %f\n", numType, decimalNum)
		return decimalNum
	}

	count0 := 0
	count1 := 0

	for _, num := range nums {
		if num[bitPos] == '0' {
			count0 += 1
		} else {
			count1 += 1
		}
	}

	newNums := make([]string, 0)

	for _, num := range nums {
		if numType == "oxygen" && count0 > count1 && num[bitPos] == '0' {
			newNums = append(newNums, num)
		}
		if numType == "oxygen" && count0 <= count1 && num[bitPos] == '1' {
			newNums = append(newNums, num)
		}
		if numType == "carbon" && count0 > count1 && num[bitPos] == '1' {
			newNums = append(newNums, num)
		}
		if numType == "carbon" && count0 <= count1 && num[bitPos] == '0' {
			newNums = append(newNums, num)
		}
	}

	bitPos += 1

	return getNumTwo(newNums, bitPos, numType)

}

func two() {
	_, currentFilePath, _, _ := runtime.Caller(0)
	content := input.ReadInput(currentFilePath)

	nums := strings.Split(string(content), "\n")

	fmt.Printf("%f\n", getNumTwo(nums, 0, "oxygen")*getNumTwo(nums, 0, "carbon"))
}
