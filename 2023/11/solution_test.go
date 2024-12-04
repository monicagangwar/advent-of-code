package _1

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

func TestSolution(t *testing.T) {
	type test struct {
		name            string
		input           string
		distance        int
		expectedPartOne int
	}

	tests := []test{
		{
			name:            "with sample with 2",
			input:           sample,
			distance:        2,
			expectedPartOne: 374,
		}, {
			name:            "with sample with 10",
			input:           sample,
			distance:        10,
			expectedPartOne: 1030,
		}, {
			name:            "with sample with 100",
			input:           sample,
			distance:        100,
			expectedPartOne: 8410,
		}, {
			name:            "with large input 2",
			input:           input,
			distance:        2,
			expectedPartOne: 0,
		}, {
			name:            "with large input 1000000",
			input:           input,
			distance:        1000000,
			expectedPartOne: 0,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			galaxyCoordinates := parseInput(tst.input, tst.distance)
			if tst.expectedPartOne != -1 {
				if got := findDistance(galaxyCoordinates); got != tst.expectedPartOne {
					t.Errorf("partOne() = %v, want %v", got, tst.expectedPartOne)
				}
			}
		})

	}

}

func parseInput(input string, distance int) [][2]int {
	lines := strings.Split(input, "\n")

	emptyRows := make([]int, 0)
	for i, line := range lines {
		emptyRow := true
		for _, char := range line {
			if char == '#' {
				emptyRow = false
			}
		}
		if emptyRow {
			emptyRows = append(emptyRows, i)
		}
	}

	colLen := len(lines[0])

	emptyColumns := make([]int, 0)
	for j := 0; j < colLen; j++ {
		emptyCol := true
		for i := 0; i < len(lines); i++ {
			if lines[i][j] == '#' {
				emptyCol = false
			}
		}
		if emptyCol {
			emptyColumns = append(emptyColumns, j)
		}
	}

	galaxyCoordinates := make([][2]int, 0)
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] == '#' {
				galaxyCoordinates = append(galaxyCoordinates, [2]int{i, j})
			}
		}
	}

	for i := 0; i < len(emptyRows); i++ {
		emptyRow := emptyRows[i]
		for i := 0; i < len(galaxyCoordinates); i++ {
			if galaxyCoordinates[i][0] > emptyRow {
				galaxyCoordinates[i][0] += distance - 1
			}
		}
		for j := i + 1; j < len(emptyRows); j++ {
			emptyRows[j] += distance - 1
		}

	}
	for i := 0; i < len(emptyColumns); i++ {
		emptyColumn := emptyColumns[i]
		for i := 0; i < len(galaxyCoordinates); i++ {
			if galaxyCoordinates[i][1] > emptyColumn {
				galaxyCoordinates[i][1] += distance - 1
			}
		}
		for j := i + 1; j < len(emptyColumns); j++ {
			emptyColumns[j] += distance - 1
		}
	}

	//fmt.Println(galaxyCoordinates)

	return galaxyCoordinates
}

func findDistance(galaxyCoordinates [][2]int) int {
	sum := 0

	//fmt.Println(galaxyCoordinates)

	for i := 0; i < len(galaxyCoordinates); i++ {
		x1 := galaxyCoordinates[i][1]
		y1 := galaxyCoordinates[i][0]
		for j := i + 1; j < len(galaxyCoordinates); j++ {
			x2 := galaxyCoordinates[j][1]
			y2 := galaxyCoordinates[j][0]

			sum += int(math.Abs(float64(x2-x1)) + math.Abs(float64(y2-y1)))

			//fmt.Printf("%d: -> %d = %d\n", i, j, int(math.Abs(float64(x2-x1))+math.Abs(float64(y2-y1))))
		}
	}
	return sum
}
