package main

// Approach: Greedy + Sorting.
// Signs vanish when squaring, so sort elements by absolute value. An array of
// length n has (n+1)/2 added slots (even indices) and n/2 subtracted slots
// (odd indices). Assign the smallest n/2 squares to the subtracted slots and
// the rest to the added slots. By an exchange argument this maximizes the score.
// Time: O(n log n), Space: O(n).

import "sort"

func maxAlternatingSum(nums []int) int64 {
	n := len(nums)
	abs := make([]int, n)
	for i, x := range nums {
		if x < 0 {
			x = -x
		}
		abs[i] = x
	}
	sort.Ints(abs)

	var score int64
	neg := n / 2
	for i, x := range abs {
		sq := int64(x) * int64(x)
		if i < neg {
			score -= sq
		} else {
			score += sq
		}
	}
	return score
}
