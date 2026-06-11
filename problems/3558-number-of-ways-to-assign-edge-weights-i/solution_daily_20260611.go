package main

// 3558. Number of Ways to Assign Edge Weights I
//
// Approach:
// We only care about the path from the root (node 1) to a node at the maximum
// depth. Every edge on that path is assigned weight 1 (odd) or 2 (even). Since
// weight 2 never changes the parity of the running sum, the total cost is odd
// exactly when an odd number of edges are set to 1. For a path with k edges,
// the number of subsets of odd size is 2^(k-1). So the answer is
// 2^(maxDepth-1) mod (1e9+7), where maxDepth is the number of edges on the
// longest root-to-leaf path. We find maxDepth with a BFS and compute the power
// with fast (binary) exponentiation.

const mod = 1_000_000_007

func assignEdgeWeights(edges [][]int) int {
	n := len(edges) + 1

	// Build adjacency list (nodes are 1-indexed).
	adj := make([][]int, n+1)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// BFS from node 1 to find the maximum depth in edges.
	visited := make([]bool, n+1)
	queue := []int{1}
	visited[1] = true
	maxDepth := 0
	for len(queue) > 0 {
		next := make([]int, 0, len(queue))
		for _, node := range queue {
			for _, nb := range adj[node] {
				if !visited[nb] {
					visited[nb] = true
					next = append(next, nb)
				}
			}
		}
		if len(next) > 0 {
			maxDepth++
		}
		queue = next
	}

	// Number of odd-cost assignments on a path of maxDepth edges is 2^(maxDepth-1).
	return powMod(2, maxDepth-1, mod)
}

// powMod computes base^exp mod m using binary exponentiation.
func powMod(base, exp, m int) int {
	base %= m
	result := 1
	for exp > 0 {
		if exp&1 == 1 {
			result = result * base % m
		}
		base = base * base % m
		exp >>= 1
	}
	return result
}
