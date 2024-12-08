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
			expectedPartOne: 14,
			expectedPartTwo: 34,
		}, {
			name:            "with large input",
			input:           input,
			expectedPartOne: 0,
			expectedPartTwo: 0,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			world := parseInput(tst.input)
			if tst.expectedPartOne != -1 {
				if got := partOne(world); got != tst.expectedPartOne {
					t.Errorf("partOne() = %v, want %v", got, tst.expectedPartOne)
				}
			}

			if tst.expectedPartTwo != -1 {
				if got := partTwo(world); got != tst.expectedPartTwo {
					t.Errorf("partTwo() = %v, want %v", got, tst.expectedPartTwo)
				}
			}
		})

	}

}

func parseInput(input string) [][]byte {
	lines := strings.Split(input, "\n")
	world := make([][]byte, len(lines))
	for i, line := range lines {
		world[i] = []byte(line)
	}
	return world
}

func partOne(world [][]byte) int {
	rowLen := len(world)
	colLen := len(world[0])
	antennas := make(map[byte][][2]int)
	for i := range rowLen {
		for j := range colLen {
			if world[i][j] != '.' {
				if _, ok := antennas[world[i][j]]; !ok {
					antennas[world[i][j]] = [][2]int{}
				}
				antennas[world[i][j]] = append(antennas[world[i][j]], [2]int{i, j})
			}
		}
	}

	foundAntinodes := make(map[[2]int]struct{})

	for _, antenna := range antennas {
		for i := 0; i < len(antenna); i++ {
			for j := i + 1; j < len(antenna); j++ {
				distx := antenna[j][0] - antenna[i][0]
				disty := antenna[j][1] - antenna[i][1]

				an1 := [2]int{antenna[i][0] - distx, antenna[i][1] - disty}
				an2 := [2]int{antenna[j][0] + distx, antenna[j][1] + disty}

				if inLimits(an1, rowLen, colLen) {
					foundAntinodes[an1] = struct{}{}
				}
				if inLimits(an2, rowLen, colLen) {
					foundAntinodes[an2] = struct{}{}
				}

			}
		}
	}

	return len(foundAntinodes)
}

func inLimits(node [2]int, rowLen, colLen int) bool {
	return node[0] >= 0 && node[0] < rowLen && node[1] >= 0 && node[1] < colLen
}

func partTwo(world [][]byte) int {
	rowLen := len(world)
	colLen := len(world[0])
	antennas := make(map[byte][][2]int)
	foundAntinodes := make(map[[2]int]struct{})
	for i := range rowLen {
		for j := range colLen {
			if world[i][j] != '.' {
				foundAntinodes[[2]int{i, j}] = struct{}{}
				if _, ok := antennas[world[i][j]]; !ok {
					antennas[world[i][j]] = [][2]int{}
				}
				antennas[world[i][j]] = append(antennas[world[i][j]], [2]int{i, j})
			}
		}
	}

	for _, antenna := range antennas {
		for i := 0; i < len(antenna); i++ {
			for j := i + 1; j < len(antenna); j++ {
				distx := antenna[j][0] - antenna[i][0]
				disty := antenna[j][1] - antenna[i][1]
				an1 := antenna[i]
				an2 := antenna[j]
				for {
					an1 = [2]int{an1[0] - distx, an1[1] - disty}
					an2 = [2]int{an2[0] + distx, an2[1] + disty}

					inLimitAn1 := inLimits(an1, rowLen, colLen)
					inLimitAn2 := inLimits(an2, rowLen, colLen)

					canContinue := false

					if inLimitAn1 {
						canContinue = true
						foundAntinodes[an1] = struct{}{}
					}
					if inLimitAn2 {
						canContinue = true
						foundAntinodes[an2] = struct{}{}
					}

					if !canContinue {
						break
					}

				}

			}
		}
	}

	//for i := 0; i < rowLen; i++ {
	//	for j := 0; j < colLen; j++ {
	//		if world[i][j] == '.' {
	//			if _, found := foundAntinodes[[2]int{i, j}]; found {
	//				fmt.Printf("#")
	//			} else {
	//				fmt.Printf(".")
	//			}
	//		} else {
	//			fmt.Printf("%c", world[i][j])
	//		}
	//	}
	//	fmt.Println()
	//}

	return len(foundAntinodes)
}
