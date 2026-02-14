package main

// Champagne Tower - DP Simulation
// Simulate the pour row by row. Track total liquid at each glass position.
// If a glass overflows (> 1 cup), distribute excess equally to the two glasses below.
// Time: O(query_row^2), Space: O(query_row^2)
func champagneTowerDaily(poured int, query_row int, query_glass int) float64 {
	// dp[i][j] = total liquid that has flowed into glass (i, j)
	dp := make([][]float64, query_row+1)
	for i := range dp {
		dp[i] = make([]float64, i+2)
	}

	dp[0][0] = float64(poured)

	for i := 0; i < query_row; i++ {
		for j := 0; j <= i; j++ {
			if dp[i][j] > 1.0 {
				overflow := (dp[i][j] - 1.0) / 2.0
				dp[i+1][j] += overflow
				dp[i+1][j+1] += overflow
			}
		}
	}

	if dp[query_row][query_glass] > 1.0 {
		return 1.0
	}
	return dp[query_row][query_glass]
}
