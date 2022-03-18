package main

import (
	"fmt"
	"strconv"
)

// Bowling Game Data
type Game2 struct {
	score     int
	scorecard string
}

// New Game - No Players Yet
func NewBowlingGame(scorecard string) Game2 {
	rune1 := []rune(scorecard)[0]
	runestring := string(rune1)
	fmt.Print(runestring)
	val, _ := strconv.Atoi(runestring)

	return Game2{
		score: val,
	}
}

//  Sean's code above

// ----------------------------------------------------------------------------------------------

type Game struct {
	TotalRolls [21]int
	current    int
}
