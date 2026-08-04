package main

// 3731. Find Missing Elements
//
// The original range is exactly [min(nums), max(nums)], because the problem
// guarantees both endpoints survived. So: scan once to recover those bounds,
// scan again to mark which values of the range are present in a boolean slice
// offset by min, then walk the range in increasing order and collect every
// unmarked value. Walking upward means the output is already sorted, so no
// sorting step is needed.
//
// Time: O(n + R) where R = max-min+1. Space: O(R) auxiliary.
func findMissingElements(nums []int) []int {
	lo, hi := nums[0], nums[0]
	for _, v := range nums {
		if v < lo {
			lo = v
		}
		if v > hi {
			hi = v
		}
	}

	seen := make([]bool, hi-lo+1)
	for _, v := range nums {
		seen[v-lo] = true
	}

	// Non-nil so that "nothing missing" returns [] rather than null.
	missing := make([]int, 0)
	for i, ok := range seen {
		if !ok {
			missing = append(missing, lo+i)
		}
	}
	return missing
}
