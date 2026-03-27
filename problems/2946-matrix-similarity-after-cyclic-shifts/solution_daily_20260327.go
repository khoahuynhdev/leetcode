package main

// areSimilar checks if the matrix is identical after k cyclic shifts
// (left for even rows, right for odd rows). Since both shift directions
// require the same periodicity condition, we simply verify that
// mat[i][j] == mat[i][(j + k%n) % n] for every element.
func areSimilar(mat [][]int, k int) bool {
	n := len(mat[0])
	shift := k % n
	if shift == 0 {
		return true
	}
	for _, row := range mat {
		for j := 0; j < n; j++ {
			if row[j] != row[(j+shift)%n] {
				return false
			}
		}
	}
	return true
}
