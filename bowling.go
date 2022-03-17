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
type Game struct {
	score     int
	scorecard string
}

// New Game - No Players Yet
func NewBowlingGame(scorecard string) Game {
	rune1 := []rune(scorecard)[0]
	runestring := string(rune1)
	fmt.Print(runestring)
	val, _ := strconv.Atoi(runestring)

	return Game{
		score: val,
	}
}
