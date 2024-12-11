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
			expectedPartOne: 36,
			expectedPartTwo: 81,
		}, {
			name:            "with large input",
			input:           input,
			expectedPartOne: 0,
			expectedPartTwo: 0,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			heightMap := parseInput(tst.input)
			if tst.expectedPartOne != -1 {
				if got := partOne(heightMap); got != tst.expectedPartOne {
					t.Errorf("partOne() = %v, want %v", got, tst.expectedPartOne)
				}
			}

			if tst.expectedPartTwo != -1 {
				if got := partTwo(heightMap); got != tst.expectedPartTwo {
					t.Errorf("partTwo() = %v, want %v", got, tst.expectedPartTwo)
				}
			}
		})

	}

}

func parseInput(input string) [][]int {
	lines := strings.Split(input, "\n")
	heightMap := make([][]int, len(lines))
	for i, line := range lines {
		heightMap[i] = make([]int, len(line))
		for j, char := range line {
			heightMap[i][j] = int(char - '0')
		}
	}
	return heightMap
}

func partOne(heightMap [][]int) int {
	trailCountSum := 0
	for i := 0; i < len(heightMap); i++ {
		for j := 0; j < len(heightMap[i]); j++ {
			if heightMap[i][j] == 0 {
				visited := make([][]bool, len(heightMap))
				for p := 0; p < len(visited); p++ {
					visited[p] = make([]bool, len(heightMap[p]))
				}

				markVisited(heightMap, i, j, visited)

				for p := 0; p < len(visited); p++ {
					for q := 0; q < len(visited[p]); q++ {
						if visited[p][q] {
							trailCountSum++
						}
					}
				}
			}
		}
	}

	return trailCountSum
}

func partTwo(heightMap [][]int) int {
	trailCountSum := 0
	for i := 0; i < len(heightMap); i++ {
		for j := 0; j < len(heightMap[i]); j++ {
			if heightMap[i][j] == 0 {
				trailCountSum += getDistinctTrailCount(heightMap, i, j)
			}
		}
	}

	return trailCountSum
}

func markVisited(heightMap [][]int, i, j int, visited [][]bool) {
	if heightMap[i][j] == 9 {
		visited[i][j] = true
		return
	}

	for _, neighbour := range getNeighbours(i, j) {
		ni, nj := neighbour[0], neighbour[1]
		if ni >= 0 && ni < len(heightMap) && nj >= 0 && nj < len(heightMap[ni]) && heightMap[ni][nj]-heightMap[i][j] == 1 {
			markVisited(heightMap, ni, nj, visited)
		}
	}
	return
}

func getDistinctTrailCount(heightMap [][]int, i, j int) int {
	if heightMap[i][j] == 9 {
		return 1
	}
	count := 0
	for _, neighbour := range getNeighbours(i, j) {
		ni, nj := neighbour[0], neighbour[1]
		if ni >= 0 && ni < len(heightMap) && nj >= 0 && nj < len(heightMap[ni]) && heightMap[ni][nj]-heightMap[i][j] == 1 {
			count += getDistinctTrailCount(heightMap, ni, nj)
		}
	}
	return count
}

func getNeighbours(i, j int) [4][2]int {
	return [4][2]int{
		{i - 1, j},
		{i + 1, j},
		{i, j - 1},
		{i, j + 1},
	}
}
