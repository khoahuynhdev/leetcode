package main

// Flip Square Submatrix Vertically
// Approach: Two-pointer row swap within the k×k submatrix starting at (x, y).
// Swap rows from top and bottom moving inward, only touching columns y..y+k-1.
// Time: O(k²), Space: O(1)

func flipSubmatrix(grid [][]int, x int, y int, k int) [][]int {
	top, bottom := x, x+k-1
	for top < bottom {
		for j := y; j < y+k; j++ {
			grid[top][j], grid[bottom][j] = grid[bottom][j], grid[top][j]
		}
		top++
		bottom--
	}
	return grid
}
