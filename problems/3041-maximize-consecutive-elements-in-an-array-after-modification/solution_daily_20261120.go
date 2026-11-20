package main

import "sort"

// Approach: Sort + DP keyed by value.
// dp[v] = length of the longest consecutive run ending exactly at value v,
// built from elements processed so far. Processing in sorted order, each
// element x can become either x or x+1:
//   - used as x+1: dp[x+1] = dp[x] + 1   (computed first, using the old dp[x])
//   - used as x:   dp[x]   = dp[x-1] + 1
// Updating x+1 before x guarantees one element is never counted twice.
// Time: O(n log n), Space: O(n).
func maxSelectedElements(nums []int) int {
	sorted := append([]int(nil), nums...)
	sort.Ints(sorted)

	dp := make(map[int]int, 2*len(sorted))
	best := 0
	for _, x := range sorted {
		dp[x+1] = dp[x] + 1
		dp[x] = dp[x-1] + 1
		best = max(best, dp[x], dp[x+1])
	}
	return best
}
