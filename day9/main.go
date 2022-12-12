package main

import (
	"fmt"
	"math"
	"runtime"
	"strconv"
	"strings"

	"github.com/monicagangwar/advent-of-code-2022/input"
)

func main() {
	partOne(false)
	partTwo(false)
}

func moveTail(posHead [2]int, posTail *[2]int) {
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

func findDel(diff int) int {
	if diff == 0 {
		return 0
	}
	if diff < 0 {
		return -1
	}
	return 1
}

func moveRopeElement(posHead [2]int, posTail *[2]int) {
	xdiff, ydiff := posHead[0]-posTail[0], posHead[1]-posTail[1]

	if math.Abs(float64(xdiff)) > 1 || math.Abs(float64(ydiff)) > 1 {
		posTail[0] += findDel(xdiff)
		posTail[1] += findDel(ydiff)
	}
}

func moveRope(posRope *[10][2]int) {
	for idx := 0; idx < 9; idx++ {
		moveRopeElement(posRope[idx], &posRope[idx+1])
	}
}

func partOne(shouldDisplay bool) {
	_, currentFilePath, _, _ := runtime.Caller(0)
	content := input.ReadInput(currentFilePath)
	lines := strings.Split(string(content), "\n")

	positionsVisited := make(map[string]struct{})
	positionsVisited["0,0"] = struct{}{}
	posHead := [2]int{0, 0}
	posTail := [2]int{0, 0}
	var minX, maxX, minY, maxY int = 0, 0, 0, 0
	if shouldDisplay {
		display(&minX, &maxX, &minY, &maxY, posHead, posTail)
		fmt.Printf("\n")
	}
	for _, line := range lines {
		instruction := strings.Split(line, " ")
		direction := instruction[0]
		delta, _ := strconv.ParseInt(instruction[1], 10, 32)
		if shouldDisplay {
			fmt.Printf("\n=============== %s ===============\n", line)
		}
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

			moveTail(posHead, &posTail)
			positionsVisited[fmt.Sprintf("%d,%d", posTail[0], posTail[1])] = struct{}{}

			if shouldDisplay {
				display(&minX, &maxX, &minY, &maxY, posHead, posTail)
				fmt.Printf("\n")
			}
			delta--
			if delta == 0 {
				break
			}
		}
	}

	if shouldDisplay {
		fmt.Printf("\n")
		displayTail(minX, maxX, minY, maxY, positionsVisited)
	}

	fmt.Println(len(positionsVisited))
}

func partTwo(shouldDisplay bool) {
	_, currentFilePath, _, _ := runtime.Caller(0)
	content := input.ReadInput(currentFilePath)
	lines := strings.Split(string(content), "\n")

	positionsVisited := make(map[string]struct{})
	positionsVisited["0,0"] = struct{}{}
	posRope := [10][2]int{
		{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0},
	}
	var minX, maxX, minY, maxY int = 0, 0, 0, 0

	if shouldDisplay {
		displayRope(&minX, &maxX, &minY, &maxY, posRope)
		fmt.Printf("\n")
	}

	for _, line := range lines {
		instruction := strings.Split(line, " ")
		direction := instruction[0]
		delta, _ := strconv.ParseInt(instruction[1], 10, 32)

		if shouldDisplay {
			fmt.Printf("\n=============== %s ===============\n", line)
		}

		for {
			switch direction {
			case "L":
				posRope[0][1]--
				break
			case "R":
				posRope[0][1]++
				break
			case "U":
				posRope[0][0]++
				break
			case "D":
				posRope[0][0]--
				break
			}

			moveRope(&posRope)
			positionsVisited[fmt.Sprintf("%d,%d", posRope[9][0], posRope[9][1])] = struct{}{}

			if shouldDisplay {
				displayRope(&minX, &maxX, &minY, &maxY, posRope)
				fmt.Printf("\n")
			}

			delta--
			if delta == 0 {
				break
			}
		}
	}

	if shouldDisplay {
		fmt.Printf("\n")
		displayTail(minX, maxX, minY, maxY, positionsVisited)
	}
	fmt.Println(len(positionsVisited))
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

func displayRope(minX *int, maxX *int, minY *int, maxY *int, posRope [10][2]int) {

	for idx := 0; idx < 10; idx++ {
		*minX = min(*minX, posRope[idx][0])
		*maxX = max(*maxX, posRope[idx][0])
		*minY = min(*minY, posRope[idx][1])
		*maxY = max(*maxY, posRope[idx][1])
	}

	for x := *maxX; x >= *minX; x-- {
		for y := *minY; y <= *maxY; y++ {
			tailFound := false
			if x == posRope[0][0] && y == posRope[0][1] {
				fmt.Printf("H")
				continue
			} else {
				for idx := 1; idx < 10; idx++ {
					if x == posRope[idx][0] && y == posRope[idx][1] {
						fmt.Printf("%d", idx)
						tailFound = true
						break
					}
				}
				if tailFound {
					continue
				}
			}
			if x == 0 && y == 0 {
				fmt.Printf("s")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}

}
