package main

// import (
// 	"fmt"
// )

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
<<<<<<< Updated upstream
func NewBowlingGame(scorecard string) Game {
	return Game{
		score: 1,
=======
func NewBowlingGame(scorecard string) Game2 {
	rune1 := []rune(scorecard)[0]
	runestring := string(rune1)
	fmt.Print(runestring)
	val, _ := strconv.Atoi(runestring)

	return Game2{
		score: val,
>>>>>>> Stashed changes
	}
}

// ^ Sean's code above

// Nav's new code below.

// I want to start with a clean slate. First, I'll create a new type called "New Game"

type Game struct {
	maxRollsPerTenFrames [21]int
	currentRoll          int
}

func (g *Game) ballsThrown(pinsKnocked int)
