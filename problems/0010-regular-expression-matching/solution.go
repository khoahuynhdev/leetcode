package main

// Dynamic Programming approach:
// dp[i][j] represents whether s[0...i-1] matches p[0...j-1]
// For each position, we check:
// 1. If current pattern char is not '*', it must match current string char (or be '.')
// 2. If current pattern char is '*', we can either skip it (zero occurrences) or use it (one or more occurrences)

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)

	// dp[i][j] = true if s[0...i-1] matches p[0...j-1]
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	// Base case: empty string matches empty pattern
	dp[0][0] = true

	// Initialize first row: empty string with patterns like a*, a*b*, etc.
	for j := 2; j <= n; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	// Fill the DP table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				// Star can match zero occurrences (skip pattern char and star)
				dp[i][j] = dp[i][j-2]

				// Or star can match one or more occurrences
				// Check if current char matches the char before '*'
				charMatch := s[i-1] == p[j-2] || p[j-2] == '.'
				if charMatch {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			} else {
				// Current pattern char must match current string char (or be '.')
				charMatch := s[i-1] == p[j-1] || p[j-1] == '.'
				if charMatch {
					dp[i][j] = dp[i-1][j-1]
				}
			}
		}
	}

	return dp[m][n]
}
