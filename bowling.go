package main

// contstants

const (
	// maximum number of pins
	totalPins     = 10
	framesPerGame = 10
	// Max Rolls Per Frame
	maxRollsPerFrame = 2
	// Max Rolls in Last Frame
	maxRollsLastFrame = 3
	// Max Rolls Per Game (Last Frame can have up to three rolls)
	maxRollsPerGame = 21
)

// Bowling Game Data
type Game struct {
	score int
}

// New Game - No Players Yet
func NewBowlingGame(scorecard string) Game {
	return Game{
		score: 1,
	}
}

// func main() {
// 	scorecard := make((map[string]int))

// 	scorecard["--"] = nil
// }
