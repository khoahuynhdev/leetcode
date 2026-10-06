package main

// Greedy balance counting: scan left to right tracking unmatched '('.
// A ')' with no pending '(' requires inserting a '(' (counted in adds).
// Any '(' still unmatched at the end requires inserting a ')'.
// Time: O(n), Space: O(1).
func minAddToMakeValid(s string) int {
	open, adds := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			open++
		} else if open > 0 {
			open--
		} else {
			adds++
		}
	}
	return adds + open
}
