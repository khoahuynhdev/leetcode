package main

// Approach: count-guided backtracking.
// First, scan once to compute the exact number of unmatched '(' (leftRem) and
// unmatched ')' (rightRem). Their sum is the minimum number of removals. Then
// DFS over the string: for each parenthesis, either remove it (if budget for
// that type remains) or keep it (if keeping doesn't make the prefix invalid).
// We prune when the running balance would go negative or when the remaining
// removal budget exceeds the characters left. A set deduplicates results.
//
// Time: O(2^n * n) worst case, heavily pruned in practice. Space: O(n) recursion.
func removeInvalidParentheses(s string) []string {
	leftRem, rightRem := 0, 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			leftRem++
		case ')':
			if leftRem > 0 {
				leftRem--
			} else {
				rightRem++
			}
		}
	}

	seen := make(map[string]bool)
	result := []string{}
	buf := make([]byte, 0, len(s))

	var dfs func(i, open, lRem, rRem int)
	dfs = func(i, open, lRem, rRem int) {
		if lRem+rRem > len(s)-i {
			return
		}
		if i == len(s) {
			if open == 0 {
				str := string(buf)
				if !seen[str] {
					seen[str] = true
					result = append(result, str)
				}
			}
			return
		}

		c := s[i]

		// Option 1: remove this parenthesis if we still have budget for its type.
		if c == '(' && lRem > 0 {
			dfs(i+1, open, lRem-1, rRem)
		} else if c == ')' && rRem > 0 {
			dfs(i+1, open, lRem, rRem-1)
		}

		// Option 2: keep this character.
		buf = append(buf, c)
		switch c {
		case '(':
			dfs(i+1, open+1, lRem, rRem)
		case ')':
			if open > 0 {
				dfs(i+1, open-1, lRem, rRem)
			}
		default:
			dfs(i+1, open, lRem, rRem)
		}
		buf = buf[:len(buf)-1]
	}

	dfs(0, 0, leftRem, rightRem)
	return result
}
