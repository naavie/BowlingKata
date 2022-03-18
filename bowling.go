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

// ----------------------------------------------------------------------------------------------

type Game struct {
	TotalRolls [21]int
	current    int
}

// trying to figure out how to iterate to the next frame. Since each frame has two pins, it needs to iterate over two int values
func (game *Game) playerScore() int {
	score := 0            //initial score when game starts
	iterateNextFrame := 0 // this variable will be used to iterate to the next frame after the existing frame has been completed
	for frameIndex := 0; frameIndex < 10; frameIndex += 1 {
		if isStrike(game.TotalRolls[iterateNextFrame]) {
			score += 10 + game.TotalRolls[iterateNextFrame+1] + game.TotalRolls[iterateNextFrame+2]
			iterateNextFrame += 1
		}

	}
	return score
}

// Strike Function
func isStrike(rollOne int) bool {
	booleanVal := false
	if rollOne == 10 {
		booleanVal = true
	}
	return booleanVal
}

// Spare Function
func isSpare(rollOne int, rollTwo int) bool {
	scorePerFrame := rollOne + rollTwo
	booleanVal := false
	if scorePerFrame == 10 {
		booleanVal = true
	}
	return booleanVal
}
