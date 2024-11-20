package _7

import (
	_ "embed"
	"sort"
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
		expectedPartOne int64
		expectedPartTwo int64
	}

	tests := []test{
		{
			name:            "with sample",
			input:           sample,
			expectedPartOne: 6440,
			expectedPartTwo: 5905,
		}, {
			name:            "with large input",
			input:           input,
			expectedPartOne: 0,
			expectedPartTwo: 0,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			hands := parseInput(tst.input)
			if tst.expectedPartOne != -1 {
				if got := getWinnings(hands, cardRank, getHandType); got != tst.expectedPartOne {
					t.Errorf("partOne() = %v, want %v", got, tst.expectedPartOne)
				}
			}

			if tst.expectedPartTwo != -1 {
				if got := getWinnings(hands, cardRankWithJoker, getHandTypeWithJoker); got != tst.expectedPartTwo {
					t.Errorf("partTwo() = %v, want %v", got, tst.expectedPartTwo)
				}
			}
		})

	}

}

func parseInput(input string) []Hand {
	lines := strings.Split(input, "\n")
	hands := make([]Hand, len(lines))
	for idx, line := range lines {
		parts := strings.Split(line, " ")
		bid, _ := strconv.Atoi(parts[1])
		hands[idx] = Hand{
			cards: parts[0],
			bid:   bid}
	}
	return hands
}

func lessThanCardOp(card1, card2 byte, cardRank map[byte]int) bool {
	return cardRank[card1] < cardRank[card2]
}

func lessThanHandOp(hand1, hand2 Hand, cardRank map[byte]int) bool {
	if hand1.handType != hand2.handType {
		return hand1.handType < hand2.handType
	}
	for i := 0; i < 5; i++ {
		if hand1.cards[i] != hand2.cards[i] {
			return lessThanCardOp(hand1.cards[i], hand2.cards[i], cardRank)
		}
	}
	return false
}

func getWinnings(hands []Hand, cardRank map[byte]int, getHandType func(s string) HandType) int64 {
	for idx, hand := range hands {
		hands[idx].handType = getHandType(hand.cards)
	}
	sort.Slice(hands, func(i, j int) bool {
		return lessThanHandOp(hands[i], hands[j], cardRank)
	})

	totalWinnings := int64(0)

	for rank, hand := range hands {
		totalWinnings += int64(rank+1) * int64(hand.bid)
	}
	return totalWinnings
}

type HandType int

const (
	HighCard HandType = iota - 1
	OnePair
	TwoPairs
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

var cardRank = map[byte]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

var cardRankWithJoker = map[byte]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
	'J': 0,
}

type Hand struct {
	cards    string
	bid      int
	handType HandType
}

func getHandTypeWithJoker(cards string) HandType {
	if !strings.Contains(cards, "J") {
		return getHandType(cards)
	}

	cardCount := make(map[byte]int)
	maxCountSameCards := 0
	for i := 0; i < 5; i++ {
		cardCount[cards[i]]++
		if cards[i] != 'J' && cardCount[cards[i]] > maxCountSameCards {
			maxCountSameCards = cardCount[cards[i]]
		}
	}

	if cardCount['J'] == 5 || cardCount['J'] == 4 {
		return FiveOfAKind
	}
	if cardCount['J'] == 3 {
		if len(cardCount) == 2 {
			return FiveOfAKind
		}
		if len(cardCount) == 3 {
			return FourOfAKind
		}
	}

	if cardCount['J'] == 2 {
		if len(cardCount) == 2 {
			return FiveOfAKind
		}
		if len(cardCount) == 3 {
			return FourOfAKind
		}
		if len(cardCount) == 4 {
			return ThreeOfAKind
		}
	}
	if cardCount['J'] == 1 {
		if len(cardCount) == 2 {
			return FiveOfAKind
		}
		if len(cardCount) == 3 {
			if maxCountSameCards == 3 {
				return FourOfAKind
			}
			return FullHouse
		}
		if len(cardCount) == 4 {
			return ThreeOfAKind
		}
		if len(cardCount) == 5 {
			return OnePair
		}
	}
	return HighCard
}

func getHandType(cards string) HandType {

	cardCount := make(map[byte]int)
	maxCountSameCards := 0
	for i := 0; i < 5; i++ {
		cardCount[cards[i]]++
		if cardCount[cards[i]] > maxCountSameCards {
			maxCountSameCards = cardCount[cards[i]]
		}
	}

	if maxCountSameCards == 5 {
		return FiveOfAKind
	}
	if maxCountSameCards == 4 {
		return FourOfAKind
	}
	if maxCountSameCards == 3 {
		if len(cardCount) == 2 {
			return FullHouse
		}
		return ThreeOfAKind
	}
	if maxCountSameCards == 2 {
		if len(cardCount) == 3 {
			return TwoPairs
		}
		return OnePair
	}
	return HighCard
}
