package main

// countBinarySubstrings counts substrings with equal consecutive 0s and 1s.
// Track previous and current group lengths; at each group boundary,
// min(prev, curr) valid substrings exist. O(n) time, O(1) space.
func countBinarySubstrings(s string) int {
	prev, curr := 0, 1
	result := 0

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			curr++
		} else {
			result += min(prev, curr)
			prev = curr
			curr = 1
		}
	}
	result += min(prev, curr)

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
