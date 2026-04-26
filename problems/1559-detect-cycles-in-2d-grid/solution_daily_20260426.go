package main

// Approach: DFS with parent tracking on each same-value connected component.
// In an undirected graph, a cycle exists iff DFS reaches an already-visited
// node that is not the immediate parent. On a 4-directional grid, the
// smallest undirected cycle has length 4, so the "length >= 4" constraint
// is satisfied automatically. A single shared visited matrix is enough
// because a component fully explored without a cycle is proven cycle-free.

func containsCycleDaily20260426(grid [][]byte) bool {
	m := len(grid)
	if m == 0 {
		return false
	}
	n := len(grid[0])
	if n == 0 {
		return false
	}

	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	var dfs func(r, c, pr, pc int, val byte) bool
	dfs = func(r, c, pr, pc int, val byte) bool {
		visited[r][c] = true
		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if nr < 0 || nr >= m || nc < 0 || nc >= n {
				continue
			}
			if grid[nr][nc] != val {
				continue
			}
			if nr == pr && nc == pc {
				continue
			}
			if visited[nr][nc] {
				return true
			}
			if dfs(nr, nc, r, c, val) {
				return true
			}
		}
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !visited[i][j] {
				if dfs(i, j, -1, -1, grid[i][j]) {
					return true
				}
			}
		}
	}
	return false
}
