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
	rolls        int
	initscore    int
	currentscore int
	playername   string
}

// New Game - No Players Yet
func NewBowlingGame() *Game {
	return &Game{
		rolls:        maxrollsperframe, // how do we handle the last frame? maxrollsperframe + 1?
		initscore:    0,
		currentscore: 0,
		playername:   "",
	}
}
