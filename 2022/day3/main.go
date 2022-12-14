package main

import (
	"runtime"
	"strings"

	"github.com/monicagangwar/advent-of-code/input"
)

func main() {
	partOne()
	partTwo()
}

func getPriority(char rune) int {
	if char >= 'a' && char <= 'z' {
		return int(char-'a') + 1
	}
	return int(char-'A') + 27
}

func partOne() {
	_, currentFilePath, _, _ := runtime.Caller(0)
	content := input.ReadInput(currentFilePath)
	lines := strings.Split(string(content), "\n")
	priority := 0
	for _, line := range lines {
		rucksackLen := len(line)
		items := make(map[rune]struct{})
		for i := 0; i < rucksackLen/2; i++ {
			items[rune(line[i])] = struct{}{}
		}
		var commonChar rune
		for i := rucksackLen / 2; i < rucksackLen; i++ {
			if _, found := items[rune(line[i])]; found {
				commonChar = rune(line[i])
				break
			}
		}
		priority += getPriority(commonChar)

	}
	println(priority)
}

func partTwo() {
	_, currentFilePath, _, _ := runtime.Caller(0)
	content := input.ReadInput(currentFilePath)
	lines := strings.Split(string(content), "\n")
	priority := 0
	items := make(map[rune]int)
	elfInGroup := 0
	for _, line := range lines {
		elfInGroup += 1
		for i := 0; i < len(line); i++ {
			char := rune(line[i])
			if elfInGroup == 1 {
				items[char] = elfInGroup
			} else {
				if count, found := items[char]; found && count == elfInGroup-1 {
					items[char] = elfInGroup
					if elfInGroup == 3 {
						priority += getPriority(char)
						break
					}
				}
			}
		}
		if elfInGroup == 3 {
			items = make(map[rune]int)
			elfInGroup = 0
		}
	}
	println(priority)
}
