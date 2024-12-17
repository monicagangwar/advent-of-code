package main

import (
	_ "embed"
	"fmt"
	tm "github.com/buger/goterm"
	"strings"
	"testing"
)

//go:embed input.txt
var input string

//go:embed sample.txt
var sample string

//go:embed sample2.txt
var sample2 string

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
			expectedPartOne: 10092,
			expectedPartTwo: 9021,
		}, {
			name:            "with sample 2",
			input:           sample2,
			expectedPartOne: 2028,
			expectedPartTwo: -1,
		}, {
			name:            "with large input",
			input:           input,
			expectedPartOne: 0,
			expectedPartTwo: 0,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			warehouse, instruction, expandedWarehouse := parseInput(tst.input)
			if tst.expectedPartOne != -1 {
				if got := partOne(warehouse, instruction); got != tst.expectedPartOne {
					t.Errorf("partOne() = %v, want %v", got, tst.expectedPartOne)
				}
			}

			if tst.expectedPartTwo != -1 {
				if got := partTwo(expandedWarehouse, instruction); got != tst.expectedPartTwo {
					t.Errorf("partTwo() = %v, want %v", got, tst.expectedPartTwo)
				}
			}
		})

	}

}

func parseInput(input string) ([][]byte, string, [][]byte) {
	lines := strings.Split(input, "\n")
	warehouse := make([][]byte, 0)
	instruction := ""
	warehouseComplete := false
	for _, line := range lines {
		if line == "" {
			warehouseComplete = true
		}
		if warehouseComplete {
			instruction = fmt.Sprintf("%s%s", instruction, line)
		} else {
			warehouse = append(warehouse, []byte(line))
		}
	}
	expandedInWarehouse := make([][]byte, 0)
	for i := 0; i < len(warehouse); i++ {
		whRow := ""
		for j := 0; j < len(warehouse[i]); j++ {
			switch warehouse[i][j] {
			case '#':
				whRow = fmt.Sprintf("%s##", whRow)
				break
			case '.':
				whRow = fmt.Sprintf("%s..", whRow)
				break
			case 'O':
				whRow = fmt.Sprintf("%s[]", whRow)
				break
			case '@':
				whRow = fmt.Sprintf("%s@.", whRow)
			}
		}
		expandedInWarehouse = append(expandedInWarehouse, []byte(whRow))
	}

	return warehouse, instruction, expandedInWarehouse
}

func partOne(warehouse [][]byte, instruction string) int {
	robotPos := [2]int{}
	for i := 0; i < len(warehouse); i++ {
		for j := 0; j < len(warehouse[i]); j++ {
			if warehouse[i][j] == '@' {
				robotPos = [2]int{i, j}
			}
		}
	}

	for _, ins := range []byte(instruction) {
		robotPos = moveRobot(ins, robotPos, warehouse)
		//printWarehouse(warehouse)
	}

	sumBoxCoord := 0

	for i := 0; i < len(warehouse); i++ {
		for j := 0; j < len(warehouse[i]); j++ {
			if warehouse[i][j] == 'O' {
				sumBoxCoord += (100 * i) + j
			}
		}
	}
	return sumBoxCoord
}

func partTwo(warehouse [][]byte, instruction string) int {
	robotPos := [2]int{}
	for i := 0; i < len(warehouse); i++ {
		for j := 0; j < len(warehouse[i]); j++ {
			if warehouse[i][j] == '@' {
				robotPos = [2]int{i, j}
			}
		}
	}

	//printWarehouse(warehouse)

	for _, ins := range []byte(instruction) {
		//insString := string(ins)
		//fmt.Println(insString)
		if strings.Contains("^v", string(ins)) {
			newWh := copyWarehouse(warehouse)
			newRobotPos := moveRobotPartTwo(ins, robotPos, robotPos, newWh)
			if newRobotPos[0] != -1 {
				robotPos = newRobotPos
				warehouse = newWh
			}
		} else {
			robotPos = moveRobot(ins, robotPos, warehouse)
		}
		//printWarehouse(warehouse)
		//time.Sleep(1 * time.Second)
	}

	//printWarehouse(warehouse)

	sumBoxCoord := 0

	for i := 0; i < len(warehouse); i++ {
		for j := 0; j < len(warehouse[i]); j++ {
			if warehouse[i][j] == '[' {
				sumBoxCoord += (100 * i) + j
			}
		}
	}
	return sumBoxCoord

}

func copyWarehouse(warehouse [][]byte) [][]byte {
	newWarehouse := make([][]byte, 0)
	for i := 0; i < len(warehouse); i++ {
		newWarehouse = append(newWarehouse, make([]byte, len(warehouse[i])))
		copy(newWarehouse[i], warehouse[i])
	}
	return newWarehouse
}

func moveRobotPartTwo(ins byte, curPos1 [2]int, curPos2 [2]int, warehouse [][]byte) [2]int {

	if curPos1 == curPos2 {
		nextPos := getCoordinate(curPos1, ins, false)
		if warehouse[nextPos[0]][nextPos[1]] == '#' {
			return [2]int{-1, -1}
		}
		if warehouse[nextPos[0]][nextPos[1]] == '[' {
			newPos := moveRobotPartTwo(ins, nextPos, [2]int{nextPos[0], nextPos[1] + 1}, warehouse)
			if newPos[0] == -1 {
				return [2]int{-1, -1}
			}
		}
		if warehouse[nextPos[0]][nextPos[1]] == ']' {
			newPos := moveRobotPartTwo(ins, [2]int{nextPos[0], nextPos[1] - 1}, nextPos, warehouse)
			if newPos[0] == -1 {
				return [2]int{-1, -1}
			}
		}

		warehouse[nextPos[0]][nextPos[1]] = warehouse[curPos1[0]][curPos1[1]]
		warehouse[curPos1[0]][curPos1[1]] = '.'
		return nextPos
	}

	nextPos1 := getCoordinate(curPos1, ins, false)
	nextPos2 := getCoordinate(curPos2, ins, false)

	if warehouse[nextPos1[0]][nextPos1[1]] == '#' || warehouse[nextPos2[0]][nextPos2[1]] == '#' {
		return [2]int{-1, -1}
	}
	if warehouse[nextPos1[0]][nextPos1[1]] == '[' {
		newPos := moveRobotPartTwo(ins, nextPos1, nextPos2, warehouse)
		if newPos[0] == -1 {
			return [2]int{-1, -1}
		}
	}
	if warehouse[nextPos1[0]][nextPos1[1]] == ']' {
		newPos := moveRobotPartTwo(ins, [2]int{nextPos1[0], nextPos1[1] - 1}, nextPos1, warehouse)
		if newPos[0] == -1 {
			return [2]int{-1, -1}
		}
	}
	if warehouse[nextPos2[0]][nextPos2[1]] == '[' {
		newPos := moveRobotPartTwo(ins, nextPos2, [2]int{nextPos2[0], nextPos2[1] + 1}, warehouse)
		if newPos[0] == -1 {
			return [2]int{-1, -1}
		}
	}

	warehouse[nextPos1[0]][nextPos1[1]] = warehouse[curPos1[0]][curPos1[1]]
	warehouse[curPos1[0]][curPos1[1]] = '.'

	warehouse[nextPos2[0]][nextPos2[1]] = warehouse[curPos2[0]][curPos2[1]]
	warehouse[curPos2[0]][curPos2[1]] = '.'

	return curPos1
}

func moveRobot(ins byte, robotPos [2]int, warehouse [][]byte) [2]int {
	curPos := robotPos
	for {
		newPos := getCoordinate(curPos, ins, false)
		if warehouse[newPos[0]][newPos[1]] == '#' {
			return robotPos
		} else if warehouse[newPos[0]][newPos[1]] == '.' {
			curPos = newPos
			break
		}
		curPos = newPos
	}

	for {
		nextPos := getCoordinate(curPos, ins, true)
		warehouse[curPos[0]][curPos[1]] = warehouse[nextPos[0]][nextPos[1]]

		if nextPos == robotPos {
			warehouse[robotPos[0]][robotPos[1]] = '.'
			warehouse[curPos[0]][curPos[1]] = '@'
			return curPos
		}
		curPos = nextPos
	}
}

func getCoordinate(curPos [2]int, dir byte, opposite bool) [2]int {
	if opposite {
		switch dir {
		case '<':
			dir = '>'
			break
		case '>':
			dir = '<'
			break
		case '^':
			dir = 'v'
			break
		case 'v':
			dir = '^'
			break
		}
	}

	switch dir {
	case '<':
		return [2]int{curPos[0], curPos[1] - 1}
	case '>':
		return [2]int{curPos[0], curPos[1] + 1}
	case '^':
		return [2]int{curPos[0] - 1, curPos[1]}
	case 'v':
		return [2]int{curPos[0] + 1, curPos[1]}
	}
	return curPos
}

func printWarehouse(warehouse [][]byte) {
	fmt.Printf("\x1b[2J")
	// from top
	//tm.MoveCursor(0, 6)
	fmt.Println()
	for _, row := range warehouse {
		for _, c := range row {
			if c == '@' {
				fmt.Printf(tm.Color("@", tm.RED))
			} else {
				fmt.Printf(string(c))
			}
		}
		fmt.Println()
	}
}
