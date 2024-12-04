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
			expectedPartOne: 18,
			expectedPartTwo: 9,
		}, {
			name:            "with large input",
			input:           input,
			expectedPartOne: 0,
			expectedPartTwo: 0,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			lines := parseInput(tst.input)
			if tst.expectedPartOne != -1 {
				if got := partOne(lines); got != tst.expectedPartOne {
					t.Errorf("partOne() = %v, want %v", got, tst.expectedPartOne)
				}
			}

			if tst.expectedPartTwo != -1 {
				if got := partTwo(lines); got != tst.expectedPartTwo {
					t.Errorf("partTwo() = %v, want %v", got, tst.expectedPartTwo)
				}
			}
		})

	}

}

func parseInput(input string) [][]byte {
	lines := strings.Split(input, "\n")
	parsedInput := make([][]byte, len(lines))
	for i, line := range lines {
		parsedInput[i] = []byte(line)
		for j, char := range line {
			parsedInput[i][j] = byte(char)
		}
	}
	return parsedInput
}

func partOne(lines [][]byte) int {
	rowLen := len(lines)
	colLen := len(lines[0])
	totalWords := 0
	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			if lines[i][j] == 'X' {
				neighbors := getNeighbors(i, j)
				for _, neighbor := range neighbors {
					word := "X"
					for _, n := range neighbor {
						if n[0] < 0 || n[0] >= rowLen || n[1] < 0 || n[1] >= colLen {
							break
						}
						word = fmt.Sprintf("%s%c", word, lines[n[0]][n[1]])
					}
					if word == "XMAS" {
						totalWords++
					}
				}
			}
		}
	}
	return totalWords
}

func partTwo(lines [][]byte) int {
	rowLen := len(lines)
	colLen := len(lines[0])
	patterns := getPatterns()
	totalWords := 0
	//found := make([][]bool, rowLen)
	//for i := 0; i < rowLen; i++ {
	//	found[i] = make([]bool, colLen)
	//}

	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			if lines[i][j] == 'M' || lines[i][j] == 'S' {
				for _, pattern := range patterns {
					patternFound := true
					for p := 0; p < len(pattern); p++ {
						for q := 0; q < len(pattern[p]); q++ {
							newi := i + p
							newj := j + q
							if newi < 0 || newi >= rowLen || newj < 0 || newj >= colLen {
								patternFound = false
								break
							}
							if pattern[p][q] != '.' && lines[newi][newj] != pattern[p][q] {
								patternFound = false
								break
							}
						}
						if !patternFound {
							break
						}
					}

					if patternFound {
						//for p := 0; p < len(pattern); p++ {
						//	for q := 0; q < len(pattern[p]); q++ {
						//		newi := i + p
						//		newj := j + q
						//		if pattern[p][q] != '.' {
						//			found[newi][newj] = true
						//		}
						//	}
						//}
						totalWords++
						break
					}
				}
			}
		}
	}
	//for i := 0; i < rowLen; i++ {
	//	for j := 0; j < colLen; j++ {
	//		if found[i][j] {
	//			fmt.Print("1 ")
	//		} else {
	//			fmt.Print("0 ")
	//		}
	//	}
	//	fmt.Println()
	//}
	return totalWords
}

func getNeighbors(i, j int) [][][2]int {
	return [][][2]int{
		{{i, j - 1}, {i, j - 2}, {i, j - 3}},
		{{i, j + 1}, {i, j + 2}, {i, j + 3}},
		{{i - 1, j}, {i - 2, j}, {i - 3, j}},
		{{i + 1, j}, {i + 2, j}, {i + 3, j}},
		{{i - 1, j - 1}, {i - 2, j - 2}, {i - 3, j - 3}},
		{{i - 1, j + 1}, {i - 2, j + 2}, {i - 3, j + 3}},
		{{i + 1, j - 1}, {i + 2, j - 2}, {i + 3, j - 3}},
		{{i + 1, j + 1}, {i + 2, j + 2}, {i + 3, j + 3}},
	}
}

func getPatterns() [][][3]byte {
	return [][][3]byte{
		{
			{'M', '.', 'S'},
			{'.', 'A', '.'},
			{'M', '.', 'S'},
		}, {
			{'S', '.', 'M'},
			{'.', 'A', '.'},
			{'S', '.', 'M'},
		},
		{
			{'M', '.', 'M'},
			{'.', 'A', '.'},
			{'S', '.', 'S'},
		}, {
			{'S', '.', 'S'},
			{'.', 'A', '.'},
			{'M', '.', 'M'},
		},
	}
}
