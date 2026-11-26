package main

// Approach: Simulation.
// Players alternate removing 10, 9, 8, ... stones. Keep subtracting the
// required amount while the pile can afford it, flipping the turn each time.
// The player who made the last successful move wins; the one who is unable
// to move loses. Since amounts strictly decrease from 10, there are at most
// 10 moves, so this runs in constant time.
func canAliceWin(n int) bool {
	take := 10
	aliceTurn := true
	for n >= take {
		n -= take
		take--
		aliceTurn = !aliceTurn
	}
	// The player whose turn it is now cannot move and loses.
	return !aliceTurn
}
