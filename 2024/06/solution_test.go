package template

import (
	_ "embed"
	"fmt"
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
			expectedPartOne: 41,
			expectedPartTwo: 6,
		}, {
			name:            "with large input",
			input:           input,
			expectedPartOne: 0,
			expectedPartTwo: 0,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			labMap, guardPos := parseInput(tst.input)
			if tst.expectedPartOne != -1 {
				if got := partOne(labMap, guardPos); got != tst.expectedPartOne {
					t.Errorf("partOne() = %v, want %v", got, tst.expectedPartOne)
				}
			}

			if tst.expectedPartTwo != -1 {
				if got := partTwo(labMap, guardPos); got != tst.expectedPartTwo {
					t.Errorf("partTwo() = %v, want %v", got, tst.expectedPartTwo)
				}
			}
		})

	}

}

func parseInput(input string) ([][]byte, [2]int) {
	lines := strings.Split(input, "\n")
	labMap := make([][]byte, len(lines))
	guardPos := [2]int{}
	for i, line := range lines {
		labMap[i] = []byte(line)
		for j := 0; j < len(labMap[i]); j++ {
			if labMap[i][j] != '.' && labMap[i][j] != '#' {
				guardPos = [2]int{i, j}
			}
		}
	}
	return labMap, guardPos
}

func partOne(labMap [][]byte, guardPos [2]int) int {
	rowLen := len(labMap)
	colLen := len(labMap[0])
	curDir := up
	posVisited := make(map[[2]int]struct{})
	for {
		posVisited[guardPos] = struct{}{}
		nextPos := getNextPos(guardPos, curDir)

		if nextPos[0] < 0 || nextPos[1] < 0 || nextPos[0] >= rowLen || nextPos[1] >= colLen {
			break
		}

		if labMap[nextPos[0]][nextPos[1]] == '#' {
			curDir = getNextDirection(curDir)
			nextPos = getNextPos(guardPos, curDir)
		}
		guardPos = nextPos

	}
	return len(posVisited)
}

func partTwo(labMap [][]byte, guardPos [2]int) int {
	rowLen := len(labMap)
	colLen := len(labMap[0])

	cycleFoundCount := 0

	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			if labMap[i][j] == '.' {
				labMap[i][j] = '#'
				//printMap(labMap)
				cycleFound := findCycle(labMap, guardPos)
				//fmt.Println(cycleFound)
				if cycleFound {
					cycleFoundCount++
				}
				labMap[i][j] = '.'
			}
		}
	}
	return cycleFoundCount
}

func printMap(labMap [][]byte) {
	for _, row := range labMap {
		fmt.Println(string(row))
	}
	fmt.Println()
}

func findCycle(labMap [][]byte, guardPos [2]int) bool {
	visitedRoute := map[[3]int]struct{}{}
	rowLen := len(labMap)
	colLen := len(labMap[0])
	curDir := up
	for {
		nextPos := getNextPos(guardPos, curDir)

		if nextPos[0] < 0 || nextPos[1] < 0 || nextPos[0] >= rowLen || nextPos[1] >= colLen {
			break
		}

		if labMap[nextPos[0]][nextPos[1]] == '#' {
			key := [3]int{int(curDir), guardPos[0], guardPos[1]}

			if _, found := visitedRoute[key]; found {
				return true
			}
			visitedRoute[key] = struct{}{}
			curDir = getNextDirection(curDir)
		} else {
			guardPos = nextPos
		}
	}
	return false
}

func getNextPos(curPos [2]int, curDir direction) [2]int {
	if curDir == up {
		return [2]int{curPos[0] - 1, curPos[1]}
	}
	if curDir == down {
		return [2]int{curPos[0] + 1, curPos[1]}
	}
	if curDir == left {
		return [2]int{curPos[0], curPos[1] - 1}
	}
	if curDir == right {
		return [2]int{curPos[0], curPos[1] + 1}
	}
	return curPos
}

func getNextDirection(curDir direction) direction {
	dirMap := map[direction]direction{
		up:    right,
		right: down,
		down:  left,
		left:  up,
	}
	return dirMap[curDir]
}

type direction int

const (
	up direction = iota
	down
	left
	right
)
