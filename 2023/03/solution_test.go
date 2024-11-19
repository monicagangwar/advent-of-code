package _3

import (
	_ "embed"
	"strings"
	"testing"
)

//go:embed input.txt
var input string

func TestSolution(t *testing.T) {
	type test struct {
		name            string
		input           string
		expectedPartOne int64
		expectedPartTwo int64
	}

	tests := []test{
		{
			name:            "with sample",
			input:           "467..114..\n...*......\n..35..633.\n......#...\n617*......\n.....+.58.\n..592.....\n......755.\n...$.*....\n.664.598..",
			expectedPartOne: 4361,
			expectedPartTwo: 467835,
		}, {
			name:            "with large input",
			input:           input,
			expectedPartOne: ***REMOVED***,
			expectedPartTwo: ***REMOVED***,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			parts, maxHeight, maxWidth := parseInput(tst.input)
			if tst.expectedPartOne != -1 {
				if got := partOne(parts, maxHeight, maxWidth); got != tst.expectedPartOne {
					t.Errorf("partOne() = %v, want %v", got, tst.expectedPartOne)
				}
			}

			if tst.expectedPartTwo != -1 {
				if got := partTwo(parts, maxHeight, maxWidth); got != tst.expectedPartTwo {
					t.Errorf("partTwo() = %v, want %v", got, tst.expectedPartTwo)
				}
			}
		})

	}

}

func parseInput(input string) ([][]rune, int, int) {
	lines := strings.Split(input, "\n")
	parts := make([][]rune, 0, len(lines))
	maxHeight := len(lines)
	maxWidth := 0
	for _, line := range lines {
		parts = append(parts, []rune(line))
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}

	//for i := 0; i < maxHeight; i++ {
	//	for j := 0; j < maxWidth; j++ {
	//		fmt.Print(string(parts[i][j]))
	//	}
	//	fmt.Println()
	//}

	return parts, maxHeight, maxWidth
}

func getNeighborCoordinates(i, j int) [8][2]int {
	return [8][2]int{
		{i - 1, j - 1},
		{i - 1, j},
		{i - 1, j + 1},
		{i, j - 1},
		{i, j + 1},
		{i + 1, j - 1},
		{i + 1, j},
		{i + 1, j + 1},
	}
}

func constructPartNumber(parts [][]rune, visited [][]bool, i, j, maxHeight, maxWidth int) int {
	if i < 0 || i >= maxHeight || j < 0 || j >= maxWidth || visited[i][j] || parts[i][j] < '0' || parts[i][j] > '9' {
		return -1
	}
	for {
		if j >= 0 && parts[i][j] >= '0' && parts[i][j] <= '9' && !visited[i][j] {
			j--
		} else {
			break
		}
	}

	j = j + 1
	num := -1

	for {
		if j >= maxWidth || parts[i][j] < '0' || parts[i][j] > '9' || visited[i][j] {
			break
		}
		visited[i][j] = true
		digit := int(parts[i][j] - '0')
		if num == -1 {
			num = digit
		} else {
			num = num*10 + digit
		}
		j++
	}
	return num
}

func isSymbol(c rune) bool {
	if c == '.' {
		return false
	}
	if c >= '0' && c <= '9' {
		return false
	}
	return true
}

func partOne(parts [][]rune, maxHeight, maxWidth int) int64 {
	visited := make([][]bool, maxHeight)
	for i := range visited {
		visited[i] = make([]bool, maxWidth)
	}

	partsSum := int64(0)

	for i := 0; i < maxHeight; i++ {
		for j := 0; j < maxWidth; j++ {
			if isSymbol(parts[i][j]) {
				neighborCoordinates := getNeighborCoordinates(i, j)
				for _, coordinates := range neighborCoordinates {
					x, y := coordinates[0], coordinates[1]
					num := constructPartNumber(parts, visited, x, y, maxHeight, maxWidth)
					if num != -1 {
						//fmt.Println(num)
						partsSum += int64(num)
					}

				}
			}
		}
	}
	return partsSum
}

func partTwo(parts [][]rune, maxHeight, maxWidth int) int64 {
	visited := make([][]bool, maxHeight)
	for i := range visited {
		visited[i] = make([]bool, maxWidth)
	}

	gearRatioSum := int64(0)

	for i := 0; i < maxHeight; i++ {
		for j := 0; j < maxWidth; j++ {
			if parts[i][j] == '*' {
				nums := make([]int, 0)
				neighborCoordinates := getNeighborCoordinates(i, j)
				for _, coordinates := range neighborCoordinates {
					x, y := coordinates[0], coordinates[1]
					num := constructPartNumber(parts, visited, x, y, maxHeight, maxWidth)
					if num != -1 {
						nums = append(nums, num)
					}
				}
				if len(nums) == 2 {
					gearRatioSum += int64(nums[0] * nums[1])
				}
			}
		}
	}
	return gearRatioSum
}
