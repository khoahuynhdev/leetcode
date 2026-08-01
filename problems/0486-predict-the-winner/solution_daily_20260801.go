package main

// 0486. Predict the Winner
//
// Interval DP over score differences (negamax). Because both players only take
// from the ends, every reachable game state is a contiguous window nums[i..j].
// Let dp[i][j] be the best value of (mover's score - opponent's score) on that
// window with optimal play. Taking nums[i] leaves the opponent as the mover on
// nums[i+1..j], and their advantage counts against us:
//
//	dp[i][j] = max(nums[i]-dp[i+1][j], nums[j]-dp[i][j-1])
//	dp[i][i] = nums[i]
//
// Player 1 wins or ties exactly when dp[0][n-1] >= 0 (ties count as a win).
//
// Since dp[i][j] reads only dp[i+1][j] and dp[i][j-1], a single rolling row of
// length n is enough: iterating i downward, dp[j] still holds the previous row
// when read, and dp[j-1] already holds the current row.
//
// Time: O(n^2), Space: O(n).
func PredictTheWinner(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return true
	}

	// dp[j] holds dp[i][j] for the current i, and dp[i+1][j] before it is written.
	dp := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		dp[i] = nums[i] // base case: dp[i][i]
		for j := i + 1; j < n; j++ {
			takeLeft := nums[i] - dp[j]    // dp[j] is still dp[i+1][j]
			takeRight := nums[j] - dp[j-1] // dp[j-1] is already dp[i][j-1]
			dp[j] = max(takeLeft, takeRight)
		}
	}

	return dp[n-1] >= 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
