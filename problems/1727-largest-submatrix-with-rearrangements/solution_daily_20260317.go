package solution

import "sort"

// largestSubmatrix computes the largest all-1 submatrix area after optimally
// rearranging columns. It builds consecutive-1 heights per column, then sorts
// each row's heights descending to greedily compute the best rectangle.
func largestSubmatrixDaily(matrix [][]int) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])

	// Build height matrix in place: matrix[i][j] = consecutive 1s ending at row i.
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 {
				matrix[i][j] += matrix[i-1][j]
			}
		}
	}

	maxArea := 0
	row := make([]int, n)

	for i := 0; i < m; i++ {
		copy(row, matrix[i])
		sort.Sort(sort.Reverse(sort.IntSlice(row)))
		for j := 0; j < n; j++ {
			if row[j] == 0 {
				break
			}
			area := row[j] * (j + 1)
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}
