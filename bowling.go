package main

import (
	"fmt"
	"strconv"
)

// Bowling Game Data
type Game struct {
	score     int
	scorecard string
}

// New Game - No Players Yet
func NewBowlingGame(scorecard string) Game {
	rune1 := []rune(scorecard)[0]
	runestring1 := string(rune1)
	fmt.Print(runestring1)
	val, _ := strconv.Atoi(runestring1)
	return Game{
		score:     val,
		scorecard: scorecard,
	}
}

//  Sean's code above
// ----------------------------------------------------------------------
