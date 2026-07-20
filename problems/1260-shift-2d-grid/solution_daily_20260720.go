package main

// Shift 2D Grid (LeetCode 1260)
//
// A single shift moves every element one position forward when the grid is
// read row by row, with the last cell wrapping to the first. Flattening the
// grid into m*n cells turns k shifts into a cyclic shift by k. An element at
// (i, j) has flat index i*n+j; after k shifts its destination flat index is
// (i*n + j + k) % (m*n), which we convert back to (row, col) via /n and %n.
// A fresh result grid avoids overwriting cells we still need to read.
//
// Time:  O(m*n) — each cell read and written once, independent of k.
// Space: O(m*n) — the returned result grid.
func shiftGrid(grid [][]int, k int) [][]int {
	m := len(grid)
	if m == 0 {
		return grid
	}
	n := len(grid[0])
	if n == 0 {
		return grid
	}

	total := m * n
	k %= total

	result := make([][]int, m)
	for i := range result {
		result[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dest := (i*n + j + k) % total
			result[dest/n][dest%n] = grid[i][j]
		}
	}

	return result
}
