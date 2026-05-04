package main

// rotate performs an in-place 90-degree clockwise rotation of an n x n matrix
// by transposing it and then reversing each row. Runs in O(n^2) time and O(1)
// extra space.
func rotate(matrix [][]int) {
	n := len(matrix)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i < n; i++ {
		for left, right := 0, n-1; left < right; left, right = left+1, right-1 {
			matrix[i][left], matrix[i][right] = matrix[i][right], matrix[i][left]
		}
	}
}
