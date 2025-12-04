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
			expectedPartOne: 13,
			expectedPartTwo: 43,
		}, {
			name:            "with large input",
			input:           input,
			expectedPartOne: 0,
			expectedPartTwo: 0,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			grid := parseInput(tst.input)
			if tst.expectedPartOne != -1 {
				if got := partOne(grid); got != tst.expectedPartOne {
					t.Errorf("partOne() = %v, want %v", got, tst.expectedPartOne)
				}
			}

			if tst.expectedPartTwo != -1 {
				if got := partTwo(grid); got != tst.expectedPartTwo {
					t.Errorf("partTwo() = %v, want %v", got, tst.expectedPartTwo)
				}
			}
		})

	}

}

func parseInput(input string) [][]bool {
	lines := strings.Split(input, "\n")
	grid := make([][]bool, len(lines))
	for i, line := range lines {
		grid[i] = make([]bool, len(line))
		for j, r := range line {
			if r == '@' {
				grid[i][j] = true
			} else {
				grid[i][j] = false
			}
		}
	}
	return grid
}

func partOne(grid [][]bool) int {
	rollsCanBeAccessed := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if canBeAccessed(grid, i, j) {
				rollsCanBeAccessed++
			}
		}
	}
	return rollsCanBeAccessed
}

func partTwo(grid [][]bool) int {
	rollsCanBeAccessed := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if canBeAccessed(grid, i, j) {
				rollsCanBeAccessed++
				grid[i][j] = false
				i = 0
				j = 0
			}
		}
	}
	return rollsCanBeAccessed
}

func canBeAccessed(grid [][]bool, i, j int) bool {
	if !grid[i][j] {
		return false
	}
	countNeighbours := 0
	lenGridX := len(grid[i])
	lenGridY := len(grid)
	for nI := i - 1; nI <= i+1; nI++ {
		for nJ := j - 1; nJ <= j+1; nJ++ {
			if nI == i && nJ == j {
				continue
			}
			if nI >= 0 && nI < lenGridX && nJ >= 0 && nJ < lenGridY && grid[nI][nJ] {
				countNeighbours++

			}
		}
	}

	return countNeighbours < 4
}
