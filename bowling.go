package main

import (
	"fmt"
	"strconv"
)

// Type: Bowling Scorecard
type GameCard struct {
	score     int
	scorecard string
}

// Type: Every Frame in Bowling Game
type Frame struct {
	frameOne   string
	frameTwo   string
	frameThree string
	frameFour  string
	frameFive  string
	frameSix   string
	frameSeven string
	frameEight string
	frameNine  string
	frameTen   string

	frameScore string

	GameCard
}

type Score interface {
	TotalScore(int, int, int) int
}

// New Game - No Players Yet
func NewBowlingGame(scorecard string) GameCard {
	// blankScorecard := [21]string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}
	rune1 := []rune(scorecard)[0]
	runestring1 := string(rune1)
	fmt.Print(runestring1)
	val, _ := strconv.Atoi(runestring1)
	return GameCard{
		score:     val,
		scorecard: scorecard,
	}
}

//  Sean's code above
// ----------------------------------------------------------------------
