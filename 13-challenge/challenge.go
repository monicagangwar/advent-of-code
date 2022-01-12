package main

import (
	"bufio"
	"fmt"
	"runtime"
	"strconv"
	"strings"

	"github.com/monicagangwar/advent-of-code-2021/input"
)

func main() {
	one()
	two()
}

func foldHorizontal(x int, graph [][]bool, maxRow int, maxCol int) {
	for delta := 1; delta != -1; delta++ {
		top := x - delta
		bottom := x + delta
		if top < 0 || bottom > maxRow {
			delta = -1
			break
		}
		//fmt.Printf("merging row: %d with row: %d\n", top, bottom)
		for col := 0; col <= maxCol; col++ {
			if graph[top][col] || graph[bottom][col] {
				graph[top][col] = true
			} else {
				graph[top][col] = false
			}
		}
	}
}

func foldVertical(y int, graph [][]bool, maxRow int, maxCol int) {
	for delta := 1; delta != -1; delta++ {
		left := y - delta
		right := y + delta
		if left < 0 || right > maxCol {
			delta = -1
			break
		}
		for row := 0; row <= maxRow; row++ {
			if graph[row][left] || graph[row][right] {
				graph[row][left] = true
			} else {
				graph[row][left] = false
			}
		}
	}
}

func printGraph(graph [][]bool, maxRow int, maxCol int) {
	for row := 0; row <= maxRow; row++ {
		for col := 0; col <= maxCol; col++ {
			if graph[row][col] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}

func one() {
	_, currentFilePath, _, _ := runtime.Caller(0)
	file := input.GetFileMarker(currentFilePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	maxRow := 0
	maxCol := 0

	graph := make([][]bool, 0)
	for row := 0; row <= 1500; row++ {
		graphRow := make([]bool, 0)
		for col := 0; col <= 1500; col++ {
			graphRow = append(graphRow, false)
		}
		graph = append(graph, graphRow)
	}

	foldInstructions := false
	instructionImplementedOnce := false

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			foldInstructions = true

			//fmt.Printf("original\n")
			//printGraph(graph, maxRow, maxCol)
			continue
		}
		if !foldInstructions {
			vertices := strings.Split(line, ",")

			col, _ := strconv.Atoi(vertices[0])
			if col > maxCol {
				maxCol = col
			}
			row, _ := strconv.Atoi(vertices[1])
			if row > maxRow {
				maxRow = row
			}

			graph[row][col] = true
		} else if !instructionImplementedOnce {
			instructionImplementedOnce = true
			//fmt.Printf("%s\n", line)
			line = strings.Replace(line, "fold along ", "", -1)
			instruction := strings.Split(line, "=")
			if instruction[0] == "y" {
				num, _ := strconv.Atoi(instruction[1])
				foldHorizontal(num, graph, maxRow, maxCol)
				maxRow = num - 1
			} else {
				num, _ := strconv.Atoi(instruction[1])
				foldVertical(num, graph, maxRow, maxCol)
				maxCol = num - 1
			}
			//printGraph(graph, maxRow, maxCol)
			break

		}
	}
	count := 0

	for row := 0; row <= maxRow; row++ {
		for col := 0; col <= maxCol; col++ {
			if graph[row][col] == true {
				count += 1
			}
		}
	}

	fmt.Printf("%d\n", count)
}

func two() {
	_, currentFilePath, _, _ := runtime.Caller(0)
	file := input.GetFileMarker(currentFilePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	maxRow := 0
	maxCol := 0

	graph := make([][]bool, 0)
	for row := 0; row <= 1500; row++ {
		graphRow := make([]bool, 0)
		for col := 0; col <= 1500; col++ {
			graphRow = append(graphRow, false)
		}
		graph = append(graph, graphRow)
	}

	foldInstructions := false

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			foldInstructions = true

			//fmt.Printf("original\n")
			//printGraph(graph, maxRow, maxCol)
			continue
		}
		if !foldInstructions {
			vertices := strings.Split(line, ",")

			col, _ := strconv.Atoi(vertices[0])
			if col > maxCol {
				maxCol = col
			}
			row, _ := strconv.Atoi(vertices[1])
			if row > maxRow {
				maxRow = row
			}

			graph[row][col] = true
		} else {
			//fmt.Printf("%s\n", line)
			line = strings.Replace(line, "fold along ", "", -1)
			instruction := strings.Split(line, "=")
			if instruction[0] == "y" {
				num, _ := strconv.Atoi(instruction[1])
				foldHorizontal(num, graph, maxRow, maxCol)
				maxRow = num - 1
			} else {
				num, _ := strconv.Atoi(instruction[1])
				foldVertical(num, graph, maxRow, maxCol)
				maxCol = num - 1
			}
			//printGraph(graph, maxRow, maxCol)
		}
	}
	// RLBCJGLU
	printGraph(graph, maxRow, maxCol)
	//count := 0
	//
	//for row := 0; row <= maxRow; row++ {
	//	for col := 0; col <= maxCol; col++ {
	//		if graph[row][col] == true {
	//			count += 1
	//		}
	//	}
	//}

	//fmt.Printf("%d\n", count)
}
