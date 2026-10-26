package main

// Approach: flood fill each island with an iterative DFS (explicit stack),
// summing its cell values as we go. Visited cells are zeroed out so each
// cell is processed once. After an island is fully explored, check whether
// its total is divisible by k. Sums can reach 1e5 * 1e6 = 1e11, which fits
// in Go's 64-bit int.
//
// Time: O(m*n), Space: O(m*n) for the stack in the worst case.
// Note: the input grid is modified in place.
func countIslands(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	dirs := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	count := 0
	stack := make([][2]int, 0)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}

			sum := grid[i][j]
			grid[i][j] = 0
			stack = append(stack[:0], [2]int{i, j})

			for len(stack) > 0 {
				cur := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				for _, d := range dirs {
					r, c := cur[0]+d[0], cur[1]+d[1]
					if r < 0 || r >= m || c < 0 || c >= n || grid[r][c] == 0 {
						continue
					}
					sum += grid[r][c]
					grid[r][c] = 0
					stack = append(stack, [2]int{r, c})
				}
			}

			if sum%k == 0 {
				count++
			}
		}
	}

	return count
}
