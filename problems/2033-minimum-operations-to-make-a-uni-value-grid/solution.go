package main

// Approach: flatten the grid, verify all cells share the same remainder mod x
// (otherwise -1), then make every cell equal to the median — the value that
// minimizes the sum of absolute deviations. Each cell needs |v - median| / x
// operations.

import "sort"

func minOperations(grid [][]int, x int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}

	flat := make([]int, 0, m*n)
	target := grid[0][0] % x
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j]%x != target {
				return -1
			}
			flat = append(flat, grid[i][j])
		}
	}

	sort.Ints(flat)
	median := flat[len(flat)/2]

	ops := 0
	for _, v := range flat {
		d := v - median
		if d < 0 {
			d = -d
		}
		ops += d / x
	}
	return ops
}
