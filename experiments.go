package main

import (
	"fmt"
)

func test() {
	newgame := make(map[string]int)

	//Frame 1: Two Rolls
	newgame["firstRollFirstFrame"] = 0
	newgame["secondRollFirstFrame"] = 0

	// Frame 2: Two Rolls
	newgame["firstRollSecondFrame"] = 0
	newgame["secondRollSecondFrame"] = 0

	// Frame 3: Two Rolls
	newgame["firstRollThirdFrame"] = 0
	newgame["secondRollThirdFrame"] = 0

	// Frame 4: Two Rolls
	newgame["firstRollFourthFrame"] = 0
	newgame["secondRollFourthFrame"] = 0

	// Frame 5: Two Rolls
	newgame["firstRollFifthFrame"] = 0
	newgame["secondRollFifthFrame"] = 0

	// Frame 6: Two Rolls
	newgame["firstRollSixthFrame"] = 0
	newgame["secondRollSixthFrame"] = 0

	// Frame 7: Two Rolls
	newgame["firstRollSeventhFrame"] = 0
	newgame["secondRollSeventhFrame"] = 0

	// Frame 8: Two Rolls
	newgame["firstRollEighthFrame"] = 0
	newgame["secondRollEiththFrame"] = 0

	// Frame 9: Two Rolls
	newgame["firstRollNinthFrame"] = 0
	newgame["secondRollNinthFrame"] = 0

	// Frame 10: Two Rolls + 1 Roll (If conditions are met)
	newgame["firstRollTenthFrame"] = 0
	newgame["secondRollTenthFrame"] = 0

	fmt.Println(newgame)
}
