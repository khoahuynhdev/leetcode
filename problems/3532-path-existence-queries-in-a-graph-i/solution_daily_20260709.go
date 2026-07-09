package main

// Path Existence Queries in a Graph I
//
// Because nums is sorted non-decreasing, an edge exists between adjacent nodes
// i and i+1 exactly when nums[i+1]-nums[i] <= maxDiff. No edge can ever cross a
// gap larger than maxDiff (any wider pair spans at least that gap), so every
// connected component is a contiguous block of indices.
//
// We label components in a single left-to-right pass, then answer each query in
// O(1) by comparing the two nodes' component labels.
//
// Time: O(n + q), Space: O(n).
func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []bool {
	comp := make([]int, n)
	id := 0
	for i := 1; i < n; i++ {
		if nums[i]-nums[i-1] > maxDiff {
			id++
		}
		comp[i] = id
	}

	answer := make([]bool, len(queries))
	for i, q := range queries {
		answer[i] = comp[q[0]] == comp[q[1]]
	}
	return answer
}
