package main

import "sort"

// Path Existence Queries in a Graph II
//
// Edges depend only on values: i and j are connected when |nums[i]-nums[j]| <= maxDiff.
// Sorting nodes by value exposes the structure:
//   - Connected components are maximal runs of the sorted values whose consecutive
//     gaps are all <= maxDiff. Two nodes are reachable iff they share a component.
//   - From a sorted position i, the greedy best single hop toward the right lands at
//     far[i] = farthest index still within maxDiff (computed with two pointers).
// The shortest path between two same-component nodes is the minimum number of greedy
// jumps from the left one to the right one, answered in O(log n) per query using a
// binary-lifting table over far.
//
// Time:  O(n log n + q log n)
// Space: O(n log n)
func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []int {
	// Sort node indices by their value.
	order := make([]int, n)
	for i := range order {
		order[i] = i
	}
	sort.Slice(order, func(a, b int) bool {
		return nums[order[a]] < nums[order[b]]
	})

	sortedVals := make([]int, n) // value at each sorted rank
	pos := make([]int, n)        // pos[node] = sorted rank of original node
	for rank, node := range order {
		sortedVals[rank] = nums[node]
		pos[node] = rank
	}

	// Component id per sorted rank: a gap larger than maxDiff starts a new component.
	comp := make([]int, n)
	for i := 1; i < n; i++ {
		comp[i] = comp[i-1]
		if sortedVals[i]-sortedVals[i-1] > maxDiff {
			comp[i]++
		}
	}

	// far[i] = farthest sorted index reachable from i in one hop (rightward).
	// far is non-decreasing, so a single two-pointer sweep suffices.
	far := make([]int, n)
	j := 0
	for i := 0; i < n; i++ {
		if j < i {
			j = i
		}
		for j+1 < n && sortedVals[j+1]-sortedVals[i] <= maxDiff {
			j++
		}
		far[i] = j
	}

	// Binary lifting: up[k][i] = position after 2^k greedy jumps from i.
	LOG := 1
	for (1 << LOG) < n {
		LOG++
	}
	LOG++ // one extra level for safety (also keeps LOG >= 2 when n == 1)
	up := make([][]int, LOG)
	up[0] = far
	for k := 1; k < LOG; k++ {
		up[k] = make([]int, n)
		prev := up[k-1]
		for i := 0; i < n; i++ {
			up[k][i] = prev[prev[i]]
		}
	}

	ans := make([]int, len(queries))
	for qi, q := range queries {
		pu, pv := pos[q[0]], pos[q[1]]
		if pu == pv {
			ans[qi] = 0
			continue
		}
		if comp[pu] != comp[pv] {
			ans[qi] = -1
			continue
		}
		l, r := pu, pv
		if l > r {
			l, r = r, l
		}
		cur, steps := l, 0
		for k := LOG - 1; k >= 0; k-- {
			if up[k][cur] < r {
				cur = up[k][cur]
				steps += 1 << k
			}
		}
		ans[qi] = steps + 1 // one final greedy jump reaches r
	}
	return ans
}
