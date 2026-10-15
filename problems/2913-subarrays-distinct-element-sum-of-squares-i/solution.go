package main

// Approach: for each start index i, extend the end index j and keep a "seen"
// array of values in the current window. The distinct count grows by one only
// when nums[j] is new. Add distinct^2 after each extension. O(n^2) time,
// O(1) extra space (values are bounded by 100).
func sumCounts(nums []int) int {
	ans := 0
	for i := range nums {
		var seen [101]bool
		distinct := 0
		for j := i; j < len(nums); j++ {
			if !seen[nums[j]] {
				seen[nums[j]] = true
				distinct++
			}
			ans += distinct * distinct
		}
	}
	return ans
}
