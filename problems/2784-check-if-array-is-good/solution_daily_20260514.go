package main

// Approach: An array is good iff it is a permutation of base[n] = [1, 2, ..., n-1, n, n].
// The candidate n is forced to be max(nums). After confirming len(nums) == n+1, sort the
// array and verify nums[i] == i+1 for i in [0, n-1]; the last slot is then guaranteed to
// be n as well. Time O(k log k) for the sort, extra space O(1).

import "sort"

func isGood(nums []int) bool {
	n := 0
	for _, v := range nums {
		if v > n {
			n = v
		}
	}
	if len(nums) != n+1 {
		return false
	}
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return false
		}
	}
	return nums[n] == n
}
