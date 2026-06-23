package main

// 3699. Number of ZigZag Arrays I
//
// A ZigZag array of length n has values in [l, r], no two adjacent elements
// equal, and no three consecutive elements strictly increasing or decreasing.
// Those two latter rules force the array to strictly alternate direction
// (up, down, up, ... or down, up, down, ...).
//
// We count with DP over the array length. For each length we track two rows
// indexed by the last value:
//
//	up[v]   = # valid arrays ending at value v whose last step went up   (prev < v)
//	down[v] = # valid arrays ending at value v whose last step went down (prev > v)
//
// Because directions must alternate, the transitions are:
//
//	newUp[v]   = sum of down[u] for u < v   (prefix sum of the old down row)
//	newDown[v] = sum of up[u]   for u > v   (suffix sum of the old up row)
//
// Using running prefix/suffix sums makes each transition O(m), where
// m = r - l + 1. Total time O(n * m), space O(m). Answer is taken mod 1e9+7.
func zigZagArrays(n int, l int, r int) int {
	const mod = 1_000_000_007
	m := r - l + 1

	// Base case: a single element is valid in either orientation.
	up := make([]int, m)
	down := make([]int, m)
	for v := 0; v < m; v++ {
		up[v] = 1
		down[v] = 1
	}

	for i := 2; i <= n; i++ {
		newUp := make([]int, m)
		newDown := make([]int, m)

		// newUp[v] = sum of down[u] for u < v (left-to-right prefix sum).
		prefix := 0
		for v := 0; v < m; v++ {
			newUp[v] = prefix
			prefix = (prefix + down[v]) % mod
		}

		// newDown[v] = sum of up[u] for u > v (right-to-left suffix sum).
		suffix := 0
		for v := m - 1; v >= 0; v-- {
			newDown[v] = suffix
			suffix = (suffix + up[v]) % mod
		}

		up, down = newUp, newDown
	}

	ans := 0
	for v := 0; v < m; v++ {
		ans = (ans + up[v] + down[v]) % mod
	}
	return ans
}
