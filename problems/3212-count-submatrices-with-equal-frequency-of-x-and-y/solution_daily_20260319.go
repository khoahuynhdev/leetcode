package main

// countSubmatrices counts submatrices anchored at (0,0) with equal frequency
// of 'X' and 'Y' and at least one 'X', using 2D prefix sums.
func countSubmatrices(grid [][]byte) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])

	// 1-indexed prefix sums for 'X' and 'Y' counts
	prefX := make([][]int, m+1)
	prefY := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		prefX[i] = make([]int, n+1)
		prefY[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			x, y := 0, 0
			if grid[i-1][j-1] == 'X' {
				x = 1
			} else if grid[i-1][j-1] == 'Y' {
				y = 1
			}
			prefX[i][j] = prefX[i-1][j] + prefX[i][j-1] - prefX[i-1][j-1] + x
			prefY[i][j] = prefY[i-1][j] + prefY[i][j-1] - prefY[i-1][j-1] + y
		}
	}

	count := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			cx := prefX[i][j]
			cy := prefY[i][j]
			if cx == cy && cx > 0 {
				count++
			}
		}
	}
	return count
}
