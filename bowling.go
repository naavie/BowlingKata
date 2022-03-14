package main

// contstants

const (
	// minimum number of pins
	minpins = 0
	// maximum number of pins
	maxpins = 10
	// Final score cannot be greater than 300 point.
	maxscore = 300
	// Frames
	frames = 10
	// Max Rolls Per Frame
	maxrollsperframe = 2
	// Max Rolls in Last Frame
	maxrollslastframe = 3
	// Max Rolls Per Game (Last Frame can have up to three rolls)
	maxrollspergame = 21
)

// Bowling Game Data
type Game struct {
	score int
}

// New Game - No Players Yet
func NewBowlingGame(scorecard string) Game {
	return Game{
		score: 10,
	}
}
