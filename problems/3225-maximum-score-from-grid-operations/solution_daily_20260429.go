package main

// Approach: reformulate the problem as choosing a height h[j] in 0..n for each
// column. A white cell (i, j) scores iff i >= h[j] and i < max(h[j-1], h[j+1]).
// DP over columns with state (p, c) = (h[j-1], h[j]); per-column prefix sums
// give O(1) contributions. The naive O(n^4) inner max-over-q collapses to
// max(M1, M2 + K_c) using the monotonicity of prefix sums, yielding O(n^3).

func maximumScore(grid [][]int) int64 {
	n := len(grid)
	if n == 1 {
		return 0
	}

	ps := make([][]int64, n)
	for j := 0; j < n; j++ {
		ps[j] = make([]int64, n+1)
		for i := 0; i < n; i++ {
			ps[j][i+1] = ps[j][i] + int64(grid[i][j])
		}
	}

	const negInf = int64(-1) << 62

	dp := make([][]int64, n+1)
	for i := range dp {
		dp[i] = make([]int64, n+1)
		for k := range dp[i] {
			dp[i][k] = negInf
		}
	}
	for c := 0; c <= n; c++ {
		dp[0][c] = 0
	}

	for j := 1; j < n; j++ {
		col := j - 1
		newDp := make([][]int64, n+1)
		for i := range newDp {
			newDp[i] = make([]int64, n+1)
			for k := range newDp[i] {
				newDp[i][k] = negInf
			}
		}

		for p := 0; p <= n; p++ {
			H := ps[col][p]
			M1, M2 := negInf, negInf
			for q := 0; q <= n; q++ {
				v := dp[q][p]
				if v == negInf {
					continue
				}
				if v > M2 {
					M2 = v
				}
				kq := ps[col][q] - H
				if kq < 0 {
					kq = 0
				}
				if v+kq > M1 {
					M1 = v + kq
				}
			}
			if M2 == negInf {
				continue
			}
			for c := 0; c <= n; c++ {
				kc := ps[col][c] - H
				if kc < 0 {
					kc = 0
				}
				best := M1
				if M2+kc > best {
					best = M2 + kc
				}
				newDp[p][c] = best
			}
		}

		dp = newDp
	}

	answer := int64(0)
	last := n - 1
	for p := 0; p <= n; p++ {
		for c := 0; c <= n; c++ {
			v := dp[p][c]
			if v == negInf {
				continue
			}
			var contrib int64
			if p > c {
				contrib = ps[last][p] - ps[last][c]
			}
			if v+contrib > answer {
				answer = v + contrib
			}
		}
	}

	return answer
}
