package main

// Approach:
// Sort both robots and factories by position. By an exchange argument, there
// exists an optimal assignment in which each factory repairs a contiguous
// block of robots (in sorted order). Define dp[i][j] as the minimum total
// distance to repair robots[i:] using factories[j:]. For each state we either
// skip factory j, or let it repair the next k robots (1..min(limit, n-i)),
// accumulating the cost incrementally. The answer is dp[0][0].
//
// Time:  O(n*m*L) where L is the max factory limit (<= n) -> O(n^2 * m).
// Space: O(n*m).

import (
	"math"
	"sort"
)

func minimumTotalDistance(robot []int, factory [][]int) int64 {
	sort.Ints(robot)
	sort.Slice(factory, func(i, j int) bool {
		return factory[i][0] < factory[j][0]
	})

	n := len(robot)
	m := len(factory)

	const inf = int64(math.MaxInt64 / 4)

	dp := make([][]int64, n+1)
	for i := range dp {
		dp[i] = make([]int64, m+1)
	}
	// dp[n][j] = 0 (no robots left); dp[i<n][m] = inf (no factory available).
	for i := 0; i < n; i++ {
		dp[i][m] = inf
	}

	for j := m - 1; j >= 0; j-- {
		pos := int64(factory[j][0])
		limit := factory[j][1]
		for i := n - 1; i >= 0; i-- {
			// Option 1: skip this factory.
			best := dp[i][j+1]

			// Option 2: assign k robots to factory j.
			maxK := limit
			if maxK > n-i {
				maxK = n - i
			}
			var cost int64
			for k := 1; k <= maxK; k++ {
				diff := pos - int64(robot[i+k-1])
				if diff < 0 {
					diff = -diff
				}
				cost += diff
				if dp[i+k][j+1] < inf {
					total := cost + dp[i+k][j+1]
					if total < best {
						best = total
					}
				}
			}
			dp[i][j] = best
		}
	}

	return dp[0][0]
}
