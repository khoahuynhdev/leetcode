package main

// Approach: greedy single pass treating "))" as one closing token.
// Track unmatched '(' in open. On ')', if the next char is also ')',
// consume both; otherwise insert a ')' to complete the pair. Then match
// the pair with an open '(' or insert a '(' if none is available.
// Finally, each remaining '(' needs two ')' inserted.
// Time: O(n), Space: O(1).
func minInsertions(s string) int {
	open, ans := 0, 0
	n := len(s)
	for i := 0; i < n; {
		if s[i] == '(' {
			open++
			i++
			continue
		}
		if i+1 < n && s[i+1] == ')' {
			i += 2
		} else {
			ans++
			i++
		}
		if open > 0 {
			open--
		} else {
			ans++
		}
	}
	return ans + 2*open
}
