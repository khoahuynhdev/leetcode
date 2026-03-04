package main

// numSpecial counts positions (i, j) where mat[i][j] == 1 and all other
// elements in row i and column j are 0. It precomputes row and column sums
// so each cell can be checked in O(1).
func numSpecial(mat [][]int) int {
	m := len(mat)
	if m == 0 {
		return 0
	}
	n := len(mat[0])

	rowSum := make([]int, m)
	colSum := make([]int, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			rowSum[i] += mat[i][j]
			colSum[j] += mat[i][j]
		}
	}

	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 && rowSum[i] == 1 && colSum[j] == 1 {
				count++
			}
		}
	}

	return count
}
