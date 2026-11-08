package main

// Approach: two-pass counting.
// First pass counts servers in each row and each column. Second pass counts
// every server whose row or column contains more than one server (i.e. at
// least one other server it can communicate with).
// Time: O(m*n), Space: O(m+n).
func countServers(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])

	rowCount := make([]int, m)
	colCount := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				rowCount[i]++
				colCount[j]++
			}
		}
	}

	result := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && (rowCount[i] > 1 || colCount[j] > 1) {
				result++
			}
		}
	}
	return result
}
