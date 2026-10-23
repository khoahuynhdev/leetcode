package main

// Approach: every misplaced element (nums[i] != i) must be swapped at least
// once, so k must be a submask of each of them, which means k <= AND of all
// misplaced values. That AND, M, is also achievable: the value M is a submask
// of every misplaced value, so it can act as a pivot that swaps with any of
// them. Return the AND of misplaced values, or 0 if the array is already sorted.
// Time: O(n), Space: O(1).
func sortPermutation(nums []int) int {
	res := -1
	for i, v := range nums {
		if v != i {
			res &= v
		}
	}
	if res == -1 {
		return 0
	}
	return res
}
