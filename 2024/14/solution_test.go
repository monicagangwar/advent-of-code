package main

import (
	_ "embed"
	"fmt"
	tm "github.com/buger/goterm"
	"regexp"
	"strconv"
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
		space           [2]int
		timeIntSeconds  int
	}

	tests := []test{
		{
			name:            "with sample",
			input:           sample,
			expectedPartOne: 12,
			expectedPartTwo: -1,
			space:           [2]int{11, 7},
			timeIntSeconds:  100,
		}, {
			name:            "with large input",
			input:           input,
			expectedPartOne: 0,
			expectedPartTwo: 0,
			space:           [2]int{101, 103},
			timeIntSeconds:  100,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			robots := parseInput(tst.input)
			if tst.expectedPartOne != -1 {
				if got := partOne(robots, tst.space, tst.timeIntSeconds); got != tst.expectedPartOne {
					t.Errorf("partOne() = %v, want %v", got, tst.expectedPartOne)
				}
			}

			if tst.expectedPartTwo != -1 {
				if got := partTwo(robots, tst.space); got != tst.expectedPartTwo {
					t.Errorf("partTwo() = %v, want %v", got, tst.expectedPartTwo)
				}
			}
		})

	}

}

type Robot struct {
	Pos      [2]int
	Velocity [2]int
}

func parseInput(input string) []Robot {
	lines := strings.Split(input, "\n")
	robots := make([]Robot, len(lines))
	robotRe := regexp.MustCompile(`p=(?P<posx>\d+),(?P<posy>\d+) v=(?P<xNeg>-)*(?P<velx>\d+),(?P<yNeg>-)*(?P<vely>\d+)`)
	for i, line := range lines {
		matches := robotRe.FindStringSubmatch(line)
		robots[i].Pos[0], _ = strconv.Atoi(matches[robotRe.SubexpIndex("posx")])
		robots[i].Pos[1], _ = strconv.Atoi(matches[robotRe.SubexpIndex("posy")])
		robots[i].Velocity[0], _ = strconv.Atoi(matches[robotRe.SubexpIndex("velx")])
		robots[i].Velocity[1], _ = strconv.Atoi(matches[robotRe.SubexpIndex("vely")])
		if matches[robotRe.SubexpIndex("xNeg")] == "-" {
			robots[i].Velocity[0] *= -1
		}
		if matches[robotRe.SubexpIndex("yNeg")] == "-" {
			robots[i].Velocity[1] *= -1
		}
	}
	//fmt.Println(robots)
	return robots
}

func partOne(robots []Robot, space [2]int, timeIntSeconds int) int {

	well := make([][]int, space[0])
	for i := range well {
		well[i] = make([]int, space[1])
	}

	quadRobotCount := make([]int, 4)

	for _, robot := range robots {

		robot.Pos[0] += timeIntSeconds * robot.Velocity[0]
		robot.Pos[1] += timeIntSeconds * robot.Velocity[1]

		robot.Pos[0] = ((robot.Pos[0] % space[0]) + space[0]) % space[0]
		robot.Pos[1] = ((robot.Pos[1] % space[1]) + space[1]) % space[1]

		quad := getQuad(robot.Pos, space)
		if quad != -1 {
			quadRobotCount[getQuad(robot.Pos, space)]++
		}

		well[robot.Pos[0]][robot.Pos[1]]++
	}

	//fmt.Println(quadRobotCount)

	safetyFactor := 1
	for _, count := range quadRobotCount {
		safetyFactor *= count
	}

	return safetyFactor
}

func partTwo(robots []Robot, space [2]int) int {
	well := make([][]int, space[0])
	for i := range well {
		well[i] = make([]int, space[1])
	}
	timeIntSeconds := 1
	for {
		for i, _ := range robots {
			robot := robots[i]
			if well[robot.Pos[0]][robot.Pos[1]] > 0 {
				well[robot.Pos[0]][robot.Pos[1]]--
			}
			robot.Pos[0] += robot.Velocity[0]
			robot.Pos[1] += robot.Velocity[1]

			robot.Pos[0] = ((robot.Pos[0] % space[0]) + space[0]) % space[0]
			robot.Pos[1] = ((robot.Pos[1] % space[1]) + space[1]) % space[1]

			robots[i] = robot

			well[robot.Pos[0]][robot.Pos[1]]++
		}
		if easterEggFound(well) {
			break
		}

		timeIntSeconds++
	}

	printEasterEgg(timeIntSeconds, well)

	return 0
}

func easterEggFound(well [][]int) bool {
	for i := 0; i < len(well[0]); i++ {
		countContRow := 0
		for j := 0; j < len(well); j++ {
			if well[j][i] > 0 {
				countContRow++
			} else {
				countContRow = 0
			}
			if countContRow >= 20 {
				return true
			}
		}
	}

	for i := 0; i < len(well); i++ {
		countContCol := 0
		for j := 0; j < len(well[0]); j++ {
			if well[i][j] > 0 {
				countContCol++
			} else {
				countContCol = 0
			}
			if countContCol >= 20 {
				return true
			}
		}
	}

	return false
}

func printEasterEgg(timeIntSeconds int, well [][]int) {
	//tm.Clear()
	//tm.MoveCursor(1, 1)
	fmt.Println("Time:", timeIntSeconds)
	fmt.Println()
	for i := 0; i < len(well[0]); i++ {
		for j := 0; j < len(well); j++ {
			if well[j][i] > 0 {
				fmt.Print(tm.Color("o", tm.BLUE))
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	//tm.Flush()
}

func printWell(well [][]int) {
	for i := 0; i < len(well[0]); i++ {
		for j := 0; j < len(well); j++ {
			if well[j][i] > 0 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func getQuad(pos, space [2]int) int {
	if pos[0] < space[0]/2 && pos[1] < space[1]/2 {
		return 0
	}
	if pos[0] > space[0]/2 && pos[1] < space[1]/2 {
		return 1
	}
	if pos[0] < space[0]/2 && pos[1] > space[1]/2 {
		return 2
	}
	if pos[0] > space[0]/2 && pos[1] > space[1]/2 {
		return 3
	}
	return -1
}
