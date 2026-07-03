package main

// 3620. Network Recovery Pathways
//
// We want the largest "bottleneck" score x such that there is a path from
// node 0 to node n-1 where every edge costs at least x, all intermediate
// nodes are online, and the total edge cost is at most k.
//
// Approach: binary search on the answer x (over the sorted distinct edge
// costs). For a fixed x, keep only edges with cost >= x and ask whether the
// cheapest 0 -> n-1 path costs at most k. Because the graph is a DAG, a single
// pass in topological order computes those minimum costs. Feasibility is
// monotonic in x, so binary search finds the largest feasible x, or -1.
//
// Time:  O((n + m) log m)   Space: O(n + m)

import "sort"

type edge struct {
	to   int
	cost int64
}

func findMaxPathScore(edges [][]int, online []bool, k int64) int64 {
	n := len(online)

	// Build adjacency list and in-degrees for topological sorting.
	adj := make([][]edge, n)
	indeg := make([]int, n)
	costSet := make(map[int64]struct{})
	for _, e := range edges {
		u, v, c := e[0], e[1], int64(e[2])
		adj[u] = append(adj[u], edge{to: v, cost: c})
		indeg[v]++
		costSet[c] = struct{}{}
	}

	if len(costSet) == 0 {
		return -1 // no edges => no path
	}

	// Kahn's algorithm: fixed topological order of the DAG.
	topo := make([]int, 0, n)
	queue := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if indeg[i] == 0 {
			queue = append(queue, i)
		}
	}
	deg := make([]int, n)
	copy(deg, indeg)
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		topo = append(topo, u)
		for _, e := range adj[u] {
			deg[e.to]--
			if deg[e.to] == 0 {
				queue = append(queue, e.to)
			}
		}
	}

	// Distinct edge costs, ascending: the only candidate scores.
	costs := make([]int64, 0, len(costSet))
	for c := range costSet {
		costs = append(costs, c)
	}
	sort.Slice(costs, func(i, j int) bool { return costs[i] < costs[j] })

	const inf = int64(1) << 62
	dist := make([]int64, n)

	// feasible reports whether some 0 -> n-1 path using only edges with
	// cost >= x has total cost <= k.
	feasible := func(x int64) bool {
		for i := range dist {
			dist[i] = inf
		}
		dist[0] = 0
		for _, u := range topo {
			if dist[u] == inf || !online[u] {
				// Unreachable, or an offline node we may not route through.
				continue
			}
			du := dist[u]
			for _, e := range adj[u] {
				if e.cost < x {
					continue
				}
				if nd := du + e.cost; nd < dist[e.to] {
					dist[e.to] = nd
				}
			}
		}
		return dist[n-1] != inf && dist[n-1] <= k
	}

	// Largest candidate cost that stays feasible.
	ans := int64(-1)
	lo, hi := 0, len(costs)-1
	for lo <= hi {
		mid := (lo + hi) / 2
		if feasible(costs[mid]) {
			ans = costs[mid]
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return ans
}
