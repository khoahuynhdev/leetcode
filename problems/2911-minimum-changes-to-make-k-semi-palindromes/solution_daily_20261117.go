package main

// Approach: interval cost + partition DP.
// 1. For every substring s[i..j] (length >= 2), compute cost[i][j]: the minimum
//    changes to make it a semi-palindrome. For each proper divisor d of the
//    length, the substring splits into d interleaved groups; each group must be
//    a palindrome, so count mismatched mirrored pairs within each group.
//    Take the minimum over all d.
// 2. dp[p][j] = minimum changes to split the prefix s[0..j) into p parts.
//    dp[p][j] = min over i of dp[p-1][i] + cost[i][j-1].
// Time: O(n^3 * D + k * n^2), where D is the max number of divisors of any
// length <= n (D <= 12 for n <= 200). Each divisor check costs O(L).
// Space: O(n^2) for the cost table.

func minimumChanges(s string, k int) int {
	n := len(s)
	const inf = 1 << 30

	// divisors[L] holds the proper divisors d of L (1 <= d < L).
	divisors := make([][]int, n+1)
	for d := 1; d <= n; d++ {
		for L := 2 * d; L <= n; L += d {
			divisors[L] = append(divisors[L], d)
		}
	}

	cost := make([][]int, n)
	for i := range cost {
		cost[i] = make([]int, n)
		for j := range cost[i] {
			cost[i][j] = inf
		}
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			L := j - i + 1
			best := inf
			for _, d := range divisors[L] {
				m := L / d // members per group
				changes := 0
				for g := 0; g < d; g++ {
					for lo, hi := 0, m-1; lo < hi; lo, hi = lo+1, hi-1 {
						if s[i+g+lo*d] != s[i+g+hi*d] {
							changes++
						}
					}
				}
				if changes < best {
					best = changes
				}
			}
			cost[i][j] = best
		}
	}

	// dp[p][j]: min changes for prefix of length j split into p parts.
	prev := make([]int, n+1)
	for j := range prev {
		prev[j] = inf
	}
	prev[0] = 0
	for p := 1; p <= k; p++ {
		cur := make([]int, n+1)
		for j := range cur {
			cur[j] = inf
		}
		for j := 2 * p; j <= n; j++ {
			for i := 2 * (p - 1); i <= j-2; i++ {
				if prev[i] == inf {
					continue
				}
				if v := prev[i] + cost[i][j-1]; v < cur[j] {
					cur[j] = v
				}
			}
		}
		prev = cur
	}
	return prev[n]
}
