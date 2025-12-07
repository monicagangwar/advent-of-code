package template

import (
	_ "embed"
	"strings"
	"testing"
)

//go:embed input.txt
var input string

//go:embed sample.txt
var sample string

func TestSolution(t *testing.T) {
	type test struct {
		name            string
		input           string
		expectedPartOne int
		expectedPartTwo int
	}

	tests := []test{
		{
			name:            "with sample",
			input:           sample,
			expectedPartOne: 21,
			expectedPartTwo: 40,
		}, {
			name:            "with large input",
			input:           input,
			expectedPartOne: 0,
			expectedPartTwo: 0,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			if tst.expectedPartOne != -1 {
				if got := partOne(parseInput(tst.input)); got != tst.expectedPartOne {
					t.Errorf("partOne() = %v, want %v", got, tst.expectedPartOne)
				}
			}

			if tst.expectedPartTwo != -1 {
				if got := partTwo(parseInput(tst.input)); got != tst.expectedPartTwo {
					t.Errorf("partTwo() = %v, want %v", got, tst.expectedPartTwo)
				}
			}
		})

	}

}

func parseInput(input string) [][]rune {
	lines := strings.Split(input, "\n")
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	return grid
}

func partOne(grid [][]rune) int {
	visited := make([][]bool, len(grid))
	nodes := make([][2]int, 0)
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, len(grid[i]))
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 'S' {
				nodes = append(nodes, [2]int{i, j})
			}
		}
	}

	countSplitters := 0
	nodesIdx := 0
	for nodesIdx < len(nodes) {
		x, y := nodes[nodesIdx][0], nodes[nodesIdx][1]
		nodesIdx++

		if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[x]) {
			continue
		}
		if visited[x][y] {
			continue
		}
		visited[x][y] = true
		if grid[x][y] == 'S' || grid[x][y] == '.' {
			nodes = append(nodes, [2]int{x + 1, y})
		} else if grid[x][y] == '^' {
			countSplitters++
			nodes = append(nodes, [2]int{x + 1, y - 1})
			nodes = append(nodes, [2]int{x + 1, y + 1})
		}

	}
	return countSplitters
}

func partTwo(grid [][]rune) int {
	count := make([][]int, len(grid))
	var startNode [2]int
	for i := 0; i < len(grid); i++ {
		count[i] = make([]int, len(grid[i]))
		for j := 0; j < len(grid[i]); j++ {
			count[i][j] = -1
		}
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 'S' {
				startNode = [2]int{i, j}
			}
		}
	}
	return countTimelines(grid, startNode, count)
}

func countTimelines(grid [][]rune, startNode [2]int, count [][]int) int {
	x, y := startNode[0], startNode[1]
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[x]) {
		return 1
	}
	if count[x][y] != -1 {
		return count[x][y]
	}
	if grid[x][y] == 'S' || grid[x][y] == '.' {
		count[x][y] = countTimelines(grid, [2]int{x + 1, y}, count)
	} else if grid[x][y] == '^' {
		count[x][y] = countTimelines(grid, [2]int{x + 1, y - 1}, count) + countTimelines(grid, [2]int{x + 1, y + 1}, count)
	}
	return count[x][y]
}
