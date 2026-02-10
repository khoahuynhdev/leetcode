package main

// Approach: Brute force with two sets to track distinct even and odd numbers.
// For each starting position, expand rightward and check if the subarray is balanced.
// Time: O(n²), Space: O(n)

func longestBalancedSubarray(nums []int) int {
	n := len(nums)
	maxLen := 0

	// Try all possible starting positions
	for i := 0; i < n; i++ {
		evens := make(map[int]bool)
		odds := make(map[int]bool)

		// Expand rightward from position i
		for j := i; j < n; j++ {
			// Add current number to appropriate set
			if nums[j]%2 == 0 {
				evens[nums[j]] = true
			} else {
				odds[nums[j]] = true
			}

			// Check if balanced (equal count of distinct evens and odds)
			if len(evens) == len(odds) {
				length := j - i + 1
				if length > maxLen {
					maxLen = length
				}
			}
		}
	}

	return maxLen
}
