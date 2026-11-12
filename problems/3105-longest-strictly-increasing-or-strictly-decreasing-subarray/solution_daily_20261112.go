package main

// Approach: single-pass run-length scan.
// Track the length of the strictly increasing run and the strictly decreasing
// run ending at the current index. Each step, one run grows (or neither, on
// equality) and the others reset to 1. The answer is the maximum run seen.
// Time: O(n), Space: O(1).
func longestMonotonicSubarray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	inc, dec, best := 1, 1, 1
	for i := 1; i < len(nums); i++ {
		switch {
		case nums[i] > nums[i-1]:
			inc++
			dec = 1
		case nums[i] < nums[i-1]:
			dec++
			inc = 1
		default:
			inc, dec = 1, 1
		}
		best = max(best, inc, dec)
	}
	return best
}
