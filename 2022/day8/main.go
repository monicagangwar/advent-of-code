package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"

	"github.com/monicagangwar/advent-of-code/input"
)

func main() {
	_, currentFilePath, _, _ := runtime.Caller(0)
	content := input.ReadInput(currentFilePath)
	lines := strings.Split(string(content), "\n")
	height := len(lines)
	width := len(lines[0])
	trees := make([][]int, height)
	visibility := make([][]bool, height)
	for rowIdx, line := range lines {
		treeRow := make([]int, width)
		visibleRow := make([]bool, width)
		for colIdx, char := range line {
			treeHeight, _ := strconv.ParseInt(string(char), 10, 32)
			treeRow[colIdx] = int(treeHeight)
		}
		trees[rowIdx] = treeRow
		visibility[rowIdx] = visibleRow
	}
	partOne(height, width, trees, visibility)
	partTwo(height, width, trees)
}

func partOne(height int, width int, trees [][]int, visibility [][]bool) {

	for rowIdx := 0; rowIdx < height; rowIdx++ {
		max := -1
		for colIdx := 0; colIdx < width; colIdx++ {
			if trees[rowIdx][colIdx] > max {
				visibility[rowIdx][colIdx] = true
				max = trees[rowIdx][colIdx]
			}
		}
		max = -1
		for colIdx := width - 1; colIdx >= 0; colIdx-- {
			if trees[rowIdx][colIdx] > max {
				visibility[rowIdx][colIdx] = true
				max = trees[rowIdx][colIdx]
			}
		}
	}

	for colIdx := 0; colIdx < width; colIdx++ {
		max := -1
		for rowIdx := 0; rowIdx < height; rowIdx++ {
			if trees[rowIdx][colIdx] > max {
				visibility[rowIdx][colIdx] = true
				max = trees[rowIdx][colIdx]
			}
		}

		max = -1
		for rowIdx := height - 1; rowIdx >= 0; rowIdx-- {
			if trees[rowIdx][colIdx] > max {
				visibility[rowIdx][colIdx] = true
				max = trees[rowIdx][colIdx]
			}
		}
	}

	visibleCount := 0

	for rowIdx := 0; rowIdx < height; rowIdx++ {
		for colIdx := 0; colIdx < width; colIdx++ {
			if visibility[rowIdx][colIdx] {
				visibleCount++
			}
		}
	}

	fmt.Println(visibleCount)
}

func partTwo(height int, width int, trees [][]int) {
	scenicScore := 0
	for rowIdx := 0; rowIdx < height; rowIdx++ {
		for colIdx := 0; colIdx < width; colIdx++ {
			scenicScoreLeft := 0
			for left := colIdx - 1; left >= 0; left-- {
				scenicScoreLeft++
				if trees[rowIdx][left] >= trees[rowIdx][colIdx] {
					break
				}
			}
			scenicScoreRight := 0
			for right := colIdx + 1; right < width; right++ {
				scenicScoreRight++
				if trees[rowIdx][right] >= trees[rowIdx][colIdx] {
					break
				}
			}

			scenicScoreTop := 0
			for top := rowIdx - 1; top >= 0; top-- {
				scenicScoreTop++
				if trees[top][colIdx] >= trees[rowIdx][colIdx] {
					break
				}
			}
			scenicScoreBottom := 0
			for bottom := rowIdx + 1; bottom < height; bottom++ {
				scenicScoreBottom++
				if trees[bottom][colIdx] >= trees[rowIdx][colIdx] {
					break
				}
			}

			treeScenicScore := scenicScoreLeft * scenicScoreRight * scenicScoreTop * scenicScoreBottom
			if treeScenicScore > scenicScore {
				scenicScore = treeScenicScore
			}
		}
	}
	fmt.Println(scenicScore)
}
