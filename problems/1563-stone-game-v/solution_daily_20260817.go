package main

// 1563. Stone Game V
//
// Interval DP in O(n^2). Let dp[i][j] be the best score Alice can still obtain
// from stoneValue[i..j], and let f(i, j) = sum(i, j) + dp[i][j] be the score she
// gets when segment i..j is the half Bob spares.
//
// For a split at k (left = i..k, right = k+1..j) Bob's move is forced, so every
// branch of the recurrence is a maximum of f over a contiguous range of split
// points: keeping the left half maximizes f(i, k) over the k where the left half
// is not heavier, and keeping the right half maximizes f(k+1, j) over the k where
// it is not lighter. A tie belongs to both ranges, which is exactly Alice picking
// the better side.
//
// Because sum(i, k) grows with k while sum(k+1, j) shrinks, the boundary between
// those two ranges is a single point that only moves right as j grows, so a
// monotone pointer locates it in amortized O(1). The range maxima are prefix
// maxima along a row (maxLeft) and suffix maxima along a column (maxRight), both
// maintained incrementally.
func stoneGameV(stoneValue []int) int {
	n := len(stoneValue)
	if n < 2 {
		return 0
	}

	prefix := make([]int, n+1)
	for i, v := range stoneValue {
		prefix[i+1] = prefix[i] + v
	}
	sum := func(i, j int) int { return prefix[j+1] - prefix[i] }

	// dp[i][j]       best score still obtainable from stoneValue[i..j]
	// maxLeft[i][j]  max over t in [i, j] of sum(i, t) + dp[i][t]
	// maxRight[i][j] max over t in [i, j] of sum(t, j) + dp[t][j]
	dp := make([][]int, n)
	maxLeft := make([][]int, n)
	maxRight := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		maxLeft[i] = make([]int, n)
		maxRight[i] = make([]int, n)
	}

	// i descends and j ascends so that dp[i][j] only reads shorter intervals:
	// row i for the left half, already-finished rows below i for the right half.
	for i := n - 1; i >= 0; i-- {
		// dp[i][i] is 0: a single stone cannot be split.
		maxLeft[i][i] = stoneValue[i]
		maxRight[i][i] = stoneValue[i]

		split := i // smallest k with 2*sum(i, k) >= sum(i, j); never moves back
		for j := i + 1; j < n; j++ {
			total := sum(i, j)
			for split < j && 2*sum(i, split) < total {
				split++
			}

			// k0: last split point where the left half may be kept.
			// k1: first split point where the right half may be kept.
			k0, k1 := split, split
			if split == j || 2*sum(i, split) > total {
				k0 = split - 1
			}

			best := 0
			if k0 >= i {
				best = maxLeft[i][k0]
			}
			if k1 <= j-1 && maxRight[k1+1][j] > best {
				best = maxRight[k1+1][j]
			}
			dp[i][j] = best

			v := total + best
			maxLeft[i][j] = max(maxLeft[i][j-1], v)
			maxRight[i][j] = max(maxRight[i+1][j], v)
		}
	}

	return dp[0][n-1]
}
