package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input []byte

func main() {
	var coords [][3]int
	parseCoord(&coords)

	var cache = make(map[[3]int]struct{})
	for _, coord := range coords {
		cache[coord] = struct{}{}
	}

	totalSurfaces := 0

	neighbors := [][3]int{
		{-1, 0, 0},
		{1, 0, 0},
		{0, -1, 0},
		{0, 1, 0},
		{0, 0, -1},
		{0, 0, 1},
	}

	for _, coord := range coords {
		for _, n := range neighbors {
			neighborCoords := [3]int{coord[0] + n[0], coord[1] + n[1], coord[2] + n[2]}
			if _, ok := cache[neighborCoords]; !ok {
				totalSurfaces++
			}
		}

	}
	fmt.Println(totalSurfaces)
}

func parseCoord(coords *[][3]int) {

	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		strCoords := strings.Split(line, ",")
		var key [3]int
		for i, c := range strCoords {
			val, _ := strconv.Atoi(c)
			key[i] = val
		}
		*coords = append(*coords, key)
	}
}
