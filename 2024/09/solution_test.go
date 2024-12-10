package _9

import (
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

//go:embed sample.txt
var sample string

func TestSolution(t *testing.T) {
	type test struct {
		name            string
		input           string
		expectedPartOne int
		expectedPartTwo int64
	}

	tests := []test{
		{
			name:            "with sample",
			input:           sample,
			expectedPartOne: 1928,
			expectedPartTwo: 2858,
		}, {
			name:            "with large input",
			input:           input,
			expectedPartOne: 0,
			expectedPartTwo: 0,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			if tst.expectedPartOne != -1 {
				if got := partOne([]byte(tst.input)); got != tst.expectedPartOne {
					t.Errorf("partOne() = %v, want %v", got, tst.expectedPartOne)
				}
			}

			if tst.expectedPartTwo != -1 {
				if got := partTwo([]byte(tst.input)); got != tst.expectedPartTwo {
					t.Errorf("partTwo() = %v, want %v", got, tst.expectedPartTwo)
				}
			}
		})

	}

}

func partOne(in []byte) int {
	expandedIn := make([]int, 0)
	fileId := 0

	for i, char := range in {
		num := int(char - '0')

		arrToAppend := make([]int, 0)
		spaceType := -1
		if i%2 == 0 {
			spaceType = fileId
			fileId++
		}
		for j := 1; j <= num; j++ {
			arrToAppend = append(arrToAppend, spaceType)
		}
		expandedIn = append(expandedIn, arrToAppend...)
	}

	freespacePtr := 0
	lastElemPtr := len(expandedIn) - 1
	for i := 0; i < len(expandedIn); i++ {
		if expandedIn[i] == -1 {
			freespacePtr = i
			break
		}
	}

	for {

		if expandedIn[freespacePtr] != -1 {
			for {
				if freespacePtr >= len(expandedIn) {
					break
				}
				if expandedIn[freespacePtr] == -1 {
					break
				}
				freespacePtr++
			}
		}

		if expandedIn[lastElemPtr] == -1 {
			for {
				if lastElemPtr < 0 {
					break
				}
				if expandedIn[lastElemPtr] != -1 {
					break
				}
				lastElemPtr--
			}
		}

		//fmt.Printf("freespacePtr: %d, lastElemPtr: %d\n", freespacePtr, lastElemPtr)

		if freespacePtr >= lastElemPtr {
			break
		}

		expandedIn[freespacePtr] = expandedIn[lastElemPtr]
		expandedIn[lastElemPtr] = -1
		freespacePtr++
		lastElemPtr--

	}

	checksum := 0

	for i := 0; i < len(expandedIn); i++ {
		if expandedIn[i] == -1 {
			break
		}

		checksum += i * expandedIn[i]
	}

	return checksum
}

type node struct {
	id    int
	size  int
	left  *node
	right *node
}

func partTwo(in []byte) int64 {
	var start *node
	var last *node
	fileID := 0
	for i := 0; i < len(in); i++ {
		size := int(in[i] - '0')
		newNode := node{
			size: size,
		}
		if i%2 == 0 { // block
			newNode.id = fileID
			fileID++
		} else { // empty space
			newNode.id = -1
		}

		if start == nil {
			start = &newNode
			last = &newNode
		} else {
			last.right = &newNode
			newNode.left = last
			last = &newNode
		}
	}

	curNode := last

	for {
		if curNode == nil {
			break
		}
		if curNode.id == -1 {
			curNode = curNode.left
			continue
		}

		startNode := start
		for {
			if startNode == nil {
				break
			}
			if startNode == curNode {
				break
			}

			if startNode.id == -1 && startNode.size == curNode.size {
				newFileNode := node{
					id:   curNode.id,
					size: curNode.size,
				}
				prevFreeSpaceNode := startNode.left
				nextFreeSpaceNode := startNode.right
				newFileNode.left = prevFreeSpaceNode
				newFileNode.right = nextFreeSpaceNode

				prevFreeSpaceNode.right = &newFileNode
				nextFreeSpaceNode.left = &newFileNode

				curNode.id = -1

				break

			} else if startNode.id == -1 && startNode.size > curNode.size {
				newFileNode := node{
					id:   curNode.id,
					size: curNode.size,
				}
				newEmptySpaceNode := node{
					id:   -1,
					size: startNode.size - curNode.size,
				}

				prevFreeSpaceNode := startNode.left
				nextFreeSpaceNode := startNode.right

				newFileNode.right = &newEmptySpaceNode
				newFileNode.left = prevFreeSpaceNode
				newEmptySpaceNode.left = &newFileNode
				newEmptySpaceNode.right = nextFreeSpaceNode

				prevFreeSpaceNode.right = &newFileNode
				nextFreeSpaceNode.left = &newEmptySpaceNode

				curNode.id = -1

				break
			}
			startNode = startNode.right
		}
		curNode = curNode.left
	}

	curNode = start
	fileArr := make([]int, 0)
	for {
		if curNode == nil {
			break
		}
		for i := 0; i < curNode.size; i++ {
			fileArr = append(fileArr, curNode.id)
		}
		curNode = curNode.right
	}

	ans := int64(0)
	for i := 0; i < len(fileArr); i++ {
		if fileArr[i] != -1 {
			ans += int64(i) * int64(fileArr[i])
		}
	}
	return ans

}
