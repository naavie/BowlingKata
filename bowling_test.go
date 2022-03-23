package main

import (
	"strconv"
	"testing"
)

//Unit Test for 1 Frame, 1 bowling pin knocked

func TestWhenNPinsKnockedInFrameOne(t *testing.T) {
	for i := 1; i <= 9; i++ {
		x := strconv.Itoa // converts int to string
		actual := NewBowlingGame(x(i) + "- -- -- -- -- -- -- -- -- --").score
		if actual != i {
			t.Errorf("Expected %v but got %v", i, actual)
		}
	}
}

// 1 to 9 pins for every bowl, every frame
// probably need to use slices. Need to know how to convert slices of strings to ints.

func TestWhenXPinsKnockedInEveryFrame(t *testing.T) {
	for i := 1; i <= 9; i++ {
		x := strconv.Itoa // converts int to string
		frameOne := NewBowlingGame(x(i) + "-").score
		frameTwo := NewBowlingGame(x(i) + "--").score
		frameThree := NewBowlingGame(x(i) + "--").score
		frameFour := NewBowlingGame(x(i) + "--").score
		frameFive := NewBowlingGame(x(i) + "--").score
		frameSix := NewBowlingGame(x(i) + "--").score
		frameSeven := NewBowlingGame(x(i) + "--").score
		frameEight := NewBowlingGame(x(i) + "--").score
		frameNine := NewBowlingGame(x(i) + "--").score
		frameTen := NewBowlingGame(x(i) + "--").score

		// tests
		if frameOne != i {
			t.Errorf("Expected %v but got %v", i, frameOne)
		}
		if frameTwo != i {
			t.Errorf("Expected %v but got %v", i, frameTwo)

		}
		if frameThree != i {
			t.Errorf("Expected %v but got %v", i, frameThree)

		}
		if frameFour != i {
			t.Errorf("Expected %v but got %v", i, frameFour)

		}
		if frameFive != i {
			t.Errorf("Expected %v but got %v", i, frameFive)

		}
		if frameSix != i {
			t.Errorf("Expected %v but got %v", i, frameSix)

		}
		if frameSeven != i {
			t.Errorf("Expected %v but got %v", i, frameSeven)

		}
		if frameEight != i {
			t.Errorf("Expected %v but got %v", i, frameEight)

		}
		if frameNine != i {
			t.Errorf("Expected %v but got %v", i, frameNine)

		}
		if frameTen != i {
			t.Errorf("Expected %v but got %v", i, frameTen)
		}
	}
}
