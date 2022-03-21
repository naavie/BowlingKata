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

}
