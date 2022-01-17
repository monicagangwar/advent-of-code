package main

import (
	"fmt"
	"regexp"
	"runtime"
	"strconv"

	"github.com/monicagangwar/advent-of-code-2021/input"
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

	target := result["ymin"]
	if result["ymax"] < result["ymin"] {
		target = result["ymax"]
	}

	target = (-1 * target) - 1

	answer := (target * (target + 1)) / 2

	fmt.Printf("%+v", answer)
}
