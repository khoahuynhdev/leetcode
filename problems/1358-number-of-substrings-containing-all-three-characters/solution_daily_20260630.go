package main

// Problem 1358: Number of Substrings Containing All Three Characters
//
// Approach: single-pass sliding window using "last seen" positions.
// For each right index, track the most recent index where 'a', 'b', and 'c'
// appeared. The number of valid substrings ending at `right` equals
// min(lastA, lastB, lastC) + 1, because any left endpoint at or before that
// minimum still keeps all three characters inside the window. Positions are
// initialized to -1 so that, until all three characters have been seen, the
// minimum is -1 and contributes nothing. Summing over every right index gives
// the answer in O(n) time and O(1) space.
func numberOfSubstrings(s string) int {
	// last[0], last[1], last[2] are the most recent indices of 'a', 'b', 'c'.
	last := [3]int{-1, -1, -1}
	total := 0

	for right := 0; right < len(s); right++ {
		last[s[right]-'a'] = right

		// Smallest of the three last-seen positions; substrings starting at
		// any index 0..minPos (inclusive) and ending at right are valid.
		minPos := last[0]
		if last[1] < minPos {
			minPos = last[1]
		}
		if last[2] < minPos {
			minPos = last[2]
		}

		total += minPos + 1
	}

	return total
}
