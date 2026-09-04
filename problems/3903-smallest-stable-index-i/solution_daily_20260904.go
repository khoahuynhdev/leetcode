package main

// 3903. Smallest Stable Index I
//
// The instability score at index i splits into a prefix fact and a suffix fact:
// max(nums[0..i]) and min(nums[i..n-1]), both inclusive of i. Precompute the
// suffix minima in one right-to-left pass, then sweep left to right carrying a
// running prefix maximum and return the first index whose score is <= k.
//
// The score is not monotonic in i (both prefMax and sufMin are non-decreasing,
// so their difference can rise and fall), which is why we scan rather than
// binary search.
//
// Time: O(n), Space: O(n).
func smallestStableIndex(nums []int, k int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}

	// sufMin[i] == min(nums[i..n-1])
	sufMin := make([]int, n)
	sufMin[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		sufMin[i] = min(nums[i], sufMin[i+1])
	}

	prefMax := nums[0]
	for i := 0; i < n; i++ {
		// Fold nums[i] in first, so prefMax == max(nums[0..i]) at the test.
		prefMax = max(prefMax, nums[i])
		if prefMax-sufMin[i] <= k {
			return i
		}
	}
	return -1
}
