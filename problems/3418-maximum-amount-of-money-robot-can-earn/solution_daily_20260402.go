package main

// 3D DP approach: dp[i][j][k] = max coins reaching (i,j) with k neutralizations used.
// At negative cells, we choose between taking the loss or spending a neutralization.
func maximumAmount(coins [][]int) int {
	m := len(coins)
	n := len(coins[0])

	const maxK = 3
	const negInf = -1<<60

	// dp[i][j][k] = max coins at (i,j) with k neutralizations used
	dp := make([][][3]int, m)
	for i := range dp {
		dp[i] = make([][3]int, n)
		for j := range dp[i] {
			dp[i][j] = [3]int{negInf, negInf, negInf}
		}
	}

	// Base case: starting cell
	dp[0][0][0] = coins[0][0]
	if coins[0][0] < 0 {
		dp[0][0][1] = 0 // neutralize starting cell
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			for k := 0; k < maxK; k++ {
				// Best value arriving at (i,j) with k neutralizations used (before processing this cell)
				best := negInf
				if i > 0 && dp[i-1][j][k] > best {
					best = dp[i-1][j][k]
				}
				if j > 0 && dp[i][j-1][k] > best {
					best = dp[i][j-1][k]
				}

				if best == negInf {
					// Can't reach (i,j) with k neutralizations
				} else {
					// Option 1: collect (or lose) coins normally
					val := best + coins[i][j]
					if val > dp[i][j][k] {
						dp[i][j][k] = val
					}
				}

				// Option 2: neutralize this cell (only useful if coins < 0, costs one neutralization)
				if coins[i][j] < 0 && k > 0 {
					prev := negInf
					if i > 0 && dp[i-1][j][k-1] > prev {
						prev = dp[i-1][j][k-1]
					}
					if j > 0 && dp[i][j-1][k-1] > prev {
						prev = dp[i][j-1][k-1]
					}
					if prev != negInf {
						val := prev // neutralize: add 0 instead of coins[i][j]
						if val > dp[i][j][k] {
							dp[i][j][k] = val
						}
					}
				}
			}
		}
	}

	ans := negInf
	for k := 0; k < maxK; k++ {
		if dp[m-1][n-1][k] > ans {
			ans = dp[m-1][n-1][k]
		}
	}
	return ans
}
