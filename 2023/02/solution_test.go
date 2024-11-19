package _2

import (
	_ "embed"
	"fmt"
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
		expectedPartTwo int64
	}

	tests := []test{
		{
			name:            "with sample",
			input:           sample,
			expectedPartOne: 8,
			expectedPartTwo: 2286,
		}, {
			name:            "with large input",
			input:           input,
			expectedPartOne: ***REMOVED***,
			expectedPartTwo: ***REMOVED***,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			games := parseGame(tst.input)
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
	id    int
	plays [][]Cube
}

type Cube struct {
	count int
	color Color
}

type Color string

const (
	red   Color = "red"
	blue  Color = "blue"
	green Color = "green"
)

func parseGame(input string) []Game {
	lines := strings.Split(input, "\n")

	games := make([]Game, len(lines))

	gameIDRegex := regexp.MustCompile(`Game (\d+):`)
	playsRegex := regexp.MustCompile(`^(?P<count>\d+) (?P<color>red|blue|green)$`)

	for gameIdx, line := range lines {
		gameID := gameIDRegex.FindStringSubmatch(line)[1]
		gameIDInt, _ := strconv.Atoi(gameID)

		game := Game{id: gameIDInt, plays: make([][]Cube, 0)}

		plays := strings.Split(strings.Replace(line, fmt.Sprintf("Game %d: ", gameIDInt), "", 1), ";")
		for _, play := range plays {
			cubes := make([]Cube, 0)
			draws := strings.Split(strings.TrimSpace(play), ",")
			for _, draw := range draws {
				matches := playsRegex.FindStringSubmatch(strings.TrimSpace(draw))
				count := matches[playsRegex.SubexpIndex("count")]
				color := matches[playsRegex.SubexpIndex("color")]

				countInt, _ := strconv.Atoi(count)
				cubes = append(cubes, Cube{count: countInt, color: Color(color)})
			}
			game.plays = append(game.plays, cubes)
		}
		games[gameIdx] = game
	}
	return games
}

func partOne(games []Game) int {
	maxCubes := map[Color]int{
		red:   12,
		blue:  14,
		green: 13,
	}
	validGamesSum := 0

	for _, game := range games {
		valid := true
		for _, play := range game.plays {
			for _, cube := range play {
				if cube.count > maxCubes[cube.color] {
					valid = false
					break
				}
			}
		}
		if valid {
			validGamesSum += game.id
		}

	}

	return validGamesSum
}

func partTwo(games []Game) int64 {
	sumPowerOfCubes := int64(0)
	for _, game := range games {
		minColorCount := map[Color]int{
			red:   -1,
			blue:  -1,
			green: -1,
		}
		for _, play := range game.plays {
			for _, cube := range play {
				if minColorCount[cube.color] < cube.count {
					minColorCount[cube.color] = cube.count
				}
			}
		}
		power := 1
		for _, count := range minColorCount {
			power = power * count
		}
		//fmt.Printf("minColor: %+v, power: %d\n", minColorCount, power)
		sumPowerOfCubes += int64(power)
	}
	return sumPowerOfCubes
}
