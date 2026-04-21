package main

// Approach: Union-Find over indices. Each allowed swap is an edge, and swaps
// are transitive, so within a connected component of indices the source values
// can be reordered freely. For each component, the maximum number of positions
// that match target equals the multiset intersection size of its source and
// target values: sum of min(srcCount[v], tgtCount[v]) over all v. The answer
// is n minus the total matches across all components.

func minimumHammingDistanceDaily20260421(source []int, target []int, allowedSwaps [][]int) int {
	n := len(source)
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	union := func(a, b int) {
		ra, rb := find(a), find(b)
		if ra != rb {
			parent[ra] = rb
		}
	}

	for _, sw := range allowedSwaps {
		union(sw[0], sw[1])
	}

	srcCount := make(map[int]map[int]int)
	tgtCount := make(map[int]map[int]int)
	for i := 0; i < n; i++ {
		r := find(i)
		if srcCount[r] == nil {
			srcCount[r] = make(map[int]int)
			tgtCount[r] = make(map[int]int)
		}
		srcCount[r][source[i]]++
		tgtCount[r][target[i]]++
	}

	matches := 0
	for r, sm := range srcCount {
		tm := tgtCount[r]
		for v, c := range sm {
			tc := tm[v]
			if tc < c {
				matches += tc
			} else {
				matches += c
			}
		}
	}
	return n - matches
}
