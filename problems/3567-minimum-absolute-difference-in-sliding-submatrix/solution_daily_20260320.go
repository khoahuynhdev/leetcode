package main

import "sort"

// Approach: For each k x k submatrix, collect elements, sort them,
// and find the minimum difference between consecutive sorted values.
func minimumAbsoluteDifference(grid [][]int, k int) [][]int {
	m := len(grid)
	n := len(grid[0])
	rows := m - k + 1
	cols := n - k + 1
	ans := make([][]int, rows)

	for i := 0; i < rows; i++ {
		ans[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			// Collect all elements in the k x k submatrix
			vals := make([]int, 0, k*k)
			for r := i; r < i+k; r++ {
				for c := j; c < j+k; c++ {
					vals = append(vals, grid[r][c])
				}
			}
			sort.Ints(vals)

			minDiff := 0
			if len(vals) > 1 {
				minDiff = vals[len(vals)-1] - vals[0] // max possible diff as initial
				if minDiff == 0 {
					// all elements are the same
					ans[i][j] = 0
					continue
				}
				for p := 1; p < len(vals); p++ {
					d := vals[p] - vals[p-1]
					if d > 0 && d < minDiff {
						minDiff = d
					}
				}
			}
			ans[i][j] = minDiff
		}
	}
	return ans
}
