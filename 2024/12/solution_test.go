package template

import (
	_ "embed"
	"sort"
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

//go:embed sample5.txt
var sample5 string

func TestSolution(t *testing.T) {
	type test struct {
		name            string
		input           string
		expectedPartOne int
		expectedPartTwo int
	}

	tests := []test{
		{
			name:            "with sample 1",
			input:           sample,
			expectedPartOne: 140,
			expectedPartTwo: 80,
		}, {
			name:            "with sample 2",
			input:           sample2,
			expectedPartOne: 772,
			expectedPartTwo: 436,
		}, {
			name:            "with sample 3",
			input:           sample3,
			expectedPartOne: 1930,
			expectedPartTwo: 1206,
		}, {
			name:            "with sample 4",
			input:           sample4,
			expectedPartOne: -1,
			expectedPartTwo: 236,
		}, {
			name:            "with sample 5",
			input:           sample5,
			expectedPartOne: -1,
			expectedPartTwo: 368,
		}, {
			name:            "with large input",
			input:           input,
			expectedPartOne: 0,
			expectedPartTwo: 0,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			farm := parseInput(tst.input)
			if tst.expectedPartOne != -1 {
				if got := partOne(farm); got != tst.expectedPartOne {
					t.Errorf("partOne() = %v, want %v", got, tst.expectedPartOne)
				}
			}

			if tst.expectedPartTwo != -1 {
				if got := partTwo(farm); got != tst.expectedPartTwo {
					t.Errorf("partTwo() = %v, want %v", got, tst.expectedPartTwo)
				}
			}
		})

	}

}

func parseInput(input string) [][]byte {
	lines := strings.Split(input, "\n")
	farm := make([][]byte, len(lines))
	for i, line := range lines {
		farm[i] = []byte(line)
	}
	return farm
}

func dfs(farm [][]byte, visited [][]bool, i, j int, sides *[][3]int) (int, int) {

	visited[i][j] = true
	neighbors := getNeighbors(i, j)
	perimeter := 0
	area := 1
	for _, n := range neighbors {
		ni, nj := n[0], n[1]
		if ni < 0 || ni >= len(farm) || nj < 0 || nj >= len(farm[ni]) || farm[ni][nj] != farm[i][j] {
			if sides != nil {
				*sides = append(*sides, [3]int{ni, nj, n[2]})
			}
			perimeter++
		} else if !visited[ni][nj] {
			a, p := dfs(farm, visited, ni, nj, sides)
			area += a
			perimeter += p
		}
	}
	return area, perimeter
}

func getNeighbors(i, j int) [][3]int {
	return [][3]int{
		{i - 1, j, 0},
		{i + 1, j, 1},
		{i, j - 1, 2},
		{i, j + 1, 3},
	}
}

func partOne(farm [][]byte) int {
	visited := make([][]bool, len(farm))
	for i := range visited {
		visited[i] = make([]bool, len(farm[i]))
	}
	sum := 0
	for i := 0; i < len(farm); i++ {
		for j := 0; j < len(farm[i]); j++ {
			if !visited[i][j] {
				area, perimeter := dfs(farm, visited, i, j, nil)
				//fmt.Printf("Char: %s = Area: %d, Perimeter: %d\n", string(farm[i][j]), area, perimeter)
				sum += area * perimeter
			}
		}
	}
	return sum
}

func getSideCount(sides [][3]int) int {
	sideMap := make(map[[3]int]bool)

	sort.Slice(sides, func(i, j int) bool {
		if sides[i][0] == sides[j][0] {
			return sides[i][1] < sides[j][1]
		}
		return sides[i][0] < sides[j][0]
	})

	//fmt.Println(sides)

	sideCount := 0

	for _, s := range sides {
		getCombinations := getNeighbors(s[0], s[1])
		combFound := false

		for _, c := range getCombinations {
			c[2] = s[2]
			if _, found := sideMap[c]; found {
				combFound = true
			}
		}
		if !combFound {
			sideCount++
		}

		sideMap[s] = true

	}

	return sideCount
}

func partTwo(farm [][]byte) int {
	visited := make([][]bool, len(farm))
	for i := range visited {
		visited[i] = make([]bool, len(farm[i]))
	}
	sum := 0
	for i := 0; i < len(farm); i++ {
		for j := 0; j < len(farm[i]); j++ {
			if !visited[i][j] {
				sides := make([][3]int, 0)
				area, _ := dfs(farm, visited, i, j, &sides)
				sideCount := getSideCount(sides)
				//fmt.Printf("Char: %s = Area: %d, Sides: %d\n", string(farm[i][j]), area, sideCount)
				sum += area * sideCount
			}
		}
	}
	return sum
}
