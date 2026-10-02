package main

// Approach: backtracking with two counters.
// Place '(' while open < n, and ')' while close < open. Every string of
// length 2n built under these rules is well-formed, so no filtering is needed.
// Time: O(4^n / sqrt(n)), Space: O(n) auxiliary.
func generateParenthesis(n int) []string {
	result := []string{}
	buf := make([]byte, 2*n)

	var backtrack func(open, close int)
	backtrack = func(open, close int) {
		pos := open + close
		if pos == 2*n {
			result = append(result, string(buf))
			return
		}
		if open < n {
			buf[pos] = '('
			backtrack(open+1, close)
		}
		if close < open {
			buf[pos] = ')'
			backtrack(open, close+1)
		}
	}

	backtrack(0, 0)
	return result
}
