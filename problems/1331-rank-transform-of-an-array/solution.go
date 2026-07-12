package main

import "sort"

// arrayRankTransform replaces each element with its rank.
//
// Approach: the rank of a value equals 1 plus the number of distinct values
// strictly smaller than it. We sort a copy of the array, then walk it once to
// build a value -> rank map, assigning a new rank only when we meet a value we
// have not seen before (so equal values share a rank and there are no gaps).
// Finally we map the original array back through the lookup, preserving order.
//
// Time: O(n log n) for the sort. Space: O(n) for the copy, map, and result.
func arrayRankTransform(arr []int) []int {
	sorted := make([]int, len(arr))
	copy(sorted, arr)
	sort.Ints(sorted)

	rankOf := make(map[int]int, len(sorted))
	rank := 1
	for _, v := range sorted {
		if _, seen := rankOf[v]; !seen {
			rankOf[v] = rank
			rank++
		}
	}

	result := make([]int, len(arr))
	for i, v := range arr {
		result[i] = rankOf[v]
	}
	return result
}
