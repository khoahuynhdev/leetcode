package main

// Approach: only core "()" pairs contribute to the score. Each one is doubled
// once per enclosing pair, so a "()" at depth d adds 2^d. Scan once while
// tracking depth and add 1<<depth whenever a ')' directly follows a '('.
// Time: O(n), Space: O(1).
func scoreOfParentheses(s string) int {
	score, depth := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			depth++
			continue
		}
		depth--
		if s[i-1] == '(' {
			score += 1 << depth
		}
	}
	return score
}
