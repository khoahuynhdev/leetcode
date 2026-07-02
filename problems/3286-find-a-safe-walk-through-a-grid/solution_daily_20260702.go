package main

// Find a Safe Walk Through a Grid (LeetCode 3286)
//
// Approach: 0-1 BFS shortest path.
// Each move into a cell costs `grid[i][j]` health (0 for safe, 1 for unsafe),
// so every edge weight is 0 or 1. We compute dist[i][j] = the minimum number of
// unsafe cells on any path reaching (i, j), counting the start cell's cost.
// A deque keeps cells ordered by distance: 0-cost moves go to the front,
// 1-cost moves go to the back. We survive iff health - dist[end] >= 1.

func findSafeWalk(grid [][]int, health int) bool {
	m := len(grid)
	n := len(grid[0])

	const inf = int(1e9)
	dist := make([][]int, m)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = inf
		}
	}

	type cell struct{ r, c int }
	deque := make([]cell, 0, m*n)
	dist[0][0] = grid[0][0]
	deque = append(deque, cell{0, 0})

	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(deque) > 0 {
		// Pop from the front.
		cur := deque[0]
		deque = deque[1:]

		for _, d := range dirs {
			nr, nc := cur.r+d[0], cur.c+d[1]
			if nr < 0 || nr >= m || nc < 0 || nc >= n {
				continue
			}
			nd := dist[cur.r][cur.c] + grid[nr][nc]
			if nd < dist[nr][nc] {
				dist[nr][nc] = nd
				if grid[nr][nc] == 0 {
					// 0-cost edge: same distance layer, push to front.
					deque = append([]cell{{nr, nc}}, deque...)
				} else {
					// 1-cost edge: next distance layer, push to back.
					deque = append(deque, cell{nr, nc})
				}
			}
		}
	}

	return health-dist[m-1][n-1] >= 1
}
