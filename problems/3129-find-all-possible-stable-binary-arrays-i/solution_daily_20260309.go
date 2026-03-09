package main

// Approach: 3D DP where dp[i][j][last] counts valid stable binary arrays
// using i zeros and j ones with last element = last (0 or 1).
// Transition uses inclusion-exclusion to subtract arrangements that
// would create a run of more than `limit` consecutive identical elements.
func numberOfStableArrays(zero int, one int, limit int) int {
	const mod = 1_000_000_007

	// dp[i][j][k]: number of valid arrays with i zeros, j ones, ending in k
	dp := make([][][2]int, zero+1)
	for i := range dp {
		dp[i] = make([][2]int, one+1)
	}

	// Base cases: arrays of all zeros or all ones (valid if length <= limit)
	for i := 1; i <= min(zero, limit); i++ {
		dp[i][0][0] = 1
	}
	for j := 1; j <= min(one, limit); j++ {
		dp[0][j][1] = 1
	}

	for i := 1; i <= zero; i++ {
		for j := 1; j <= one; j++ {
			// Ending with 0: append a 0 to any array with (i-1) zeros and j ones
			dp[i][j][0] = (dp[i-1][j][0] + dp[i-1][j][1]) % mod
			// Subtract invalid: run of limit+1 consecutive zeros
			if i > limit {
				dp[i][j][0] = (dp[i][j][0] - dp[i-limit-1][j][1]%mod + mod) % mod
			}

			// Ending with 1: append a 1 to any array with i zeros and (j-1) ones
			dp[i][j][1] = (dp[i][j-1][0] + dp[i][j-1][1]) % mod
			// Subtract invalid: run of limit+1 consecutive ones
			if j > limit {
				dp[i][j][1] = (dp[i][j][1] - dp[i][j-limit-1][0]%mod + mod) % mod
			}
		}
	}

	return (dp[zero][one][0] + dp[zero][one][1]) % mod
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
