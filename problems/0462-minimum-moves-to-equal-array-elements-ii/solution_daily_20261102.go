package main

import "sort"

// minMoves2 returns the minimum number of +1/-1 moves to make all elements equal.
// Approach: the sum of absolute deviations is minimized at the median. After
// sorting, pair the smallest with the largest element; each pair (lo, hi)
// costs hi - lo for any target between them, and the median lies between
// every such pair. Summing these differences gives the answer.
// Time: O(n log n), Space: O(1) extra.
func minMoves2(nums []int) int {
	sort.Ints(nums)
	moves := 0
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		moves += nums[j] - nums[i]
	}
	return moves
}
