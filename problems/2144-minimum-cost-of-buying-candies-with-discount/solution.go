package main

// Greedy: sort costs in descending order, then iterate. For every group of three
// (indices 0-1-2, 3-4-5, ...) we pay for the first two and take the third for free.
// The freebie is always the smallest of its triple, satisfying the discount rule,
// and no other pairing can free up a more expensive candy.

import "sort"

func minimumCost(cost []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(cost)))
	total := 0
	for i, c := range cost {
		if i%3 != 2 {
			total += c
		}
	}
	return total
}
