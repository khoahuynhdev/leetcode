package main

// 0877. Stone Game
//
// Interval DP on score difference. dp[i][j] is the largest lead (my stones
// minus the opponent's) that the player to move can guarantee on the window
// piles[i..j]. Taking an end pile hands the smaller window to the opponent,
// whose own best lead counts against us:
//
//	dp[i][j] = max(piles[i]-dp[i+1][j], piles[j]-dp[i][j-1])
//	dp[i][i] = piles[i]
//
// Alice moves first on the whole row, so she wins iff dp[0][n-1] > 0.
// Windows are processed by increasing length, and length L only reads length
// L-1, so a single rolling array suffices: O(n^2) time, O(n) space.
//
// (There is also an O(1) answer — with an even pile count and an odd total,
// Alice can force all even-indexed or all odd-indexed piles and picks the
// bigger group, so the result is always true. The DP is kept because it also
// solves the general game, including an odd number of piles.)
func stoneGame(piles []int) bool {
	n := len(piles)
	if n == 0 {
		return false
	}

	// dp[i] holds dp[i][i+length-1] for the length currently being processed.
	dp := make([]int, n)
	copy(dp, piles) // length 1: dp[i][i] = piles[i]

	for length := 2; length <= n; length++ {
		// Sweep i upward so dp[i+1] is still the length-1 shorter value.
		for i := 0; i+length-1 < n; i++ {
			j := i + length - 1
			dp[i] = max(piles[i]-dp[i+1], piles[j]-dp[i])
		}
	}

	// Strictly greater: a tie is not a win for Alice.
	return dp[0] > 0
}
