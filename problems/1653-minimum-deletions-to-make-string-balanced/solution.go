package main

// Dynamic Programming approach:
// Track minimum deletions and count of 'b's seen so far.
// For each 'a', we either delete it or delete all previous 'b's.
// Time: O(n), Space: O(1)

func minimumDeletions(s string) int {
	deletions := 0
	bCount := 0

	for i := 0; i < len(s); i++ {
		if s[i] == 'b' {
			bCount++
		} else {
			// s[i] == 'a'
			// Option 1: delete this 'a' (cost: deletions + 1)
			// Option 2: delete all previous 'b's (cost: bCount)
			deletions = min(deletions+1, bCount)
		}
	}

	return deletions
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
