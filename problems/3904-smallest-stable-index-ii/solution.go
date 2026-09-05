package main

// 3904. Smallest Stable Index II
//
// The instability score at index i is max(nums[0..i]) - min(nums[i..n-1]).
// Both windows change by one element per step, so instead of recomputing them
// we precompute a suffix-minimum array in one right-to-left pass, then sweep
// left to right carrying a running prefix maximum and return the first index
// whose score is <= k.
//
// Note that prefMax and sufMin are both non-decreasing, so their difference is
// not monotonic (e.g. [1,3,2] gives scores 0,1,1) and binary search does not
// apply -- the linear scan is required.
//
// Time: O(n), Space: O(n).
func smallestStableIndex(nums []int, k int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}

	// sufMin[i] = min(nums[i..n-1])
	sufMin := make([]int, n)
	sufMin[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		sufMin[i] = min(nums[i], sufMin[i+1])
	}

	prefMax := nums[0]
	for i := 0; i < n; i++ {
		prefMax = max(prefMax, nums[i])
		if prefMax-sufMin[i] <= k {
			return i
		}
	}
	return -1
}
