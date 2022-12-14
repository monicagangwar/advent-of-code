package main

import (
	"fmt"
	"math"
	"regexp"
	"runtime"
	"strconv"

	"github.com/monicagangwar/advent-of-code/input"
)

func main() {
	_, currentWorkingDirectory, _, _ := runtime.Caller(0)
	content := input.ReadInput(currentWorkingDirectory)

	inputRegex := regexp.MustCompile("target area: x=(?P<xmin>([0-9-]+))..(?P<xmax>([0-9-]+)), y=(?P<ymin>([0-9-]+))..(?P<ymax>([0-9-]+))")
	match := inputRegex.FindStringSubmatch(string(content))

	result := make(map[string]int64)
	for i, name := range inputRegex.SubexpNames() {
		if i != 0 && name != "" {
			matchNum, _ := strconv.ParseInt(match[i], 10, 64)
			result[name] = matchNum
		}
	}

	fmt.Printf("%+v", one(result["xmin"], result["xmax"], result["ymin"], result["ymax"]))

	fmt.Printf("\n%+v", two(result["xmin"], result["xmax"], result["ymin"], result["ymax"]))

}

func one(xmin, xmax, ymin, ymax int64) int64 {

	target := ymin
	if ymax < ymin {
		target = ymax
	}

	target = (-1 * target) - 1

	return (target * (target + 1)) / 2
}

func canHit(velx, vely, xmin, xmax, ymin, ymax int64) int {
	curX, curY := int64(0), int64(0)
	initialVelX, initialVelY := velx, vely
	for {
		if curX > xmax {
			return 0
		}
		if velx == 0 && !(xmin <= curX && curX <= xmax) {
			return 0
		}
		if velx == 0 && curY < ymin {
			return 0
		}

		if xmin <= curX && curX <= xmax && ymin <= curY && curY <= ymax {
			fmt.Printf("\n %d %d", initialVelX, initialVelY)
			return 1
		}

		curX += velx
		curY += vely

		if velx > 0 {
			velx--
		}

		vely--
	}
}

func two(xmin, xmax, ymin, ymax int64) int {
	count := 0
	for velx := int64(1); velx <= xmax+1; velx++ {
		for vely := ymin; vely <= int64(math.Max(math.Abs(float64(ymin)), math.Abs(float64(ymax)))); vely++ {
			count += canHit(velx, vely, xmin, xmax, ymin, ymax)
		}
	}
	return count
}
