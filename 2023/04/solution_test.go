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
			expectedPartOne: 13,
			expectedPartTwo: 30,
		}, {
			name:            "with large input",
			input:           input,
			expectedPartOne: ***REMOVED***,
			expectedPartTwo: ***REMOVED***,
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

func partTwo(cards []Card) int {
	cardCount := make(map[int]int, len(cards))

	for _, card := range cards {
		cardCount[card.ID] = 1
	}

	for _, card := range cards {
		curCardCount := cardCount[card.ID]
		matchCount := card.getMatchCount()
		for i := card.ID + 1; i <= card.ID+matchCount; i++ {
			cardCount[i] += curCardCount
		}
	}

	totalCards := 0

	for _, count := range cardCount {
		totalCards += count
	}

	return totalCards
}
