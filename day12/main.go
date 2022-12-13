package main

import (
	"container/heap"
	"fmt"
	"math"
	"runtime"
	"strings"

	"github.com/monicagangwar/advent-of-code-2022/input"
)

func main() {
	_, currentFilePath, _, _ := runtime.Caller(0)
	content := input.ReadInput(currentFilePath)
	lines := strings.Split(string(content), "\n")
	var partOne, partTwo = math.MaxInt32, math.MaxInt32
	for row, line := range lines {
		for col, ch := range line {
			if string(ch) == "S" {
				partOne = compute(row, col, lines)
				//partTwo = partOne
			}
			if string(ch) == "a" {
				output := compute(row, col, lines)
				if output < partTwo {
					partTwo = output
				}
			}
		}
	}
	fmt.Println(partOne)
	fmt.Println(partTwo)
}

func getDiff(ch1 byte, ch2 byte) int {
	if ch1 == 'S' {
		ch1 = 'a'
	}
	if ch1 == 'E' {
		ch1 = 'z'
	}
	if ch2 == 'S' {
		ch2 = 'a'
	}
	if ch2 == 'E' {
		ch2 = 'z'
	}
	//fmt.Printf("\n %s(%d) -> %s(%d) = %d", string(ch1), int(ch1), string(ch2), int(ch2), int(ch2)-int(ch1))
	return int(ch2) - int(ch1)
}

func compute(rowIdx int, colIdx int, lines []string) int {
	height := len(lines)
	width := len(lines[0])
	visited := make([][]bool, 0)
	distance := make([][]int, 0)
	queue := make(PriorityQueue, 0)
	presentInQ := make(map[string]struct{})

	for _, line := range lines {
		visitedRow := make([]bool, len(line))
		distanceRow := make([]int, len(line))
		visited = append(visited, visitedRow)
		distance = append(distance, distanceRow)
	}

	queue.Push(&Item{
		row:      rowIdx,
		col:      colIdx,
		distance: 0,
	})
	presentInQ[fmt.Sprintf("%d,%d", rowIdx, colIdx)] = struct{}{}

	for {
		if len(queue) == 0 {
			break
		}
		item := heap.Pop(&queue).(*Item)
		visited[item.row][item.col] = true
		left, right, top, bottom := item.col-1, item.col+1, item.row-1, item.row+1

		if left >= 0 && left < width && !visited[item.row][left] &&
			getDiff(lines[item.row][item.col], lines[item.row][left]) <= 1 &&
			(distance[item.row][left] == 0 || item.distance+1 < distance[item.row][left]) {
			distance[item.row][left] = item.distance + 1
			if _, found := presentInQ[fmt.Sprintf("%d,%d", item.row, left)]; !found {
				queue.Push(&Item{
					row:      item.row,
					col:      left,
					distance: distance[item.row][left],
				})
			}
		}

		if right >= 0 && right < width && !visited[item.row][right] &&
			getDiff(lines[item.row][item.col], lines[item.row][right]) <= 1 &&
			(distance[item.row][right] == 0 || item.distance+1 < distance[item.row][right]) {
			distance[item.row][right] = item.distance + 1
			if _, found := presentInQ[fmt.Sprintf("%d,%d", item.row, right)]; !found {
				queue.Push(&Item{
					row:      item.row,
					col:      right,
					distance: distance[item.row][right],
				})
			}
		}

		if top >= 0 && top < height && !visited[top][item.col] &&
			getDiff(lines[item.row][item.col], lines[top][item.col]) <= 1 &&
			(distance[top][item.col] == 0 || item.distance+1 < distance[top][item.col]) {
			distance[top][item.col] = item.distance + 1
			if _, found := presentInQ[fmt.Sprintf("%d,%d", top, item.col)]; !found {
				queue.Push(&Item{
					row:      top,
					col:      item.col,
					distance: distance[top][item.col],
				})
			}
		}

		if bottom >= 0 && bottom < height && !visited[bottom][item.col] &&
			getDiff(lines[item.row][item.col], lines[bottom][item.col]) <= 1 &&
			(distance[bottom][item.col] == 0 || item.distance+1 < distance[bottom][item.col]) {
			distance[bottom][item.col] = item.distance + 1
			if _, found := presentInQ[fmt.Sprintf("%d,%d", bottom, item.col)]; !found {
				queue.Push(&Item{
					row:      bottom,
					col:      item.col,
					distance: distance[bottom][item.col],
				})
			}
		}

		//display(distance)
	}

	for row, line := range lines {
		for col, ch := range line {
			if ch == 'E' && distance[row][col] != 0 {
				return distance[row][col]
			}
		}
	}
	return math.MaxInt32
}

func display(distance [][]int) {
	fmt.Println()
	for _, row := range distance {
		for _, dist := range row {
			fmt.Printf("%d ", dist)
		}
		fmt.Println()
	}
}
