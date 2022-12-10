package main

import (
	"fmt"
	"runtime"

	"github.com/monicagangwar/advent-of-code-2022/input"
)

func main() {
	compute(4)
	compute(14)
}

func compute(marker int) {
	_, currentFilePath, _, _ := runtime.Caller(0)
	content := input.ReadInput(currentFilePath)
	charFoundIdx := make(map[string]int, 0)
	charConsumed := 0
	startingIdx := 0
	for idx, char := range string(content) {
		charConsumed++
		foundIdx, found := charFoundIdx[string(char)]
		if found && startingIdx <= foundIdx {
			startingIdx = foundIdx + 1
			charConsumed = idx - foundIdx
		}
		charFoundIdx[string(char)] = idx
		if charConsumed == marker {
			fmt.Println(idx + 1)
			break
		}
	}
}
