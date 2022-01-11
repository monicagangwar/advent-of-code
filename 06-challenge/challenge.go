package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"

	"github.com/monicagangwar/advent-of-code-2021/input"
)

func main() {
	challenge(80)
	challenge(256)
}

func challenge(days int) {
	_, currentFilePath, _, _ := runtime.Caller(0)
	content := input.ReadInput(currentFilePath)

	countByStages := make([]int, 10)

	for idx := 0; idx < 10; idx++ {
		countByStages = append(countByStages, 0)
	}

	for _, strNum := range strings.Split(string(content), ",") {
		num, _ := strconv.Atoi(strNum)
		countByStages[num] += 1
	}

	for day := 1; day <= days; day++ {
		count0 := countByStages[0]
		for idx := 1; idx <= 8; idx++ {
			countByStages[idx-1] = countByStages[idx]
		}
		countByStages[8] = count0
		countByStages[6] += count0

		//fmt.Printf("day %d: %+v\n", day, countByStages)

	}

	total := 0

	for i := 0; i <= 8; i++ {
		total += countByStages[i]
	}

	fmt.Printf("%d\n", total)

}
