package main

// 3D DP on the grid: dp[i][j][c] = max score reaching cell (i, j) using
// exactly cost c. From each cell we extend down or right, paying cost 1 for
// any non-zero cell and adding its value to the score. We cap c at
// min(k, m+n-1) because no path can spend more than its length.
// A rolling pair of 2D slices keeps memory at O(n * maxCost).

func maxPathScoreDaily20260430(grid [][]int, k int) int {
	m := len(grid)
	n := len(grid[0])

	maxCost := m + n - 1
	if k < maxCost {
		maxCost = k
	}

	const neg = -1 << 30

	prev := make([][]int, n)
	cur := make([][]int, n)
	for j := 0; j < n; j++ {
		prev[j] = make([]int, maxCost+1)
		cur[j] = make([]int, maxCost+1)
		for c := 0; c <= maxCost; c++ {
			prev[j][c] = neg
			cur[j][c] = neg
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for c := 0; c <= maxCost; c++ {
				cur[j][c] = neg
			}
		}
		for j := 0; j < n; j++ {
			v := grid[i][j]
			costInc := 0
			if v > 0 {
				costInc = 1
			}
			scoreInc := v

			if i == 0 && j == 0 {
				cur[0][0] = 0
				continue
			}

			for c := costInc; c <= maxCost; c++ {
				best := neg
				if i > 0 && prev[j][c-costInc] > best {
					best = prev[j][c-costInc]
				}
				if j > 0 && cur[j-1][c-costInc] > best {
					best = cur[j-1][c-costInc]
				}
				if best != neg {
					cur[j][c] = best + scoreInc
				}
			}
		}
		prev, cur = cur, prev
	}

	ans := neg
	for c := 0; c <= maxCost; c++ {
		if prev[n-1][c] > ans {
			ans = prev[n-1][c]
		}
	}
	if ans == neg {
		return -1
	}
	return ans
}
