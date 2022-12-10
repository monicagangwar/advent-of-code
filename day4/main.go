package main

import (
	"fmt"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"github.com/monicagangwar/advent-of-code-2022/input"
)

func main() {
	_, currentFilePath, _, _ := runtime.Caller(0)
	content := input.ReadInput(currentFilePath)
	lines := strings.Split(string(content), "\n")
	coordinatesRegex := regexp.MustCompile(`^(\d+)-(\d+),(\d+)-(\d+)$`)
	completeOverlaps := 0
	partialOverlaps := 0
	for _, line := range lines {
		matches := coordinatesRegex.FindAllStringSubmatch(line, -1)
		points := make([]int64, 0)
		for _, point := range matches[0][1:] {
			convertedPoint, _ := strconv.ParseInt(point, 10, 32)
			points = append(points, convertedPoint)
		}
		completeOverlaps += isCompleteOverlap(points[0], points[1], points[2], points[3])
		partialOverlaps += isPartialOverlap(points[0], points[1], points[2], points[3])
	}
	fmt.Println(completeOverlaps)
	fmt.Println(partialOverlaps)
}

func isCompleteOverlap(x1 int64, x2 int64, y1 int64, y2 int64) int {
	if x1 <= y1 && y2 <= x2 {
		return 1
	}
	if y1 <= x1 && x2 <= y2 {
		return 1
	}
	return 0
}
func isPartialOverlap(x1 int64, x2 int64, y1 int64, y2 int64) int {
	if x1 <= y1 && y1 <= x2 {
		return 1
	}
	if x1 <= y2 && y2 <= x2 {
		return 1
	}
	if y1 <= x1 && x1 <= y2 {
		return 1
	}
	if y1 <= x2 && x2 <= y2 {
		return 1
	}
	return 0
}
