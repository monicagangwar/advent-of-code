package main

import (
	"bufio"
	"fmt"
	"runtime"
	"strconv"
	"strings"

	"github.com/monicagangwar/advent-of-code-2021/input"
)

func main() {
	one()
	two()
}

type plotPoint struct {
	row       int
	col       int
	diagonalL int
	diagonalR int
}

type pointData struct {
	start plotPoint
	stop  plotPoint
	sum   plotPoint
}

func parsePoint(strPoint string) plotPoint {
	pt := strings.Split(strPoint, ",")
	row, _ := strconv.Atoi(pt[0])
	col, _ := strconv.Atoi(pt[1])
	return plotPoint{row: row, col: col}
}

func swap(lp *plotPoint, rp *plotPoint) {
	temp := *lp
	*lp = *rp
	*rp = temp
}

func one() {
	_, currentFilePath, _, _ := runtime.Caller(0)
	file := input.GetFileMarker(currentFilePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	plot := make([][]pointData, 0)

	for row := 0; row <= 1000; row++ {
		plotRow := make([]pointData, 0)
		for col := 0; col <= 1000; col++ {
			plotRow = append(plotRow, pointData{})
		}
		plot = append(plot, plotRow)
	}

	maxRow := 0
	maxCol := 0

	for scanner.Scan() {
		line := scanner.Text()
		points := strings.Split(line, " ")
		leftPoint := parsePoint(points[0])
		rightPoint := parsePoint(points[2])

		if leftPoint.row == rightPoint.row {
			if rightPoint.col < leftPoint.col {
				swap(&leftPoint, &rightPoint)
			}

			plot[leftPoint.row][leftPoint.col].start.row += 1
			plot[rightPoint.row][rightPoint.col].stop.row += 1
		}

		if leftPoint.col == rightPoint.col {
			if rightPoint.row < leftPoint.row {
				swap(&leftPoint, &rightPoint)
			}

			plot[leftPoint.row][leftPoint.col].start.col += 1
			plot[rightPoint.row][rightPoint.col].stop.col += 1
		}

		if leftPoint.row >= rightPoint.row && leftPoint.row >= maxRow {
			maxRow = leftPoint.row
		}
		if rightPoint.row >= leftPoint.row && rightPoint.row >= maxRow {
			maxRow = rightPoint.row
		}
		if leftPoint.col >= rightPoint.col && leftPoint.col >= maxCol {
			maxCol = leftPoint.col
		}
		if rightPoint.col >= leftPoint.col && rightPoint.col >= maxCol {
			maxCol = rightPoint.col
		}

	}

	intersection := 0

	for row := 0; row <= maxRow; row++ {
		for col := 0; col <= maxCol; col++ {

			point := plot[row][col]

			if row-1 >= 0 {
				up := plot[row-1][col]

				if up.sum.col-up.stop.col > 0 {
					point.sum.col += up.sum.col - up.stop.col
				} else {
					point.sum.col = 0
				}
			}

			if col-1 >= 0 {
				left := plot[row][col-1]

				if left.sum.row-left.stop.row > 0 {
					point.sum.row += left.sum.row - left.stop.row
				} else {
					point.sum.row = 0
				}
			}

			point.sum.row += point.start.row
			point.sum.col += point.start.col

			plot[row][col] = point

			if point.sum.row+point.sum.col > 1 {
				intersection += 1
			}

		}
	}

	fmt.Printf("%d\n", intersection)
}

func two() {
	_, currentFilePath, _, _ := runtime.Caller(0)
	file := input.GetFileMarker(currentFilePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	plot := make([][]pointData, 0)

	for row := 0; row <= 1000; row++ {
		plotRow := make([]pointData, 0)
		for col := 0; col <= 1000; col++ {
			plotRow = append(plotRow, pointData{})
		}
		plot = append(plot, plotRow)
	}

	maxRow := 0
	maxCol := 0

	for scanner.Scan() {
		line := scanner.Text()
		points := strings.Split(line, " ")
		leftPoint := parsePoint(points[0])
		rightPoint := parsePoint(points[2])

		if leftPoint.row == rightPoint.row {
			if rightPoint.col < leftPoint.col {
				swap(&leftPoint, &rightPoint)
			}

			plot[leftPoint.row][leftPoint.col].start.row += 1
			plot[rightPoint.row][rightPoint.col].stop.row += 1
		} else if leftPoint.col == rightPoint.col {
			if rightPoint.row < leftPoint.row {
				swap(&leftPoint, &rightPoint)
			}

			plot[leftPoint.row][leftPoint.col].start.col += 1
			plot[rightPoint.row][rightPoint.col].stop.col += 1
		} else if leftPoint.row > rightPoint.row && leftPoint.col > rightPoint.col {

			plot[rightPoint.row][rightPoint.col].start.diagonalR += 1
			plot[leftPoint.row][leftPoint.col].stop.diagonalR += 1

		} else if leftPoint.row > rightPoint.row && leftPoint.col < rightPoint.col {

			plot[rightPoint.row][rightPoint.col].start.diagonalL += 1
			plot[leftPoint.row][leftPoint.col].stop.diagonalL += 1

		} else if leftPoint.row < rightPoint.row && leftPoint.col > rightPoint.col {

			plot[leftPoint.row][leftPoint.col].start.diagonalL += 1
			plot[rightPoint.row][rightPoint.col].stop.diagonalL += 1

		} else if leftPoint.row < rightPoint.row && leftPoint.col < rightPoint.col {

			plot[leftPoint.row][leftPoint.col].start.diagonalR += 1
			plot[rightPoint.row][rightPoint.col].stop.diagonalR += 1

		}

		if leftPoint.row >= rightPoint.row && leftPoint.row >= maxRow {
			maxRow = leftPoint.row
		}
		if rightPoint.row >= leftPoint.row && rightPoint.row >= maxRow {
			maxRow = rightPoint.row
		}
		if leftPoint.col >= rightPoint.col && leftPoint.col >= maxCol {
			maxCol = leftPoint.col
		}
		if rightPoint.col >= leftPoint.col && rightPoint.col >= maxCol {
			maxCol = rightPoint.col
		}

	}

	intersection := 0

	for row := 0; row <= maxRow; row++ {
		for col := 0; col <= maxCol; col++ {

			point := plot[row][col]

			if row-1 >= 0 {
				up := plot[row-1][col]

				if up.sum.col-up.stop.col > 0 {
					point.sum.col += up.sum.col - up.stop.col
				} else {
					point.sum.col = 0
				}
			}

			if col-1 >= 0 {
				left := plot[row][col-1]

				if left.sum.row-left.stop.row > 0 {
					point.sum.row += left.sum.row - left.stop.row
				} else {
					point.sum.row = 0
				}
			}

			if row-1 >= 0 && col-1 >= 0 {
				diagonalR := plot[row-1][col-1]

				if diagonalR.sum.diagonalR-diagonalR.stop.diagonalR > 0 {
					point.sum.diagonalR += diagonalR.sum.diagonalR - diagonalR.stop.diagonalR
				} else {
					point.sum.diagonalR = 0
				}
			}

			if row-1 >= 0 && col+1 <= maxCol {
				diagonalL := plot[row-1][col+1]

				if diagonalL.sum.diagonalL-diagonalL.stop.diagonalL > 0 {
					point.sum.diagonalL += diagonalL.sum.diagonalL - diagonalL.stop.diagonalL
				} else {
					point.sum.diagonalL = 0
				}
			}

			point.sum.row += point.start.row
			point.sum.col += point.start.col
			point.sum.diagonalR += point.start.diagonalR
			point.sum.diagonalL += point.start.diagonalL

			plot[row][col] = point

			if point.sum.row+point.sum.col+point.sum.diagonalL+point.sum.diagonalR > 1 {
				intersection += 1
			}

		}
	}

	fmt.Printf("%d\n", intersection)
}
