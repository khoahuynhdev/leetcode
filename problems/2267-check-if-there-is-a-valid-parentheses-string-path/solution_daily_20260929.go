package main

// Approach: DP over (row, col, balance).
// A path string is valid iff its running balance ('(' = +1, ')' = -1) never
// drops below zero and ends at exactly zero. Every path has length m+n-1, so
// it must be even. For each cell we track the set of balances reachable by
// some path from (0,0), using a rolling row of boolean sets. A balance above
// the number of remaining cells can never return to zero, so it is pruned.
// Time: O(m*n*(m+n)), Space: O(n*(m+n)).
func hasValidPath(grid [][]byte) bool {
	m, n := len(grid), len(grid[0])
	total := m + n - 1
	if total%2 == 1 || grid[0][0] == ')' || grid[m-1][n-1] == '(' {
		return false
	}

	maxBal := total/2 + 1
	// dp[j][b] is true if balance b is reachable at (current row, j).
	dp := make([][]bool, n)
	for j := range dp {
		dp[j] = make([]bool, maxBal+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			delta := 1
			if grid[i][j] == ')' {
				delta = -1
			}
			remaining := total - (i + j + 1) // cells left after this one

			next := make([]bool, maxBal+1)
			if i == 0 && j == 0 {
				next[1] = true
			} else {
				for b := 0; b <= maxBal; b++ {
					// dp[j] still holds the cell above; dp[j-1] is already the cell to the left.
					fromTop := i > 0 && dp[j][b]
					fromLeft := j > 0 && dp[j-1][b]
					if !fromTop && !fromLeft {
						continue
					}
					nb := b + delta
					if nb < 0 || nb > remaining || nb > maxBal {
						continue
					}
					next[nb] = true
				}
			}
			dp[j] = next
		}
	}

	return dp[n-1][0]
}
