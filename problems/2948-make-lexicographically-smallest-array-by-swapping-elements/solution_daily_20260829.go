package main

import "sort"

// 2948. Make Lexicographically Smallest Array by Swapping Elements
//
// Repeated swaps make "swappable" transitive, so the elements split into
// connected components. Sorting the values makes each component a contiguous
// run: a new run starts wherever consecutive sorted values differ by more than
// limit. Within a run every permutation of its members is reachable, so we take
// the run's original indices, sort them ascending, and drop the run's ascending
// values into them in order -- smallest value into the earliest position, which
// is exactly what lexicographic minimality requires.
//
// Time: O(n log n), Space: O(n).
func lexicographicallySmallestArrayDaily20260829(nums []int, limit int) []int {
	n := len(nums)
	res := make([]int, n)
	if n == 0 {
		return res
	}

	// order holds the original indices sorted by their value ascending.
	order := make([]int, n)
	for i := range order {
		order[i] = i
	}
	sort.Slice(order, func(a, b int) bool {
		return nums[order[a]] < nums[order[b]]
	})

	// Sweep the sorted order, cutting a group whenever the value gap exceeds limit.
	start := 0
	for k := 1; k <= n; k++ {
		if k == n || nums[order[k]]-nums[order[k-1]] > limit {
			group := order[start:k]

			// group's values are already ascending; pair them with ascending indices.
			idx := make([]int, len(group))
			copy(idx, group)
			sort.Ints(idx)

			for t, pos := range idx {
				res[pos] = nums[group[t]]
			}
			start = k
		}
	}

	return res
}
