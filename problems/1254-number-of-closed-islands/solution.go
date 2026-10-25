package main

// Approach: two-pass flood fill.
// 1. Flood-fill every land cell on the border to water; those islands can't be closed.
// 2. Count the remaining islands. Each is fully surrounded by water.
// Time: O(m*n), Space: O(m*n) recursion stack in the worst case.

func closedIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	var sink func(r, c int)
	sink = func(r, c int) {
		if r < 0 || r >= m || c < 0 || c >= n || grid[r][c] != 0 {
			return
		}
		grid[r][c] = 1
		sink(r+1, c)
		sink(r-1, c)
		sink(r, c+1)
		sink(r, c-1)
	}

	for r := 0; r < m; r++ {
		sink(r, 0)
		sink(r, n-1)
	}
	for c := 0; c < n; c++ {
		sink(0, c)
		sink(m-1, c)
	}

	count := 0
	for r := 1; r < m-1; r++ {
		for c := 1; c < n-1; c++ {
			if grid[r][c] == 0 {
				count++
				sink(r, c)
			}
		}
	}
	return count
}
