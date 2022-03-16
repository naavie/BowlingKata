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
	// Max Rolls Per Game (Last Frame can have up to three rolls)
	maxRollsPerGame = 21
)

// Bowling Game Data
type Game struct {
	score     int
	scorecard string
	newgame   map[string]int
}

// New Game - No Players Yet
func NewBowlingGame(scorecard string) Game {
	return Game{
		score: 1,
	}
}

// // New Game Version 2 - No Score
// func BowlingGameVersion2(scorecard string, score int) Game {
// 	newgame := make(map[string]int)

// 	//Frame 1: Two Rolls
// 	newgame["firstRollFirstFrame"] = 0
// 	newgame["secondRollFirstFrame"] = 0

// 	// Frame 2: Two Rolls
// 	newgame["firstRollSecondFrame"] = 0
// 	newgame["secondRollSecondFrame"] = 0

// 	// Frame 3: Two Rolls
// 	newgame["firstRollThirdFrame"] = 0
// 	newgame["secondRollThirdFrame"] = 0

// 	// Frame 4: Two Rolls
// 	newgame["firstRollFourthFrame"] = 0
// 	newgame["secondRollFourthFrame"] = 0

// 	// Frame 5: Two Rolls
// 	newgame["firstRollFifthFrame"] = 0
// 	newgame["secondRollFifthFrame"] = 0

// 	// Frame 6: Two Rolls
// 	newgame["firstRollSixthFrame"] = 0
// 	newgame["secondRollSixthFrame"] = 0

// 	// Frame 7: Two Rolls
// 	newgame["firstRollSeventhFrame"] = 0
// 	newgame["secondRollSeventhFrame"] = 0

// 	// Frame 8: Two Rolls
// 	newgame["firstRollEighthFrame"] = 0
// 	newgame["secondRollEiththFrame"] = 0

// 	// Frame 9: Two Rolls
// 	newgame["firstRollNinthFrame"] = 0
// 	newgame["secondRollNinthFrame"] = 0

// 	// Frame 10: Two Rolls + 1 Roll (If conditions are met)
// 	newgame["firstRollTenthFrame"] = 0
// 	newgame["secondRollTenthFrame"] = 0

// 	// for key, value := range newgame {
// 	// 	fmt.Println("key:", key, "value:", value)

// 	// 	// if newgame["firstRollTenthFrame"]+newgame["secondRollTenthFrame"] == 10 {
// 	// 	// 	// Conditions: If newgame["firstRollTenthFrame"] + newgame["secondRollTenthFrame"] == 10, then this for loop appends a new key-pair to "newgame" for the third roll in frame 10
// 	// 	// }
// 	// }
// 	return Game{
// 		score,
// 		scorecard,
// 	}
// }
