package template

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
			expectedPartOne: 480,
			expectedPartTwo: 875318608908,
		}, {
			name:            "with large input",
			input:           input,
			expectedPartOne: 0,
			expectedPartTwo: 0,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			games := parseInput(tst.input)
			if tst.expectedPartOne != -1 {
				if got := partOne(games); got != tst.expectedPartOne {
					t.Errorf("partOne() = %v, want %v", got, tst.expectedPartOne)
				}
			}

			if tst.expectedPartTwo != -1 {
				if got := partTwo(games); got != tst.expectedPartTwo {
					t.Errorf("partTwo() = %v, want %v", got, tst.expectedPartTwo)
				}
			}
		})

	}

}

type Game struct {
	AInc  [2]int
	BInc  [2]int
	Prize [2]int
}

func parseInput(input string) []Game {
	lines := strings.Split(input, "\n")

	buttonRegex := regexp.MustCompile(`Button (?P<button>A|B): X\+(?P<xInc>\d+), Y\+(?P<yInc>\d+)`)
	prizeRegex := regexp.MustCompile(`Prize: X=(?P<x>\d+), Y=(?P<y>\d+)`)

	games := make([]Game, 0)
	game := Game{}
	for _, line := range lines {
		matches := buttonRegex.FindStringSubmatch(line)

		if len(matches) > 0 {
			button := matches[1]
			xInc, _ := strconv.Atoi(matches[2])
			yInc, _ := strconv.Atoi(matches[3])
			if button == "A" {
				game.AInc = [2]int{xInc, yInc}
			} else {
				game.BInc = [2]int{xInc, yInc}
			}
		} else {
			matches = prizeRegex.FindStringSubmatch(line)
			if len(matches) > 0 {
				xPos, _ := strconv.Atoi(matches[1])
				yPos, _ := strconv.Atoi(matches[2])
				game.Prize = [2]int{xPos, yPos}
				games = append(games, game)
				game = Game{}
			}
		}
	}

	return games
}

func partOne(games []Game) int {
	tokens := 0
	for _, game := range games {
		tokenA := (game.Prize[0]*game.BInc[1] - game.Prize[1]*game.BInc[0]) / (game.AInc[0]*game.BInc[1] - game.AInc[1]*game.BInc[0])
		tokenB := (game.AInc[0]*game.Prize[1] - game.AInc[1]*game.Prize[0]) / (game.AInc[0]*game.BInc[1] - game.AInc[1]*game.BInc[0])

		if tokenA*game.AInc[0]+tokenB*game.BInc[0] == game.Prize[0] && tokenA*game.AInc[1]+tokenB*game.BInc[1] == game.Prize[1] {
			//fmt.Printf("A: %d B: %d\n", tokenA, tokenB)
			tokens += (tokenA * 3) + tokenB
		}

	}
	return tokens
}

func partTwo(games []Game) int {
	tokens := 0
	for _, game := range games {
		game.Prize[0] += 10000000000000
		game.Prize[1] += 10000000000000
		tokenA := (game.Prize[0]*game.BInc[1] - game.Prize[1]*game.BInc[0]) / (game.AInc[0]*game.BInc[1] - game.AInc[1]*game.BInc[0])
		tokenB := (game.AInc[0]*game.Prize[1] - game.AInc[1]*game.Prize[0]) / (game.AInc[0]*game.BInc[1] - game.AInc[1]*game.BInc[0])

		if tokenA*game.AInc[0]+tokenB*game.BInc[0] == game.Prize[0] && tokenA*game.AInc[1]+tokenB*game.BInc[1] == game.Prize[1] {
			//fmt.Printf("A: %d B: %d\n", tokenA, tokenB)
			tokens += (tokenA * 3) + tokenB
		}

	}
	return tokens
}
