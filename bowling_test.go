package main

import (
	"testing"
)

func TestBowling(t *testing.T) {
	if 1+1 != 2 {
		t.Error("1+1 is not equal to 2.")

	}
}

//Unit Test for 1 Frame, 1 bowling pin knocked

func TestWhenOnePinIsKnockedAtTheBegginingOfTheGameTheScoreShouldBeOne(t *testing.T) {
	actual := NewBowlingGame("1- -- -- -- -- -- -- -- -- --").score
	if actual != 1 {
		t.Errorf("Expected %v but got %v", 1, actual)
	}

}

// Unit Test for 1 Frame, 2 bowling pin knocked
func TestWhenTwoPinsAreKnockedAtTheBegginingOfTheGameTheScoreShouldBeTwo(t *testing.T) {
	actual := NewBowlingGame("2- -- -- -- -- -- -- -- -- --").score
	if actual != 2 {
		t.Errorf("Expected %v but got %v", 2, actual)
	}

}

// Unit Test for 1 Frame, 3 bowling pin knocked
func TestWhenThreePinsAreKnockedAtTheBegginingOfTheGameTheScoreShouldBeThree(t *testing.T) {
	actual := NewBowlingGame("3- -- -- -- -- -- -- -- -- --").score
	if actual != 3 {
		t.Errorf("Expected %v but got %v", 3, actual)
	}
}

// Unit Test for 1 Frame, 4 bowling pin knocked
func TestWhenFourPinsAreKnockedAtTheBegginingOfTheGameTheScoreShouldBeFour(t *testing.T) {
	actual := NewBowlingGame("4- -- -- -- -- -- -- -- -- --").score
	if actual != 4 {
		t.Errorf("Expected %v but got %v", 4, actual)
	}
}

// Unit Test for 1 Frame, 5 bowling pin knocked
func TestWhenFivePinsAreKnockedAtTheBegginingOfTheGameTheScoreShouldBeFive(t *testing.T) {
	actual := NewBowlingGame("5- -- -- -- -- -- -- -- -- --").score
	if actual != 5 {
		t.Errorf("Expected %v but got %v", 5, actual)
	}
}

// Unit Test for 1 Frame, 6 bowling pin knocked
func TestWhenSixPinsAreKnockedAtTheBegginingOfTheGameTheScoreShouldBeSix(t *testing.T) {
	actual := NewBowlingGame("6- -- -- -- -- -- -- -- -- --").score
	if actual != 6 {
		t.Errorf("Expected %v but got %v", 6, actual)
	}
}

// Unit Test for 1 Frame, 7 bowling pin knocked
func TestWhenSevenPinsAreKnockedAtTheBegginingOfTheGameTheScoreShouldBeSeven(t *testing.T) {
	actual := NewBowlingGame("7- -- -- -- -- -- -- -- -- --").score
	if actual != 7 {
		t.Errorf("Expected %v but got %v", 7, actual)
	}
}

// Unit Test for 1 Frame, 8 bowling pin knocked
func TestWhenEightPinsAreKnockedAtTheBegginingOfTheGameTheScoreShouldBeEight(t *testing.T) {
	actual := NewBowlingGame("8- -- -- -- -- -- -- -- -- --").score
	if actual != 8 {
		t.Errorf("Expected %v but got %v", 8, actual)
	}
}

// Unit Test for 1 Frame, 9 bowling pin knocked
func TestWhenNinePinsAreKnockedAtTheBegginingOfTheGameTheScoreShouldBeNine(t *testing.T) {
	actual := NewBowlingGame("9- -- -- -- -- -- -- -- -- --").score
	if actual != 9 {
		t.Errorf("Expected %v but got %v", 9, actual)
	}
}

// Unit Test for 1 Frame, 10 bowling pin knocked
func TestWhenTenPinsAreKnockedAtTheBegginingOfTheGameTheScoreShouldBeTen(t *testing.T) {
	actual := NewBowlingGame("X- -- -- -- -- -- -- -- -- --").score
	if actual != 10 {
		t.Errorf("Expected %v but got %v", 10, actual)
	}
}
func TestWhenXPinsAreKnockedDownInFrameOne(t *testing.T) {
	TotalRolls := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for i := 0; i <= 9; i++ {
		t.Errorf("Expected %v but got %v", i, TotalRolls)
	}
}
