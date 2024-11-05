package main

import (
	_ "embed"
	"fmt"
	"time"

	"github.com/buger/goterm"
)

//go:embed input.txt
var jetBlows []byte

type rockCacheKey struct {
	x          int
	rockType   int
	jetBlowIdx int
}

type rockCacheValue struct {
	rockIdx int
	height  int
}

func main() {
	fmt.Println(findMaxHeight(2022, nil))
}

func findMaxHeight(maxRockCount int, cache map[rockCacheKey]rockCacheValue) int {
	jetIdx, maxHeight, rockCount := 0, 0, 0
	x, y := 3, 4
	well := make([][8]int, 0)

	for i := 0; i < 1000000; i++ {
		var row [8]int
		well = append(well, row)
	}
	for {
		rockType := rockCount % 5
		if jetBlows[jetIdx] == '<' && canGoToPos(x, y, x-1, y, rockType, well) {
			movePiece(x, y, x-1, y, rockType, well)
			x = x - 1
		} else if jetBlows[jetIdx] == '>' && canGoToPos(x, y, x+1, y, rockType, well) {
			movePiece(x, y, x+1, y, rockType, well)
			x = x + 1
		}
		jetIdx = (jetIdx + 1) % len(jetBlows)

		if canGoToPos(x, y, x, y-1, rockType, well) {
			movePiece(x, y, x, y-1, rockType, well)
			y = y - 1
		} else {
			maxHeight = max(maxHeight, getHeight(x, y, rockType))
			x, y = 3, maxHeight+4
			rockCount++
			if rockCount == maxRockCount {
				break
			}
		}
	}
	return maxHeight

}

func getHeight(x, y, rockType int) int {
	coordinates := getCoordinatesForPiece(x, y, rockType)
	height := 0
	for _, coord := range coordinates {
		height = max(height, coord[1])
	}
	return height
}

func movePiece(curX, curY, x, y, rockType int, well [][8]int) {
	curCoords := getCoordinatesForPiece(curX, curY, rockType)
	for _, coord := range curCoords {
		well[coord[1]][coord[0]] = 0
	}

	newCoords := getCoordinatesForPiece(x, y, rockType)
	for _, coord := range newCoords {
		well[coord[1]][coord[0]] = 1
	}

	//printWell(well)
}

func printWell(well [][8]int) {
	time.Sleep(100 * time.Millisecond)
	goterm.Flush()
	goterm.MoveCursor(1, 6)
	for y := 4000; y >= 1; y-- {
		for x := 1; x < 8; x++ {
			if well[y][x] == 1 {
				fmt.Printf(goterm.Color("#", goterm.BLUE))
			} else {
				fmt.Printf(goterm.Color("o", goterm.YELLOW))
			}
		}
		fmt.Println()
	}
	//goterm.Clear()
}

func canGoToPos(curX, curY, x, y, rockType int, well [][8]int) bool {
	curCoordinates := getCoordinatesForPiece(curX, curY, rockType)
	curCoordMap := make(map[[2]int]struct{})
	for _, coord := range curCoordinates {
		curCoordMap[coord] = struct{}{}
	}
	coordinates := getCoordinatesForPiece(x, y, rockType)
	for _, coord := range coordinates {
		if _, ok := curCoordMap[coord]; ok {
			continue
		}
		newX := coord[0]
		newY := coord[1]
		//fmt.Printf("x: %d, y: %d\n", newX, newY)
		if newX < 1 || newX > 7 || newY < 1 {
			return false
		}
		if well[newY][newX] == 1 {
			return false
		}
	}
	return true
}

func getCoordinatesForPiece(x, y int, rockType int) [][2]int {
	switch rockType {
	case 0:
		// ----
		return [][2]int{{x, y}, {x + 1, y}, {x + 2, y}, {x + 3, y}}
	case 1:
		// +
		return [][2]int{{x, y + 1}, {x + 1, y + 1}, {x + 2, y + 1}, {x + 1, y}, {x + 1, y + 2}}
	case 2:
		// _|
		return [][2]int{{x, y}, {x + 1, y}, {x + 2, y}, {x + 2, y + 1}, {x + 2, y + 2}}
	case 3:
		// |
		return [][2]int{{x, y}, {x, y + 1}, {x, y + 2}, {x, y + 3}}
	case 4:
		// ::
		return [][2]int{{x, y}, {x, y + 1}, {x + 1, y}, {x + 1, y + 1}}
	}
	return nil
}
