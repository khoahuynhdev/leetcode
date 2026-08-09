package main

// 1140. Stone Game II
//
// Approach: minimax dynamic programming over a suffix of the row plus a suffix-sum table.
//
// Because every move takes a prefix of what remains, the only state that matters is
// (i, m): the index of the first untaken pile and the current value of M. The game is
// zero-sum, so instead of tracking both players we compute dp[i][m] = the maximum number
// of stones the player *about to move* can secure from piles[i:].
//
// If that player takes X piles, the opponent then faces state (i+X, max(m, X)) and, by the
// same definition, collects dp[i+X][max(m,X)] out of the suffix[i+X] stones remaining. Since
// every remaining stone goes to somebody, the mover ends up with:
//
//	suffix[i] - dp[i+X][max(m, X)]
//
// so dp[i][m] = max over 1 <= X <= 2m (with i+X <= n) of that expression. When i+2m >= n the
// mover can sweep the rest of the row, so dp[i][m] = suffix[i]. Alice moves first with M = 1,
// hence the answer is dp[0][1].
func stoneGameII(piles []int) int {
	n := len(piles)
	if n == 0 {
		return 0
	}

	// suffix[i] = piles[i] + piles[i+1] + ... + piles[n-1]
	suffix := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		suffix[i] = suffix[i+1] + piles[i]
	}

	// M never needs to exceed n: taking more than n piles is impossible.
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := n - 1; i >= 0; i-- {
		for m := n; m >= 1; m-- {
			// The mover can take every remaining pile in one turn.
			if i+2*m >= n {
				dp[i][m] = suffix[i]
				continue
			}
			best := 0
			for x := 1; x <= 2*m; x++ {
				nextM := m
				if x > nextM {
					nextM = x
				}
				if got := suffix[i] - dp[i+x][nextM]; got > best {
					best = got
				}
			}
			dp[i][m] = best
		}
	}

	return dp[0][1]
}
