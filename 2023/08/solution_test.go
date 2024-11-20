package _8

import (
	_ "embed"
	"regexp"
	"strings"
	"testing"
)

//go:embed input.txt
var input string

//go:embed sample.txt
var sample string

//go:embed sample2.txt
var sample2 string

//go:embed sample3.txt
var sample3 string

func TestSolution(t *testing.T) {
	type test struct {
		name            string
		input           string
		expectedPartOne int
		expectedPartTwo int
	}

	tests := []test{
		{
			name:            "with sample",
			input:           sample,
			expectedPartOne: 2,
			expectedPartTwo: -1,
		}, {
			name:            "with sample 2",
			input:           sample2,
			expectedPartOne: 6,
			expectedPartTwo: -1,
		}, {
			name:            "with sample 3",
			input:           sample3,
			expectedPartOne: -1,
			expectedPartTwo: 6,
		},
		{
			name:            "with large input",
			input:           input,
			expectedPartOne: 0,
			expectedPartTwo: 0,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			instruction, nodeMap := parseInput(tst.input)
			if tst.expectedPartOne != -1 {
				if got := partOne(instruction, nodeMap); got != tst.expectedPartOne {
					t.Errorf("partOne() = %v, want %v", got, tst.expectedPartOne)
				}
			}

			if tst.expectedPartTwo != -1 {
				if got := partTwo(instruction, nodeMap); got != tst.expectedPartTwo {
					t.Errorf("partTwo() = %v, want %v", got, tst.expectedPartTwo)
				}
			}
		})

	}

}

func partOne(instruction string, nodeMap map[string]*Node) int {
	startNode := nodeMap["AAA"]
	endNode := nodeMap["ZZZ"]
	return findPath(instruction, startNode, endNode)
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(nums []int) int {
	result := (nums[0] * nums[1]) / GCD(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		result = (result * nums[i]) / GCD(result, nums[i])
	}
	return result
}

func partTwo(instruction string, nodeMap map[string]*Node) int {
	allAs := make([]*Node, 0)
	for _, node := range nodeMap {
		if node.id[2] == 'A' {
			allAs = append(allAs, node)
		}
	}

	allPaths := make([]int, len(allAs))
	for i, startNode := range allAs {
		allPaths[i] = findPath(instruction, startNode, nil)
	}

	return LCM(allPaths)
}

func findPath(instruction string, startNode, endNode *Node) int {
	instructionCount := 0
	var curNode *Node
	curNode = startNode
	for {
		if endNode == nil && curNode.id[2] == 'Z' {
			break
		}
		if endNode != nil && curNode.id == endNode.id {
			break
		}

		instructionIdx := instructionCount % len(instruction)

		curInstruction := instruction[instructionIdx]
		if curInstruction == 'R' {
			curNode = curNode.right
		} else {
			curNode = curNode.left
		}
		instructionCount++
	}

	return instructionCount
}

func parseInput(input string) (string, map[string]*Node) {
	lines := strings.Split(input, "\n")

	instruction := lines[0]

	networkRegex := regexp.MustCompile(`^(?P<node>[A-Z]*) = \((?P<left>[A-Z]*), (?P<right>[A-Z]*)\)$`)

	nodeMap := make(map[string]*Node)

	for i := 2; i < len(lines); i++ {
		matches := networkRegex.FindStringSubmatch(lines[i])
		curNodeId := matches[1]
		left := matches[2]
		right := matches[3]
		for _, nodeID := range []string{curNodeId, left, right} {
			node, ok := nodeMap[nodeID]
			if !ok {
				node = &Node{id: nodeID}
				nodeMap[nodeID] = node
			}
		}
		nodeMap[curNodeId].left = nodeMap[left]
		nodeMap[curNodeId].right = nodeMap[right]
	}

	return instruction, nodeMap
}

type Node struct {
	id    string
	left  *Node
	right *Node
}
