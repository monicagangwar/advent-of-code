package main

import (
	"runtime"
	"strings"

	"github.com/monicagangwar/advent-of-code-2022/input"
)

type play int

const (
	rock play = iota + 1
	paper
	scissor
)

type outcome int

const (
	lose outcome = 0
	draw outcome = 3
	win  outcome = 6
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	strategyMappings := map[string]play{
		"A": rock,
		"B": paper,
		"C": scissor,
		"X": rock,
		"Y": paper,
		"Z": scissor,
	}
	type round struct {
		play1 play
		play2 play
	}

	getOutcomeForPlayer2 := map[round]outcome{
		round{play1: rock, play2: rock}:       draw,
		round{play1: rock, play2: paper}:      win,
		round{play1: rock, play2: scissor}:    lose,
		round{play1: paper, play2: rock}:      lose,
		round{play1: paper, play2: paper}:     draw,
		round{play1: paper, play2: scissor}:   win,
		round{play1: scissor, play2: rock}:    win,
		round{play1: scissor, play2: paper}:   lose,
		round{play1: scissor, play2: scissor}: draw,
	}

	_, currentFilePath, _, _ := runtime.Caller(0)
	content := input.ReadInput(currentFilePath)
	lines := strings.Split(string(content), "\n")
	score := 0
	for _, line := range lines {
		roundStrategy := strings.Split(line, " ")
		player1 := strategyMappings[roundStrategy[0]]
		player2 := strategyMappings[roundStrategy[1]]

		roundOutcome := getOutcomeForPlayer2[round{play1: player1, play2: player2}]
		score += int(player2) + int(roundOutcome)
	}
	println(score)
}

func partTwo() {
	strategyMappings := map[string]play{
		"A": rock,
		"B": paper,
		"C": scissor,
	}
	outcomeMappings := map[string]outcome{
		"X": lose,
		"Y": draw,
		"Z": win,
	}
	type round struct {
		play1   play
		outcome outcome
	}

	getPlayForPlayer2 := map[round]play{
		round{play1: rock, outcome: lose}:    scissor,
		round{play1: rock, outcome: draw}:    rock,
		round{play1: rock, outcome: win}:     paper,
		round{play1: paper, outcome: lose}:   rock,
		round{play1: paper, outcome: draw}:   paper,
		round{play1: paper, outcome: win}:    scissor,
		round{play1: scissor, outcome: lose}: paper,
		round{play1: scissor, outcome: draw}: scissor,
		round{play1: scissor, outcome: win}:  rock,
	}

	_, currentFilePath, _, _ := runtime.Caller(0)
	content := input.ReadInput(currentFilePath)
	lines := strings.Split(string(content), "\n")
	score := 0
	for _, line := range lines {
		roundStrategy := strings.Split(line, " ")
		player1 := strategyMappings[roundStrategy[0]]
		roundOutcome := outcomeMappings[roundStrategy[1]]

		player2 := getPlayForPlayer2[round{play1: player1, outcome: roundOutcome}]
		score += int(player2) + int(roundOutcome)
	}
	println(score)
}
