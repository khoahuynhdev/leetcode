package main

// Rotate mat by 90° clockwise up to 4 times and check if any rotation equals target.
// Time: O(n²), Space: O(n²)
func findRotation(mat [][]int, target [][]int) bool {
	n := len(mat)
	for k := 0; k < 4; k++ {
		if equal(mat, target) {
			return true
		}
		mat = rotate90(mat, n)
	}
	return false
}

func rotate90(mat [][]int, n int) [][]int {
	rotated := make([][]int, n)
	for i := range rotated {
		rotated[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			rotated[j][n-1-i] = mat[i][j]
		}
	}
	return rotated
}

func equal(a, b [][]int) bool {
	for i := range a {
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
