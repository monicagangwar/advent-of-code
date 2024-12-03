package _0

import (
	_ "embed"
	"math"
	"strings"
	"testing"
)

//go:embed input.txt
var input string

//go:embed sample.txt
var sample string

//go:embed sample2.txt
var sample2 string

//go:embed sample3.txt
var sample3 string

//go:embed sample4.txt
var sample4 string

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
			expectedPartOne: 8,
			expectedPartTwo: -1,
		}, {
			name:            "with sample 2",
			input:           sample2,
			expectedPartOne: -1,
			expectedPartTwo: 4,
		}, {
			name:            "with sample 3",
			input:           sample3,
			expectedPartOne: -1,
			expectedPartTwo: 8,
		}, {
			name:            "with sample 4",
			input:           sample4,
			expectedPartOne: -1,
			expectedPartTwo: 10,
		}, {
			name:            "with large input",
			input:           input,
			expectedPartOne: 0,
			expectedPartTwo: 0,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			ground := parseInput(tst.input)
			if tst.expectedPartOne != -1 {
				if got := partOne(ground); got != tst.expectedPartOne {
					t.Errorf("partOne() = %v, want %v", got, tst.expectedPartOne)
				}
			}

			if tst.expectedPartTwo != -1 {
				if got := partTwo(ground); got != tst.expectedPartTwo {
					t.Errorf("partTwo() = %v, want %v", got, tst.expectedPartTwo)
				}
			}
		})

	}

}

func partOne(ground []string) int {
	startingPos := [2]int{0, 0}
	for i := 0; i < len(ground); i++ {
		for j := 0; j < len(ground[i]); j++ {
			if ground[i][j] == 'S' {
				startingPos = [2]int{i, j}
				break
			}
		}
	}

	visited := make([][]bool, 0)

	for i := 0; i < len(ground); i++ {
		visitedRow := make([]bool, 0)
		for j := 0; j < len(ground[i]); j++ {
			visitedRow = append(visitedRow, false)
		}
		visited = append(visited, visitedRow)
	}

	loopSize, _ := findMaxLoopSize(ground, startingPos, visited, 0, nil)
	return loopSize / 2
}

func partTwo(ground []string) int {
	startingPos := [2]int{0, 0}
	for i := 0; i < len(ground); i++ {
		for j := 0; j < len(ground[i]); j++ {
			if ground[i][j] == 'S' {
				startingPos = [2]int{i, j}
				break
			}
		}
	}

	visited := make([][]bool, 0)

	for i := 0; i < len(ground); i++ {
		visitedRow := make([]bool, 0)
		for j := 0; j < len(ground[i]); j++ {
			visitedRow = append(visitedRow, false)
		}
		visited = append(visited, visitedRow)
	}

	loopSize, pathVisited := findMaxLoopSize(ground, startingPos, visited, 0, [][2]int{})

	area := 0
	for i := 0; i < len(pathVisited); i++ {
		nextIdx := (i + 1) % len(pathVisited)
		x1 := pathVisited[i][0]
		y1 := pathVisited[i][1]
		x2 := pathVisited[nextIdx][0]
		y2 := pathVisited[nextIdx][1]
		area += (x1 * y2) - (x2 * y1)
	}

	area = int(math.Abs(float64(area)))
	return (area / 2) - (loopSize / 2) + 1
}

func findMaxLoopSize(ground []string, currentPos [2]int, visited [][]bool, loopLen int, path [][2]int) (int, [][2]int) {
	curChar := ground[currentPos[0]][currentPos[1]]

	if curChar == 'S' && loopLen > 0 {
		return loopLen, path
	}

	if curChar == '.' {
		return -1, nil
	}

	visited[currentPos[0]][currentPos[1]] = true
	path = append(path, [2]int{currentPos[1], currentPos[0]})
	var maxPath [][2]int
	neighbours := getNeighbours(curChar, currentPos)
	maxLoopLen := -1
	for _, neighbour := range neighbours {
		canGoToNeighbour := canGoTo(ground, currentPos, neighbour, visited)
		if canGoToNeighbour {
			newVisited := copyVisited(visited)
			newPath := copyPath(path)
			newLoopLen, pathVisited := findMaxLoopSize(ground, neighbour, newVisited, loopLen+1, newPath)
			if newLoopLen > maxLoopLen {
				maxLoopLen = newLoopLen
				maxPath = pathVisited
			}
		}
	}

	return maxLoopLen, maxPath
}

func copyVisited(visited [][]bool) [][]bool {
	newVisited := make([][]bool, 0)
	for i := 0; i < len(visited); i++ {
		newVisitedRow := make([]bool, 0)
		for j := 0; j < len(visited[i]); j++ {
			newVisitedRow = append(newVisitedRow, visited[i][j])
		}
		newVisited = append(newVisited, newVisitedRow)
	}
	return newVisited
}

func copyPath(path [][2]int) [][2]int {
	newPath := make([][2]int, 0)
	for i := 0; i < len(path); i++ {
		newPath = append(newPath, path[i])
	}
	return newPath
}

func canGoTo(ground []string, curPos, neighbour [2]int, visited [][]bool) bool {
	rowLen := len(ground)
	colLen := len(ground[0])
	ni := neighbour[0]
	nj := neighbour[1]

	if ni < 0 || ni >= rowLen || nj < 0 || nj >= colLen {
		return false
	}

	curChar := ground[curPos[0]][curPos[1]]
	nextChar := ground[ni][nj]

	if nextChar == 'S' && visited[ni][nj] {
		return true
	}
	if visited[ni][nj] {
		return false
	}

	if curChar == 'S' {
		direction := getDirection(curPos, neighbour)
		if direction == "up" && !(nextChar == '7' || nextChar == 'F' || nextChar == '|') {
			return false
		}
		if direction == "down" && !(nextChar == 'L' || nextChar == 'J' || nextChar == '|') {
			return false
		}
		if direction == "left" && !(nextChar == 'F' || nextChar == 'L' || nextChar == '-') {
			return false
		}
		if direction == "right" && !(nextChar == '7' || nextChar == 'J' || nextChar == '-') {
			return false
		}
	}
	return true
}

func getDirection(curPos, neighbour [2]int) string {
	if curPos[0] == neighbour[0] {
		if curPos[1] > neighbour[1] {
			return "left"
		} else {
			return "right"
		}
	} else {
		if curPos[0] > neighbour[0] {
			return "up"
		} else {
			return "down"
		}
	}
}

func getNeighbours(curChar byte, curPos [2]int) [][2]int {
	i := curPos[0]
	j := curPos[1]

	if curChar == 'S' {
		return [][2]int{
			{i - 1, j},
			{i + 1, j},
			{i, j - 1},
			{i, j + 1}}
	}

	if curChar == '|' {
		return [][2]int{{i - 1, j}, {i + 1, j}}
	}
	if curChar == '-' {
		return [][2]int{{i, j - 1}, {i, j + 1}}
	}
	if curChar == 'L' {
		return [][2]int{{i - 1, j}, {i, j + 1}}
	}
	if curChar == 'J' {
		return [][2]int{{i - 1, j}, {i, j - 1}}
	}
	if curChar == '7' {
		return [][2]int{{i + 1, j}, {i, j - 1}}
	}
	if curChar == 'F' {
		return [][2]int{{i + 1, j}, {i, j + 1}}
	}
	return [][2]int{}
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}
