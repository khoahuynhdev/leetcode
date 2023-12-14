package solution

func onesMinusZeros(grid [][]int) [][]int {
	n := len(grid)
	m := len(grid[0])
	row := make([]int, n)
	col := make([]int, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			row[i] += grid[i][j]
			col[j] += grid[i][j]
		}
	}
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			ans[i][j] = row[i] + col[j] - (m - row[i]) - (n - col[j])
		}
	}

	return ans
}
