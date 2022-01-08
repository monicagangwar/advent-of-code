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

func performOp(op string, steps int, hPos *int, vPos *int) {
	switch op {
	case "forward":
		*hPos = *hPos + steps
		break
	case "down":
		*vPos = *vPos + steps
		break
	case "up":
		*vPos = *vPos - steps
	}
}

func performOpTwo(op string, steps int, hPos *int, vPos *int, aim *int) {
	switch op {
	case "forward":
		*hPos = *hPos + steps
		*vPos = *vPos + (*aim * steps)
		break
	case "down":
		*aim = *aim + steps
		break
	case "up":
		*aim = *aim - steps
	}
}

func one() {
	_, currentFilePath, _, _ := runtime.Caller(0)
	file := input.GetFileMarker(currentFilePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	horizontal := 0
	vertical := 0

	for scanner.Scan() {
		instruction := strings.Split(scanner.Text(), " ")
		steps, _ := strconv.Atoi(instruction[1])
		performOp(instruction[0], steps, &horizontal, &vertical)
	}

	fmt.Printf("%d\n", horizontal*vertical)

}

func two() {
	_, currentFilePath, _, _ := runtime.Caller(0)
	file := input.GetFileMarker(currentFilePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	horizontal := 0
	vertical := 0
	aim := 0

	for scanner.Scan() {
		instruction := strings.Split(scanner.Text(), " ")
		steps, _ := strconv.Atoi(instruction[1])
		performOpTwo(instruction[0], steps, &horizontal, &vertical, &aim)
	}

	fmt.Printf("%d\n", horizontal*vertical)
}
