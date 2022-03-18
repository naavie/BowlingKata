package main

import (
	"fmt"
	"strconv"
)

// contstants

const (
	// Number of pins in a game of bowling
	totalPins = 10
	// Maximum number of frames in a standard bowling game
	framesPerGame = 10
	// Max Rolls Per Frame
	maxRollsPerFrame = 2
	// Max Rolls in Last Frame
	maxRollsLastFrame = 3
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

// Spare Function
func isSpare(rollOne int, rollTwo int) bool {
	scorePerFrame := rollOne + rollTwo
	booleanVal := false
	if scorePerFrame == 10 {
		booleanVal = true
	}
	return booleanVal
}

// Strike Function
func isStrike(rollOne int) bool {
	booleanVal := false
	if rollOne == 10 {
		booleanVal = true
	}
	return booleanVal
}
