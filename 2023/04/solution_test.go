package _4

import (
	_ "embed"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

//go:embed input.txt
var input string

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
			input:           "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53\nCard 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19\nCard 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1\nCard 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83\nCard 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36\nCard 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
			expectedPartOne: 13,
			expectedPartTwo: -1,
		}, {
			name:            "with large input",
			input:           input,
			expectedPartOne: ***REMOVED***,
			expectedPartTwo: -1,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			cards := parseInput(tst.input)
			if tst.expectedPartOne != -1 {
				if got := partOne(cards); got != tst.expectedPartOne {
					t.Errorf("partOne() = %v, want %v", got, tst.expectedPartOne)
				}
			}

			if tst.expectedPartTwo != -1 {
				if got := partTwo(cards); got != tst.expectedPartTwo {
					t.Errorf("partTwo() = %v, want %v", got, tst.expectedPartTwo)
				}
			}
		})

	}

}

type Card struct {
	ID             int
	WinningNumbers []int
	CardNumbers    []int
}

func (c Card) getMatchCount() int {
	matchCount := 0
	winningNumberMap := make(map[int]struct{})
	for _, num := range c.WinningNumbers {
		winningNumberMap[num] = struct{}{}
	}
	for _, num := range c.CardNumbers {
		if _, ok := winningNumberMap[num]; ok {
			//fmt.Printf("%d is a match\n", num)
			matchCount++
		}
	}
	//fmt.Println()
	return matchCount
}

func convertStrListToInt(strList string) []int {
	nums := make([]int, 0)
	for _, numStr := range strings.Split(strList, " ") {
		num, err := strconv.Atoi(strings.TrimSpace(numStr))
		if err == nil {
			nums = append(nums, num)
		}

	}
	return nums
}

func parseInput(input string) []Card {
	lines := strings.Split(input, "\n")

	cardEntryRegex := regexp.MustCompile(`^Card *(?P<id>\d+): (?P<winning>[\d+ ]*) \| (?P<num>[\d+ ]*)$`)

	cards := make([]Card, len(lines))

	for idx, line := range lines {
		matches := cardEntryRegex.FindStringSubmatch(strings.TrimSpace(line))
		cardID, _ := strconv.Atoi(matches[1])
		cards[idx] = Card{
			ID:             cardID,
			WinningNumbers: convertStrListToInt(matches[2]),
			CardNumbers:    convertStrListToInt(matches[3]),
		}
	}
	return cards
}

func partOne(cards []Card) int {
	totalPoints := 0

	for _, card := range cards {
		matchCount := card.getMatchCount()
		if matchCount > 0 {
			totalPoints += 1 << (matchCount - 1)
			//fmt.Printf("Card %d: %d = %d\n", card.ID, matchCount, 1<<(matchCount-1))
		}

	}
	return totalPoints
}

func partTwo(cards []Card) int64 {
	return -1
}
