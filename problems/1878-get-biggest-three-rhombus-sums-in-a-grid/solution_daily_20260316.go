package main

import "sort"

// getBiggestThree enumerates all rhombuses by center and size,
// collects distinct border sums, and returns the top 3 in descending order.
func getBiggestThree(grid [][]int) []int {
	m, n := len(grid), len(grid[0])
	set := map[int]bool{}

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			// Size 0: single cell
			set[grid[r][c]] = true

			// Size s >= 1: walk the four sides of the rhombus border
			for s := 1; r-s >= 0 && r+s < m && c-s >= 0 && c+s < n; s++ {
				sum := 0
				for i := 0; i < s; i++ {
					sum += grid[r-s+i][c+i]   // top to right
					sum += grid[r+i][c+s-i]   // right to bottom
					sum += grid[r+s-i][c-i]   // bottom to left
					sum += grid[r-i][c-s+i]   // left to top
				}
				set[sum] = true
			}
		}
	}

	sums := make([]int, 0, len(set))
	for v := range set {
		sums = append(sums, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sums)))
	if len(sums) > 3 {
		sums = sums[:3]
	}
	return sums
}
