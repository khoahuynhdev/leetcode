package main

// 3559. Number of Ways to Assign Edge Weights II
//
// Approach:
// For a path of k edges, each edge is assigned weight 1 (flips parity) or 2
// (keeps parity). The cost is odd exactly when an odd number of edges get
// weight 1, and the number of odd-sized subsets of k items is 2^(k-1) for
// k >= 1 (and 0 for k == 0). So every query reduces to:
//
//	answer = (k == 0) ? 0 : 2^(k-1) mod (1e9+7),  where k = dist(u, v).
//
// The distance in a rooted tree is depth[u] + depth[v] - 2*depth[lca(u, v)].
// We compute depths/parents with an iterative DFS, build a binary-lifting
// table for O(log n) LCA queries, and precompute powers of two.

const modAssignEdgeWeights = 1_000_000_007

func assignEdgeWeights(edges [][]int, queries [][]int) []int {
	n := len(edges) + 1

	// Build adjacency list (nodes are 1-indexed).
	adj := make([][]int, n+1)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// Binary-lifting table dimensions.
	LOG := 1
	for (1 << LOG) < n {
		LOG++
	}
	LOG++ // headroom so the highest jump covers the deepest path

	up := make([][]int, LOG)
	for j := range up {
		up[j] = make([]int, n+1)
	}
	depth := make([]int, n+1)

	// Iterative DFS from the root (node 1) to avoid stack overflow on deep
	// trees. Record each node's immediate parent in up[0] and its depth.
	visited := make([]bool, n+1)
	stack := []int{1}
	visited[1] = true
	up[0][1] = 1 // root is its own 2^0 ancestor (sentinel)
	depth[1] = 0
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for _, next := range adj[node] {
			if !visited[next] {
				visited[next] = true
				up[0][next] = node
				depth[next] = depth[node] + 1
				stack = append(stack, next)
			}
		}
	}

	// Fill the rest of the binary-lifting table.
	for j := 1; j < LOG; j++ {
		for v := 1; v <= n; v++ {
			up[j][v] = up[j-1][up[j-1][v]]
		}
	}

	lca := func(u, v int) int {
		if depth[u] < depth[v] {
			u, v = v, u
		}
		// Lift u up to the depth of v.
		diff := depth[u] - depth[v]
		for j := 0; j < LOG; j++ {
			if diff&(1<<j) != 0 {
				u = up[j][u]
			}
		}
		if u == v {
			return u
		}
		// Lift both until just below the LCA.
		for j := LOG - 1; j >= 0; j-- {
			if up[j][u] != up[j][v] {
				u = up[j][u]
				v = up[j][v]
			}
		}
		return up[0][u]
	}

	// Precompute powers of two mod M up to n.
	pow2 := make([]int, n+1)
	pow2[0] = 1
	for i := 1; i <= n; i++ {
		pow2[i] = pow2[i-1] * 2 % modAssignEdgeWeights
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		u, v := q[0], q[1]
		k := depth[u] + depth[v] - 2*depth[lca(u, v)]
		if k == 0 {
			ans[i] = 0
		} else {
			ans[i] = pow2[k-1]
		}
	}
	return ans
}
