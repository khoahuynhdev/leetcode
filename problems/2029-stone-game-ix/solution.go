package main

// 2029. Stone Game IX
//
// Only each stone's value mod 3 matters, so the whole array collapses into
// three counters. Stones with remainder 0 never change the running sum; they
// are pure "pass" moves, so only their parity matters. Once the sum is
// nonzero, the residues the players must take are forced to alternate
// (1,2,1,2,... or 2,1,2,1,...), which reduces the game to a counting rule:
//
//	c0 even -> Alice wins iff she holds at least one stone of each residue.
//	c0 odd  -> the leftover pass flips the tempo to Bob, so Alice needs one
//	           residue class to outnumber the other by more than 2.
//
// Time: O(n), Space: O(1).
func stoneGameIX(stones []int) bool {
	var count [3]int
	for _, s := range stones {
		count[s%3]++
	}

	// Remainder-0 stones only swap the turn; an even number of them cancels out.
	if count[0]%2 == 0 {
		return count[1] > 0 && count[2] > 0
	}

	// One unpaired pass move hands Bob the tempo unless a residue class is
	// lopsided enough to exhaust him anyway. Note the strict ">": a difference
	// of exactly 2 is still a Bob win.
	diff := count[1] - count[2]
	if diff < 0 {
		diff = -diff
	}
	return diff > 2
}
