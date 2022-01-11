package main

import (
	"fmt"
	"runtime"
	"sort"
	"strconv"
	"strings"

	"github.com/monicagangwar/advent-of-code-2021/input"
)

func main() {
	one()
	two()
}

func one() {
	_, currentFilePath, _, _ := runtime.Caller(0)
	content := input.ReadInput(currentFilePath)

	floor := make([][]int, 0)

	for _, rowStr := range strings.Split(string(content), "\n") {
		row := make([]int, 0)
		for _, char := range rowStr {
			num, _ := strconv.Atoi(string(char))
			row = append(row, num)
		}
		floor = append(floor, row)
	}

	maxRow := len(floor)
	maxCol := len(floor[0])

	sum := 0

	for row := 0; row < maxRow; row++ {
		for col := 0; col < maxCol; col++ {
			curPointMin := true

			// top
			if row-1 >= 0 && floor[row-1][col] <= floor[row][col] {
				curPointMin = false
			}
			// right
			if curPointMin && col+1 < maxCol && floor[row][col+1] <= floor[row][col] {
				curPointMin = false
			}
			// bottom
			if curPointMin && row+1 < maxRow && floor[row+1][col] <= floor[row][col] {
				curPointMin = false
			}
			// left
			if curPointMin && col-1 >= 0 && floor[row][col-1] <= floor[row][col] {
				curPointMin = false
			}

			if curPointMin {
				//fmt.Printf("%d %d : %d\n", row, col, floor[row][col])
				sum += floor[row][col] + 1
			}
		}
	}

	fmt.Printf("%d\n", sum)

}

type point struct {
	row int
	col int
}

func two() {
	_, currentFilePath, _, _ := runtime.Caller(0)
	content := input.ReadInput(currentFilePath)

	floor := make([][]int, 0)
	visited := make([][]bool, 0)

	for _, rowStr := range strings.Split(string(content), "\n") {
		row := make([]int, 0)
		visitedRow := make([]bool, 0)
		for _, char := range rowStr {
			num, _ := strconv.Atoi(string(char))
			row = append(row, num)
			visitedRow = append(visitedRow, false)
		}
		floor = append(floor, row)
		visited = append(visited, visitedRow)
	}

	maxRow := len(floor)
	maxCol := len(floor[0])
	basins := make([]int, 0)
	minPoints := make([]point, 0)

	for row := 0; row < maxRow; row++ {
		for col := 0; col < maxCol; col++ {
			curPointMin := true

			// top
			if row-1 >= 0 && floor[row-1][col] <= floor[row][col] {
				curPointMin = false
			}
			// right
			if curPointMin && col+1 < maxCol && floor[row][col+1] <= floor[row][col] {
				curPointMin = false
			}
			// bottom
			if curPointMin && row+1 < maxRow && floor[row+1][col] <= floor[row][col] {
				curPointMin = false
			}
			// left
			if curPointMin && col-1 >= 0 && floor[row][col-1] <= floor[row][col] {
				curPointMin = false
			}

			if curPointMin {
				minPoints = append(minPoints, point{row, col})
			}
		}
	}

	for _, minPoint := range minPoints {
		queue := make([]point, 0)
		queue = append(queue, minPoint)
		basinSize := 1
		visited[minPoint.row][minPoint.col] = true
		for {
			if len(queue) == 0 {
				break
			}
			pt := queue[0]
			queue = queue[1:]

			//fmt.Printf("%d %d: %d -> ", pt.row, pt.col, floor[pt.row][pt.col])
			// top
			if pt.row-1 >= 0 && floor[pt.row-1][pt.col] != 9 && !visited[pt.row-1][pt.col] && floor[pt.row-1][pt.col] > floor[pt.row][pt.col] {
				queue = append(queue, point{pt.row - 1, pt.col})
				visited[pt.row-1][pt.col] = true
				basinSize += 1
			}
			// right
			if pt.col+1 < maxCol && floor[pt.row][pt.col+1] != 9 && !visited[pt.row][pt.col+1] && floor[pt.row][pt.col+1] > floor[pt.row][pt.col] {
				queue = append(queue, point{pt.row, pt.col + 1})
				visited[pt.row][pt.col+1] = true
				basinSize += 1
			}
			// bottom
			if pt.row+1 < maxRow && floor[pt.row+1][pt.col] != 9 && !visited[pt.row+1][pt.col] && floor[pt.row+1][pt.col] > floor[pt.row][pt.col] {
				queue = append(queue, point{pt.row + 1, pt.col})
				visited[pt.row+1][pt.col] = true
				basinSize += 1
			}
			// left
			if pt.col-1 >= 0 && floor[pt.row][pt.col-1] != 9 && !visited[pt.row][pt.col-1] && floor[pt.row][pt.col-1] > floor[pt.row][pt.col] {
				queue = append(queue, point{pt.row, pt.col - 1})
				visited[pt.row][pt.col-1] = true
				basinSize += 1
			}
		}
		//fmt.Printf("\n")

		basins = append(basins, basinSize)
	}

	sort.Slice(basins, func(i, j int) bool {
		return basins[i] > basins[j]
	})

	fmt.Printf("%d", basins[0]*basins[1]*basins[2])

}
