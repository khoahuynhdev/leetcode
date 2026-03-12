package main

// Binary search on the answer + Union-Find greedy check.
// For a candidate stability mid, verify that we can build a spanning tree
// where every edge has effective strength >= mid using at most k upgrades.

func maxStability(n int, edges [][]int, k int) int {
	// Check if must-edges form a cycle (always invalid).
	uf := newUF(n)
	for _, e := range edges {
		if e[3] == 1 {
			if !uf.union(e[0], e[1]) {
				return -1
			}
		}
	}

	// Binary search on the stability value.
	lo, hi := 1, 200001
	ans := -1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if canAchieve(n, edges, k, mid) {
			ans = mid
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return ans
}

func canAchieve(n int, edges [][]int, k int, mid int) bool {
	uf := newUF(n)

	// 1. Add all must-edges. If any has strength < mid, fail.
	for _, e := range edges {
		if e[3] == 1 {
			if e[2] < mid {
				return false
			}
			uf.union(e[0], e[1])
		}
	}

	// 2. Add optional edges that meet threshold without upgrade (free).
	for _, e := range edges {
		if e[3] == 0 && e[2] >= mid {
			uf.union(e[0], e[1])
		}
	}

	// 3. Add optional edges that need an upgrade (2*s >= mid).
	upgrades := 0
	for _, e := range edges {
		if e[3] == 0 && e[2] < mid && 2*e[2] >= mid {
			if uf.union(e[0], e[1]) {
				upgrades++
			}
		}
	}

	return uf.count == 1 && upgrades <= k
}

type unionFind struct {
	parent []int
	rank   []int
	count  int
}

func newUF(n int) *unionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &unionFind{parent: parent, rank: rank, count: n}
}

func (u *unionFind) find(x int) int {
	for u.parent[x] != x {
		u.parent[x] = u.parent[u.parent[x]]
		x = u.parent[x]
	}
	return x
}

func (u *unionFind) union(x, y int) bool {
	rx, ry := u.find(x), u.find(y)
	if rx == ry {
		return false
	}
	if u.rank[rx] < u.rank[ry] {
		rx, ry = ry, rx
	}
	u.parent[ry] = rx
	if u.rank[rx] == u.rank[ry] {
		u.rank[rx]++
	}
	u.count--
	return true
}
