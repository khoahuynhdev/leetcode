package solution

import (
	"sort"
)

// https://leetcode.com/problems/largest-submatrix-with-rearrangements/editorial/?envType=daily-question&envId=2023-11-26
func largestSubmatrix(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])
	ans := 0
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if matrix[row][col] != 0 && row > 0 {
				matrix[row][col] += matrix[row-1][col]
			}
		}
		curr_row := make([]int, len(matrix[row]))
		copy(curr_row, matrix[row])
		sort.Sort(sort.Reverse(sort.IntSlice(curr_row)))
		for i := 0; i < n; i++ {
			if ans < curr_row[i]*(i+1) {
				ans = curr_row[i] * (i + 1)
			}
		}
	}
	return ans
}
