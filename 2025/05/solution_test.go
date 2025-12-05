package template

import (
	"cmp"
	_ "embed"
	"regexp"
	"slices"
	"strconv"
	"strings"
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
		expectedPartTwo int
	}

	tests := []test{
		{
			name:            "with sample",
			input:           sample,
			expectedPartOne: 3,
			expectedPartTwo: 14,
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
				if got := partOne(parseInput(tst.input)); got != tst.expectedPartOne {
					t.Errorf("partOne() = %v, want %v", got, tst.expectedPartOne)
				}
			}

			if tst.expectedPartTwo != -1 {
				if got := partTwo(parseInput(tst.input)); got != tst.expectedPartTwo {
					t.Errorf("partTwo() = %v, want %v", got, tst.expectedPartTwo)
				}
			}
		})

	}

}

func parseInput(input string) ([][2]int, []int) {
	lines := strings.Split(input, "\n")
	ranges := make([][2]int, 0)
	ingredientIDs := make([]int, 0)
	rangeRegex := regexp.MustCompile(`^(\d+)-(\d+)$`)
	for _, line := range lines {
		line = strings.TrimSpace(line)
		matches := rangeRegex.FindStringSubmatch(line)
		if len(matches) > 1 {
			num1, _ := strconv.Atoi(matches[1])
			num2, _ := strconv.Atoi(matches[2])
			ranges = append(ranges, [2]int{num1, num2})
		} else if line != "" {
			num, _ := strconv.Atoi(line)
			ingredientIDs = append(ingredientIDs, num)
		}
	}
	return ranges, ingredientIDs
}

func partOne(ranges [][2]int, ingredientIDs []int) int {
	freshCount := 0
	for _, id := range ingredientIDs {
		fresh := false
		for _, rng := range ranges {
			if rng[0] <= id && id <= rng[1] {
				fresh = true
				break
			}
		}
		if fresh {
			freshCount++
		}
	}
	return freshCount
}

/*
3-5
10-14
16-20
12-18
*/

type node struct {
	rng  [2]int
	next *node
}

func partTwo(ranges [][2]int, _ []int) int {
	for i := 0; i < len(ranges); i++ {
		if ranges[i][0] > ranges[i][1] {
			tmp := ranges[i][0]
			ranges[i][0] = ranges[i][1]
			ranges[i][1] = tmp
		}
	}

	slices.SortFunc(ranges, func(a, b [2]int) int {
		return cmp.Compare(a[0], b[0])
	})

	var start *node
	var curNode *node

	for _, rng := range ranges {
		newNode := node{
			rng: rng,
		}
		if start == nil {
			start = &newNode
			curNode = &newNode
		} else {
			curNode.next = &newNode
			curNode = &newNode
		}
	}

	curNode = start

	for curNode.next != nil {
		x1, y1 := curNode.rng[0], curNode.rng[1]
		x2, y2 := curNode.next.rng[0], curNode.next.rng[1]

		// r1 = x1, y1; r2 = x2, y2
		// case 1 = x1 <= x2 <= y2 <= y1
		// case 2 = x1 <= x2 <= y1 <= y2
		if x1 <= x2 && y2 <= y1 {
			curNode.next = curNode.next.next
		} else if x2 <= y1 && y1 <= y2 {
			curNode.rng[1] = y2
			curNode.next = curNode.next.next
		} else {
			curNode = curNode.next
		}
	}

	totalFresh := 0
	curNode = start
	for curNode != nil {
		total := curNode.rng[1] - curNode.rng[0] + 1
		//fmt.Printf("%d %d = %d\n", curNode.rng[0], curNode.rng[1], total)
		totalFresh += total
		curNode = curNode.next
	}

	return totalFresh
}
