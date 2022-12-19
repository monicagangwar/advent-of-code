package main

import (
	"container/heap"
	"fmt"
	"math"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/monicagangwar/advent-of-code/input"
)

func main() {
	startTime := time.Now()
	_, currentFilePath, _, _ := runtime.Caller(0)
	content := input.ReadInput(currentFilePath)
	lines := strings.Split(string(content), "\n")
	cave := make(map[string][]string)
	flowRate := make(map[string]int)
	distances := make(map[[2]string]int)
	locationRegex := regexp.MustCompile(`Valve (\w+) has flow rate=(\d+); tunnel(s)? lead(s)? to valve(s)? ((\w+)(, \w+)*)`)
	for _, line := range lines {
		matches := locationRegex.FindAllStringSubmatch(line, -1)
		if tunnels, found := cave[matches[0][1]]; found {
			tunnels = append(tunnels, strings.Split(matches[0][6], ", ")...)
			cave[matches[0][1]] = tunnels
		} else {
			cave[matches[0][1]] = strings.Split(matches[0][6], ", ")
		}
		flowRate[matches[0][1]], _ = strconv.Atoi(matches[0][2])
	}

	computeDistances(cave, distances)
	openedValves := make([]string, 0)
	for valve, rate := range flowRate {
		if rate == 0 {
			openedValves = append(openedValves, valve)
		}
	}
	fmt.Println(findMaxVal(distances, flowRate, 30, "AA", strings.Join(openedValves, ",")))
	fmt.Printf("total time taken %s", time.Now().Sub(startTime).String())
}

func findMaxVal(distances map[[2]string]int, flowRate map[string]int, timeRemaining int, curValve string,
	openedValve string) int {
	//fmt.Println(fmt.Sprintf("%d %s %s", timeRemaining, curValve, openedValve))
	if timeRemaining < 0 {
		return 0
	}
	if len(strings.Split(openedValve, ",")) == len(flowRate) {
		return 0
	}
	maxVal := 0
	if flowRate[curValve] > 0 {
		timeRemaining--
		maxVal = timeRemaining * flowRate[curValve]
		openedValve = openedValve + "," + curValve
	}
	maxDestVal := 0
	for destValve := range flowRate {
		if destValve != curValve && !strings.Contains(openedValve, destValve) {
			curTimeRemaining := timeRemaining - distances[getKey(curValve, destValve)]
			maxDestVal = max(maxDestVal, findMaxVal(distances, flowRate, curTimeRemaining, destValve, openedValve))
		}
	}
	return maxVal + maxDestVal
}

func computeDistances(cave map[string][]string, distances map[[2]string]int) {
	for srcValve := range cave {
		for destValve := range cave {
			if srcValve == destValve {
				distances[getKey(srcValve, destValve)] = 0
			} else {
				distances[getKey(srcValve, destValve)] = math.MaxInt32
			}
		}
	}

	for srcValve := range cave {
		for destValve := range cave {
			if srcValve == destValve {
				continue
			}
			queue := make(PriorityQueue, 0)
			queue.Push(&Item{valve: srcValve, distance: 0})
			visited := make(map[string]struct{})
			for {
				if queue.Len() == 0 {
					break
				}
				item := heap.Pop(&queue).(*Item)
				visited[item.valve] = struct{}{}
				for _, neighbor := range cave[item.valve] {
					distances[getKey(srcValve, neighbor)] = min(distances[getKey(srcValve, neighbor)], distances[getKey(srcValve, item.valve)]+1)
					if _, found := visited[neighbor]; !found {
						queue.Push(&Item{valve: neighbor, distance: distances[getKey(srcValve, neighbor)]})
					}
				}
			}
		}
	}
}

func getKey(srcValve, destValve string) [2]string {
	return [2]string{srcValve, destValve}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
