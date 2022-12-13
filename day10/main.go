package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"

	"github.com/monicagangwar/advent-of-code-2022/input"
)

func main() {
	partOne()
	partTwo()
}

func computeSignalStrength(cycle int, x int) int {
	if cycle == 20 || (cycle-20)%40 == 0 {
		return cycle * x
	}
	return 0
}

func draw(cycle int, x int) {
	curSpritePos := cycle % 40
	//fmt.Printf("%d %d %d", cycle, x, curSpritePos)

	if curSpritePos >= x && curSpritePos < x+3 {
		fmt.Printf("#")
	} else {
		fmt.Printf(".")
	}
	if cycle%40 == 0 {
		fmt.Println()
	}
	//fmt.Println()
}

func partOne() {
	_, currentFilePath, _, _ := runtime.Caller(0)
	content := input.ReadInput(currentFilePath)
	lines := strings.Split(string(content), "\n")

	x := 1
	cycle := 1
	signalStrength := 0
	for _, line := range lines {
		//fmt.Printf("%d %d %s\n", cycle, x, line)
		if line == "noop" {
			cycle++
			signalStrength += computeSignalStrength(cycle, x)
		} else {
			instruction := strings.Split(line, " ")
			num, _ := strconv.Atoi(instruction[1])
			cycle++
			signalStrength += computeSignalStrength(cycle, x)
			//fmt.Printf("%d %d %s\n", cycle, x, line)
			cycle++
			x += num
			signalStrength += computeSignalStrength(cycle, x)
		}
	}
	fmt.Println(signalStrength)
}

func partTwo() {
	_, currentFilePath, _, _ := runtime.Caller(0)
	content := input.ReadInput(currentFilePath)
	lines := strings.Split(string(content), "\n")

	x := 1
	cycle := 1
	draw(cycle, x)
	for _, line := range lines {
		if line == "noop" {
			cycle++
			draw(cycle, x)
		} else {
			instruction := strings.Split(line, " ")
			num, _ := strconv.Atoi(instruction[1])
			cycle++
			draw(cycle, x)
			cycle++
			x += num
			draw(cycle, x)
		}
	}
}
