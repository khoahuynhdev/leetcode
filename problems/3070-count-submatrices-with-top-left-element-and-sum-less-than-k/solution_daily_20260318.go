package main

// countSubmatrices counts submatrices containing the top-left element (0,0)
// whose sum is <= k. We compute 2D prefix sums in-place and count valid cells.
// Since all values are non-negative, we break early when a row's running sum exceeds k.
func countSubmatrices(grid [][]int, k int) int {
	m := len(grid)
	n := len(grid[0])
	count := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i > 0 {
				grid[i][j] += grid[i-1][j]
			}
			if j > 0 {
				grid[i][j] += grid[i][j-1]
			}
			if i > 0 && j > 0 {
				grid[i][j] -= grid[i-1][j-1]
			}
			if grid[i][j] <= k {
				count++
			} else {
				break // values are non-negative, so wider submatrices only grow
			}
		}
	}

	return count
}
