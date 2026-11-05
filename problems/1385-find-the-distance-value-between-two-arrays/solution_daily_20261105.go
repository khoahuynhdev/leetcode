package main

import "sort"

// findTheDistanceValue sorts arr2, then for each x in arr1 binary-searches
// for the first element >= x-d. If no such element exists or it is > x+d,
// no element of arr2 lies within distance d of x, so x is counted.
// Time: O((n+m) log m), Space: O(1) extra.
func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	sort.Ints(arr2)
	count := 0
	for _, x := range arr1 {
		i := sort.SearchInts(arr2, x-d)
		if i == len(arr2) || arr2[i] > x+d {
			count++
		}
	}
	return count
}
