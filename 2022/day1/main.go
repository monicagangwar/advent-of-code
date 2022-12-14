package main

import (
	"fmt"
	"runtime"
	"sort"
	"strconv"
	"strings"

	"github.com/monicagangwar/advent-of-code/input"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	_, currentFilePath, _, _ := runtime.Caller(0)
	content := input.ReadInput(currentFilePath)
	lines := strings.Split(string(content), "\n")
	maxCalories := int64(0)
	calories := int64(0)
	for _, line := range lines {
		if line == "" {
			if calories > maxCalories {
				maxCalories = calories
			}
			calories = int64(0)
		}
		calorie, _ := strconv.ParseInt(line, 10, 64)
		calories += calorie
	}
	fmt.Println(maxCalories)
}

func partTwo() {
	_, currentFilePath, _, _ := runtime.Caller(0)
	content := input.ReadInput(currentFilePath)
	lines := strings.Split(string(content), "\n")
	caloriesList := make([]int, 0)
	calories := 0
	for _, line := range lines {
		if line == "" {
			caloriesList = append(caloriesList, calories)
			calories = 0
		}
		calorie, _ := strconv.ParseInt(line, 10, 32)
		calories += int(calorie)
	}
	caloriesList = append(caloriesList, calories)
	sort.Ints(caloriesList)
	elvesCount := len(caloriesList)
	fmt.Println(caloriesList[elvesCount-1] + caloriesList[elvesCount-2] + caloriesList[elvesCount-3])
}
