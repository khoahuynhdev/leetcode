package main

// Approach: 1-D dynamic programming with a rolling window.
// Let dp[j] be the minimum cost to reach step j. Step j can only be reached
// from j-1, j-2, or j-3, with jump cost costs[j] + (j-i)^2, so
//
//	dp[j] = costs[j] + min(dp[j-1]+1, dp[j-2]+4, dp[j-3]+9)
//
// Only the last three dp values are needed, giving O(n) time and O(1) space.
// Steps before 0 are treated as unreachable.
func climbStairs(n int, costs []int) int {
	const inf = 1 << 60
	// a = dp[j-3], b = dp[j-2], c = dp[j-1]
	a, b, c := inf, inf, 0
	for j := 1; j <= n; j++ {
		best := c + 1
		if b+4 < best {
			best = b + 4
		}
		if a+9 < best {
			best = a + 9
		}
		a, b, c = b, c, best+costs[j-1]
	}
	return c
}
