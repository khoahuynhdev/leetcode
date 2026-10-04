package main

// Approach: greedy range tracking.
// Track the minimum (lo) and maximum (hi) number of unmatched '(' possible
// after each prefix. '*' widens the range in both directions. If hi < 0 the
// string has too many ')' and can never be valid. lo is clamped at 0 since a
// negative balance is never a valid state. The string is valid iff lo == 0 at
// the end. O(n) time, O(1) space.
func checkValidString(s string) bool {
	lo, hi := 0, 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			lo++
			hi++
		case ')':
			lo--
			hi--
		default: // '*'
			lo--
			hi++
		}
		if hi < 0 {
			return false
		}
		if lo < 0 {
			lo = 0
		}
	}
	return lo == 0
}
