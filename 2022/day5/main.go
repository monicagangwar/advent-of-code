package main

import (
	"fmt"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"github.com/monicagangwar/advent-of-code/input"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	_, currentFilePath, _, _ := runtime.Caller(0)
	content := input.ReadInput(currentFilePath)
	lines := strings.Split(string(content), "\n")
	storingStack := true
	revStacks := make([][]string, 0)
	stacks := make([][]string, 0)
	stackLengths := make([]int, 0)
	instructionRegex := regexp.MustCompile(`^move (\d+) from (\d+) to (\d+)$`)
	for _, line := range lines {
		if storingStack && string(line[1]) == "1" {
			storingStack = false
			for _, revStack := range revStacks {
				stackLengths = append(stackLengths, len(revStack))
				stack := make([]string, 0)
				for i := len(revStack) - 1; i >= 0; i-- {
					stack = append(stack, revStack[i])
				}
				stacks = append(stacks, stack)
			}
			//fmt.Println(stacks)
			continue
		}
		if line == "" {
			continue
		}
		if storingStack {
			line = strings.ReplaceAll(line, "] [", "][")
			line = strings.ReplaceAll(line, "    ", "#")
			line = strings.ReplaceAll(line, "# ", "#")
			line = strings.ReplaceAll(line, "[", "")
			line = strings.ReplaceAll(line, "]", "")
			totalStacks := len(line)
			for len(revStacks) < totalStacks {
				revStacks = append(revStacks, make([]string, 0))
			}

			for stackIdx, char := range line {
				if string(char) != "#" {
					revStacks[stackIdx] = append(revStacks[stackIdx], string(char))
				}
			}
		} else {
			matches := instructionRegex.FindAllStringSubmatch(line, -1)
			count, _ := strconv.ParseInt(matches[0][1], 10, 32)
			from, _ := strconv.ParseInt(matches[0][2], 10, 32)
			to, _ := strconv.ParseInt(matches[0][3], 10, 32)

			// since we are using 0 index
			from--
			to--

			for idx := stackLengths[from] - 1; idx >= 0 && count > 0; count-- {
				if len(stacks[to]) > stackLengths[to] {
					stacks[to][stackLengths[to]] = stacks[from][idx]
				} else {
					stacks[to] = append(stacks[to], stacks[from][idx])
				}
				stackLengths[from]--
				stackLengths[to]++
				idx--
			}
			//fmt.Println(stacks)
		}
	}

	outputStr := ""
	for idx, stackLength := range stackLengths {
		if stackLength > 0 {
			outputStr += stacks[idx][stackLength-1]
		}
	}
	fmt.Println(outputStr)
}

func partTwo() {
	_, currentFilePath, _, _ := runtime.Caller(0)
	content := input.ReadInput(currentFilePath)
	lines := strings.Split(string(content), "\n")
	storingStack := true
	revStacks := make([][]string, 0)
	stacks := make([][]string, 0)
	stackLengths := make([]int, 0)
	instructionRegex := regexp.MustCompile(`^move (\d+) from (\d+) to (\d+)$`)
	for _, line := range lines {
		if storingStack && string(line[1]) == "1" {
			storingStack = false
			for _, revStack := range revStacks {
				stackLengths = append(stackLengths, len(revStack))
				stack := make([]string, 0)
				for i := len(revStack) - 1; i >= 0; i-- {
					stack = append(stack, revStack[i])
				}
				stacks = append(stacks, stack)
			}
			//fmt.Println(stacks)
			continue
		}
		if line == "" {
			continue
		}
		if storingStack {
			line = strings.ReplaceAll(line, "] [", "][")
			line = strings.ReplaceAll(line, "    ", "#")
			line = strings.ReplaceAll(line, "# ", "#")
			line = strings.ReplaceAll(line, "[", "")
			line = strings.ReplaceAll(line, "]", "")
			totalStacks := len(line)
			for len(revStacks) < totalStacks {
				revStacks = append(revStacks, make([]string, 0))
			}

			for stackIdx, char := range line {
				if string(char) != "#" {
					revStacks[stackIdx] = append(revStacks[stackIdx], string(char))
				}
			}
		} else {
			matches := instructionRegex.FindAllStringSubmatch(line, -1)
			count, _ := strconv.ParseInt(matches[0][1], 10, 32)
			from, _ := strconv.ParseInt(matches[0][2], 10, 32)
			to, _ := strconv.ParseInt(matches[0][3], 10, 32)

			// since we are using 0 index
			from--
			to--

			for idx := int64(stackLengths[from]) - count; count > 0; count-- {
				if len(stacks[to]) > stackLengths[to] {
					stacks[to][stackLengths[to]] = stacks[from][idx]
				} else {
					stacks[to] = append(stacks[to], stacks[from][idx])
				}
				stackLengths[from]--
				stackLengths[to]++
				idx++
			}
			//fmt.Println(stacks)
		}
	}

	outputStr := ""
	for idx, stackLength := range stackLengths {
		if stackLength > 0 {
			outputStr += stacks[idx][stackLength-1]
		}
	}
	fmt.Println(outputStr)
}
