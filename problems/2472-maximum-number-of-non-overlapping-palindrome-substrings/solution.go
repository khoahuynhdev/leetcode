package main

// 2472. Maximum Number of Non-overlapping Palindrome Substrings
//
// Two phases:
//  1. Precompute isPal[i][j] ("s[i..j] is a palindrome") with the recurrence
//     isPal[i][j] = s[i] == s[j] && (j-i < 2 || isPal[i+1][j-1]), filling i
//     downward so the inner substring is always ready. Stored as one flat
//     []bool indexed i*n+j to keep the 2000x2000 worst case in a single 4 MB
//     allocation.
//  2. DP over suffixes: dp[i] = best answer for s[i:], with
//     dp[i] = max(dp[i+1], 1+dp[j+1]) over palindromic s[i..j] of length >= k.
//     Since dp is non-increasing, the shortest palindrome starting at i is
//     always at least as good as a longer one, so the scan breaks on the first
//     match.
func maxPalindromes(s string, k int) int {
	n := len(s)
	if n == 0 || k > n {
		return 0
	}

	isPal := make([]bool, n*n)
	for i := n - 1; i >= 0; i-- {
		row := i * n
		for j := i; j < n; j++ {
			if s[i] == s[j] && (j-i < 2 || isPal[row+n+j-1]) {
				isPal[row+j] = true
			}
		}
	}

	dp := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		dp[i] = dp[i+1]
		row := i * n
		for j := i + k - 1; j < n; j++ {
			if isPal[row+j] {
				// Shortest valid palindrome at i wins; no need to look further.
				if best := 1 + dp[j+1]; best > dp[i] {
					dp[i] = best
				}
				break
			}
		}
	}

	return dp[0]
}
