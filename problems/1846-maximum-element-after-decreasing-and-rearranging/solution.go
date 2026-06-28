package main

import "sort"

// maximumElementAfterDecreasing solves LeetCode 1846.
//
// Approach: Rearranging is free, so sort arr ascending. Then walk left to right
// keeping prev, the value committed to the previous position (starting at 1, the
// required first value). At each position the best legal value is
// min(arr[i], prev+1): we may decrease arr[i] down to prev+1, but we can never
// increase it, and we can never exceed prev+1. The answer is prev after the
// scan, which lands on the largest (last) sorted element.
//
// Time: O(n log n) for the sort. Space: O(1) extra.
func maximumElementAfterDecreasing(arr []int) int {
	sort.Ints(arr)

	prev := 1
	for i := 1; i < len(arr); i++ {
		if arr[i] < prev+1 {
			prev = arr[i]
		} else {
			prev++
		}
	}
	return prev
}
