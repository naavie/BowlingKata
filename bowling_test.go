package main

import (
	"testing"
)

func TestBowling(t *testing.T) {
	if 1+1 != 2 {
		t.Error("1+1 is not equal to 2.")

	}
}

// // Perfect Score Test
// func PerfectScore(t *testing.T) {
// 	if 300 != maxscore { //placeholder
// 		t.Error("You did not get a perfect score this game but good try!")
// 	}
// }

//Unit Test for 1 Frame, 1 bowling pin knocked

func TestWhenOnePinIsKnockedAtTheBegginingOfTheGameTheScoreShouldBeOne(t *testing.T) {
	actual := NewBowlingGame("1- -- -- -- -- -- -- -- -- --").score
	if actual != 1 {
		t.Errorf("Expected %v but got %v", 1, actual)
	}

}

// Unit Test for 1 Frame, 2 bowling pin knocked
func TestWhenTwoPinsAreKnockedAtTheBegginingOfTheGameTheScoreShouldBeTwo(t *testing.T) {
	actual := NewBowlingGame("1- -- -- -- -- -- -- -- -- --").score
	if actual != 2 {
		t.Errorf("Expected %v but got %v", 2, actual)
	}

}

// Unit Test for 1 Frame, 3 bowling pin knocked
func TestWhenThreePinsAreKnockedAtTheBegginingOfTheGameTheScoreShouldBeThree(t *testing.T) {
	actual := NewBowlingGame("1- -- -- -- -- -- -- -- -- --").score
	if actual != 3 {
		t.Errorf("Expected %v but got %v", 3, actual)
	}
}

// Unit Test for 1 Frame, 4 bowling pin knocked
func TestWhenFourPinsAreKnockedAtTheBegginingOfTheGameTheScoreShouldBeFour(t *testing.T) {
	actual := NewBowlingGame("1- -- -- -- -- -- -- -- -- --").score
	if actual != 4 {
		t.Errorf("Expected %v but got %v", 4, actual)
	}
}

// Unit Test for 1 Frame, 5 bowling pin knocked
func TestWhenFivePinsAreKnockedAtTheBegginingOfTheGameTheScoreShouldBeFive(t *testing.T) {
	actual := NewBowlingGame("1- -- -- -- -- -- -- -- -- --").score
	if actual != 5 {
		t.Errorf("Expected %v but got %v", 5, actual)
	}
}

// Unit Test for 1 Frame, 6 bowling pin knocked
func TestWhenSixPinsAreKnockedAtTheBegginingOfTheGameTheScoreShouldBeSix(t *testing.T) {
	actual := NewBowlingGame("1- -- -- -- -- -- -- -- -- --").score
	if actual != 6 {
		t.Errorf("Expected %v but got %v", 6, actual)
	}
}

// Unit Test for 1 Frame, 7 bowling pin knocked
func TestWhenSevenPinsAreKnockedAtTheBegginingOfTheGameTheScoreShouldBeSeven(t *testing.T) {
	actual := NewBowlingGame("1- -- -- -- -- -- -- -- -- --").score
	if actual != 7 {
		t.Errorf("Expected %v but got %v", 7, actual)
	}
}

// Unit Test for 1 Frame, 8 bowling pin knocked
func TestWhenEightPinsAreKnockedAtTheBegginingOfTheGameTheScoreShouldBeEight(t *testing.T) {
	actual := NewBowlingGame("1- -- -- -- -- -- -- -- -- --").score
	if actual != 8 {
		t.Errorf("Expected %v but got %v", 8, actual)
	}
}

// Unit Test for 1 Frame, 9 bowling pin knocked
func TestWhenNinePinsAreKnockedAtTheBegginingOfTheGameTheScoreShouldBeNine(t *testing.T) {
	actual := NewBowlingGame("1- -- -- -- -- -- -- -- -- --").score
	if actual != 9 {
		t.Errorf("Expected %v but got %v", 9, actual)
	}
}

// Unit Test for 1 Frame, 10 bowling pin knocked
func TestWhenTenPinsAreKnockedAtTheBegginingOfTheGameTheScoreShouldBeTen(t *testing.T) {
	actual := NewBowlingGame("1- -- -- -- -- -- -- -- -- --").score
	if actual != 10 {
		t.Errorf("Expected %v but got %v", 10, actual)
	}
}

func Test_V2_OnePinOneFrame(t *testing.T) {
	actual := BowlingGameVersion2("firstRollFirstFrame", 0).newgame()
	if actual != 5 {
		t.Errorf("Expected %v but got %v", 5, actual)
	}
}

// // These are not tests, they are features.
// func ScoreTests() {
// 	if NewBowlingGame().currentscore == maxscore {
// 		fmt.Println("Congratulations", NewBowlingGame().playername, "you had a perfect game!")
// 	}

// 	if NewBowlingGame().currentscore == 0 {
// 		fmt.Println("Please play the game again.")
// 	}
// }
