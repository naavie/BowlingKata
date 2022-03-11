package main

import (
	"fmt"
	"testing"
)

func TestBowling(t *testing.T) {
	if 1+1 != 2 {
		t.Error("1+1 is not equal to 2.")

	}
}

// Perfect Score Test
func PerfectScore(t *testing.T) {
	if 300 != maxscore { //placeholder
		t.Error("You did not get a perfect score this game but good try!")
	}
}

func ScoreTests() {
	if NewBowlingGame().currentscore == maxscore {
		fmt.Println("Congratulations", NewBowlingGame().playername, "you had a perfect game!")
	}

	if NewBowlingGame().currentscore == 0 {
		fmt.Println("Please play the game again.")
	}
}
