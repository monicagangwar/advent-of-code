package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"

	"github.com/monicagangwar/advent-of-code/input"
)

func main() {
	one()
	two()
}

type point struct {
	row int
	col int
}

func one() {
	_, currentFilePath, _, _ := runtime.Caller(0)
	content := input.ReadInput(currentFilePath)

	octopus := make([][]int, 0)

	for _, rowStr := range strings.Split(string(content), "\n") {
		octopusRow := make([]int, 0)
		for _, char := range rowStr {
			num, _ := strconv.Atoi(string(char))
			octopusRow = append(octopusRow, num)
		}
		octopus = append(octopus, octopusRow)
	}

	flashes := 0
	pos9 := make([]point, 0)

	for step := 1; step <= 100; step++ {
		for row := 0; row <= 9; row++ {
			for col := 0; col <= 9; col++ {
				octopus[row][col] += 1
				if octopus[row][col] > 9 {
					pos9 = append(pos9, point{row, col})
				}
			}
		}

		for {
			//fmt.Printf("step: %d, flashes: %d, pos9: %+v \n", step, flashes, pos9)

			if len(pos9) == 0 {
				break
			}

			pt := pos9[0]
			pos9 = pos9[1:]
			octopus[pt.row][pt.col] = 0
			flashes += 1

			for row := -1; row <= 1; row++ {
				for col := -1; col <= 1; col++ {
					newPt := point{pt.row + row, pt.col + col}
					if newPt.row >= 0 && newPt.row <= 9 && newPt.col >= 0 && newPt.col <= 9 {
						if octopus[newPt.row][newPt.col] != 0 && octopus[newPt.row][newPt.col] <= 9 {
							octopus[newPt.row][newPt.col] += 1
							if octopus[newPt.row][newPt.col] > 9 {
								pos9 = append(pos9, newPt)
							}
						}
					}
				}
			}
		}

		//for row := 0; row <= 9; row++ {
		//	for col := 0; col <= 9; col++ {
		//		fmt.Printf("%d", octopus[row][col])
		//	}
		//	fmt.Printf("\n")
		//}
	}

	fmt.Printf("%d\n", flashes)

}

func two() {
	_, currentFilePath, _, _ := runtime.Caller(0)
	content := input.ReadInput(currentFilePath)

	octopus := make([][]int, 0)

	for _, rowStr := range strings.Split(string(content), "\n") {
		octopusRow := make([]int, 0)
		for _, char := range rowStr {
			num, _ := strconv.Atoi(string(char))
			octopusRow = append(octopusRow, num)
		}
		octopus = append(octopus, octopusRow)
	}

	flashes := 0
	pos9 := make([]point, 0)

	for step := 1; step <= 600; step++ {
		for row := 0; row <= 9; row++ {
			for col := 0; col <= 9; col++ {
				octopus[row][col] += 1
				if octopus[row][col] > 9 {
					pos9 = append(pos9, point{row, col})
				}
			}
		}

		allFlash := 0

		for {
			//fmt.Printf("step: %d, flashes: %d, pos9: %+v \n", step, flashes, pos9)

			if len(pos9) == 0 {
				break
			}

			pt := pos9[0]
			pos9 = pos9[1:]
			octopus[pt.row][pt.col] = 0
			flashes += 1
			allFlash += 1

			for row := -1; row <= 1; row++ {
				for col := -1; col <= 1; col++ {
					newPt := point{pt.row + row, pt.col + col}
					if newPt.row >= 0 && newPt.row <= 9 && newPt.col >= 0 && newPt.col <= 9 {
						if octopus[newPt.row][newPt.col] != 0 && octopus[newPt.row][newPt.col] <= 9 {
							octopus[newPt.row][newPt.col] += 1
							if octopus[newPt.row][newPt.col] > 9 {
								pos9 = append(pos9, newPt)
							}
						}
					}
				}
			}
		}

		//fmt.Printf("step : %d, allFlash: %d\n", step, allFlash)

		if allFlash == 100 {
			fmt.Printf("%d\n", step)
			return
		}

		//for row := 0; row <= 9; row++ {
		//	for col := 0; col <= 9; col++ {
		//		fmt.Printf("%d", octopus[row][col])
		//	}
		//	fmt.Printf("\n")
		//}
	}
}
