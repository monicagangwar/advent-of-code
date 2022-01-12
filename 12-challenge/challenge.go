package main

import (
	"bufio"
	"fmt"
	"runtime"
	"strings"
	"unicode"

	"github.com/monicagangwar/advent-of-code-2021/input"
)

func main() {
	one()
	two()
}

func getAllPathCount(vertex string, graph map[string][]string, visited map[string]struct{}) int {
	//fmt.Printf("%s %+v", vertex, visited)
	if vertex == "end" {
		//fmt.Printf("\n")
		return 1
	}
	sum := 0
	for _, neighbors := range graph[vertex] {
		_, isVisited := visited[neighbors]
		if !isVisited {
			if unicode.IsLower(rune(neighbors[0])) {
				newVisited := make(map[string]struct{})
				for k, v := range visited {
					newVisited[k] = v
				}
				newVisited[neighbors] = struct{}{}
				sum += getAllPathCount(neighbors, graph, newVisited)
			} else {
				sum += getAllPathCount(neighbors, graph, visited)
			}
		}
	}
	//fmt.Printf("\n")
	return sum
}

func getAllPathCountTwo(vertex string, graph map[string][]string, visited map[string]int, twiceVisited bool) int {
	//fmt.Printf("%s,", vertex)
	if vertex == "end" {
		return 1
	}
	sum := 0
	for _, neighbors := range graph[vertex] {
		if unicode.IsUpper(rune(neighbors[0])) {
			//fmt.Printf("vertex: %s is not lowercase\n", neighbors)
			sum += getAllPathCountTwo(neighbors, graph, visited, twiceVisited)
		} else {
			visitedCount, _ := visited[neighbors]
			newVisited := make(map[string]int)
			for k, v := range visited {
				newVisited[k] = v
			}
			//fmt.Printf("vertex: %s, visitedCount: %d\n", neighbors, visitedCount)
			if visitedCount == 0 {
				newVisited[neighbors] = 1
				sum += getAllPathCountTwo(neighbors, graph, newVisited, twiceVisited)
			} else if visitedCount == 1 && !twiceVisited && neighbors != "start" && neighbors != "end" {
				newVisited[neighbors] = 2
				sum += getAllPathCountTwo(neighbors, graph, newVisited, true)
			}
		}
	}
	return sum
}

func one() {
	_, currentFilePath, _, _ := runtime.Caller(0)
	file := input.GetFileMarker(currentFilePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	graph := make(map[string][]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		vertices := strings.Split(line, "-")

		leftVertex := vertices[0]
		leftVertexList := []string{vertices[1]}
		rightVertex := vertices[1]
		rightVertexList := []string{vertices[0]}

		if verticesAttached, found := graph[leftVertex]; found {
			leftVertexList = append(verticesAttached, leftVertexList...)
		}
		if verticesAttached, found := graph[rightVertex]; found {
			rightVertexList = append(verticesAttached, rightVertexList...)
		}

		graph[leftVertex] = leftVertexList
		graph[rightVertex] = rightVertexList
	}

	visited := make(map[string]struct{}, 0)

	visited["start"] = struct{}{}

	fmt.Printf("%d\n", getAllPathCount("start", graph, visited))

}

func two() {
	_, currentFilePath, _, _ := runtime.Caller(0)
	file := input.GetFileMarker(currentFilePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	graph := make(map[string][]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		vertices := strings.Split(line, "-")

		leftVertex := vertices[0]
		leftVertexList := []string{vertices[1]}
		rightVertex := vertices[1]
		rightVertexList := []string{vertices[0]}

		if verticesAttached, found := graph[leftVertex]; found {
			leftVertexList = append(verticesAttached, leftVertexList...)
		}
		if verticesAttached, found := graph[rightVertex]; found {
			rightVertexList = append(verticesAttached, rightVertexList...)
		}

		graph[leftVertex] = leftVertexList
		graph[rightVertex] = rightVertexList
	}

	visited := make(map[string]int, 0)

	visited["start"] = 1

	fmt.Printf("%d\n", getAllPathCountTwo("start", graph, visited, false))

}
