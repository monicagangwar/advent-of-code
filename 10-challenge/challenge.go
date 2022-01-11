package main

import (
	"bufio"
	"fmt"
	"runtime"
	"sort"

	"github.com/monicagangwar/advent-of-code-2021/input"
)

func main() {
	one()
	two()
}

func one() {
	_, currentFilePath, _, _ := runtime.Caller(0)
	file := input.GetFileMarker(currentFilePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	penalty := map[byte]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	count := map[byte]int{
		')': 0,
		']': 0,
		'}': 0,
		'>': 0,
	}

	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
		'>': '<',
	}

	//lineCount := 1

	for scanner.Scan() {
		line := scanner.Text()
		stack := make([]byte, 0)
		for _, char := range line {
			//fmt.Printf("%s\n", stack)
			if char == '(' || char == '[' || char == '{' || char == '<' {
				stack = append(stack, byte(char))
			} else if stack[len(stack)-1] == pairs[byte(char)] {
				//fmt.Printf("line %d: char: %s removing: %s\n", lineCount, string(byte(char)), string(stack[len(stack)-1]))
				stack = stack[:len(stack)-1]
			} else {
				//fmt.Printf("line %d: char expected : %s char found: %s\n", lineCount, string(pairs[byte(char)]), string(stack[len(stack)-1]))
				count[byte(char)] += 1
				break
			}
		}
		//fmt.Printf("%+v", count)
		//lineCount += 1
	}

	totalPenalty := 0

	for _, char := range []byte{')', ']', '}', '>'} {
		totalPenalty += count[char] * penalty[char]
	}

	fmt.Printf("%d\n", totalPenalty)

}

func two() {
	_, currentFilePath, _, _ := runtime.Caller(0)
	file := input.GetFileMarker(currentFilePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	penalty := map[byte]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}

	pairs := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}

	pairsOpposite := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
		'>': '<',
	}

	scores := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		stack := make([]byte, 0)
		incorrect := false
		for _, char := range line {
			if char == '(' || char == '[' || char == '{' || char == '<' {
				stack = append(stack, byte(char))
			} else if stack[len(stack)-1] == pairsOpposite[byte(char)] {
				stack = stack[:len(stack)-1]
			} else {
				incorrect = true
				break
			}
		}
		if incorrect {
			continue
		}

		score := 0

		completedChunk := ""

		for idx := len(stack) - 1; idx >= 0; idx-- {
			score = (score * 5) + penalty[pairs[stack[idx]]]
			completedChunk = fmt.Sprintf("%s%s", completedChunk, string(pairs[stack[idx]]))
		}

		//fmt.Printf("%s: %d\n", completedChunk, score)

		scores = append(scores, score)
	}

	sort.Slice(scores, func(i, j int) bool {
		return scores[i] < scores[j]
	})

	fmt.Printf("%d\n", scores[len(scores)/2])

}
