package main

import "sort"

// minimizeSum returns the minimum score after changing two elements.
//
// Approach: changing two elements to values already present in the array
// makes the low score 0 for free, so we only need to minimize the high score
// (max - min). After sorting, the best we can do is drop two extreme elements:
// the two smallest, the two largest, or one of each. Take the smallest range.
func minimizeSum(nums []int) int {
	a := append([]int(nil), nums...)
	sort.Ints(a)
	n := len(a)

	return min(
		a[n-1]-a[2], // drop the two smallest
		a[n-3]-a[0], // drop the two largest
		a[n-2]-a[1], // drop the smallest and the largest
	)
}
