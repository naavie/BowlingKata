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
	frameOne   int
	frameTwo   int
	frameThree int
	frameFour  int
	frameFive  int
	frameSix   int
	frameSeven int
	frameEight int
	frameNine  int
	frameTen   int

	GameCard string
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

func (f *Frame) playerScore(currentScore int, rollOneScore int, rollTwoScore int) {
	currentScore = 0

	for i := 1; i <= 10; i++ {
		currentScore = rollOneScore + rollTwoScore
		if rollOneScore == 10 { //strike
			rollTwoScore = 0
			currentScore = rollOneScore
			f.frameOne = currentScore

		} else if rollOneScore+rollTwoScore == 10 { // spare
			currentScore = rollOneScore + rollTwoScore
			f.frameOne = currentScore

		} else if rollOneScore+rollTwoScore != 10 {
			currentScore = rollOneScore + rollTwoScore
			f.frameOne = currentScore

		}

		for i := 1; i <= 10; i++ {
			currentScore = f.frameOne
			if rollOneScore == 10 { //strike
				rollTwoScore = 0
				currentScore = rollOneScore + rollTwoScore
				f.frameTwo = f.frameOne + currentScore

			} else if rollOneScore+rollTwoScore == 10 { // spare
				currentScore = rollOneScore + rollTwoScore
				f.frameTwo = f.frameOne + currentScore

			} else if rollOneScore+rollTwoScore != 10 {
				currentScore = rollOneScore + rollTwoScore
				f.frameTwo = f.frameOne + currentScore
			}
		} // etc.. for the remaining frames
	}
}
