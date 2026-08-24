package main

// 1872. Stone Game VIII
//
// Key observation: after any sequence of moves the row is always
// [p[j], stones[j+1], ..., stones[n-1]] where p is the prefix-sum array of the
// original stones. So a single index j describes the whole game state, and a
// move that ends at original index k scores exactly p[k] -- the merged left
// stone already carries every earlier value.
//
// The game is zero sum, so a single DP from the mover's point of view suffices:
//
//	dp[i] = best (mover - opponent) difference when the mover may finish their
//	        take at any index >= i
//	dp[n-1] = p[n-1]
//	dp[i]   = max(p[i]-dp[i+1], dp[i+1])
//
// The second branch says "don't stop at i, reuse the best option from i+1",
// which collapses an O(n) inner scan into O(1) work per index.
//
// Alice must remove x > 1 stones on her first move, so she cannot stop at
// index 0 and the answer is dp[1].
//
// Prefix sums are walked backwards by subtracting from the running total, so no
// auxiliary array is needed and the input is left untouched.
//
// Time: O(n), Space: O(1).
func stoneGameVIII(stones []int) int {
	n := len(stones)

	// prefix holds p[i] as i walks down from n-1 to 1.
	prefix := 0
	for _, v := range stones {
		prefix += v
	}

	// Base case: at the last index the mover is forced to sweep everything.
	dp := prefix
	for i := n - 1; i >= 2; i-- {
		prefix -= stones[i] // now prefix == p[i-1]
		dp = max(prefix-dp, dp)
	}
	return dp
}
