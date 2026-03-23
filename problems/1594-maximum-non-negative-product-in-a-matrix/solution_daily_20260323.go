package main

// Maximum Non Negative Product in a Matrix
//
// DP approach tracking both max and min products at each cell,
// since multiplying by a negative flips the extremes.
// At the end, return maxProd[m-1][n-1] mod 1e9+7 if non-negative, else -1.

func maxProductPath(grid [][]int) int {
	const mod = 1_000_000_007
	m, n := len(grid), len(grid[0])

	maxDP := make([][]int64, m)
	minDP := make([][]int64, m)
	for i := range maxDP {
		maxDP[i] = make([]int64, n)
		minDP[i] = make([]int64, n)
	}

	maxDP[0][0] = int64(grid[0][0])
	minDP[0][0] = int64(grid[0][0])

	// First row
	for j := 1; j < n; j++ {
		v := int64(grid[0][j])
		maxDP[0][j] = maxDP[0][j-1] * v
		minDP[0][j] = minDP[0][j-1] * v
	}

	// First column
	for i := 1; i < m; i++ {
		v := int64(grid[i][0])
		maxDP[i][0] = maxDP[i-1][0] * v
		minDP[i][0] = minDP[i-1][0] * v
	}

	// Fill rest
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			v := int64(grid[i][j])
			// Candidates from top and left
			a := maxDP[i-1][j] * v
			b := minDP[i-1][j] * v
			c := maxDP[i][j-1] * v
			d := minDP[i][j-1] * v

			maxDP[i][j] = max64(max64(a, b), max64(c, d))
			minDP[i][j] = min64(min64(a, b), min64(c, d))
		}
	}

	result := maxDP[m-1][n-1]
	if result < 0 {
		return -1
	}
	return int(result % mod)
}

func max64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func min64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
