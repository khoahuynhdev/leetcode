package main

// Approach: Union-Find (Disjoint Set Union).
//
// 1. Union every edge so vertices in the same connected component share a root.
//    Track the number of vertices in each component via a size array.
// 2. Count how many edges lie inside each component by tallying edges per root.
// 3. A component with k vertices is "complete" iff it holds exactly
//    k*(k-1)/2 edges (every pair connected). Since the input has no repeated
//    edges, matching this maximum guarantees completeness.
//
// Time:  O(V + E * α(V))   Space: O(V)

type dsu struct {
	parent []int
	size   []int
}

func newDSU(n int) *dsu {
	d := &dsu{
		parent: make([]int, n),
		size:   make([]int, n),
	}
	for i := 0; i < n; i++ {
		d.parent[i] = i
		d.size[i] = 1
	}
	return d
}

func (d *dsu) find(x int) int {
	for d.parent[x] != x {
		d.parent[x] = d.parent[d.parent[x]] // path compression
		x = d.parent[x]
	}
	return x
}

func (d *dsu) union(a, b int) {
	ra, rb := d.find(a), d.find(b)
	if ra == rb {
		return
	}
	if d.size[ra] < d.size[rb] {
		ra, rb = rb, ra
	}
	d.parent[rb] = ra
	d.size[ra] += d.size[rb]
}

func countCompleteComponents(n int, edges [][]int) int {
	d := newDSU(n)
	for _, e := range edges {
		d.union(e[0], e[1])
	}

	// Count edges belonging to each component root.
	edgeCount := make([]int, n)
	for _, e := range edges {
		edgeCount[d.find(e[0])]++
	}

	complete := 0
	for v := 0; v < n; v++ {
		if d.find(v) != v {
			continue // only inspect component roots
		}
		k := d.size[v]
		if edgeCount[v] == k*(k-1)/2 {
			complete++
		}
	}
	return complete
}
