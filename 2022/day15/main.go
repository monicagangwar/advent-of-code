package main

import (
	"fmt"
	"math"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"

	"github.com/monicagangwar/advent-of-code/input"
)

const limit = 20

func main() {
	_, currentFilePath, _, _ := runtime.Caller(0)
	content := input.ReadInput(currentFilePath)
	lines := strings.Split(string(content), "\n")
	sensors := make([][2]int, len(lines))
	closestBeacons := make([][2]int, len(lines))
	distance := make([]int, len(lines))
	locationRegex := regexp.MustCompile(`^Sensor at x=(-)?(\d+), y=(-)?(\d+): closest beacon is at x=(-)?(\d+), y=(-)?(\d+)$`)

	for idx, line := range lines {
		matches := locationRegex.FindAllStringSubmatch(line, -1)
		sensors[idx] = [2]int{findNum(matches[0][1], matches[0][2]), findNum(matches[0][3], matches[0][4])}
		closestBeacons[idx] = [2]int{findNum(matches[0][5], matches[0][6]), findNum(matches[0][7], matches[0][8])}
		distance[idx] = calculateDistance(sensors[idx], closestBeacons[idx])
	}
	//fmt.Println(computePointsWhereBeaconCannotBePresent(sensors, closestBeacons, distance, 10))
	fmt.Println(computePointsWhereBeaconCannotBePresent(sensors, closestBeacons, distance, 2000000))

	for y := 0; y <= limit; y++ {
		x := findBeacon(sensors, closestBeacons, distance, y)
		if x != -1 {
			fmt.Println(x*limit + y)
			break
		}
	}
}
func calculateDistance(s, b [2]int) int {
	return abs(s[0]-b[0]) + abs(s[1]-b[1])
}

func computePointsWhereBeaconCannotBePresent(sensors, closestBeacons [][2]int, distance []int, y int) int {
	pointsRange := make([][2]int, 0)
	for idx, sensor := range sensors {
		sensorVerticalRange := [2]int{sensor[1] - distance[idx], sensor[1] + distance[idx]}
		if y >= sensorVerticalRange[0] && y <= sensorVerticalRange[1] {
			distanceFromSensor := abs(sensor[1] - y)
			sensorRangeAtRow := [2]int{sensor[0] - (distance[idx] - distanceFromSensor), sensor[0] + (distance[idx] - distanceFromSensor)}
			if closestBeacons[idx][1] == y && sensorRangeAtRow[0] == closestBeacons[idx][0] {
				sensorRangeAtRow[0]++
			}
			if closestBeacons[idx][1] == y && sensorRangeAtRow[1] == closestBeacons[idx][0] {
				sensorRangeAtRow[1]--
			}
			if sensorRangeAtRow[0] <= sensorRangeAtRow[1] {
				pointsRange = append(pointsRange, sensorRangeAtRow)
			}
		}
	}
	sort.Slice(pointsRange, func(i, j int) bool {
		return pointsRange[i][0] < pointsRange[j][0]
	})
	for idx := 0; idx < len(pointsRange)-1; idx++ {
		point1 := pointsRange[idx]
		point2 := pointsRange[idx+1]
		if (point2[0] <= point1[0] && point1[0] <= point2[1]) ||
			(point2[0] <= point1[1] && point1[1] <= point2[1]) ||
			(point1[0] <= point2[0] && point2[0] <= point1[1]) ||
			(point1[0] <= point2[1] && point2[1] <= point1[1]) {
			pointsRange[idx+1] = [2]int{
				min(point1[0], min(point1[1], min(point2[0], point2[1]))),
				max(point1[0], max(point1[1], max(point2[0], point2[1]))),
			}
			pointsRange[idx] = [2]int{math.MinInt32, math.MinInt32}
		}
	}
	points := 0
	for _, point := range pointsRange {
		if !(point[0] == math.MinInt32 && point[1] == math.MinInt32) {
			fmt.Printf("[%d %d] ", point[0], point[1])
			points += point[1] - point[0] + 1
		}
	}
	return points
}

func findBeacon(sensors, closestBeacons [][2]int, distance []int, y int) int {
	pointsRange := make([][2]int, 0)
	for idx, sensor := range sensors {
		sensorVerticalRange := [2]int{sensor[1] - distance[idx], sensor[1] + distance[idx]}
		if y >= sensorVerticalRange[0] && y <= sensorVerticalRange[1] {
			distanceFromSensor := abs(sensor[1] - y)
			sensorRangeAtRow := [2]int{inLimit(sensor[0] - (distance[idx] - distanceFromSensor)),
				inLimit(sensor[0] + (distance[idx] - distanceFromSensor))}
			if closestBeacons[idx][1] == y && sensorRangeAtRow[0] == closestBeacons[idx][0] {
				sensorRangeAtRow[0]++
			}
			if closestBeacons[idx][1] == y && sensorRangeAtRow[1] == closestBeacons[idx][0] {
				sensorRangeAtRow[1]--
			}
			if sensorRangeAtRow[0] <= sensorRangeAtRow[1] {
				pointsRange = append(pointsRange, sensorRangeAtRow)
			}
		}
	}
	sort.Slice(pointsRange, func(i, j int) bool {
		return pointsRange[i][0] < pointsRange[j][0]
	})
	for idx := 0; idx < len(pointsRange)-1; idx++ {
		point1 := pointsRange[idx]
		point2 := pointsRange[idx+1]
		if (point2[0] <= point1[0] && point1[0] <= point2[1]) ||
			(point2[0] <= point1[1] && point1[1] <= point2[1]) ||
			(point1[0] <= point2[0] && point2[0] <= point1[1]) ||
			(point1[0] <= point2[1] && point2[1] <= point1[1]) {
			pointsRange[idx+1] = [2]int{
				min(point1[0], min(point1[1], min(point2[0], point2[1]))),
				max(point1[0], max(point1[1], max(point2[0], point2[1]))),
			}
			pointsRange[idx] = [2]int{math.MinInt32, math.MinInt32}
		}
	}
	points := 0
	for _, point := range pointsRange {
		if !(point[0] == math.MinInt32 && point[1] == math.MinInt32) {
			fmt.Printf("[%d %d] ", point[0], point[1])
			points += point[1] - point[0] + 1
		}
	}
	fmt.Println()
	return -1
}
func inLimit(num int) int {
	if num < 0 {
		return 0
	}
	if num > limit {
		return limit
	}
	return num
}
func findNum(sign string, digitStr string) int {
	digit, _ := strconv.Atoi(digitStr)
	if sign == "-" {
		digit *= -1
	}
	return digit
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
