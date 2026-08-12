package main

// 2958. Length of Longest Subarray With at Most K Frequency
//
// Sliding window with a frequency map. A subarray is "good" if no value occurs
// more than k times; that property is downward closed, so shrinking a bad window
// from the left eventually makes it good and the left pointer never rewinds.
//
// After adding nums[right] the only count that grew is count[nums[right]], and
// the window was good one step earlier, so it is the only possible violator.
// Shrink from the left until it fits, then record the window length.
//
// Time: O(n), Space: O(n) for the frequency map.
func maxSubarrayLength(nums []int, k int) int {
	count := make(map[int]int)
	best, left := 0, 0

	for right, v := range nums {
		count[v]++
		for count[v] > k {
			count[nums[left]]--
			left++
		}
		if length := right - left + 1; length > best {
			best = length
		}
	}

	return best
}
