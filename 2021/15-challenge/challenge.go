package main

import (
	"bufio"
	"fmt"
	"runtime"
	"strconv"

	"github.com/golang-collections/go-datastructures/queue"

	"github.com/monicagangwar/advent-of-code/input"
)

func main() {
	grid := getGrid()
	fmt.Printf("%d\n", computeMinPath(grid))
	grid = increaseSizeBy(grid, 5)
	fmt.Printf("%d\n", computeMinPath(grid))
}

type node struct {
	row  int
	col  int
	risk int
}

func (n node) Compare(other queue.Item) int {
	otherNode := other.(node)

	if n.risk < otherNode.risk {
		return -1
	}
	return 1
}

func computeMinPath(grid [][]int) int {
	gridSize := len(grid)
	prioQ := queue.NewPriorityQueue(1000)
	visited := map[[2]int]bool{}
	prioQ.Put(node{0, 0, 0})

	for !prioQ.Empty() {
		items, _ := prioQ.Get(1)
		top := items[0].(node)
		coords := [2]int{top.row, top.col}

		if top.row == gridSize-1 && top.col == gridSize-1 {
			return top.risk
		}

		if visited[coords] {
			continue
		}
		visited[coords] = true

		//fmt.Printf("%d %d : %d\n", top.row, top.col, top.risk)

		for _, neighbor := range [][2]int{
			{0, 1},
			{1, 0},
			{0, -1},
			{-1, 0},
		} {
			nr, nc := top.row+neighbor[0], top.col+neighbor[1]

			if nr >= 0 && nr < gridSize && nc >= 0 && nc < gridSize {
				prioQ.Put(node{nr, nc, top.risk + grid[nr][nc]})
			}
		}
	}

	return 0

}

func getGrid() [][]int {
	_, currentFilePath, _, _ := runtime.Caller(0)
	file := input.GetFileMarker(currentFilePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	grid := make([][]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, 0)
		for _, char := range line {
			num, _ := strconv.Atoi(string(char))
			row = append(row, num)
		}
		grid = append(grid, row)
	}
	return grid
}

func increaseSizeBy(grid [][]int, increaseBy int) [][]int {
	gridSize := len(grid)

	for idx := 1; idx < increaseBy; idx++ {
		newGrid := make([][]int, 0)
		for row := 0; row < gridSize; row++ {
			newGridRow := make([]int, 0)
			for col := 0; col < gridSize; col++ {
				newNum := grid[row][col] + idx
				if newNum > 9 {
					newNum = newNum % 9
				}
				newGridRow = append(newGridRow, newNum)
			}
			newGrid = append(newGrid, newGridRow)
		}

		for row := 0; row < gridSize*(idx+1); row++ {
			newGridRow := row % gridSize
			if row < len(grid) {
				grid[row] = append(grid[row], newGrid[newGridRow]...)
			} else {
				grid = append(grid, newGrid[newGridRow])
			}
		}
	}

	for idx := 1; idx < increaseBy; idx++ {
		newGrid := make([][]int, 0)
		for row := 0; row < gridSize; row++ {
			newGridRow := make([]int, 0)
			for col := 0; col < gridSize; col++ {
				newNum := grid[row][col] + idx + increaseBy - 1
				if newNum > 9 {
					newNum = newNum % 9
				}
				newGridRow = append(newGridRow, newNum)
			}
			newGrid = append(newGrid, newGridRow)
		}

		for row := gridSize * idx; row < len(grid); row++ {
			grid[row] = append(grid[row], newGrid[row%gridSize]...)
		}
	}

	return grid
}
