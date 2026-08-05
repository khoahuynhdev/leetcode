package main

// 3310. Remove Methods From Project
//
// Approach: treat invocations as directed edges a -> b.
//
//  1. Mark the suspicious set with a single iterative DFS from k. Because the
//     graph may contain cycles, a visited array is required.
//  2. The group is removable iff no edge crosses into the set from outside,
//     i.e. there is no edge (a, b) with a not suspicious and b suspicious.
//     The reverse direction cannot happen: reachability from k is closed under
//     following edges, so a suspicious tail always has a suspicious head.
//  3. If some edge blocks removal, nothing is removed and every method remains;
//     otherwise the answer is every non-suspicious method.
//
// Time: O(n + m), Space: O(n + m) where m = len(invocations).
func remainingMethods(n int, k int, invocations [][]int) []int {
	adj := make([][]int, n)
	for _, e := range invocations {
		adj[e[0]] = append(adj[e[0]], e[1])
	}

	suspicious := make([]bool, n)
	suspicious[k] = true
	stack := []int{k}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for _, next := range adj[node] {
			if !suspicious[next] {
				suspicious[next] = true
				stack = append(stack, next)
			}
		}
	}

	for _, e := range invocations {
		if !suspicious[e[0]] && suspicious[e[1]] {
			// An outside method invokes the group: remove nothing.
			all := make([]int, n)
			for i := range all {
				all[i] = i
			}
			return all
		}
	}

	remaining := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if !suspicious[i] {
			remaining = append(remaining, i)
		}
	}
	return remaining
}
