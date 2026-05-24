package main

// Jump Game V — maximum number of indices reachable from any starting index.
//
// Approach: memoized DFS / dynamic programming. Define dp[i] as the maximum
// number of indices visitable on a path that starts at index i (including i).
// Every legal jump goes from a taller bar to a strictly shorter one, so the
// jump graph is acyclic and dp[i] only depends on shorter neighbors — which
// makes memoization safe and termination guaranteed.
//
// From index i we scan outward up to d steps in each direction. The blocking
// rule ("arr[i] must exceed every bar between i and j") means the first bar we
// meet that is >= arr[i] blocks everything beyond it, so we break the scan
// immediately on hitting such a bar.
//
// Time:  O(n*d) — each index does two scans of length <= d, each dp computed once.
// Space: O(n) for the memo plus recursion depth.
func maxJumps(arr []int, d int) int {
	n := len(arr)
	memo := make([]int, n)

	var dfs func(i int) int
	dfs = func(i int) int {
		if memo[i] != 0 {
			return memo[i]
		}
		best := 1
		// Scan to the right.
		for x := 1; x <= d && i+x < n; x++ {
			j := i + x
			if arr[j] >= arr[i] {
				break
			}
			if v := 1 + dfs(j); v > best {
				best = v
			}
		}
		// Scan to the left.
		for x := 1; x <= d && i-x >= 0; x++ {
			j := i - x
			if arr[j] >= arr[i] {
				break
			}
			if v := 1 + dfs(j); v > best {
				best = v
			}
		}
		memo[i] = best
		return best
	}

	ans := 0
	for i := 0; i < n; i++ {
		if v := dfs(i); v > ans {
			ans = v
		}
	}
	return ans
}
