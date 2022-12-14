package main

import (
	"fmt"
	"math"
	"runtime"
	"strconv"
	"strings"

	"github.com/monicagangwar/advent-of-code/input"
)

func main() {
	challenge("one")
	challenge("two")
}

func challenge(part string) {
	_, currentFilePath, _, _ := runtime.Caller(0)
	content := input.ReadInput(currentFilePath)

	maxPos := 0

	ships := make([]int, 0)
	for _, shipPosStr := range strings.Split(string(content), ",") {
		shipPos, _ := strconv.Atoi(shipPosStr)
		ships = append(ships, shipPos)
		if shipPos > maxPos {
			maxPos = shipPos
		}
	}

	minFuelNeeded := float64(100000000000)

	for idx := 1; idx < maxPos; idx++ {
		fuel := float64(0)
		for jidx := 0; jidx < len(ships); jidx++ {
			diff := math.Abs(float64(idx - ships[jidx]))
			if part == "one" {
				fuel += diff
			} else {
				computedFuel := (diff * (diff + 1)) / float64(2)
				fuel += computedFuel
				//if idx == 5 {
				//	fmt.Printf("Move from %d to %d: %f fuel\n", ships[jidx], idx, computedFuel)
				//}
			}
		}
		if fuel < minFuelNeeded {
			minFuelNeeded = fuel
		}
	}

	fmt.Printf("%f\n", minFuelNeeded)

}
