package main

import "sort"

// Approach: Sort array and use sliding window to find longest balanced subarray
// A balanced subarray after sorting is contiguous with max <= k * min
// Answer is n - (length of longest valid subarray)

func minimumRemovals(nums []int, k int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}
	
	// Sort to make any valid balanced subarray contiguous
	sort.Ints(nums)
	
	maxKeep := 1 // At minimum, we can always keep one element
	left := 0
	
	// Sliding window: find longest subarray where max <= k * min
	for right := 0; right < n; right++ {
		// nums[right] is max, nums[left] is min in current window
		// Check if current window is balanced
		for left <= right && nums[right] > k*nums[left] {
			left++
		}
		
		// Update maximum elements we can keep
		if left <= right {
			maxKeep = max(maxKeep, right-left+1)
		}
	}
	
	return n - maxKeep
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}