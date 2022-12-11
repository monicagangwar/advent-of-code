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

func moveTail(posHead *[2]int, posTail *[2]int) {
	xdiff, ydiff := posHead[0]-posTail[0], posHead[1]-posTail[1]
	if xdiff > 1 {
		posTail[0]++
		posTail[1] += ydiff
	} else if xdiff < -1 {
		posTail[0]--
		posTail[1] += ydiff
	}

	if ydiff > 1 {
		posTail[1]++
		posTail[0] += xdiff
	} else if ydiff < -1 {
		posTail[1]--
		posTail[0] += xdiff
	}
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func display(minX *int, maxX *int, minY *int, maxY *int, posHead [2]int, posTail [2]int) {

	*minX = min(*minX, min(posHead[0], posTail[0]))
	*maxX = max(*maxX, max(posHead[0], posTail[0]))
	*minY = min(*minY, min(posHead[1], posTail[1]))
	*maxY = max(*maxY, max(posHead[1], posTail[1]))

	for x := *maxX; x >= *minX; x-- {
		for y := *minY; y <= *maxY; y++ {

			if x == posHead[0] && y == posHead[1] {
				fmt.Printf("H")
			} else if x == posTail[0] && y == posTail[1] {
				fmt.Printf("T")
			} else if x == 0 && y == 0 {
				fmt.Printf("s")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}

}

func displayTail(minX int, maxX int, minY int, maxY int, positionsVisited map[string]struct{}) {
	for x := maxX; x >= minX; x-- {
		for y := minY; y <= maxY; y++ {

			if x == 0 && y == 0 {
				fmt.Printf("s")
			} else if _, found := positionsVisited[fmt.Sprintf("%d,%d", x, y)]; found {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}

func partOne() {
	_, currentFilePath, _, _ := runtime.Caller(0)
	content := input.ReadInput(currentFilePath)
	lines := strings.Split(string(content), "\n")

	positionsVisited := make(map[string]struct{})
	positionsVisited["0,0"] = struct{}{}
	posHead := [2]int{0, 0}
	posTail := [2]int{0, 0}
	//var minX, maxX, minY, maxY int = 0, 0, 0, 0

	//display(&minX, &maxX, &minY, &maxY, posHead, posTail)
	//fmt.Printf("\n")

	for _, line := range lines {
		instruction := strings.Split(line, " ")
		direction := instruction[0]
		delta, _ := strconv.ParseInt(instruction[1], 10, 32)
		//fmt.Printf("\n=============== %s ===============\n", line)

		for {
			switch direction {
			case "L":
				posHead[1]--
				break
			case "R":
				posHead[1]++
				break
			case "U":
				posHead[0]++
				break
			case "D":
				posHead[0]--
				break
			}

			moveTail(&posHead, &posTail)
			positionsVisited[fmt.Sprintf("%d,%d", posTail[0], posTail[1])] = struct{}{}

			//display(&minX, &maxX, &minY, &maxY, posHead, posTail)
			//fmt.Printf("\n")

			delta--
			if delta == 0 {
				break
			}
		}
	}

	//display(&minX, &maxX, &minY, &maxY, posHead, posTail)
	//fmt.Printf("\n")
	//displayTail(minX, maxX, minY, maxY, positionsVisited)

	fmt.Println(len(positionsVisited))
}

func partTwo() {}
