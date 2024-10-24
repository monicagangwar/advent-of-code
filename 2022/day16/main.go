package main

import (
	"fmt"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"github.com/monicagangwar/advent-of-code/input"
)

type Valve struct {
	name     string
	flowRate int
	tunnels  map[string]*Valve
}

type Path struct {
	flow  int
	route []string
}

var cave map[string]*Valve
var valvesWithFlowRate []string
var distances map[string]map[string]int

func main() {
	_, currentFilePath, _, _ := runtime.Caller(0)
	content := input.ReadInput(currentFilePath)
	lines := strings.Split(string(content), "\n")
	locationRegex := regexp.MustCompile(`Valve (\w+) has flow rate=(\d+); tunnels? leads? to valves? ((\w+)(, \w+)*)`)

	cave = make(map[string]*Valve)
	valvesWithFlowRate = make([]string, 0)

	for _, line := range lines {
		matches := locationRegex.FindAllStringSubmatch(line, -1)
		name := matches[0][1]
		flowRate, _ := strconv.Atoi(matches[0][2])
		tunnels := strings.Split(matches[0][3], ", ")

		valve, found := cave[name]
		if !found {
			valve = &Valve{name: name, flowRate: flowRate, tunnels: map[string]*Valve{}}
			cave[name] = valve
		}

		valve.flowRate = flowRate
		if flowRate > 0 {
			valvesWithFlowRate = append(valvesWithFlowRate, name)
		}

		for _, tunnel := range tunnels {
			neighborNode, found := cave[tunnel]
			if !found {
				neighborNode = &Valve{name: tunnel, tunnels: map[string]*Valve{}}
				cave[tunnel] = neighborNode
			}
			valve.tunnels[neighborNode.name] = neighborNode
		}
	}
	distances = computeDistances(cave)
	//printDistances()

	fmt.Println(valvesWithFlowRate)
	fmt.Println(partOne())
	fmt.Println(partTwo())

}

func partOne() int {
	allPaths := DFS("AA", 30, Path{0, []string{}}, make(map[string]struct{}))

	maxPressure := 0
	var bestPath Path
	for _, path := range allPaths {
		if path.flow > maxPressure {
			maxPressure = path.flow
			bestPath = path
		}
	}

	fmt.Println(bestPath.route)

	return maxPressure
}

func partTwo() int {
	allPaths := DFS("AA", 26, Path{0, []string{"AA"}}, make(map[string]struct{}))

	maxPressure := 0
	//for _, path := range allPaths {
	//	maxPressure = max(maxPressure, path.flow)
	//}

	var r1, r2 Path

	for i := 0; i < len(allPaths); i++ {
		visitedA := make(map[string]struct{})
		for _, r := range allPaths[i].route {
			visitedA[r] = struct{}{}
		}
		for j := i + 1; j < len(allPaths); j++ {
			routesMutuallyExclusive := mutuallyExclusive(visitedA, allPaths[j].route)
			//fmt.Println(allPressures[i].route, allPressures[j].route, routesMutuallyExclusive)
			combinedPressure := allPaths[i].flow + allPaths[j].flow
			if routesMutuallyExclusive && combinedPressure > maxPressure {
				maxPressure = combinedPressure
				r1 = allPaths[i]
				r2 = allPaths[j]
			}
		}
	}

	fmt.Println(r1.route, r2.route)

	return maxPressure
}

func mutuallyExclusive(visited map[string]struct{}, route2 []string) bool {
	for _, r2 := range route2 {
		if _, found := visited[r2]; r2 != "AA" && found {
			return false
		}
	}
	return true
}

func DFS(currentValve string, timeRemaining int, path Path, visited map[string]struct{}) []Path {
	paths := []Path{path}

	if timeRemaining == 0 {
		return paths
	}

	for _, nextValve := range valvesWithFlowRate {
		_, nextValveVisited := visited[nextValve]
		newTimeRemaining := timeRemaining - distances[currentValve][nextValve] - 1

		if !nextValveVisited && newTimeRemaining > 0 {
			newVisited := copyVisited(visited)
			newPath := Path{
				flow:  path.flow,
				route: []string{},
			}
			newVisited[nextValve] = struct{}{}
			for _, r := range path.route {
				newPath.route = append(newPath.route, r)
			}

			newPath.flow += cave[nextValve].flowRate * newTimeRemaining
			newPath.route = append(newPath.route, nextValve)

			paths = append(paths, DFS(nextValve, newTimeRemaining, newPath, newVisited)...)
		}
	}

	return paths

}

func copyVisited(visited map[string]struct{}) map[string]struct{} {
	newVisited := make(map[string]struct{})
	for k, v := range visited {
		newVisited[k] = v
	}
	return newVisited
}

func computeDistances(cave map[string]*Valve) map[string]map[string]int {
	dist := make(map[string]map[string]int)

	for i, _ := range cave {
		if _, found := dist[i]; !found {
			dist[i] = map[string]int{}
		}
		for j, _ := range cave {
			if i == j {
				dist[i][j] = 0
			} else if _, found := cave[i].tunnels[j]; found {
				dist[i][j] = 1
			} else {
				dist[i][j] = 99999999
			}
		}
	}

	for k, _ := range cave {
		for i, _ := range cave {
			for j, _ := range cave {
				dist[i][j] = min(dist[i][j], dist[i][k]+dist[k][j])
			}
		}
	}

	return dist
}
