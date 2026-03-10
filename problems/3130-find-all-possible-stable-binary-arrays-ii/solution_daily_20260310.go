package main

// numberOfStableArrays counts binary arrays with exactly `zero` 0s and `one` 1s
// where no subarray of length > limit is all-same. Uses DP with a virtual base
// state to handle start-of-array boundaries in the inclusion-exclusion correction.
func numberOfStableArrays(zero int, one int, limit int) int {
	const mod = 1_000_000_007

	// dp[i][j][k]: number of stable arrays using i zeros, j ones, ending with k
	dp := make([][][2]int, zero+1)
	for i := range dp {
		dp[i] = make([][2]int, one+1)
	}

	// getV returns dp[i][j][k] but treats (0,0) as a virtual state returning 1.
	// This handles the boundary where a run of identical values starts at the
	// beginning of the array (no preceding opposite value in the DP table).
	getV := func(i, j, k int) int {
		if i == 0 && j == 0 {
			return 1
		}
		return dp[i][j][k]
	}

	for i := 0; i <= zero; i++ {
		for j := 0; j <= one; j++ {
			if i == 0 && j == 0 {
				continue
			}
			// Ending with 0: append a 0 to arrays with (i-1) zeros and j ones
			if i >= 1 {
				dp[i][j][0] = (dp[i-1][j][0] + getV(i-1, j, 1)) % mod
				if i > limit {
					dp[i][j][0] = (dp[i][j][0] - getV(i-limit-1, j, 1) + mod) % mod
				}
			}
			// Ending with 1: append a 1 to arrays with i zeros and (j-1) ones
			if j >= 1 {
				dp[i][j][1] = (dp[i][j-1][1] + getV(i, j-1, 0)) % mod
				if j > limit {
					dp[i][j][1] = (dp[i][j][1] - getV(i, j-limit-1, 0) + mod) % mod
				}
			}
		}
	}

	return (dp[zero][one][0] + dp[zero][one][1]) % mod
}
