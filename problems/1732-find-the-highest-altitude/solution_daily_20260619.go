package main

// Find the Highest Altitude (LeetCode 1732)
//
// Approach: prefix sum in a single pass.
// The altitudes are the running (prefix) sum of `gain`, starting from 0.
// We keep one running total and the maximum seen so far, initializing the
// maximum to 0 because the starting point's altitude (0) is a valid answer.
// Time: O(n), Space: O(1).
func largestAltitude(gain []int) int {
	highest := 0
	current := 0
	for _, g := range gain {
		current += g
		if current > highest {
			highest = current
		}
	}
	return highest
}
