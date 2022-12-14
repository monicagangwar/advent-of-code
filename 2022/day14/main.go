package main

import (
	"fmt"
	"math"
	"runtime"
	"strconv"
	"strings"

	"github.com/buger/goterm"

	"github.com/monicagangwar/advent-of-code/input"
)

func main() {
	_, currentFilePath, _, _ := runtime.Caller(0)
	content := input.ReadInput(currentFilePath)
	lines := strings.Split(string(content), "\n")
	var cave [1000][1000]int
	minX, maxX, minY, maxY := math.MaxInt32, 0, math.MaxInt32, 0
	for _, line := range lines {
		points := strings.Split(line, " -> ")
		for idx := 0; idx < len(points)-1; idx++ {
			x1, y1 := getPoint(points[idx])
			x2, y2 := getPoint(points[idx+1])
			minX = min(minX, min(x1, x2))
			maxX = max(maxX, max(x1, x2))
			minY = min(minY, min(y1, y2))
			maxY = max(maxY, max(y1, y2))

			if x1 == x2 {
				start, end := y1, y2
				if y2 < y1 {
					start, end = y2, y1
				}
				for y := start; y <= end; y++ {
					cave[x1][y] = 1
				}
			} else {
				start, end := x1, x2
				if x2 < x1 {
					start, end = x2, x1
				}
				for x := start; x <= end; x++ {
					cave[x][y1] = 1
				}
			}
		}
	}
	// part one
	compute(cave, minX, maxX, minY, maxY, true)

	// part two
	for x := 0; x < 1000; x++ {
		cave[x][maxY+2] = 1
	}
	compute(cave, minX, maxX, minY, maxY+2, false)
}

func compute(cave [1000][1000]int, minX, maxX, minY, maxY int, abyss bool) {
	sandParticles := 0
	breakPoint := false
	for {
		sandPosX, sandPosY := 500, 0
		for {
			if abyss && (sandPosX < minX || sandPosX > maxX || sandPosY > maxY) {
				breakPoint = true
				break
			}
			//display(cave, minX, maxX, maxY, sandPosX, sandPosY, true)
			if cave[sandPosX][sandPosY+1] == 0 {
				sandPosY++
			} else if cave[sandPosX-1][sandPosY+1] == 0 {
				sandPosX--
				sandPosY++
			} else if cave[sandPosX+1][sandPosY+1] == 0 {
				sandPosX++
				sandPosY++
			} else {
				sandParticles++
				cave[sandPosX][sandPosY] = 2
				if !abyss {
					minX = min(minX, sandPosX)
					maxX = max(maxX, sandPosX)
					if sandPosX == 500 && sandPosY == 0 {
						breakPoint = true
					}
				}
				//display(cave, minX, maxX, maxY, -1, -1, true)
				break
			}
		}
		if breakPoint {
			break
		}
	}
	display(cave, minX, maxX, maxY, -1, -1, false)
	fmt.Println(sandParticles)
}

func display(cave [1000][1000]int, minX, maxX, maxY, sandPosX, sandPosY int, flush bool) {
	if flush {
		goterm.MoveCursor(1, 6)
	}
	for y := 0; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if cave[x][y] == 1 {
				fmt.Printf(goterm.Color("#", goterm.BLUE))
			} else if (x == sandPosX && y == sandPosY) || cave[x][y] == 2 {
				fmt.Printf(goterm.Color("o", goterm.YELLOW))
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
	if flush {
		goterm.Flush()
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

func getPoint(pointStr string) (int, int) {
	points := strings.Split(pointStr, ",")
	x, _ := strconv.Atoi(points[0])
	y, _ := strconv.Atoi(points[1])
	return x, y
}
