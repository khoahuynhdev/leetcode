package main

// Approach: Any even-length uniform block can be split into uniform blocks of
// length 2, so a string is beautiful iff every aligned pair (s[i], s[i+1]) with
// even i has equal characters. The pairs are independent, and each mismatched
// pair needs exactly one change, so the answer is the number of mismatched pairs.
// Time: O(n), Space: O(1).
func minChanges(s string) int {
	count := 0
	for i := 0; i+1 < len(s); i += 2 {
		if s[i] != s[i+1] {
			count++
		}
	}
	return count
}
